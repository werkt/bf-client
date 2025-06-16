load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _buildfarm_extension_impl(_ctx):
    http_archive(
        name = "buildfarm",
        build_file = "@//:BUILD.buildfarm",
        sha256 = "bcbe87b95c3b6fd842a479081bf6d176663851e14224d7f9422c2a4265c5e40d",
        strip_prefix = "buildfarm-0b27cd1143e1386dd1bb7f0261e88867901de336/src/main/protobuf/build/buildfarm/v1test/",
        urls = [
            "https://github.com/buildfarm/buildfarm/archive/0b27cd1143e1386dd1bb7f0261e88867901de336.zip",
        ],
    )

build_deps = module_extension(
    implementation = _buildfarm_extension_impl,
)
