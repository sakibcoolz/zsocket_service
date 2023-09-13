FROM golang:alpine AS builder
RUN apk add --no-cache git gcc libc-dev
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fire_service ./cmd/
RUN ls -lrth

# final stage
FROM alpine:latest
LABEL Name=fire_service Version=0.0.1
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/fire_service fire_service
EXPOSE 8080
ENTRYPOINT ["/fire_service"]

