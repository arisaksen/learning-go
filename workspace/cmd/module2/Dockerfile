#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app

COPY cmd/module2 ./cmd/module2
COPY common ./common

RUN go work init
RUN go work use cmd/module2
RUN go work use common

RUN go mod download && go mod verify
WORKDIR /go/src/app/cmd/module2
RUN go build -o /app

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app /app

ENTRYPOINT /app