FROM golang:1.20-alpine
RUN mkdir /bootsmgmt
WORKDIR /booksmgmt
COPY . .
RUN go mod tidy && go mod download
RUN go build -o app

CMD ["./app"]
