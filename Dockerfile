FROM golang:latest

WORKDIR /gotinyurl

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .    

RUN go build -o ./go-tiny-url ./cmd/tiny-url

CMD [ "./go-tiny-url"]