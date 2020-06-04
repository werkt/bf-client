load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/bazelbuild/bazel-buildfarm",
    visibility = ["//visibility:private"],
    deps = [
        "@buildfarm//:go_default_library",
        "@com_github_gizak_termui//:go_default_library",
        "@com_github_gizak_termui//widgets:go_default_library",
        "@com_github_redis_go_redis_v9//:go_default_library",
        "@com_github_golang_protobuf//jsonpb:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_golang_protobuf//ptypes:go_default_library",
        "@com_github_nsf_termbox_go//:go_default_library",
        "@go_googleapis//google/longrunning:longrunning_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes:go_default_library",
        "@org_golang_google_grpc//status:go_default_library",
        "@remoteapis//build/bazel/remote/execution/v2:go_default_library",
    ],
)

go_binary(
    name = "client",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

go_binary(
    name = "test",
    srcs = ["test.go"],
    deps = [
        "@com_github_golang_protobuf//proto:go_default_library",
        "@remoteapis//build/bazel/remote/execution/v2:go_default_library",
    ],
)

go_binary(
    name = "builds",
    srcs = ["builds.go"],
    visibility = ["//visibility:public"],
    deps = [
        "@buildfarm//:go_default_library",
        "@com_github_redis_go_redis_v9//:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_golang_protobuf//ptypes:go_default_library",
        "@go_googleapis//google/longrunning:longrunning_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "@remoteapis//build/bazel/remote/execution/v2:go_default_library",
    ],
)
