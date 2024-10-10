load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_proto_library(
    name = "buildfarm_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/buildfarm/buildfarm/build/buildfarm/v1test",
    proto = "@buildfarm//:buildfarm_proto",
    deps = [
        "@com_google_cloud_go_longrunning//autogen/longrunningpb:go_default_library",
        "@org_golang_google_genproto//googleapis/longrunning:go_default_library",
        "@org_golang_google_genproto_googleapis_api//annotations:go_default_library",
        "@org_golang_google_genproto_googleapis_rpc//status:go_default_library",
        "@remoteapis//build/bazel/remote/execution/v2:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    embed = [":buildfarm_go_proto"],
    importpath = "github.com/buildfarm/buildfarm/build/buildfarm/v1test",
    visibility = ["//visibility:public"],
)
