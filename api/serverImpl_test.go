package api

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func CreateTestServer(body []byte) *httptest.Server {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write(body)
	}))
	return testServer
}

func getTestResponseInByte() []byte {
	resp := `{
		"status" : 200,
		"errmsg" : "OK",
		"data" : {
			"total" : 798,
			"items" : [ {
				"id" : 2,
				"name" : "172.16.150.61",
				"displayName" : "NOCESX02",
				"deviceType" : 0,
				"relatedDeviceId" : -1,
				"currentCollectorId" : 33,
				"preferredCollectorId" : 33,
				"preferredCollectorGroupId" : 2,
				"preferredCollectorGroupName" : "Magna5",
				"description" : "",
				"createdOn" : 1539714647,
				"updatedOn" : 1599650239,
				"disableAlerting" : false,
				"autoPropsAssignedOn" : 1599621663,
				"autoPropsUpdatedOn" : 1599621664,
				"scanConfigId" : 0,
				"link" : "",
				"enableNetflow" : false,
				"netflowCollectorId" : 0,
				"netflowCollectorGroupId" : 0,
				"netflowCollectorGroupName" : null,
				"lastDataTime" : 0,
				"lastRawdataTime" : 0,
				"hostGroupIds" : "20,6",
				"sdtStatus" : "none-none-none",
				"userPermission" : "write",
				"hostStatus" : "normal",
				"alertStatus" : "unconfirmed-error-none",
				"alertStatusPriority" : 10,
				"awsState" : 1,
				"azureState" : 1,
				"gcpState" : 1,
				"alertDisableStatus" : "none-none-none",
				"alertingDisabledOn" : null,
				"collectorDescription" : "WPAPROBE2",
				"netflowCollectorDescription" : null,
				"customProperties" : [ {
					"name" : "snmp.version",
					"value" : "v2c"
				}, {
					"name" : "esx.user",
					"value" : "nocadmin"
				}, {
					"name" : "esx.pass",
					"value" : "********"
				}, {
					"name" : "system.categories",
					"value" : "snmpTCPUDP,VMWareHost"
				} ],
				"upTimeInSeconds" : 0,
				"deletedTimeInMs" : 0,
				"toDeleteTimeInMs" : 0,
				"hasDisabledSubResource" : false,
				"ancestorHasDisabledLogicModule" : false,
				"systemProperties" : [ {
					"name" : "system.sysinfo",
					"value" : "VMware ESXi 6.0.0 build-5050593 VMware, Inc. x86_64"
				}, {
					"name" : "system.enablenetflow",
					"value" : "false"
				}, {
					"name" : "system.collectorplatform",
					"value" : "windows"
				}, {
					"name" : "system.sysoid",
					"value" : "1.3.6.1.4.1.6876.4.1"
				}, {
					"name" : "system.vendor",
					"value" : "Dell Inc."
				}, {
					"name" : "system.collectorid",
					"value" : "33"
				}, {
					"name" : "system.deviceId",
					"value" : "2"
				}, {
					"name" : "system.virtualization",
					"value" : "VMWare ESX host"
				}, {
					"name" : "system.prefcollectordesc",
					"value" : "WPAPROBE2"
				}, {
					"name" : "system.model",
					"value" : "PowerEdge R520"
				}, {
					"name" : "system.version",
					"value" : "VMware ESXi 6.0.0 build-5050593"
				}, {
					"name" : "system.collectordesc",
					"value" : "WPAPROBE2"
				}, {
					"name" : "system.groups",
					"value" : "Customers/Magna5/Servers/VMWare ESXi/ESXi NOC,Devices by Type/VMware Hosts"
				}, {
					"name" : "system.deviceGroupId",
					"value" : "20,6"
				}, {
					"name" : "system.sysname",
					"value" : "NOCESX02"
				}, {
					"name" : "system.collector",
					"value" : "false"
				}, {
					"name" : "system.ips",
					"value" : "172.16.150.61"
				}, {
					"name" : "system.resourceCreatedOn",
					"value" : "1539714647"
				}, {
					"name" : "system.devicetype",
					"value" : "0"
				}, {
					"name" : "system.CPU",
					"value" : "Intel(R) Xeon(R) CPU E5-2407 0 @ 2.20GHz 2 CPU X 8 core"
				}, {
					"name" : "system.memory",
					"value" : "65490MB"
				}, {
					"name" : "system.collectorversion",
					"value" : "28005"
				}, {
					"name" : "system.staticgroups",
					"value" : "Customers/Magna5/Servers/VMWare ESXi/ESXi NOC"
				}, {
					"name" : "system.prefcollectorid",
					"value" : "33"
				}, {
					"name" : "system.displayname",
					"value" : "NOCESX02"
				}, {
					"name" : "system.categories",
					"value" : "snmpTCPUDP,VMWareHost"
				}, {
					"name" : "system.hostname",
					"value" : "172.16.150.61"
				} ],
				"autoProperties" : [ {
					"name" : "auto.lmstatus",
					"value" : "normal"
				} ],
				"inheritedProperties" : [ {
					"name" : "api.pass",
					"value" : "********"
				}, {
					"name" : "customer.name",
					"value" : "Magna5"
				}, {
					"name" : "snmp.community",
					"value" : "********"
				}, {
					"name" : "cw_type",
					"value" : "ESXi"
				}, {
					"name" : "netapp.ssl",
					"value" : "true"
				}, {
					"name" : "api.user",
					"value" : "3n44M8XBBX8qgQX83qe2"
				}, {
					"name" : "api.account",
					"value" : "magna5global"
				} ]
			} ],
			"searchId" : null,
			"isMin" : false
		}
	}`
	return []byte(resp)
}

