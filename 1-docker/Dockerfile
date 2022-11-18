# Start by building the application.
FROM golang:1.18

WORKDIR /go/src/app

COPY ../app.go .

#RUN go mod download
RUN go env -w GO111MODULE=off
RUN CGO_ENABLED=0 go build .
ENV IMAGE BIG
# # Now copy it into our base image.
# FROM gcr.io/distroless/static-debian11
# COPY --from=build /go/bin/app /
CMD ["/go/src/app/app"]