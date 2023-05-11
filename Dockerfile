FROM golang:1.17-alpine AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# set default env
ARG NAME=$NAME
ENV USER=appuser
ENV UID=10001 
ARG VERSION=$VERSION
ARG TAG=$TAG
ARG BUILD=$BUILD

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

# set dependencies
WORKDIR "${GOPATH}/src/${NAME}/"
COPY . .

# Fetch dependencies.
# Using go get.
RUN go mod download

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags="-w -s -extldflags \"-static\" -X main.Version=${VERSION} -X main.Build=${BUILD} -X main.TAG=${TAG}" -a \
    -o "/go/bin/app" cmd/todolist/app-http/main.go

FROM scratch

ARG VERSION=$VERSION
ARG TAG=$TAG
ARG BUILD=$BUILD

# Import the user and group files from the builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
ENV TZ=Asia/Jakarta
ENV ZONEINFO=/zoneinfo.zip

# Copy our static executable.
COPY --from=builder "/go/bin/app" "/go/bin/app"

# Use an unprivileged user.
USER appuser:appuser

# Run the binary.
ENTRYPOINT ["/go/bin/app"]