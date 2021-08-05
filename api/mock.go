package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"

	log "github.com/ringsq/go-logger"
)

func testCwTypesByName(name string) ([]byte, error) {
	respo := `[{"id":52,"name":"Firewall","inactiveFlag":false,"systemFlag":false,"_info":{"lastUpdated":"2019-10-09T18:54:36Z","updatedBy":"zAdmin","questions_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/types/52/questions"}}]`

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(respo))
	}

	req := httptest.NewRequest("GET", "/company/configurations/types?conditions=name="+"'"+"name"+"'", nil)
	req.Header.Add("Authorization", "TWFnbmE1K21iczlIUmFxTjRreXpGRTU6Wlo5NFZKYU5XVEJwTjdySQ==")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("clientID", "Magna5")

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Unable to connect to server")
	}
	return body, err
}

func testGetCwCompaniesByName(name string) ([]byte, error) {
	respo := `[{"id":19310,"identifier":"Ameri100","name":"Ameri100","status":{"id":1,"name":"Active","_info":{"status_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/statuses/1"}},"addressLine1":"200 Spectrum Center Drive","city":"Irving","state":"CA","zip":"92618","phoneNumber":"9497202550","faxNumber":"","website":"","territory":{"id":11,"name":"Magna5 MS","_info":{"location_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/system/locations/11"}},"accountNumber":"19310","dateAcquired":"2019-11-04T16:41:20Z","annualRevenue":0.00,"numberOfEmployees":0,"leadFlag":false,"unsubscribeFlag":false,"vendorIdentifier":"","taxIdentifier":"","taxCode":{"id":8,"name":"AVATAX","_info":{"taxCode_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/finance/taxCodes/8"}},"billingTerms":{"id":1,"name":"Net 30 days"},"invoiceTemplate":{"id":12,"name":"M5 Custom Invoice - Standard - Time - No Detail","_info":{"billingTerms_Href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/finance/invoiceTemplateSetups/12"}},"billToCompany":{"id":19310,"identifier":"Ameri100","name":"Ameri100","_info":{"company_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19310"}},"billingSite":{"id":1035,"name":"Irving","_info":{"site_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19310/sites/1035"}},"billingContact":{"id":544,"name":"Accounts Payable","_info":{"contact_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/contacts/544"}},"invoiceDeliveryMethod":{"id":2,"name":"E-Mail"},"invoiceToEmailAddress":"AP@ameri100.com","deletedFlag":false,"dateDeleted":"2019-11-04T16:41:20Z","mobileGuid":"1d4a8841-c5e1-4da6-a59c-9f5bc6ef3d83","types":[{"id":1,"name":"Client","_info":{"type_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/types/1"}}],"site":{"id":1035,"name":"Irving","_info":{"site_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19310/sites/1035"}},"_info":{"lastUpdated":"2020-04-30T01:46:24Z","updatedBy":"RWeber","dateEntered":"2019-11-04T16:41:20Z","enteredBy":"Import1","contacts_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/contacts?conditions=company/id=19310","agreements_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/finance/agreements?conditions=company/id=19310","tickets_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/service/tickets?conditions=company/id=19310","opportunities_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/sales/opportunities?conditions=company/id=19310","activities_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/sales/activities?conditions=company/id=19310","projects_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/project/projects?conditions=company/id=19310","configurations_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations?conditions=company/id=19310","orders_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/sales/orders?conditions=company/id=19310","documents_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/system/documents?recordType=Company&recordId=19310","sites_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19310/sites","teams_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19310/teams","reports_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19310/managementSummaryReports","notes_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19310/notes"}}]`

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(respo))
	}

	req := httptest.NewRequest("GET", "/company/companies?conditions=name="+"'"+name+"'", nil)
	req.Header.Add("Authorization", "TWFnbmE1K21iczlIUmFxTjRreXpGRTU6Wlo5NFZKYU5XVEJwTjdySQ==")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("clientID", "Magna5")

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Unable to connect to server")
	}

	return body, err
}

