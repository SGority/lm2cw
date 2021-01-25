package api

import (
	log "github.com/magna5/go-logger"
	"github.com/magna5/go-logger/shims/zerolog"

	sentry "github.com/getsentry/sentry-go"
	flags "github.com/jessevdk/go-flags"
	"github.com/joeshaw/envdecode"
	"github.com/magna5/godotenv"
)

// Cfg configuration structure
type Cfg struct {
	Port                 int    `env:"PORT,default=8080" short:"p" long:"port" description:"HTTP Port"`
	JwksCertRenewMinutes int    `env:"JWKS_RENEW_MINUTES,default=60" description:"Number of minutes to wait before renewing JWKS certificates"`
	JWTIssuer            string `env:"JWT_ISSUER" description:"The URL to the JWT issuing server"`
	JobInterval          uint64 `env:"JOB_INTERVAL,default=180" description:"The interval at which the scheduler is invoked"`
	CWUser               string `env:"CW_USER,default=" description:"Connectwise username"`
	CWPass               string `env:"CW_PASS,default=" description:"Connectwise password"`
	CWURL                string `env:"CW_URL,default=" description:"The base url to connectwise."`
	CWCompany            string `env:"CW_COMPANY,default=" description:"Connectwise company name"`
	CWCompanyID          string `env:"CW_COMPANY_ID,default=" description:"Connectwise company ID"`

	SMTPHost         string   `env:"SMTP_HOST,default=" description:"SMTP host"`
	SMTPPort         string   `env:"SMTP_PORT,default=587" description:"SMTP port"`
	MailTo           []string `env:"MAIL_TO,default=" description:"Mail address of the recipient to whom the mail is being sent."`
	MailFrom         string   `env:"MAIL_FROM,default=" description:"Mail address of the sender."`
	MailPass         string   `env:"MAIL_PASS,default=" description:"Password of the sender."`
	BaseURL          string   `env:"BASE_URL,default=https://magna5global.logicmonitor.com/santaba/rest" description:"The URL to the JWT issuing server."`
	LmAccessID       string   `env:"LM_ACCESS_ID,default=fWTk7rvkN8dqqaT3stPB" description:"company ID"`
	LmAccessKey      string   `env:"LM_ACCESS_KEY,default=85pc+h2K9547}]cY8hNsR)^4%)x(9sc~4qdI(M+{" description:"company key"`
	DeviceSourcePath string   `env:"DEVICE_SOURCE_PATH,default=/device/devices" description:"device Source Path"`
	DeviceOffsetSize int      `env:"DEVICE_OFFSET_SIZE,default=100" description:"Device offsetSize"`
}

// Config is the current application configuration
var Config = &Cfg{}

// Configure func
func Configure(args []string) *Cfg {
	log.RootLogger = zerolog.New(nil)

	err := sentry.Init(sentry.ClientOptions{AttachStacktrace: true})
	if err != nil {
		log.Fatal("Error Initializing sentry: ", "error", err.Error())
	}

	err = godotenv.LoadOpt()
	if err != nil {
		panic(err)
	}
	err = envdecode.Decode(Config)
	if err != nil {
		panic(err)
	}
	_, err = flags.Parse(Config)
	if err != nil {
		panic(err)
	}
	return Config
}
