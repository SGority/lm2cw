package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	sentry "github.com/getsentry/sentry-go"
	log "github.com/magna5/go-logger"
)

// func (s *server) ReadUser(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	log.Ctx(ctx).Debug().Msg("ReadUser")
// 	uid := chi.URLParam(r, "userID")
// 	u, err := s.db.User.Query().Where(user.UserIDEQ(uid)).Only(ctx)
// 	if err != nil {
// 		render.Render(w, r, ErrServerError(r, err))
// 		return
// 	}
// 	render.JSON(w, r, u)
// 	return
// }

//DeviceDetails struct
type DeviceDetails struct {
	ID   float64
	Name string
}

//Device struct
type Device struct {
	Devices      []DeviceDetails
	CompanyNames []string
}

// GetSync runs a triggered sync
func (s *server) GetSync(w http.ResponseWriter, r *http.Request) {
	go func() {
		err := LM2CW(s.config)
		if err != nil {
			log.Errorf("Error running triggered sync: %v\n", err)
		} else {
			log.Info("Triggered sync: All devices synchronized successfully")
		}
	}()
	w.Write([]byte("Device sync has started"))
	w.WriteHeader(http.StatusAccepted)
	return
}

func deviceAuth(conf *Cfg, timestamp int64) string {
	accessID := conf.LmAccessID
	accessKey := conf.LmAccessKey
	httpMethod := "GET"
	resourcePath := conf.DeviceSourcePath

	payloadBody := fmt.Sprintf("%s%d%s%s", httpMethod, timestamp, "", resourcePath)
	token := generateHmacToken(accessKey, payloadBody)
	signature := b64.StdEncoding.EncodeToString([]byte(token))
	return fmt.Sprintf("LMv1 %s:%s:%d", accessID, signature, timestamp)
}

func generateHmacToken(accessKey, payloadBody string) string {
	mac := hmac.New(sha256.New, []byte(accessKey))
	mac.Write([]byte(payloadBody))
	expectedMAC := mac.Sum(nil)
	return "" + hex.EncodeToString(expectedMAC)
}

//FetchDevices func
func FetchDevices(conf *Cfg, page int64) ([]map[string]interface{}, error) {
	responseItems := make([]map[string]interface{}, 0)
	timestamp := time.Now().UnixNano() / 1000000
	authorization := deviceAuth(conf, timestamp)
	offset := int64(conf.DeviceOffsetSize)
	url := fmt.Sprintf("%s%s?&offset=%d&size=%d", conf.BaseURL, conf.DeviceSourcePath, page*offset, offset)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		sentry.CaptureException(err)
		return responseItems, err
	}

	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}}
	resp, err := client.Do(req)

	if err != nil {
		sentry.CaptureException(err)
		return responseItems, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sentry.CaptureException(err)
		return responseItems, err
	}

	responseItems, err = getDeviceItems(body)
	if err != nil {
		sentry.CaptureException(err)
		return responseItems, err
	}

	responseItems = flattenDeviceItems(responseItems)

	return responseItems, err
}

func flattenDeviceItems(responseItems []map[string]interface{}) []map[string]interface{} {
	responseItems = flattenJSON(responseItems, "inheritedProperties")
	responseItems = flattenJSON(responseItems, "autoProperties")
	responseItems = flattenJSON(responseItems, "systemProperties")
	responseItems = flattenJSON(responseItems, "customProperties")
	return responseItems
}

func getDeviceItems(body []byte) ([]map[string]interface{}, error) {
	responseItems := make([]map[string]interface{}, 0)

	var response map[string]interface{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		sentry.CaptureException(err)
		return responseItems, nil
	}

	var data map[string]interface{}
	data = response["data"].(map[string]interface{})

	var items []interface{}
	items = data["items"].([]interface{})

	responseItems = make([]map[string]interface{}, 0)
	for _, result := range items {
		responseItems = append(responseItems, result.(map[string]interface{}))
	}
	return responseItems, nil
}

