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
	Port        int    `env:"PORT,default=8080" short:"p" long:"port" description:"The port on which the service listens"`
	JobInterval uint64 `env:"JOB_INTERVAL,default=180" description:"The interval at which the scheduler is invoked specified in minutes"`
	CWUser      string `env:"CW_USER" description:"The username required for connectwise authentication"`
	CWPass      string `env:"CW_PASS" description:"The password required for connectwise authentication"`
	CWURL       string `env:"CW_URL,default=https://na.myconnectwise.net/v4_6_release/apis/3.0" description:"The base url to connect to connectwise."`
	CWCompany   string `env:"CW_COMPANY" description:"The company name in connectwise"`
	CWCompanyID string `env:"CW_COMPANY_ID" description:"The company ID for connectwise authentication"`

	SMTPHost         string   `env:"SMTP_HOST" description:"SMTP host"`
	SMTPPort         string   `env:"SMTP_PORT,default=587" description:"SMTP port"`
	MailTo           []string `env:"MAIL_TO" description:"Mail address of the recipient to whom the mail is being sent."`
	MailFrom         string   `env:"MAIL_FROM" description:"Mail address of the sender."`
	MailPass         string   `env:"MAIL_PASS" description:"Password of the sender."`
	BaseURL          string   `env:"BASE_URL,default=https://magna5global.logicmonitor.com/santaba/rest" description:"The base url to connect to logicmonitor."`
	LmAccessID       string   `env:"LM_ACCESS_ID,default=fWTk7rvkN8dqqaT3stPB" description:"The access id required for logicmonitor authentication"`
	LmAccessKey      string   `env:"LM_ACCESS_KEY,default=85pc+h2K9547}]cY8hNsR)^4%)x(9sc~4qdI(M+{" description:"The access key required for logicmonitor authentication"`
	DeviceSourcePath string   `env:"DEVICE_SOURCE_PATH,default=/device/devices" description:"The path to fetch the device resource from Logicmonitor"`
	DeviceOffsetSize int      `env:"DEVICE_OFFSET_SIZE,default=100" description:"The number of devices to return when fetching from logicmonitor"`
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
