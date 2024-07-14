ARG GO_VERSION=1.22
ARG ALPINE_VERSION=3.18

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as go-builder

WORKDIR /go/src/kotatuneko-rest

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o api cmd/main.go


FROM alpine:${ALPINE_VERSION}

WORKDIR /usr/src/kotatuneko-rest

COPY --from=go-builder /go/src/kotatuneko-rest/api api
COPY --from=go-builder /go/src/kotatuneko-rest/api .env
RUN chmod +x "/usr/src/kotatuneko-rest/api"

ENTRYPOINT ["/usr/src/kotatuneko-rest/api"]
