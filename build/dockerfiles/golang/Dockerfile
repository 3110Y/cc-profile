# ------------------------------------------------------------------------------
# Base image
# ------------------------------------------------------------------------------
FROM golang:1.21-alpine AS golang_base
ENV GO111MODULE="on"
ENV GOOS=linux
WORKDIR /app
COPY . .
RUN go mod tidy

# ------------------------------------------------------------------------------
# Test image
# ------------------------------------------------------------------------------
FROM golang_base AS golang_test
ENV GO111MODULE="on"
ENV GOOS=linux
WORKDIR /app
CMD go test ./...


# ------------------------------------------------------------------------------
# Build image
# ------------------------------------------------------------------------------
FROM golang_base AS golang_build
ENV GO111MODULE="on"
ENV GOOS=linux
WORKDIR /app
RUN go build -o /var/www/server /app/cmd/main.go

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------
FROM golang_base AS golang_dev
ENV APP_PORT=8080
VOLUME /app
WORKDIR /app
EXPOSE ${APP_PORT}
CMD go run /app/cmd/main.go


# ------------------------------------------------------------------------------
# Production image
# ------------------------------------------------------------------------------
FROM alpine:3.7 AS golang_prod
ENV APP_PORT=5300
COPY --from=golang_base /etc/passwd /etc/passwd
COPY --from=golang_build /var/www/server /var/www/server
COPY --from=golang_build /app/.env /var/www/.env
RUN chmod 0755 /var/www/server
WORKDIR /var/www
EXPOSE ${APP_PORT}
CMD ["/var/www/server"]