FROM golang:1.19-alpine3.16 AS builder

RUN apk -U --no-cache add build-base git gcc

WORKDIR /go/src/github.com/ory/defcon-30-ctf

ADD go.mod go.mod
ADD go.sum go.sum

RUN go mod download

ADD . .

RUN go build -tags sqlite -o /usr/bin/backend .

FROM alpine:3.16

RUN addgroup -S ory; \
    adduser -S ory -G ory -D  -h /home/ory -s /bin/nologin; \
    chown -R ory:ory /home/ory

COPY --from=builder /usr/bin/backend /usr/bin/backend

USER ory

ENTRYPOINT ["backend"]
