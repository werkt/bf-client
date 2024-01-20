load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "view",
    srcs = [
        "action.go",
        "command.go",
        "input.go",
        "queue.go",
        "operation.go",
        "operation_list.go",
        "view.go",
    ],
    importpath = "github.com/werkt/bf-client/view",
    visibility = ["//visibility:public"],
    deps = [
        "//client",
        "@buildfarm//:go_default_library",
        "@com_github_gammazero_deque//:go_default_library",
        "@com_github_gizak_termui_v3//:go_default_library",
        "@com_github_gizak_termui_v3//widgets:go_default_library",
        "@com_github_golang_protobuf//jsonpb:go_default_library_gen",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_golang_protobuf//ptypes:go_default_library_gen",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes:go_default_library",
        "@org_golang_google_grpc//status:go_default_library",
        "@com_github_hashicorp_golang_lru_v2//:go_default_library",
        "@go_googleapis//google/longrunning:longrunning_go_proto",
        "@remoteapis//build/bazel/remote/execution/v2:go_default_library",
    ],
)
