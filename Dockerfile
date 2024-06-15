FROM golang:1.22.3-alpine

WORKDIR /app


COPY . .
RUN apk add --no-cache python3
RUN apk add --no-cache curl
RUN go build -o main .

CMD ["./main"]