load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _buildfarm_extension_impl(_ctx):
    http_archive(
        name = "buildfarm",
        build_file = "@//:BUILD.buildfarm",
        sha256 = "64759237ce8f8e827f76df114db84d1bdb23ebcca623da7d2ad04eeb2eb264c6",
        strip_prefix = "buildfarm-ea377e435be6c1575282a6c5586c021c41909783/src/main/protobuf/build/buildfarm/v1test/",
        urls = [
            "https://github.com/buildfarm/buildfarm/archive/ea377e435be6c1575282a6c5586c021c41909783.zip",
        ],
    )

build_deps = module_extension(
    implementation = _buildfarm_extension_impl,
)
