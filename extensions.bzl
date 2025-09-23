load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _buildfarm_extension_impl(_ctx):
    http_archive(
        name = "buildfarm",
        build_file = "@//:BUILD.buildfarm",
        sha256 = "cf47ea9674bde436cfafdcfc772be0fb5998482185c0e339ee4170370f025cf8",
        strip_prefix = "buildfarm-6be2f5e33ca9e3a0c7a2be253d52a53c3df4eddc/src/main/protobuf/build/buildfarm/v1test/",
        urls = [
            "https://github.com/buildfarm/buildfarm/archive/6be2f5e33ca9e3a0c7a2be253d52a53c3df4eddc.zip",
        ],
    )

build_deps = module_extension(
    implementation = _buildfarm_extension_impl,
)