func testGetCwConfigurationsByName(name string) ([]byte, error) {

	response := `[{"id":1,"name":"Server123","type":{"id":25,"name":"Managed Server","_info":{"type_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/types/25"}},"status":{"id":2,"name":"Active","_info":{"status_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/statuses/2"}},"company":{"id":19297,"identifier":"XYZTestCompany","name":"XYZ Test Company","_info":{"company_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19297"}},"contact":{"id":7,"name":"Fred Stone","_info":{"contact_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/contacts/7"}},"site":{"id":1002,"name":"Tampa Office","_info":{"site_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19297/sites/1002"}},"locationId":11,"businessUnitId":1,"deviceIdentifier":"","serialNumber":"","modelNumber":"","tagNumber":"","vendorNotes":"","notes":"","macAddress":"","lastLoginName":"","billFlag":true,"backupSuccesses":0,"backupIncomplete":0,"backupFailed":0,"backupRestores":0,"backupServerName":"","backupBillableSpaceGb":0.00,"backupProtectedDeviceList":"","backupYear":0,"backupMonth":0,"ipAddress":"","defaultGateway":"","osType":"","osInfo":"","cpuSpeed":"","ram":"","localHardDrives":"","activeFlag":true,"mobileGuid":"7546e39c-c000-4148-8e90-bd4b84eb7a9e","_info":{"lastUpdated":"2019-10-09T18:41:33Z","updatedBy":"zAdmin","dateEntered":"2019-10-09T18:36:41Z","enteredBy":"zAdmin"},"companyLocationId":2,"showRemoteFlag":false,"showAutomateFlag":false,"needsRenewalFlag":false}]`

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
	}

	req := httptest.NewRequest("GET", "/company/companies?conditions=name="+"'"+name+"'", nil)
	req.Header.Add("Authorization", "TWFnbmE1K21iczlIUmFxTjRreXpGRTU6Wlo5NFZKYU5XVEJwTjdySQ==")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("clientID", "Magna5")

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Unable to connect to server")
	}

	return body, err

}

func testAddDeviceToCw(data []byte) ([]byte, error) {
	response := `[{"id":1,"name":"Server123","type":{"id":25,"name":"Firewall","_info":{"type_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/types/25"}},"status":{"id":2,"name":"Active","_info":{"status_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/statuses/2"}},"company":{"id":19297,"identifier":"TestCompany","name":"Test Company","_info":{"company_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19297"}},"contact":{"id":7,"name":"Fred Stone","_info":{"contact_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/contacts/7"}},"site":{"id":1002,"name":"Tampa Office","_info":{"site_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19297/sites/1002"}},"locationId":11,"businessUnitId":1,"deviceIdentifier":"","serialNumber":"","modelNumber":"","tagNumber":"","vendorNotes":"","notes":"","macAddress":"","lastLoginName":"","billFlag":true,"backupSuccesses":0,"backupIncomplete":0,"backupFailed":0,"backupRestores":0,"backupServerName":"","backupBillableSpaceGb":0.00,"backupProtectedDeviceList":"","backupYear":0,"backupMonth":0,"ipAddress":"","defaultGateway":"","osType":"","osInfo":"","cpuSpeed":"","ram":"","localHardDrives":"","activeFlag":true,"mobileGuid":"7546e39c-c000-4148-8e90-bd4b84eb7a9e","_info":{"lastUpdated":"2019-10-09T18:41:33Z","updatedBy":"zAdmin","dateEntered":"2019-10-09T18:36:41Z","enteredBy":"zAdmin"},"companyLocationId":2,"showRemoteFlag":false,"showAutomateFlag":false,"needsRenewalFlag":false}]`

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
	}

	req := httptest.NewRequest("GET", "/company/configurations", bytes.NewBuffer(data))
	req.Header.Add("Authorization", "TWFnbmE1K21iczlIUmFxTjRreXpGRTU6Wlo5NFZKYU5XVEJwTjdySQ==")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("clientID", "Magna5")

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Unable to connect to server")
	}

	return body, err

}

