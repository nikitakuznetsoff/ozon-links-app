FROM golang:latest

WORKDIR /go/src/github.com/nikitakuznetsoff/ozon-links-app
COPY . /go/src/github.com/nikitakuznetsoff/ozon-links-app/

RUN go build -o ./bin/linksapp ./cmd/linksapp/


CMD [ "/go/src/github.com/nikitakuznetsoff/ozon-links-app/bin/linksapp" ]

