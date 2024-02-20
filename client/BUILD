load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "client",
    srcs = [
        "app.go",
        "bytestream.go",
        "digest.go",
        "hashtag.go",
        "list.go",
        "queue.go",
        "unified_redis.go",
    ],
    importpath = "github.com/werkt/bf-client/client",
    visibility = ["//visibility:public"],
    deps = [
        "@buildfarm//:go_default_library",
        "@com_github_gizak_termui_v3//:go_default_library",
        "@com_github_golang_protobuf//jsonpb:go_default_library_gen",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_mattn_go_runewidth//:go_default_library",
        "@com_github_redis_go_redis_v9//:go_default_library",
        "@go_googleapis//google/bytestream:bytestream_go_proto",
        "@go_googleapis//google/longrunning:longrunning_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//credentials:go_default_library",
        "@remoteapis//build/bazel/remote/execution/v2:go_default_library",
    ],
)
