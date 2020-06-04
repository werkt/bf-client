load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "com_google_protobuf",
    sha256 = "dd513a79c7d7e45cbaeaf7655289f78fd6b806e52dbbd7018ef4e3cf5cff697a",
    strip_prefix = "protobuf-3.15.8",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.15.8.zip"],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "727f3e4edd96ea20c29e8c2ca9e8d2af724d8c7778e7923a854b2c80952bc405",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.30.0/bazel-gazelle-v0.30.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.30.0/bazel-gazelle-v0.30.0.tar.gz",
    ],
)

http_archive(
    name = "remoteapis",
    sha256 = "3e26758ce8ee82078ed991ef6362308a0933bd5cfd3c78aff7ef88065e44b45c",
    strip_prefix = "remote-apis-068363a3625e166056c155f6441cfb35ca8dfbf2",
    urls = ["https://github.com/bazelbuild/remote-apis/archive/068363a3625e166056c155f6441cfb35ca8dfbf2.tar.gz"],
)

http_archive(
    name = "googleapis",
    build_file = "@remoteapis//:external/BUILD.googleapis",
    patch_cmds = ["find google -name 'BUILD.bazel' -type f -delete"],
    patch_cmds_win = ["Remove-Item google -Recurse -Include *.bazel"],
    sha256 = "7b6ea252f0b8fb5cd722f45feb83e115b689909bbb6a393a873b6cbad4ceae1d",
    strip_prefix = "googleapis-143084a2624b6591ee1f9d23e7f5241856642f4d",
    urls = ["https://github.com/googleapis/googleapis/archive/143084a2624b6591ee1f9d23e7f5241856642f4d.zip"],
)

http_archive(
    name = "com_github_grpc_grpc",
    sha256 = "8da7f32cc8978010d2060d740362748441b81a34e5425e108596d3fcd63a97f2",
    strip_prefix = "grpc-1.21.0",
    urls = [
        "https://github.com/grpc/grpc/archive/v1.21.0.tar.gz",
        "https://mirror.bazel.build/github.com/grpc/grpc/archive/v1.21.0.tar.gz",
    ],
)

http_archive(
    name = "buildfarm",
    build_file = "@//:BUILD.buildfarm",
    sha256 = "e7c4646b94d4f533888327b0e60ad91687123ca91a000485d02c6428b4cada89",
    strip_prefix = "bazel-buildfarm-2.5.0/src/main/protobuf/build/buildfarm/v1test/",
    urls = [
        "https://github.com/bazelbuild/bazel-buildfarm/archive/refs/tags/2.5.0.tar.gz",
    ],
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "51dc53293afe317d2696d4d6433a4c33feedb7748a9e352072e2ec3c0dafd2c6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
    ],
)

http_archive(
    name = "build_stack_rules_proto",
    sha256 = "c62f0b442e82a6152fcd5b1c0b7c4028233a9e314078952b6b04253421d56d61",
    strip_prefix = "rules_proto-b93b544f851fdcd3fc5c3d47aee3b7ca158a8841",
    urls = ["https://github.com/stackb/rules_proto/archive/b93b544f851fdcd3fc5c3d47aee3b7ca158a8841.tar.gz"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.20.5")

load("@build_stack_rules_proto//go:deps.bzl", "go_grpc_library")

go_grpc_library()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "com_github_cespare_xxhash_v2",
    importpath = "github.com/cespare/xxhash/v2",
    sum = "h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=",
    version = "v2.2.0",
)

go_repository(
    name = "com_github_dgryski_go_rendezvous",
    importpath = "github.com/dgryski/go-rendezvous",
    sum = "h1:lO4WD4F/rVNCu3HqELle0jiPLLBs70cWOduZpkS1E78=",
    version = "v0.0.0-20200823014737-9f7001d12a5f",
)

gazelle_dependencies()

go_repository(
    name = "com_github_redis_go_redis_v9",
    importpath = "github.com/redis/go-redis/v9",
    sum = "h1:CuQcn5HIEeK7BgElubPP8CGtE0KakrnbBSTLjathl5o=",
    version = "v9.0.5",
)

load("@io_bazel_rules_go//tests:grpc_repos.bzl", "grpc_dependencies")

grpc_dependencies()

go_repository(
    name = "com_github_mitchellh_go_wordwrap",
    importpath = "github.com/mitchellh/go-wordwrap",
    sum = "h1:DpOJ2HYzCv8LZP15IdmG+YdwD2luVPHITV96TkirNBM=",
    version = "v0.0.0-20150314170334-ad45545899c7",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    importpath = "github.com/mattn/go-runewidth",
    sum = "h1:UnlwIPBGaTZfPQ6T1IGzPI0EkYAQmT9fAEJ/poFC63o=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_nsf_termbox_go",
    importpath = "github.com/nsf/termbox-go",
    sum = "h1:x3S6kxmy49zXVVyhcnrFqxvNVCBPb2KZ9hV2RBdS840=",
    version = "v0.0.0-20190121233118-02980233997d",
)

go_repository(
    name = "com_github_gizak_termui",
    importpath = "github.com/gizak/termui/v3",
    sum = "h1:ZZmVDgwHl7gR7elfKf1xc4IudXZ5qqfDh4wExk4Iajc=",
    version = "v3.1.0",
)

go_repository(
    name = "com_github_go_redis_redis",
    importpath = "github.com/go-redis/redis",
    sum = "h1:x7dvutgTPxsHoeUboXlOp9fEUMdIRO4OJV5ly1LYE40=",
    version = "v6.15.3",
)
