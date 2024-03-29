package api

import (
	"bytes"
	b64 "encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/getsentry/sentry-go"

	log "github.com/ringsq/go-logger"
)

//CWType struct
type CWType []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Company struct
type Company []struct {
	ID         int    `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

//Type struct
type Type struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Comp struct
type Comp struct {
	ID         int    `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

//CwConfig struct
type CwConfig []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Comp Comp   `json:"company"`
}

func cwAuth(conf *Cfg) string {
	token := conf.CWCompany + "+" + conf.CWUser + ":" + conf.CWPass
	enc := b64.StdEncoding.EncodeToString([]byte(token))
	return enc
}

func getCwTypesByName(conf *Cfg, name string) ([]byte, error) {
	client := &http.Client{}
	encName := url.PathEscape(name)
	req, err := http.NewRequest(http.MethodGet, conf.CWURL+"/company/configurations/types?conditions=name="+"'"+encName+"'", nil)
	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Basic "+cwAuth(conf))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("clientID", conf.CWCompanyID)
	resp, err := client.Do(req)
	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return nil, err
	}
	return body, err
}

func getCwCompaniesByName(conf *Cfg, name string) ([]byte, error) {
	client := &http.Client{}

	var encName string
	if strings.Contains(name, "&") {
		compStr := strings.ReplaceAll(name, "&", "'&'")
		encName = compStr
	} else {
		encName = url.PathEscape(name)
	}

	req, err := http.NewRequest(http.MethodGet, conf.CWURL+"/company/companies?conditions=name="+"'"+encName+"'", nil)
	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Basic "+cwAuth(conf))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("clientID", conf.CWCompanyID)
	resp, err := client.Do(req)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return nil, err
	}
	return body, err
}

func getCwConfigurationsByName(conf *Cfg, name string) ([]byte, error) {
	client := &http.Client{}
	encName := url.PathEscape(name)
	req, err := http.NewRequest(http.MethodGet, conf.CWURL+"/company/configurations?conditions=name="+"'"+encName+"'", nil)
	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Basic "+cwAuth(conf))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("clientID", conf.CWCompanyID)
	resp, err := client.Do(req)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return nil, err
	}
	return body, err
}

func addDeviceToCw(conf *Cfg, data []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, conf.CWURL+"/company/configurations", bytes.NewBuffer(data))

	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Basic "+cwAuth(conf))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("clientID", conf.CWCompanyID)
	resp, err := client.Do(req)

	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return nil, err
	}

	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return nil, err
	}
	if resp.StatusCode != 201 {
		log.Error(resp.StatusCode, resp.Status)
		err = errors.New(string(rbody))
	}
	return rbody, err
}

func updateDeviceInCw(conf *Cfg, id string, body []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, conf.CWURL+"/company/configurations/"+id, bytes.NewBuffer(body))

	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Basic "+cwAuth(conf))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("clientID", conf.CWCompanyID)
	resp, err := client.Do(req)

	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return nil, err
	}

	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Error(resp.StatusCode, resp.Status)
		err = errors.New(string(rbody))
	}
	return rbody, err
}
