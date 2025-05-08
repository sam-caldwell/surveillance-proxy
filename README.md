Surveillance-Proxy
==================

This is a golang microservice which receives events from Ubiquity Dream Machine devices, transforms the event into
target API requests (e.g. JIRA ticket-create API calls).

## Build
```shell
make clean build
```

## Package
```shell
make docker
```

## Parameters
```shell
RECVR_PORT=8080
AUTH_TOKEN=your_webhook_token
JIRA_USER=email@example.com
JIRA_TOKEN=jira_api_token
JIRA_BASE_URL=https://yourdomain.atlassian.net
```
