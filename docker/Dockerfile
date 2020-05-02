FROM l.gcr.io/google/bazel:latest AS builder

WORKDIR /apellicon/build
ENV BUILD_DIR=/apellicon/build

COPY . ${BUILD_DIR}

RUN bazel test //...
RUN bazel build
RUN ls bazel-bin/service/linux_amd64_stripped
RUN cp bazel-bin/service/linux_amd64_stripped/apellicon .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk update
RUN apk upgrade
RUN apk add --no-cache \
        libc6-compat

WORKDIR /app
COPY --from=builder /apellicon/build/service/site /app/site
COPY --from=builder /apellicon/build/apellicon /app
ENV ELASTICSEARCH_URL=http://elastic:9200
CMD ["/app/apellicon"]
