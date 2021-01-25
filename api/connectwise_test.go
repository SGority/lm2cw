package api

import (
	"testing"
)

func TestCwAuth(t *testing.T) {

	cfg := &Cfg{CWCompany: "Magna5", CWUser: "agddhhjnmfkf", CWPass: "nsnihiihuwuf"}
	token := cwAuth(cfg)
	tokenstring := "TWFnbmE1K2FnZGRoaGpubWZrZjpuc25paGlpaHV3dWY="
	if token != tokenstring {
		t.Errorf("Autherization string does not match got'%v'; want '%v'", token, tokenstring)
	}

}

func TestCwTypesByName(t *testing.T) {
	t.Run("Test cwAuth function", func(t *testing.T) {
		body := `[{"id":52,"name":"Firewall","inactiveFlag":false,"systemFlag":false,"_info":{"lastUpdated":"2019-10-09T18:54:36Z","updatedBy":"zAdmin","questions_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/types/52/questions"}}]`
		name := "Firewall"
		resp, err := testCwTypesByName(name)
		if err != nil {
			t.Errorf("Error occured")

		}

		if string(resp) != body {
			t.Errorf("Response body was '%v'; want '%v'", body, resp)
		}

	})
}

func TestCwCompanyByName(t *testing.T) {
	t.Run("Test get company by name", func(t *testing.T) {
		body := `[{"id":19310,"identifier":"Ameri100","name":"Ameri100","status":{"id":1,"name":"Active","_info":{"status_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/statuses/1"}},"addressLine1":"200 Spectrum Center Drive","city":"Irving","state":"CA","zip":"92618","phoneNumber":"9497202550","faxNumber":"","website":"","territory":{"id":11,"name":"Magna5 MS","_info":{"location_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/system/locations/11"}},"accountNumber":"19310","dateAcquired":"2019-11-04T16:41:20Z","annualRevenue":0.00,"numberOfEmployees":0,"leadFlag":false,"unsubscribeFlag":false,"vendorIdentifier":"","taxIdentifier":"","taxCode":{"id":8,"name":"AVATAX","_info":{"taxCode_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/finance/taxCodes/8"}},"billingTerms":{"id":1,"name":"Net 30 days"},"invoiceTemplate":{"id":12,"name":"M5 Custom Invoice - Standard - Time - No Detail","_info":{"billingTerms_Href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/finance/invoiceTemplateSetups/12"}},"billToCompany":{"id":19310,"identifier":"Ameri100","name":"Ameri100","_info":{"company_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19310"}},"billingSite":{"id":1035,"name":"Irving","_info":{"site_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19310/sites/1035"}},"billingContact":{"id":544,"name":"Accounts Payable","_info":{"contact_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/contacts/544"}},"invoiceDeliveryMethod":{"id":2,"name":"E-Mail"},"invoiceToEmailAddress":"AP@ameri100.com","deletedFlag":false,"dateDeleted":"2019-11-04T16:41:20Z","mobileGuid":"1d4a8841-c5e1-4da6-a59c-9f5bc6ef3d83","types":[{"id":1,"name":"Client","_info":{"type_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/types/1"}}],"site":{"id":1035,"name":"Irving","_info":{"site_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19310/sites/1035"}},"_info":{"lastUpdated":"2020-04-30T01:46:24Z","updatedBy":"RWeber","dateEntered":"2019-11-04T16:41:20Z","enteredBy":"Import1","contacts_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/contacts?conditions=company/id=19310","agreements_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/finance/agreements?conditions=company/id=19310","tickets_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/service/tickets?conditions=company/id=19310","opportunities_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/sales/opportunities?conditions=company/id=19310","activities_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/sales/activities?conditions=company/id=19310","projects_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/project/projects?conditions=company/id=19310","configurations_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations?conditions=company/id=19310","orders_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/sales/orders?conditions=company/id=19310","documents_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/system/documents?recordType=Company&recordId=19310","sites_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19310/sites","teams_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19310/teams","reports_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19310/managementSummaryReports","notes_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19310/notes"}}]`
		name := "Ameri100"
		resp, err := testGetCwCompaniesByName(name)
		if err != nil {
			t.Errorf("Error occured")

		}

		if string(resp) != body {
			t.Errorf("Response body was '%v'; want '%v'", resp, body)
		}

	})
}