func flattenJSON(results []map[string]interface{}, property string) []map[string]interface{} {
	for _, result := range results {
		properties := result[property].([]interface{})
		for _, v := range properties {
			p := v.(map[string]interface{})
			result[p["name"].(string)] = p["value"].(string)
		}
		delete(result, property)
	}
	return results
}

//CWAddUpdate func
func CWAddUpdate(conf *Cfg, lmres []map[string]interface{}) error {

	var DevDetails DeviceDetails
	var DevMail Device

	var cwerr error
	for i := 0; i < len(lmres); i++ {
		deviceType := lmres[i]["deviceType"]

		//Set CW attributes
		CWAttributes := setCWAttributes(lmres[i])
		llog := log.WithFields(map[string]interface{}{
			"deviceType": deviceType,
			"device":     CWAttributes["name"],
			"company":    CWAttributes["company"],
			"cw_type":    lmres[i]["cw_type"],
		})

		if deviceType.(float64) == 0 {

			if lmres[i]["cw_type"] != nil {
				var cwtype CWType
				tagName := (lmres[i]["cw_type"]).(string)

				//check CW type
				getType, err := getCwTypesByName(conf, tagName)
				if err != nil {
					sentry.CaptureException(err)
					llog.Error("Unable to get CW type response", err)
					ErrorCounter.Inc()
					cwerr = err
					continue
				}

				err = json.Unmarshal(getType, &cwtype)
				if err != nil {
					sentry.CaptureException(err)
					llog.Error(err)
					ErrorCounter.Inc()
					cwerr = err
					continue
				}

				if len(cwtype) != 0 {
					//set cw type field
					CWAttributes["type"] = make(map[string]interface{})
					CWAttributes["type"].(map[string]interface{})["id"] = cwtype[0].ID
					CWAttributes["type"].(map[string]interface{})["name"] = cwtype[0].Name

					//check CW company
					var comp Company
					if CWAttributes["company"] != nil {
						com := (CWAttributes["company"]).(string)
						res, err := getCwCompaniesByName(conf, com)
						if err != nil {
							llog.Error("Unable to get CW company response", err)
							ErrorCounter.Inc()
							cwerr = err
							continue
						}

						err = json.Unmarshal(res, &comp)
						if err != nil {
							sentry.CaptureException(err)
							llog.Error(err)
							ErrorCounter.Inc()
							cwerr = err
							continue
						}
						if len(comp) != 0 {

							//set CW company field
							CWAttributes["company"] = make(map[string]interface{})
							CWAttributes["company"].(map[string]interface{})["id"] = comp[0].ID
							CWAttributes["company"].(map[string]interface{})["identifier"] = comp[0].Identifier
							CWAttributes["company"].(map[string]interface{})["name"] = comp[0].Name

							upattr := updateAttributes(CWAttributes)
							compName := comp[0].Name
							deviceName := (CWAttributes["name"]).(string)
							llog.Debug("Updating ConnectWise")
							_, err := AddOrUpdate(conf, deviceName, compName, CWAttributes, upattr)
							if err != nil {
								llog.Error(err)
								ErrorCounter.Inc()
								cwerr = err
								continue
							}
							DevicesSynchronizedGauge.WithLabelValues(compName).Inc()
						} else {
							llog.Warn("Company not found in Connectwise")
							DevMail.CompanyNames = append(DevMail.CompanyNames, com)
							CompanyNotFound.WithLabelValues(com).Inc()
						}

					} else {
						DevDetails.ID = (lmres[i]["id"].(float64))
						DevDetails.Name = (lmres[i]["displayName"]).(string)
						DevMail.Devices = append(DevMail.Devices, DevDetails)
						CompanyNotSet.WithLabelValues(DevDetails.Name).Inc()
						llog.Warnf("No company set for %s\n", DevDetails.Name)
					}

				} else { // len(cwtype == 0)
					llog.Warn("Type not found in ConnectWise")
				}
			} else { // cw_type == nil
				llog.Info("No cw_type specified, skipping")
			}
		} else { // deviceType != 0
			llog.Info("DeviceType != 0, skipping")
		}
	}

	if DevMail.CompanyNames != nil || DevMail.Devices != nil {
		err := SendMail(conf, DevMail)
		if err != nil {
			sentry.CaptureException(err)
			log.Error(err)
			return err
		}
	}

	return cwerr
}

