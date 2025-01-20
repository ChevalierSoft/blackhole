FROM golang:1.23.4-alpine3.21 AS builder
WORKDIR /go/src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app .
ENTRYPOINT ["/go/bin/app"]

FROM alpine:3.21 AS runner
COPY --from=builder /go/bin/app /bin/app
COPY *.yaml /bin/
ENTRYPOINT ["/bin/app"]

# ENV BH_PORT=443
# ENV BH_DOMAIN_NAME=???
# ENV BH_EMAIL="my_email@it.oui"