func testUpdateDeviceInCw(id string, data []byte) ([]byte, error) {
	response := `[{"id":1,"name":"Server123","type":{"id":25,"name":"Managed Server","_info":{"type_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/types/25"}},"status":{"id":2,"name":"Active","_info":{"status_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/statuses/2"}},"company":{"id":19297,"identifier":"XYZTestCompany","name":"XYZ Test Company","_info":{"company_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19297"}},"contact":{"id":7,"name":"Fred Stone","_info":{"contact_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/contacts/7"}},"site":{"id":1002,"name":"Tampa Office","_info":{"site_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19297/sites/1002"}},"locationId":11,"businessUnitId":1,"deviceIdentifier":"","serialNumber":"","modelNumber":"","tagNumber":"","vendorNotes":"","notes":"","macAddress":"","lastLoginName":"","billFlag":true,"backupSuccesses":0,"backupIncomplete":0,"backupFailed":0,"backupRestores":0,"backupServerName":"","backupBillableSpaceGb":0.00,"backupProtectedDeviceList":"","backupYear":0,"backupMonth":0,"ipAddress":"","defaultGateway":"","osType":"","osInfo":"","cpuSpeed":"","ram":"","localHardDrives":"","activeFlag":true,"mobileGuid":"7546e39c-c000-4148-8e90-bd4b84eb7a9e","_info":{"lastUpdated":"2019-10-09T18:41:33Z","updatedBy":"zAdmin","dateEntered":"2019-10-09T18:36:41Z","enteredBy":"zAdmin"},"companyLocationId":2,"showRemoteFlag":false,"showAutomateFlag":false,"needsRenewalFlag":false}]`

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
	}

	req := httptest.NewRequest(http.MethodPatch, "/company/configurations"+id, bytes.NewBuffer(data))
	req.Header.Add("Authorization", "TWFnbmE1K21iczlIUmFxTjRreXpGRTU6Wlo5NFZKYU5XVEJwTjdySQ==")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("clientID", "Magna5")

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Unable to connect to server")
	}

	return body, err

}

func testaddOrUpdate(devicename, compname string, data, updata map[string]interface{}) ([]byte, error) {
	var config CwConfig
	//var cfg *Cfg
	res, err := testGetCwConfigurationsByName(devicename)

	if err != nil {
		log.Error("Unable to get the CW configuration response", err)
		return nil, err
	}

	err = json.Unmarshal(res, &config)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if len(config) != 0 && config[0].Comp.Name == compname {
		id := strconv.Itoa(config[0].ID)
		jsonData, err := json.Marshal(updata)
		if err != nil {
			log.Error(err)
		}

		res, err := testUpdateDeviceInCw(id, jsonData)
		if err != nil {
			log.Error("Unable to update device", err)
			return res, err
		}
		log.Info("Successfully updated the device")
	} else {
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Error(err)
		}

		res, err := testAddDeviceToCw(jsonData)
		if err != nil {
			log.Error("Unable to add the device", err)
			return res, err
		}
		log.Info("Successfully added the device")
	}
	return res, err
}

func testCWAddUpdate(lmresp []map[string]interface{}) ([]byte, error) {
	var deviceData []byte
	var cwerr error
	for i := 0; i < len(lmresp); i++ {
		deviceType := lmresp[i]["deviceType"]

		//Set CW attributes
		cwAttributes := setCWAttributes(lmresp[i])

		if deviceType.(int) == 0 {

			if (lmresp[i]["cw_type"]).(string) != "" {
				var cwtype CWType
				tagName := lmresp[i]["cw_type"].(string)

				//check CW type
				getType, err := testCwTypesByName(tagName)
				if err != nil {
					log.Error("Unable to get CW type response", err)
					return nil, err
				}

				err = json.Unmarshal(getType, &cwtype)
				if err != nil {
					log.Error(err)
					return nil, err
				}

				if len(cwtype) != 0 {
					//set cw type field
					cwAttributes["type"] = make(map[string]interface{})
					cwAttributes["type"].(map[string]interface{})["id"] = cwtype[0].ID
					cwAttributes["type"].(map[string]interface{})["name"] = cwtype[0].Name

					//check CW company
					var comp Company
					if lmresp[i]["company_name"] != nil {
						com := lmresp[i]["company_name"].(string)
						res, err := testGetCwCompaniesByName(com)
						if err != nil {
							log.Error("Unable to get CW company response", err)
							return nil, err
						}

						err = json.Unmarshal(res, &comp)
						if err != nil {
							log.Error(err)
							return nil, err
						}

						if len(comp) != 0 {

							//set CW company field
							cwAttributes["company"] = make(map[string]interface{})
							cwAttributes["company"].(map[string]interface{})["id"] = comp[0].ID
							cwAttributes["company"].(map[string]interface{})["identifier"] = comp[0].Identifier
							cwAttributes["company"].(map[string]interface{})["name"] = comp[0].Name

							upattr := updateAttributes(cwAttributes)
							deviceName := (cwAttributes["name"]).(string)
							compname := comp[0].Name
							deviceData, err = testaddOrUpdate(deviceName, compname, cwAttributes, upattr)
							if err != nil {
								log.Error(err)
								return deviceData, err
							}

						}

					}

				}

			}

		}

	}
	return deviceData, cwerr

}
