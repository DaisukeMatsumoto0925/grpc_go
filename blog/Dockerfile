FROM golang:1.17.8-buster as gobuilder
WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build -ldflags="-w -s" -o /go/bin/app
RUN go get github.com/google/wire/cmd/wire@latest && \
  go get -u github.com/kyoh86/richgo
RUN go generate
RUN go get github.com/cosmtrek/air
CMD ["air"]

FROM gcr.io/distroless/base
COPY --from=gobuilder /go/bin/app /
CMD ["/app"]
