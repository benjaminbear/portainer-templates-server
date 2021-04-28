FROM golang:latest AS builder

ENV GO111MODULE=on
WORKDIR /go/src/server

COPY . .
RUN TAG=$(curl --silent "https://api.github.com/repos/benjaminbear/portainer-templates-server/releases/latest" | grep -Po '"tag_name": "\K.*?(?=")') \
    && CGO_ENABLED=0 go build -o server -ldflags "-s -w -X main.Version=${TAG}"

FROM scratch

COPY --from=builder /go/src/server/server /server
COPY templates.json .

EXPOSE 8080

ENTRYPOINT [ "/server" ]