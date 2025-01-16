FROM golang:1.22 as builder
WORKDIR /go/src
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app .

FROM alpine:latest as runner
COPY --from=builder /go/bin/app /bin/app
CMD [ "ls", "-l", "/bin/app" ]
ENTRYPOINT ["/bin/app"]