func TestGetCwConfigurationsByName(t *testing.T) {
	t.Run("Test get configurations by name", func(t *testing.T) {
		body := `[{"id":1,"name":"Server123","type":{"id":25,"name":"Managed Server","_info":{"type_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/types/25"}},"status":{"id":2,"name":"Active","_info":{"status_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/statuses/2"}},"company":{"id":19297,"identifier":"XYZTestCompany","name":"XYZ Test Company","_info":{"company_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19297"}},"contact":{"id":7,"name":"Fred Stone","_info":{"contact_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/contacts/7"}},"site":{"id":1002,"name":"Tampa Office","_info":{"site_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19297/sites/1002"}},"locationId":11,"businessUnitId":1,"deviceIdentifier":"","serialNumber":"","modelNumber":"","tagNumber":"","vendorNotes":"","notes":"","macAddress":"","lastLoginName":"","billFlag":true,"backupSuccesses":0,"backupIncomplete":0,"backupFailed":0,"backupRestores":0,"backupServerName":"","backupBillableSpaceGb":0.00,"backupProtectedDeviceList":"","backupYear":0,"backupMonth":0,"ipAddress":"","defaultGateway":"","osType":"","osInfo":"","cpuSpeed":"","ram":"","localHardDrives":"","activeFlag":true,"mobileGuid":"7546e39c-c000-4148-8e90-bd4b84eb7a9e","_info":{"lastUpdated":"2019-10-09T18:41:33Z","updatedBy":"zAdmin","dateEntered":"2019-10-09T18:36:41Z","enteredBy":"zAdmin"},"companyLocationId":2,"showRemoteFlag":false,"showAutomateFlag":false,"needsRenewalFlag":false}]`
		name := "Server123"
		resp, err := testGetCwConfigurationsByName(name)
		if err != nil {
			t.Errorf("Error occured")

		}

		if string(resp) != body {
			t.Errorf("Response body was '%v'; want '%v'", resp, body)
		}

	})

}

func TestAddDeviceToCw(t *testing.T) {
	t.Run("Test add configurations by name", func(t *testing.T) {
		body := []byte(`[{"id":1,"name":"Server123","type":{"id":25,"name":"Firewall","_info":{"type_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/types/25"}},"status":{"id":2,"name":"Active","_info":{"status_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/statuses/2"}},"company":{"id":19297,"identifier":"TestCompany","name":"Test Company","_info":{"company_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19297"}},"contact":{"id":7,"name":"Fred Stone","_info":{"contact_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/contacts/7"}},"site":{"id":1002,"name":"Tampa Office","_info":{"site_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19297/sites/1002"}},"locationId":11,"businessUnitId":1,"deviceIdentifier":"","serialNumber":"","modelNumber":"","tagNumber":"","vendorNotes":"","notes":"","macAddress":"","lastLoginName":"","billFlag":true,"backupSuccesses":0,"backupIncomplete":0,"backupFailed":0,"backupRestores":0,"backupServerName":"","backupBillableSpaceGb":0.00,"backupProtectedDeviceList":"","backupYear":0,"backupMonth":0,"ipAddress":"","defaultGateway":"","osType":"","osInfo":"","cpuSpeed":"","ram":"","localHardDrives":"","activeFlag":true,"mobileGuid":"7546e39c-c000-4148-8e90-bd4b84eb7a9e","_info":{"lastUpdated":"2019-10-09T18:41:33Z","updatedBy":"zAdmin","dateEntered":"2019-10-09T18:36:41Z","enteredBy":"zAdmin"},"companyLocationId":2,"showRemoteFlag":false,"showAutomateFlag":false,"needsRenewalFlag":false}]`)
		resp, err := testAddDeviceToCw(body)
		if err != nil {
			t.Errorf("Error occured")

		}

		response := string(resp)
		data := string(body)

		if response != data {
			t.Errorf("Response body was '%v'; want '%v'", response, data)
		}

	})

}

func TestUpdateDeviceInCw(t *testing.T) {
	t.Run("Test update configurations", func(t *testing.T) {
		body := []byte(`[{"id":1,"name":"Server123","type":{"id":25,"name":"Managed Server","_info":{"type_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/types/25"}},"status":{"id":2,"name":"Active","_info":{"status_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/configurations/statuses/2"}},"company":{"id":19297,"identifier":"XYZTestCompany","name":"XYZ Test Company","_info":{"company_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19297"}},"contact":{"id":7,"name":"Fred Stone","_info":{"contact_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/contacts/7"}},"site":{"id":1002,"name":"Tampa Office","_info":{"site_href":"https://api-na.myconnectwise.net/v4_6_release/apis/3.0/company/companies/19297/sites/1002"}},"locationId":11,"businessUnitId":1,"deviceIdentifier":"","serialNumber":"","modelNumber":"","tagNumber":"","vendorNotes":"","notes":"","macAddress":"","lastLoginName":"","billFlag":true,"backupSuccesses":0,"backupIncomplete":0,"backupFailed":0,"backupRestores":0,"backupServerName":"","backupBillableSpaceGb":0.00,"backupProtectedDeviceList":"","backupYear":0,"backupMonth":0,"ipAddress":"","defaultGateway":"","osType":"","osInfo":"","cpuSpeed":"","ram":"","localHardDrives":"","activeFlag":true,"mobileGuid":"7546e39c-c000-4148-8e90-bd4b84eb7a9e","_info":{"lastUpdated":"2019-10-09T18:41:33Z","updatedBy":"zAdmin","dateEntered":"2019-10-09T18:36:41Z","enteredBy":"zAdmin"},"companyLocationId":2,"showRemoteFlag":false,"showAutomateFlag":false,"needsRenewalFlag":false}]`)
		resp, err := testUpdateDeviceInCw("1", body)
		if err != nil {
			t.Errorf("Error occured")

		}

		response := string(resp)
		data := string(body)

		if response != data {
			t.Errorf("Response body was '%v'; want '%v'", response, data)
		}

	})

}
