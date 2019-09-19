FROM golang:1.13 as build
WORKDIR /go/src/app
COPY . .

ENV GO111MODULE on
RUN go build -v -o /app .

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=build /app /app
COPY --from=build /go/src/app/views ./views
COPY --from=build /go/src/app/assets ./assets
CMD ["/app"]