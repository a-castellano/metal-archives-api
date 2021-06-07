# Metal Archives Wrapper

[Project's page](https://musicmanager.gitpages.windmaker.net/metal-archives-wrapper)

[![pipeline status](https://git.windmaker.net/musicmanager/metal-archives-wrapper/badges/master/pipeline.svg)](https://git.windmaker.net/musicmanager/metal-archives-wrapper/pipelines)[![coverage report](https://git.windmaker.net/musicmanager/metal-archives-wrapper/badges/master/coverage.svg)](https://musicmanager.gitpages.windmaker.net/metal-archives-wrapper/coverage.html)[![Quality Gate Status](https://sonarqube.windmaker.net/api/project_badges/measure?project=metal-archives-wrapper&metric=alert_status)](https://sonarqube.windmaker.net/dashboard?id=metal-archives-wrapper)

This service is part of [Music Manager](https://git.windmaker.net/musicmanager) Project, see [docs](https://git.windmaker.net/musicmanager/Music-Manager-Docs) for more details.

## Behavior

This service retrieves Artists and Records from [Metal Archives](https://www.metal-archives.com/).

The service receives jobs sent from [Job Manager](https://git.windmaker.net/musicmanager/Job-Manager), and process them. For each processed job this service will generate a new job containing process status and result.

### Config example

This service will look for its config in **/etc/music-manager-service/config.toml**, parent folder can be changed setting the environment variable **MUSIC_MANAGER_SERVICE_CONFIG_FILE_LOCATION**.

Here is a config example:

```toml
[server]
host = "localhost"
port = 5672
user = "guest"
password = "pass"

[incoming]
name = "incoming"

[outgoing]
name = "outgoing"
```

## Testing

### Unit tests

Use make to run unit tests:
```bash
make test
```

### Integration tests

Docker is required for make the following tests run, user must be sudoer too.
```bash
bash scripts/start_rabbitmq_test_server.sh
make test_integration
```
