# ─────────────────────────────────────────────────────────────
# Stage 1: Builder base
# ─────────────────────────────────────────────────────────────
FROM ubuntu:latest AS builder

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y curl git build-essential ca-certificates && \
    curl -LO https://go.dev/dl/go1.24.1.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.24.1.linux-amd64.tar.gz && \
    rm go1.24.1.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:$PATH"

# ─────────────────────────────────────────────────────────────
# Stage 2: Build application
# ─────────────────────────────────────────────────────────────
FROM builder AS buildapp

WORKDIR /app
COPY . .

RUN make build
#RUN go mod tidy
#RUN go build -o surveillance-proxy ./main


# ─────────────────────────────────────────────────────────────
# Stage 3: Runtime (distroless)
# ─────────────────────────────────────────────────────────────
FROM gcr.io/distroless/static:nonroot

WORKDIR /opt
COPY --from=buildapp /app/build/surveillance-proxy /usr/bin/

ENTRYPOINT ["/usr/bin/surveillance-proxy"]
