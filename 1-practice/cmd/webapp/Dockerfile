FROM golang:alpine AS builder

WORKDIR /build
COPY cmd/webapp .
RUN go mod download
RUN go build -o main .

WORKDIR /app
RUN cp /build/main .

FROM alpine
COPY --from=builder /app/main .

ENTRYPOINT [ "/main" ]
