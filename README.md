# Metal Archives Wrapper

[Project's page](https://musicmanager.gitpages.windmaker.net/metal-archives-wrapper)

[![pipeline status](https://git.windmaker.net/musicmanager/metal-archives-wrapper/badges/master/pipeline.svg)](https://git.windmaker.net/musicmanager/metal-archives-wrapper/pipelines)[![coverage report](https://git.windmaker.net/musicmanager/metal-archives-wrapper/badges/master/coverage.svg)](https://musicmanager.gitpages.windmaker.net/metal-archives-wrapper/coverage.html)

This service is part of [Music Manager](https://git.windmaker.net/musicmanager) Project, see [docs](https://git.windmaker.net/musicmanager/Music-Manager-Docs) for more details.

## Behavior

This service retrieves Artists and Records from [Metal Archives](https://www.metal-archives.com/).

If there is no info, this service publish a new job regarding a not found request.

### Config example

To do

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
