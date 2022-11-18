# Start by building the application.
FROM golang:1.18 as build

WORKDIR /go/src/app
COPY . .

#RUN go mod download
RUN go env -w GO111MODULE=off
RUN CGO_ENABLED=0 go build .

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian11
COPY --from=build /go/src/app/app /
ENV IMAGE IN_KUBERNETES
CMD ["/app"]