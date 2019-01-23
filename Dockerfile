FROM golang:latest as builder

WORKDIR /go
RUN go get -u gopkg.in/resty.v1
ADD server.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o server server.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/server .
ENV PORT 9090
ENV CODE 200
ENV CONTENT 'Hello from tiny mock server'
ENTRYPOINT ["sh", "-c", "PORT=$PORT CODE=$CODE CONTENT=$CONTENT ./server"]