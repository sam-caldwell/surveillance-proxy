# Surveillance Proxy

A lightweight Go-based webhook server that receives alerts from Ubiquiti Dream Machine Pro Max, creates corresponding
Jira tickets, and uploads associated images.

## Features

- Listens for authenticated HTTP POST webhooks from Ubiquiti cameras
- Generates unique event IDs
- Downloads associated image thumbnails
- Creates Jira tickets in the `SRV` project
- Uploads event thumbnails as attachments
- Fully containerized with a distroless production image

## Docker image
```bash
docker pull ghcr.io/sam-caldwell/surveillance-proxy:latest
```

## Requirements

- Go 1.22+
- Docker (for container builds)
- Atlassian Jira Cloud API token
- Valid Ubiquiti alert payload (JSON)

## Environment Variables

| Name            | Description                                             |
|-----------------|---------------------------------------------------------|
| `RECVR_ADDRESS` | Listener address (e.g. `0.0.0.0:8080`)                  |
| `AUTH_TOKEN`    | Bearer token required by sender                         |
| `JIRA_USER`     | Jira user email                                         |
| `JIRA_TOKEN`    | Jira API token                                          |
| `JIRA_BASE_URL` | Jira base URL (e.g. `https://yourdomain.atlassian.net`) |
| `JIRA_PROJECT`  | Jira project key (e.g. `SRV`)                           |

## Build

```bash
make clean build
```

## Run

```bash
RECVR_ADDRESS=0.0.0.0:8080 AUTH_TOKEN=yourtoken \
JIRA_USER=you@example.com JIRA_TOKEN=xxx \
JIRA_BASE_URL=https://yourdomain.atlassian.net \
JIRA_PROJECT=SRV \
./build/surveillance-proxy
```

## Docker

Build the container:

```bash
make docker
```

Run it:

```bash
docker run -p 8080:8080 \
  -e RECVR_ADDRESS=0.0.0.0:8080 \
  -e AUTH_TOKEN=yourtoken \
  -e JIRA_USER=you@example.com \
  -e JIRA_TOKEN=xxx \
  -e JIRA_BASE_URL=https://yourdomain.atlassian.net \
  -e JIRA_PROJECT=SRV \
  surveillance-proxy:latest
```

## License

MIT

