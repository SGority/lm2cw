# LogicMonitor 2 Connectwise (LM2CW)

This service is responsible for adding/synchronizing devices in LogicMonitor to Connectwise.The service runs at defined interval, But can also be triggered manually for synchronizing devices.

![Container Diagram](diagrams/images/Container.png)

<br/>

![Component Diagram](diagrams/images/Component.png)

## Synchronizing devices

The synchronizing of devices from logicmonitor to connectwise happens in two ways.

1. Automatic synching with cron in defined intervals. Currently the interval is set to 180 minutes.
2. Manually triggering the endpoint responsible for synchronizing devices.

The device synchronization
1. Gets all devices from logicmonitor.
2. Parse the devices to be used in connectwise.
3. Add or update those device in connectwise.

### Required Custom fields for LogicMonitor and/or Connectwise

|Logic Monitor  | Connectwise 
------------- | ------------- 
|customer.name   | company 
|cw_type  | type

For the sync to be successful, the following must be true:

* customer.name is set in LogicMonitor and the name exactly matches a company name in Connectwise
* cw_type  has been set on the device in LogicMonitor.  The value *must* match a type in Connectwise
* The `devicetype` in LogicMonitor *must* be 0
