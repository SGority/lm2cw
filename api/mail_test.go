package api

import (
	"io"
	"net/smtp"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/gomicro/postal"
)

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

			copyTempFile()
			defer os.RemoveAll("templates")

			body, err := loadTemplate(testCase.device)
			if err != nil {
				t.Error("Error occured", err)
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

func copyTempFile() error {
	src, err := os.Open("../templates/mail.tmpl")
	if err != nil {
		return err
	}
	defer src.Close()

	err = os.MkdirAll("templates", 0777)
	if err != nil {
		return err
	}

	destination := filepath.Join("templates", "mail.tmpl")

	dest, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}
	return err
}
