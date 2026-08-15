# syntax=docker/dockerfile:1

# -----------------------------------------------------------------------------
# Build stage: compile a static, stripped Go binary.
# -----------------------------------------------------------------------------
FROM golang:1.25-alpine AS build

WORKDIR /src

# Warm the module cache first so source changes don't invalidate this layer.
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
        -trimpath \
        -ldflags="-s -w" \
        -o /out/server \
        ./cmd/server

# -----------------------------------------------------------------------------
# Runtime stage: minimal, non-root distroless image (no shell, no package
# manager). Includes CA certificates for TLS to PostgreSQL.
# -----------------------------------------------------------------------------
FROM gcr.io/distroless/static-debian12:nonroot

# The app listens on REST_PORT (default 8080).
EXPOSE 8080

COPY --from=build /out/server /server

# The :nonroot tag already runs as the "nonroot" user (uid 65532).
ENTRYPOINT ["/server"]
