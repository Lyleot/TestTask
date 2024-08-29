FROM golang:1.22.2 AS build-env
ADD . /src
RUN cd /src && make build

FROM alpine:3.13

RUN apk --no-cache add gettext tzdata
RUN apk --no-cache add libc6-compat
WORKDIR /app
COPY --from=build-env /src/TestTask /app/
COPY --from=build-env /src/migrations/ /app/migrations
RUN chmod +x /app/TestTask
CMD ["/app/TestTask"]