//AddOrUpdate func
func AddOrUpdate(conf *Cfg, devname, compname string, data, updata map[string]interface{}) ([]byte, error) {
	var config CwConfig
	updlog := log.WithFields(map[string]interface{}{
		"device":  devname,
		"company": compname,
	})
	res, err := getCwConfigurationsByName(conf, devname)
	if err != nil {
		sentry.CaptureException(err)
		updlog.Error("Unable to get the CW configuration response", err)
		return nil, err
	}

	err = json.Unmarshal(res, &config)
	if err != nil {
		sentry.CaptureException(err)
		updlog.Error(err)
		return nil, err
	}

	if len(config) != 0 && config[0].Comp.Name == compname {
		id := strconv.Itoa(config[0].ID)

		patchMap := []map[string]interface{}{}

		for key, value := range updata {
			patchData := map[string]interface{}{"op": "replace",
				"path":  key,
				"value": value}
			patchMap = append(patchMap, patchData)
		}

		jsonData, err := json.Marshal(patchMap)
		if err != nil {
			sentry.CaptureException(err)
			updlog.Error(err)
			return nil, err
		}

		res, err := updateDeviceInCw(conf, id, jsonData)
		if err != nil {
			updlog.Error("Unable to update device", err)
			ErrorCounter.Inc()
			return res, err
		}
		updlog.Info("Successfully updated device")
	} else {
		jsonData, err := json.Marshal(data)
		if err != nil {
			sentry.CaptureException(err)
			updlog.Error(err)
			return nil, err
		}

		res, err := addDeviceToCw(conf, jsonData)
		if err != nil {
			updlog.Error("Unable to add the device", err)
			ErrorCounter.Inc()
			return res, err
		}
		updlog.Info("Successfully added device")
	}

	return res, err

}

func setCWAttributes(lmap map[string]interface{}) map[string]interface{} {
	CWAttributes := make(map[string]interface{})

	res := lmap
	for i := 0; i < len(res); i++ {

		CWAttributes["name"] = res["displayName"]
		CWAttributes["modelNumber"] = res["system.model"]
		CWAttributes["type"] = res["cw_type"]
		CWAttributes["notes"] = res["description"]
		CWAttributes["osInfo"] = ""
		if osInfo, ok := res["system.sysinfo"].(string); ok {
			if len(osInfo) > 250 {
				CWAttributes["osInfo"] = osInfo[:250]
			} else {
				CWAttributes["osInfo"] = osInfo
			}
		}
		CWAttributes["company"] = res["customer.name"]
		if res["system.ips"] != nil {
			cwIP := (res["system.ips"]).(string)
			ipvalue := strings.Split(cwIP, ",")
			CWAttributes["ipAddress"] = ipvalue[0]
		}

		return CWAttributes

	}

	return CWAttributes

}

func updateAttributes(attrmap map[string]interface{}) map[string]interface{} {

	UpAttributes := make(map[string]interface{})
	UpAttributes["name"] = attrmap["name"]
	UpAttributes["modelNumber"] = attrmap["modelNumber"]
	UpAttributes["type"] = attrmap["type"]
	UpAttributes["ipAddress"] = attrmap["ipAddress"]
	UpAttributes["company"] = attrmap["company"]

	return UpAttributes
}

//LM2CW func
func LM2CW(conf *Cfg) error {

	var page int64
	page = 1
	var err error
	var errCount int
	startTime := time.Now()

	for {
		items, err := FetchDevices(conf, page)
		if err != nil {
			log.Error(err)
			errCount++
			break
		}

		if len(items) == 0 {
			break
		}

		err = CWAddUpdate(conf, items)
		if err != nil {
			errCount++
		}

		page = page + 1
	}
	status := "success"
	if errCount > 0 {
		status = "error"
	}
	setTimeMetrics(status, startTime, conf)
	log.Info("Sync run complete")
	return err
}
