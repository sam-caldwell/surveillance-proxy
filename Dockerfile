# ─────────────────────────────────────────────────────────────
# Stage 1: Builder base
# ─────────────────────────────────────────────────────────────
FROM ubuntu:latest AS builder

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y curl git build-essential ca-certificates && \
    curl -LO https://go.dev/dl/go1.22.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz && \
    rm go1.22.2.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:$PATH"

WORKDIR /app
COPY . .


# ─────────────────────────────────────────────────────────────
# Stage 2: Build application
# ─────────────────────────────────────────────────────────────
FROM builder AS buildapp

RUN go mod download
RUN go build -o surveillance-proxy ./main


# ─────────────────────────────────────────────────────────────
# Stage 3: Runtime (distroless)
# ─────────────────────────────────────────────────────────────
FROM gcr.io/distroless/static:nonroot

WORKDIR /opt
COPY --from=buildapp /app/surveillance-proxy /opt/

ENTRYPOINT ["/opt/surveillance-proxy"]
