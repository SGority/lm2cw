# Configuration

Configuration is typically done with environment variables or command-line flags.  Not all settings have command-line flags or default values.  If there's no value set the feature is disabled.  

The order of preference is:

   1. command-line flag
   1. environment variable
   1. variable defined in `.env` file
   1. default

## Configuration Variables and Flags

| Variable           | Flag            | Default | Description                                                         |
| ------------------ | --------------- | ------- | ------------------------------------------------------------------- |
| PORT               | -p,<br/> --port | 8080    | The port on which the service listens                               |
| JOB_INTERVAL       |                 | 180     | The interval at which the scheduler is invoked specified in minutes |
| CW_USER            |                 |         | The username required for connectwise authentication                |
| CW_PASS            |                 |         | The password required for connectwise authentication                |
| CW_URL             |                |   https://na.myconnectwise.net/v4_6_release/apis/3.0       | The base url to connect to connectwise.                                               |
| CW_COMPANY         |                 |         | The company name in connectwise                                     |
| CW_COMPANY_ID      |                 |         | The company ID for connectwise authentication                                     |
| BASE_URL           |                 |  https://magna5global.logicmonitor.com/santaba/rest       | The base URL used to connect to logicmonitor                 |
| LM_ACCESS_ID       |                 |         | The access id required for logicmonitor authentication              |
| LM_ACCESS_KEY      |                 |         | The access key required for logicmonitor authentication             |
| DEVICE_SOURCE_PATH |                 |  /device/devices       | The path to fetch the device resource from Logicmonitor                                                  |
| DEVICE_OFFSET_SIZE        |                 | 100     | The number of devices to return from logicmonitor           |
| SMTP_HOST           |                 |         |SMTP host|
| SMTP_PORT           |                 | 587     |SMTP port|
| MAIL_TO             |                 |         |Mail address of the recipient to whom the mail is being sent.|
| MAIL_FROM           |                 |         |Mail address of the sender.|
| MAIL_PASS           |                 |         |Password of the sender.|