func TestDeviceAuth(t *testing.T) {
	t.Run("Test deviceAuth function", func(t *testing.T) {
		//cfg := Configure(os.Args)
		cfg := &Cfg{BaseURL: "https://magna5global.logicmonitor.com/santaba/rest", LmAccessID: "fWTk7rvkN8dqqaT3stPB",
			LmAccessKey: "85pc+h2K9547}]cY8hNsR)^4%)x(9sc~4qdI(M+{", DeviceSourcePath: "/device/devices", DeviceOffsetSize: 100}
		deviceAuthString := deviceAuth(cfg, 1599664441385)
		if deviceAuthString != "LMv1 fWTk7rvkN8dqqaT3stPB:NmI1NmVlZTc0MzBmMmQ0NzIyY2NiYzY0NGU1YjgyNDk3MGJiMTAwYjVlMzI1ZWIyN2MxYTk4NmQ2YzBmNDU5NQ==:1599664441385" {
			t.Errorf("deviceAuthString does not match")
		}
	})
}

func TestGenerateHmacToken(t *testing.T) {
	t.Run("Test generateHmacToken function", func(t *testing.T) {
		deviceAuthString := generateHmacToken("1234567890", "Test")
		if deviceAuthString != "03a81c66ab956495cfb0f4d4a34218a981ff467db8a45e0b8344e726c30e47ff" {
			t.Errorf("HmacToken does not match")
		}
	})
}

func TestFetchDevices(t *testing.T) {
	t.Run("Test  FetchDevices function", func(t *testing.T) {
		responseBody := getTestResponseInByte()
		testServer := CreateTestServer(responseBody)
		req, _ := http.NewRequest("GET", testServer.URL, nil)
		req.Header.Add("Authorization", "LMv1 fWTk7rvkN8dqqaT3stPB:NmI1NmVlZTc0MzBmMmQ0NzIyY2NiYzY0NGU1YjgyNDk3MGJiMTAwYjVlMzI1ZWIyN2MxYTk4NmQ2YzBmNDU5NQ==:1599664441385")
		res := httptest.NewRecorder()

		if res.Code != http.StatusOK {
			t.Errorf("Response code was %v; want 200", res.Code)
		}
		testServer.Close()
	})
}

func TestGetDeviceItems(t *testing.T) {
	t.Run("Test getDevicesItems  function", func(t *testing.T) {
		responseBody := getTestResponseInByte()
		items, err := getDeviceItems(responseBody)
		if err != nil {
			t.Errorf("Error occured %v", err)
		}
		if items[0]["id"] != float64(2) {
			t.Errorf("Wrong Item selected %s", items[0]["id"])
			t.Errorf("Wrong Item selected")
		}
	})
}

