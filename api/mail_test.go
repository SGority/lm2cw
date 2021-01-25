package api

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"reflect"
	"testing"

	"github.com/gomicro/postal"
	log "github.com/magna5/go-logger"
)

func mockSendMailBody(data Device) (bytes.Buffer, error) {
	var body bytes.Buffer
	t, _ := template.ParseFiles("template.html")

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: LM to connectwise sync details \n%s\n\n", mimeHeaders)))

	err := t.Execute(&body, data)
	if err != nil {
		log.Error(err)
	}

	return body, err

}

func TestSendMail(t *testing.T) {

	cfg := &Cfg{MailFrom: "968d451360a34c",
		MailPass: "ea1b31c6362fef",
		MailTo:   []string{"968d451360a34c"},
		SMTPHost: "mailhost.com",
		SMTPPort: "587"}

	tableTest := []struct {
		name    string
		device  Device
		expResp string
	}{
		{
			name: "when devices do not have comapany name set and companies are not found in connectwise",
			device: Device{
				Devices: []DeviceDetails{{ID: 1, Name: "Device Name"}}, CompanyNames: []string{"Test Company"}},
			expResp: `Subject: LM to connectwise sync details 
            MIME-version: 1.0;
            Content-Type: text/html; charset="UTF-8";
            
            
            
            <!DOCTYPE html>
            <html>
            <body>
            <h3>The following devices did not have the company_name set:</h3>
            <ul>
                  
                    
                        </li>ID: 1, Name: Device Name</li></br>
                    
                
            </ul>
               
            <h3>The following companies were not found in connectwise :</h3>   
            <ul>
                  
                    
                        </li>Test Company</li></br>
                    
                
            </ul>
            </body>
            </html>
            
            `,
		},
		{
			name: "when company name is not set for a device",
			device: Device{
				Devices: []DeviceDetails{{ID: 1, Name: "Device Name"}}, CompanyNames: []string{}},
			expResp: `Subject: LM to connectwise sync details
		    MIME-version: 1.0;
		    Content-Type: text/html; charset="UTF-8";

		    <!DOCTYPE html>
		    <html>
		    <body>
		    <h3>The following devices did not have the company_name set:</h3>
		    <ul>

		                </li>ID: 1, Name: Device Name</li></br>

		    </ul>

		    <h3>The following companies were not found in connectwise :</h3>
		    <ul>
		         Not Found

		    </ul>
		    </body>
			</html>
			
			`,
		},
		{
			name: "when company is not found in connectwise",
			device: Device{
				Devices: []DeviceDetails{}, CompanyNames: []string{"Test Company"}},
			expResp: `Subject: LM to connectwise sync details
		    MIME-version: 1.0;
		    Content-Type: text/html; charset="UTF-8";

		    <!DOCTYPE html>
		    <html>
		    <body>
		    <h3>The following devices did not have the company_name set:</h3>
		    <ul>
		         No Devices Found

		    </ul>

		    <h3>The following companies were not found in connectwise :</h3>
		    <ul>

		                </li>Test Company</li></br>

		    </ul>
		    </body>
			</html>
			
			`,
		},
	}

	for _, testCase := range tableTest {
		t.Run(testCase.name, func(t *testing.T) {
			body, err := mockSendMailBody(testCase.device)
			if err != nil {
				t.Error("Error occured")
			}
			p := postal.New()
			auth := smtp.PlainAuth("", cfg.MailFrom, cfg.MailPass, cfg.SMTPHost)
			SendMail := p.Mailer()

			err = SendMail(cfg.SMTPHost, auth, cfg.MailFrom, cfg.MailTo, body.Bytes())
			if err != nil {
				t.Error("Error occured")
			}

			records := p.MailRecords()

			r := records[0]

			mailBody := string(r.Body)

			if p.Mailed() != 1 {
				t.Errorf("expected 1 mailed, got %v mailed", p.Mailed())
			}

			if reflect.DeepEqual(mailBody, testCase.expResp) {
				t.Errorf("Response body was '%v'; want '%v'", mailBody, testCase.expResp)
			}

		})
	}

}
