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
| JWKS_RENEW_MINUTES |                 | 60      | Number of minutes to wait before renewing JWKS certificates         |
| JWT_ISSUER         |                 |         | The URL to the JWT issuing server                                   |
| JOB_INTERVAL       |                 | 180     | The interval at which the scheduler is invoked Specified in minutes |
| CW_USER            |                 |         | The username required for connectwise authentication                |
| CW_PASS            |                 |         | The password required for connectwise authentication                |
| CW_URL             |                 |         | The connectwise url                                                 |
| CW_COMPANY         |                 |         | The company name in connectwise                                     |
| CW_COMPANY_ID      |                 |         | The company ID for connectwise                                      |
| BASE_URL           |                 |         | The base URL used in making request to logicmonitor                 |
| LM_ACCESS_ID       |                 |         | The access id required for logicmonitor authentication              |
| LM_ACCESS_KEY      |                 |         | The access key required for logicmonitor authentication             |
| DEVICE_SOURCE_PATH |                 |         | The resource path                                                   |
| OFFSET_SIZE        |                 | 100     | The offset size when fetching any resources            |
| SMTP_HOST           |                 |         |SMTP host|
| SMTP_PORT           |                 | 587     |SMTP port|
| MAIL_TO             |                 |         |Mail address of the recipient to whom the mail is being sent.|
| MAIL_FROM           |                 |         |Mail address of the sender.|
| MAIL_PASS           |                 |         |Password of the sender.|