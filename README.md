# LogicMonitor 2 Connectwise

This service is responsible for adding/synchronizing devices in LogicMonitor to Connectwise.The service runs at defined interval, But can also be triggered manually for synchronizing devices.

--- 
### Required Custom fields for LogicMonitor and/or Connectwise

|Logic Monitor  | Connectwise 
------------- | ------------- 
|customer.name   | company 
|cw_type  | type

---

## Deployment Pipeline

This application is deployed to an Azure Container Instance via Github Actions.  The action that performs this is found in .github/workflows/deploy.yml.  The deployment action is run anytime a `release` is created in Github.  The tag for the release should follow symantec versioning.  The environment variables that configure the deployment can be found in the deploy.yml file.  The secrets can be updated in `settings` for the repo (and the org settings).

# Documentation

* [Overview](docs/README.md)
* [Configuration](docs/config.md)
* [Metrics](docs/metrics.md)
* [API Overview](docs/api/README.md)
* [Routes](docs/routes.md)
* [External Libraries](docs/external_libraries.md)
