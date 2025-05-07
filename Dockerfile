FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN CGO_ENABLE=0 GOOS=linux go build -o /docker-app

FROM alpine:latest

WORKDIR /app
COPY --from=builder /docker-app /app/docker-app

ENV PORT=8080
ENV NAME="CLOUD COMPUTING STUDENT"

EXPOSE 8080

CMD ["/app/docker-app"]