func TestFlattenDeviceItems(t *testing.T) {
	t.Run("Test FlattenDeviceItems  function", func(t *testing.T) {
		responseBody := getTestResponseInByte()
		items, err := getDeviceItems(responseBody)
		if err != nil {
			t.Errorf("Error occured %v", err)
		}
		response := flattenDeviceItems(items)
		if response[0]["system.enablenetflow"] != "false" {
			t.Errorf("system.enablenetflow value should be false")
		}
	})
}

func TestCWAddUpdate(t *testing.T) {

	tableTest := []struct {
		name    string
		lmresp  []map[string]interface{}
		expResp string
	}{
		{
			name:    "when deviceType is equal to zero",
			lmresp:  []map[string]interface{}{{"displayName": "Test", "osType": "Windows", "company_name": "Ameri100", "cw_type": "Firewall", "deviceType": 2}, {"displayName": "Server123", "osType": "Windows", "company_name": "XYZTestCompany", "cw_type": "Managed Server", "deviceType": 2}},
			expResp: ``,
		},
		{
			name:    "when cw_type is equal to zero",
			lmresp:  []map[string]interface{}{{"displayName": "Server123", "osType": "Windows", "company_name": "Ameri100", "cw_type": "", "deviceType": 0}, {"displayName": "Server123", "osType": "Windows", "company_name": "XYZTestCompany", "cw_type": "", "deviceType": 0}},
			expResp: ``,
		},
		{
			name:    "when deviceType and cw_type is valid",
			lmresp:  []map[string]interface{}{{"displayName": "Server123", "osType": "Windows", "company_name": "XYZTestCompany", "cw_type": "Managed Server", "deviceType": 0}, {"displayName": "Server123", "osType": "Windows", "company_name": "XYZTestCompany", "cw_type": "ESXi", "deviceType": 0}},
			expResp: `[{"id":1,"name":"Server123","type":{"id":25,"name":"Managed Server","_info":{"type_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/types/25"}},"status":{"id":2,"name":"Active","_info":{"status_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/statuses/2"}},"company":{"id":19297,"identifier":"XYZTestCompany","name":"XYZ Test Company","_info":{"company_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19297"}},"contact":{"id":7,"name":"Fred Stone","_info":{"contact_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/contacts/7"}},"site":{"id":1002,"name":"Tampa Office","_info":{"site_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19297/sites/1002"}},"locationId":11,"businessUnitId":1,"deviceIdentifier":"","serialNumber":"","modelNumber":"","tagNumber":"","vendorNotes":"","notes":"","macAddress":"","lastLoginName":"","billFlag":true,"backupSuccesses":0,"backupIncomplete":0,"backupFailed":0,"backupRestores":0,"backupServerName":"","backupBillableSpaceGb":0.00,"backupProtectedDeviceList":"","backupYear":0,"backupMonth":0,"ipAddress":"","defaultGateway":"","osType":"","osInfo":"","cpuSpeed":"","ram":"","localHardDrives":"","activeFlag":true,"mobileGuid":"7546e39c-c000-4148-8e90-bd4b84eb7a9e","_info":{"lastUpdated":"2019-10-09T18:41:33Z","updatedBy":"zAdmin","dateEntered":"2019-10-09T18:36:41Z","enteredBy":"zAdmin"},"companyLocationId":2,"showRemoteFlag":false,"showAutomateFlag":false,"needsRenewalFlag":false}]`,
		},
	}

	for _, testCase := range tableTest {
		t.Run(testCase.name, func(t *testing.T) {

			resp, err := testCWAddUpdate(testCase.lmresp)
			if err != nil {
				t.Errorf("Error occured %v", err)
			}

			response := (string(resp))
			res := (string(testCase.expResp))

			if res != response {
				t.Errorf("Response body was '%v'; want '%v'", response, res)
			}
		})
	}
}

func Test_setCWAttributes(t *testing.T) {
	type args struct {
		lmap map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{name: "Test osType slicing < 250",
			args: args{lmap: map[string]interface{}{
				"system.sysinfo": "this is a short test",
			}},
			want: map[string]interface{}{
				"osInfo":      "this is a short test",
				"notes":       nil,
				"company":     nil,
				"name":        nil,
				"modelNumber": nil,
				"type":        nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setCWAttributes(tt.args.lmap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setCWAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}
