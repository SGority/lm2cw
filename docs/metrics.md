# Metrics

The API exposes a variety of metrics in standard Prometheus format.  These metrics are exposed at the `/metrics` endpoint.

| Name                       | Type    | Labels                                   | Description                                        |
| -------------------------- | ------- | ---------------------------------------- | -------------------------------------------------- |
| api_processing_ops_total   | gauge   |                                          | The number of events processing                    |
| api_responses_total        | counter | - HTTP Status<br/>- method<br>- endpoint | The number of responses by endpoint and status     |
| error_counter              | counter |                                          | Total error count                                  |
| devices_synchronized_gauge | gauge   | company                                  | Number of devices synchronized                     |
| last_sync_time             | gauge   | status                                   | Last time when devices where synched               |
| last_sync_duration         | gauge   |                                          | Time taken when last time the devices were synched |
| next_sync_time             | gauge   |                                          | Next time the devices would be synched             |