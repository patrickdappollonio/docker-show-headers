FROM golang as builder
WORKDIR app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o /server

FROM scratch
COPY --from=builder /server /server
EXPOSE 8080

ENTRYPOINT ["/server"]
