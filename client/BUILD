load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "app.go",
        "bytestream.go",
        "cas.go",
        "digest.go",
        "hashtag.go",
        "list.go",
        "queue.go",
        "tree.go",
        "unified_redis.go",
    ],
    importpath = "github.com/werkt/bf-client/client",
    visibility = ["//visibility:public"],
    deps = [
        "//third_party/buildfarm:go_default_library",
        "@com_github_gizak_termui_v3//:go_default_library",
        "@com_github_golang_protobuf//jsonpb:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_mattn_go_runewidth//:go_default_library",
        "@com_github_redis_go_redis_v9//:go_default_library",
        "@org_golang_google_genproto//googleapis/longrunning:go_default_library",
        "@org_golang_google_genproto_googleapis_bytestream//:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//credentials:go_default_library",
        "@remoteapis//build/bazel/remote/execution/v2:go_default_library",
    ],
)
