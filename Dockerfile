# Start by building the application.
FROM golang:1.12-alpine AS build
RUN apk add --no-cache git
COPY . /src
USER 1000
WORKDIR /src
ENV GOPATH=/tmp/go
ENV GOCACHE=/tmp/go-cache
ENV CGO_ENABLED=0
RUN go build -o /tmp/bisect-dates

# Now copy it into our base image.
FROM alpine
COPY --from=build /tmp/bisect-dates /bisect-dates
RUN apk add --no-cache tzdata
USER 1000
ENV TZ=UTC
ENTRYPOINT ["/bisect-dates"]
