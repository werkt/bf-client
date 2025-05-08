load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _buildfarm_extension_impl(_ctx):
    http_archive(
        name = "buildfarm",
        build_file = "@//:BUILD.buildfarm",
        sha256 = "ed8907a3c20efd1f74f21329cd05f174399e19f65b3d8578bcad246014c88338",
        strip_prefix = "buildfarm-2f939dfcc71e19ab06cabe8aee78671118ec0e16/src/main/protobuf/build/buildfarm/v1test/",
        urls = [
            "https://github.com/buildfarm/buildfarm/archive/2f939dfcc71e19ab06cabe8aee78671118ec0e16.zip",
        ],
    )

build_deps = module_extension(
    implementation = _buildfarm_extension_impl,
)
