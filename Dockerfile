FROM l.gcr.io/google/bazel:latest AS builder

WORKDIR /scheme/build
ENV BUILD_DIR=/scheme/build

COPY . ${BUILD_DIR}

#RUN bazel test //...
RUN bazel build //...
RUN ls bazel-bin/service/
RUN cp bazel-bin/service/scheme_/scheme .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk update
RUN apk upgrade
RUN apk add --no-cache \
        libc6-compat

WORKDIR /app
COPY --from=builder /scheme/build/service/site /app/site
COPY --from=builder /scheme/build/scheme /app
ENV ELASTICSEARCH_URL=http://elastic:9200
CMD ["/app/scheme"]
