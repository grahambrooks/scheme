
go get -u ./...

bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

bazel run //:gazelle
