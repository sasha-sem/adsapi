FROM golang:1.15.8-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc musl-dev
        
WORKDIR /app
        
COPY go.mod go.sum ./
        
RUN go mod download
        
COPY . .
        
RUN go build -o main .
        
CMD ["./main"]