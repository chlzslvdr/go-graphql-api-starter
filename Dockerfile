FROM golang
ENV PORT 4000
EXPOSE 4000

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o app

CMD ["./app"]
