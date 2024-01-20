load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "co_honnef_go_tools",
        importpath = "honnef.co/go/tools",
        sum = "h1:/hemPrYIhOhy8zYrNj+069zDB68us2sMGsfkFJO0iZs=",
        version = "v0.0.0-20190523083050-ea95bdfd59fc",
    )
    go_repository(
        name = "com_github_alauda_go_redis_client",
        importpath = "github.com/alauda/go-redis-client",
        sum = "h1:q1Iu9uJZMlgGL3TfB7ZgK7As+z6PMDzvhltnEUdlqDo=",
        version = "v0.0.0-20190714061509-7d5ca20d0820",
    )
    go_repository(
        name = "com_github_armon_go_metrics",
        importpath = "github.com/armon/go-metrics",
        sum = "h1:hR91U9KYmb6bLBYLQjyM+3j+rcd/UhE+G78SFnF8gJA=",
        version = "v0.4.1",
    )

    go_repository(
        name = "com_github_bazelbuild_remote_apis",
        importpath = "github.com/bazelbuild/remote-apis",
        sum = "h1:pYzFHP6FWkA8FjSyx+Kx99zeyVjiRB1epllcRIMDIbc=",
        version = "v0.0.0-20231221155620-d20ae8b97fd3",
    )
    go_repository(
        name = "com_github_bsm_ginkgo_v2",
        importpath = "github.com/bsm/ginkgo/v2",
        sum = "h1:Ny8MWAHyOepLGlLKYmXG4IEkioBysk6GpaRTLC8zwWs=",
        version = "v2.12.0",
    )
    go_repository(
        name = "com_github_bsm_gomega",
        importpath = "github.com/bsm/gomega",
        sum = "h1:yeMWxP2pV2fG3FgAODIY8EiRE3dy0aeFYt4l7wh6yKA=",
        version = "v1.27.10",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=",
        version = "v2.2.0",
    )

    go_repository(
        name = "com_github_client9_misspell",
        importpath = "github.com/client9/misspell",
        sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_cncf_udpa_go",
        importpath = "github.com/cncf/udpa/go",
        sum = "h1:QQ3GSy+MqSHxm/d8nCtnAiZdYFd45cYZPs8vOOIYKfk=",
        version = "v0.0.0-20220112060539-c52dc94e7fbe",
    )
    go_repository(
        name = "com_github_cncf_xds_go",
        importpath = "github.com/cncf/xds/go",
        sum = "h1:/inchEIKaYC1Akx+H+gqO04wryn5h75LSazbRlnya1k=",
        version = "v0.0.0-20230607035331-e9ce68804cb4",
    )
    go_repository(
        name = "com_github_coreos_go_semver",
        importpath = "github.com/coreos/go-semver",
        sum = "h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_coreos_go_systemd_v22",
        importpath = "github.com/coreos/go-systemd/v22",
        sum = "h1:D9/bQk5vlXQFZ6Kwuu6zaiXJ9oTPe68++AzAJc1DzSI=",
        version = "v22.3.2",
    )

    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:ZDRjVQ15GmhC3fiQ8ni8+OwkZQO4DARzQgrnXU1Liz8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_dgryski_go_rendezvous",
        importpath = "github.com/dgryski/go-rendezvous",
        sum = "h1:lO4WD4F/rVNCu3HqELle0jiPLLBs70cWOduZpkS1E78=",
        version = "v0.0.0-20200823014737-9f7001d12a5f",
    )

    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:wSUXTlLfiAQRWs2F+p+EKOY9rUyis1MyGqJ2DIk5HpM=",
        version = "v0.11.1",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:QkIBuU5k+x7/QXPvPPnWXWlCdaBFApVqftFV6k087DA=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_fatih_color",
        importpath = "github.com/fatih/color",
        sum = "h1:qfhVLaG5s+nCROl1zJsZRxFeYrHLqWroPOQ8BWiNb4w=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_github_frankban_quicktest",
        importpath = "github.com/frankban/quicktest",
        sum = "h1:7Xjx+VpznH+oBnejlPUj8oUpdxnVs4f8XU8WnHkI4W8=",
        version = "v1.14.6",
    )
    go_repository(
        name = "com_github_fsnotify_fsnotify",
        importpath = "github.com/fsnotify/fsnotify",
        sum = "h1:8JEhPFa5W2WU7YfeZzPNqzMP6Lwt7L2715Ggo0nosvA=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_gammazero_deque",
        importpath = "github.com/gammazero/deque",
        sum = "h1:qSdsbG6pgp6nL7A0+K/B7s12mcCY/5l5SIUpMOl+dC0=",
        version = "v0.2.1",
    )

    go_repository(
        name = "com_github_gizak_termui_v3",
        importpath = "github.com/gizak/termui/v3",
        sum = "h1:ZZmVDgwHl7gR7elfKf1xc4IudXZ5qqfDh4wExk4Iajc=",
        version = "v3.1.0",
    )
    go_repository(
        name = "com_github_go_logr_logr",
        importpath = "github.com/go-logr/logr",
        sum = "h1:g01GSCwiDw2xSZfjJ2/T9M+S6pFdcNtFYsp+Y43HYDQ=",
        version = "v1.2.4",
    )
    go_repository(
        name = "com_github_go_redis_redis",
        importpath = "github.com/go-redis/redis",
        sum = "h1:K0pv1D7EQUjfyoMql+r/jZqCLizCGKFlFgcHWWmHQjg=",
        version = "v6.15.9+incompatible",
    )
    go_repository(
        name = "com_github_go_task_slim_sprig",
        importpath = "github.com/go-task/slim-sprig",
        sum = "h1:tfuBGBXKqDEevZMzYi5KSi8KkcZtzBcTgAUUtapy0OI=",
        version = "v0.0.0-20230315185526-52ccab3ef572",
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
        version = "v1.3.2",
    )

    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:DVjP2PbBOzHyzA+dn3WhHIq4NdVu3Q+pvivFICf/7fo=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_golang_groupcache",
        importpath = "github.com/golang/groupcache",
        sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
        version = "v0.0.0-20210331224755-41bb18bfe9da",
    )
    go_repository(
        name = "com_github_golang_mock",
        importpath = "github.com/golang/mock",
        sum = "h1:G5FRp8JnTd7RQH5kemVNlMeyXQAztQ3mOWV95KxsXH8=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:KhyjKVUg7Usr/dYsdSqoFveMYd5ko72D+zANwlG1mmg=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_google_pprof",
        importpath = "github.com/google/pprof",
        sum = "h1:yAJXTCF9TqKcTiHJAE8dj7HMvPfh66eeA2JYW7eFpSE=",
        version = "v0.0.0-20210407192527-94a9f03dee38",
    )

    go_repository(
        name = "com_github_google_s2a_go",
        importpath = "github.com/google/s2a-go",
        sum = "h1:60BLSyTrOV4/haCDW4zb1guZItoSq8foHCXrAnjBo/o=",
        version = "v0.1.7",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:KjJaJ9iWZ3jOFZIf1Lqf4laDRCasjl0BCmnEGxkdLb4=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_googleapis_enterprise_certificate_proxy",
        importpath = "github.com/googleapis/enterprise-certificate-proxy",
        sum = "h1:Vie5ybvEvT75RniqhfFxPRy3Bf7vr3h0cechB90XaQs=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_googleapis_gax_go_v2",
        importpath = "github.com/googleapis/gax-go/v2",
        sum = "h1:A+gCJKdRfqXkr+BIRGtZLibNXf0m1f9E4HG56etFpas=",
        version = "v2.12.0",
    )
    go_repository(
        name = "com_github_googleapis_google_cloud_go_testing",
        importpath = "github.com/googleapis/google-cloud-go-testing",
        sum = "h1:zC34cGQu69FG7qzJ3WiKW244WfhDC3xxYMeNOX2gtUQ=",
        version = "v0.0.0-20210719221736-1c9a4c676720",
    )
    go_repository(
        name = "com_github_hashicorp_consul_api",
        importpath = "github.com/hashicorp/consul/api",
        sum = "h1:CqrdhYzc8XZuPnhIYZWH45toM0LB9ZeYr/gvpLVI3PE=",
        version = "v1.25.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_cleanhttp",
        importpath = "github.com/hashicorp/go-cleanhttp",
        sum = "h1:035FKYIWjmULyFRBKPs8TBQoi0x6d9G4xc9neXJWAZQ=",
        version = "v0.5.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_hclog",
        importpath = "github.com/hashicorp/go-hclog",
        sum = "h1:bI2ocEMgcVlz55Oj1xZNBsVi900c7II+fWDyV9o+13c=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_immutable_radix",
        importpath = "github.com/hashicorp/go-immutable-radix",
        sum = "h1:DKHmCUm2hRBK510BaiZlwvpD40f8bJFeZnpfm2KLowc=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_rootcerts",
        importpath = "github.com/hashicorp/go-rootcerts",
        sum = "h1:jzhAVGtqPKbwpyCPELlgNWhE1znq+qwJtW5Oi2viEzc=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru",
        importpath = "github.com/hashicorp/golang-lru",
        sum = "h1:YDjusn29QI/Das2iO9M0BHnIbxPeyuCHsjMW+lJfyTc=",
        version = "v0.5.4",
    )

    go_repository(
        name = "com_github_hashicorp_golang_lru_v2",
        importpath = "github.com/hashicorp/golang-lru/v2",
        sum = "h1:a+bsQ5rvGLjzHuww6tVxozPZFVghXaHOwFs4luLUK2k=",
        version = "v2.0.7",
    )
    go_repository(
        name = "com_github_hashicorp_hcl",
        importpath = "github.com/hashicorp/hcl",
        sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_serf",
        importpath = "github.com/hashicorp/serf",
        sum = "h1:Z1H2J60yRKvfDYAOZLd2MU0ND4AH/WDz7xYHDWQsIPY=",
        version = "v0.10.1",
    )
    go_repository(
        name = "com_github_hpcloud_tail",
        importpath = "github.com/hpcloud/tail",
        sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_json_iterator_go",
        importpath = "github.com/json-iterator/go",
        sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
        version = "v1.1.12",
    )
    go_repository(
        name = "com_github_klauspost_compress",
        importpath = "github.com/klauspost/compress",
        sum = "h1:Rnbp4K9EjcDuVuHtd0dgA4qNuv9yKDYKK1ulpJwgrqM=",
        version = "v1.17.0",
    )
    go_repository(
        name = "com_github_kr_fs",
        importpath = "github.com/kr/fs",
        sum = "h1:Jskdu9ieNAYnjxsi0LbQp1ulIKZV1LAFgK1tWhpZgl8=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_magiconair_properties",
        importpath = "github.com/magiconair/properties",
        sum = "h1:IeQXZAiQcpL9mgcAe1Nu6cX9LLw6ExEHKjN0VQdvPDY=",
        version = "v1.8.7",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=",
        version = "v0.1.13",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:BTarxUcIeDqL27Mc+vyvdWYSL28zpIhv3RoTdsLMPng=",
        version = "v0.0.17",
    )

    go_repository(
        name = "com_github_mattn_go_runewidth",
        importpath = "github.com/mattn/go-runewidth",
        sum = "h1:Lm995f3rfxdpd6TSmuVCHVb/QhupuXlYr8sCI/QdE+0=",
        version = "v0.0.9",
    )
    go_repository(
        name = "com_github_mitchellh_go_homedir",
        importpath = "github.com/mitchellh/go-homedir",
        sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
        version = "v1.1.0",
    )

    go_repository(
        name = "com_github_mitchellh_go_wordwrap",
        importpath = "github.com/mitchellh/go-wordwrap",
        sum = "h1:DpOJ2HYzCv8LZP15IdmG+YdwD2luVPHITV96TkirNBM=",
        version = "v0.0.0-20150314170334-ad45545899c7",
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        importpath = "github.com/mitchellh/mapstructure",
        sum = "h1:jeMsZIYE/09sWLaz43PL7Gy6RuMjD2eJVyuac5Z2hdY=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_modern_go_concurrent",
        importpath = "github.com/modern-go/concurrent",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        importpath = "github.com/modern-go/reflect2",
        sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_nats_io_nats_go",
        importpath = "github.com/nats-io/nats.go",
        sum = "h1:/WFBHEc/dOKBF6qf1TZhrdEfTmOZ5JzdJ+Y3m6Y/p7E=",
        version = "v1.31.0",
    )
    go_repository(
        name = "com_github_nats_io_nkeys",
        importpath = "github.com/nats-io/nkeys",
        sum = "h1:IzVe95ru2CT6ta874rt9saQRkWfe2nFj1NtvYSLqMzY=",
        version = "v0.4.6",
    )
    go_repository(
        name = "com_github_nats_io_nuid",
        importpath = "github.com/nats-io/nuid",
        sum = "h1:5iA8DT8V7q8WK2EScv2padNa/rTESc1KdnPw4TC2paw=",
        version = "v1.0.1",
    )

    go_repository(
        name = "com_github_nsf_termbox_go",
        importpath = "github.com/nsf/termbox-go",
        sum = "h1:nksUPLCb73Q++DwbYUBEglYBRPZyoXJdrj5L+TkjyZY=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_nxadm_tail",
        importpath = "github.com/nxadm/tail",
        sum = "h1:nPr65rt6Y5JFSKQO7qToXr7pePgD6Gwiw05lkbyAQTE=",
        version = "v1.4.8",
    )
    go_repository(
        name = "com_github_onsi_ginkgo",
        importpath = "github.com/onsi/ginkgo",
        sum = "h1:8xi0RTUf59SOSfEtZMvwTvXYMzG4gV23XVHOZiXNtnE=",
        version = "v1.16.5",
    )
    go_repository(
        name = "com_github_onsi_ginkgo_v2",
        importpath = "github.com/onsi/ginkgo/v2",
        sum = "h1:0jY9lJquiL8fcf3M4LAXN5aMlS/b2BV86HFFPCPMgE4=",
        version = "v2.13.0",
    )
    go_repository(
        name = "com_github_onsi_gomega",
        importpath = "github.com/onsi/gomega",
        sum = "h1:hvMK7xYz4D3HapigLTeGdId/NcfQx1VHMJc60ew99+8=",
        version = "v1.30.0",
    )
    go_repository(
        name = "com_github_pelletier_go_toml_v2",
        importpath = "github.com/pelletier/go-toml/v2",
        sum = "h1:FnwAJ4oYMvbT/34k9zzHuZNrhlz48GB3/s6at6/MHO4=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_pkg_sftp",
        importpath = "github.com/pkg/sftp",
        sum = "h1:JFZT4XbOU7l77xGSpOdW+pwIMqP044IyjXX6FGyEKFo=",
        version = "v1.13.6",
    )

    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:gQz4mCbXsO+nc9n1hCxHcGA3Zx3Eo+UHZoInFGUIXNM=",
        version = "v0.0.0-20190812154241-14fe0d1b01d4",
    )
    go_repository(
        name = "com_github_redis_go_redis_v9",
        importpath = "github.com/redis/go-redis/v9",
        sum = "h1:Yzoz33UZw9I/mFhx4MNrB6Fk+XHO1VukNcCa1+lwyKk=",
        version = "v9.4.0",
    )
    go_repository(
        name = "com_github_rogpeppe_go_internal",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:73kH8U+JUqXU8lRuOHeVHaa/SZPifC7BkcraZVejAe8=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_sagikazarmark_crypt",
        importpath = "github.com/sagikazarmark/crypt",
        sum = "h1:ZA/7pXyjkHoK4bW4mIdnCLvL8hd+Nrbiw7Dqk7D4qUk=",
        version = "v0.17.0",
    )
    go_repository(
        name = "com_github_sagikazarmark_locafero",
        importpath = "github.com/sagikazarmark/locafero",
        sum = "h1:HApY1R9zGo4DBgr7dqsTH/JJxLTTsOt7u6keLGt6kNQ=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_sagikazarmark_slog_shim",
        importpath = "github.com/sagikazarmark/slog-shim",
        sum = "h1:diDBnUNK9N/354PgrxMywXnAwEr1QZcOr6gto+ugjYE=",
        version = "v0.1.0",
    )

    go_repository(
        name = "com_github_sirupsen_logrus",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_github_sourcegraph_conc",
        importpath = "github.com/sourcegraph/conc",
        sum = "h1:OQTbbt6P72L20UqAkXXuLOj79LfEanQ+YQFNpLA9ySo=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_spf13_afero",
        importpath = "github.com/spf13/afero",
        sum = "h1:WJQKhtpdm3v2IzqG8VMqrr6Rf3UYpEF239Jy9wNepM8=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_github_spf13_cast",
        importpath = "github.com/spf13/cast",
        sum = "h1:GEiTHELF+vaR5dhz3VqZfFSzZjYbgeKDpBxQVS4GYJ0=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        importpath = "github.com/spf13/pflag",
        sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_spf13_viper",
        importpath = "github.com/spf13/viper",
        sum = "h1:V7OZpY8i3C1x/pDmU0zNNlfVoDz112fSYvtWMjjS3f4=",
        version = "v1.1.0",
    )

    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:4G4v2dO3VZwixGIRoQ5Lfboy6nUhCyYzaqnIAPPhYs4=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:nOGnQDM7FYENwehXlg/kFVnos3rEvtKTjRvOWSzb6H4=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_subosito_gotenv",
        importpath = "github.com/subosito/gotenv",
        sum = "h1:9NlTDc1FTs4qu0DDq7AEtTPNw6SVm7uBMsUCUjABIf8=",
        version = "v1.6.0",
    )

    go_repository(
        name = "com_github_yuin_goldmark",
        importpath = "github.com/yuin/goldmark",
        sum = "h1:ruQGxdhGHe7FWOJPT0mKs5+pD2Xs1Bm/kdGlHO04FmM=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_google_cloud_go",
        importpath = "cloud.google.com/go",
        sum = "h1:YHLKNupSD1KqjDbQ3+LVdQ81h/UJbJyZG203cEfnQgM=",
        version = "v0.111.0",
    )
    go_repository(
        name = "com_google_cloud_go_accessapproval",
        importpath = "cloud.google.com/go/accessapproval",
        sum = "h1:ZvLvJ952zK8pFHINjpMBY5k7LTAp/6pBf50RDMRgBUI=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_accesscontextmanager",
        importpath = "cloud.google.com/go/accesscontextmanager",
        sum = "h1:Yo4g2XrBETBCqyWIibN3NHNPQKUfQqti0lI+70rubeE=",
        version = "v1.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_aiplatform",
        importpath = "cloud.google.com/go/aiplatform",
        sum = "h1:xyCAfpI4yUMOQ4VtHN/bdmxPQ8xoEkTwFM1nbVmuQhs=",
        version = "v1.58.0",
    )
    go_repository(
        name = "com_google_cloud_go_analytics",
        importpath = "cloud.google.com/go/analytics",
        sum = "h1:fnV7B8lqyEYxCU0LKk+vUL7mTlqRAq4uFlIthIdr/iA=",
        version = "v0.21.6",
    )
    go_repository(
        name = "com_google_cloud_go_apigateway",
        importpath = "cloud.google.com/go/apigateway",
        sum = "h1:VVIxCtVerchHienSlaGzV6XJGtEM9828Erzyr3miUGs=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeconnect",
        importpath = "cloud.google.com/go/apigeeconnect",
        sum = "h1:jSoGITWKgAj/ssVogNE9SdsTqcXnryPzsulENSRlusI=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeregistry",
        importpath = "cloud.google.com/go/apigeeregistry",
        sum = "h1:DSaD1iiqvELag+lV4VnnqUUFd8GXELu01tKVdWZrviE=",
        version = "v0.8.2",
    )
    go_repository(
        name = "com_google_cloud_go_appengine",
        importpath = "cloud.google.com/go/appengine",
        sum = "h1:Qub3fqR7iA1daJWdzjp/Q0Jz0fUG0JbMc7Ui4E9IX/E=",
        version = "v1.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_area120",
        importpath = "cloud.google.com/go/area120",
        sum = "h1:YnSO8m02pOIo6AEOgiOoUDVbw4pf+bg2KLHi4rky320=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_artifactregistry",
        importpath = "cloud.google.com/go/artifactregistry",
        sum = "h1:/hQaadYytMdA5zBh+RciIrXZQBWK4vN7EUsrQHG+/t8=",
        version = "v1.14.6",
    )
    go_repository(
        name = "com_google_cloud_go_asset",
        importpath = "cloud.google.com/go/asset",
        sum = "h1:VjwWNtEVsbpXfJqZbb2RLOBzSgAjN69vf2UJADHnkxk=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_google_cloud_go_assuredworkloads",
        importpath = "cloud.google.com/go/assuredworkloads",
        sum = "h1:FsLSkmYYeNuzDm8L4YPfLWV+lQaUrJmH5OuD37t1k20=",
        version = "v1.11.4",
    )
    go_repository(
        name = "com_google_cloud_go_automl",
        importpath = "cloud.google.com/go/automl",
        sum = "h1:i9tOKXX+1gE7+rHpWKjiuPfGBVIYoWvLNIGpWgPtF58=",
        version = "v1.13.4",
    )
    go_repository(
        name = "com_google_cloud_go_baremetalsolution",
        importpath = "cloud.google.com/go/baremetalsolution",
        sum = "h1:oQiFYYCe0vwp7J8ZmF6siVKEumWtiPFJMJcGuyDVRUk=",
        version = "v1.2.3",
    )
    go_repository(
        name = "com_google_cloud_go_batch",
        importpath = "cloud.google.com/go/batch",
        sum = "h1:AxuSPoL2fWn/rUyvWeNCNd0V2WCr+iHRCU9QO1PUmpY=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_beyondcorp",
        importpath = "cloud.google.com/go/beyondcorp",
        sum = "h1:VXf9SnrnSmj2BF2cHkoTHvOUp8gjsz1KJFOMW7czdsY=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_google_cloud_go_bigquery",
        importpath = "cloud.google.com/go/bigquery",
        sum = "h1:FiULdbbzUxWD0Y4ZGPSVCDLvqRSyCIO6zKV7E2nf5uA=",
        version = "v1.57.1",
    )
    go_repository(
        name = "com_google_cloud_go_billing",
        importpath = "cloud.google.com/go/billing",
        sum = "h1:GvKy4xLy1zF1XPbwP5NJb2HjRxhnhxjjXxvyZ1S/IAo=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_google_cloud_go_binaryauthorization",
        importpath = "cloud.google.com/go/binaryauthorization",
        sum = "h1:PHS89lcFayWIEe0/s2jTBiEOtqghCxzc7y7bRNlifBs=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_certificatemanager",
        importpath = "cloud.google.com/go/certificatemanager",
        sum = "h1:5YMQ3Q+dqGpwUZ9X5sipsOQ1fLPsxod9HNq0+nrqc6I=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_channel",
        importpath = "cloud.google.com/go/channel",
        sum = "h1:Rd4+fBrjiN6tZ4TR8R/38elkyEkz6oogGDr7jDyjmMY=",
        version = "v1.17.3",
    )
    go_repository(
        name = "com_google_cloud_go_cloudbuild",
        importpath = "cloud.google.com/go/cloudbuild",
        sum = "h1:9IHfEMWdCklJ1cwouoiQrnxmP0q3pH7JUt8Hqx4Qbck=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_google_cloud_go_clouddms",
        importpath = "cloud.google.com/go/clouddms",
        sum = "h1:xe/wJKz55VO1+L891a1EG9lVUgfHr9Ju/I3xh1nwF84=",
        version = "v1.7.3",
    )
    go_repository(
        name = "com_google_cloud_go_cloudtasks",
        importpath = "cloud.google.com/go/cloudtasks",
        sum = "h1:5xXuFfAjg0Z5Wb81j2GAbB3e0bwroCeSF+5jBn/L650=",
        version = "v1.12.4",
    )
    go_repository(
        name = "com_google_cloud_go_compute",
        importpath = "cloud.google.com/go/compute",
        sum = "h1:6sVlXXBmbd7jNX0Ipq0trII3e4n1/MsADLK6a+aiVlk=",
        version = "v1.23.3",
    )
    go_repository(
        name = "com_google_cloud_go_compute_metadata",
        importpath = "cloud.google.com/go/compute/metadata",
        sum = "h1:mg4jlk7mCAj6xXp9UJ4fjI9VUI5rubuGBW5aJ7UnBMY=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_google_cloud_go_contactcenterinsights",
        importpath = "cloud.google.com/go/contactcenterinsights",
        sum = "h1:EiGBeejtDDtr3JXt9W7xlhXyZ+REB5k2tBgVPVtmNb0=",
        version = "v1.12.1",
    )
    go_repository(
        name = "com_google_cloud_go_container",
        importpath = "cloud.google.com/go/container",
        sum = "h1:jIltU529R2zBFvP8rhiG1mgeTcnT27KhU0H/1d6SQRg=",
        version = "v1.29.0",
    )
    go_repository(
        name = "com_google_cloud_go_containeranalysis",
        importpath = "cloud.google.com/go/containeranalysis",
        sum = "h1:5rhYLX+3a01drpREqBZVXR9YmWH45RnML++8NsCtuD8=",
        version = "v0.11.3",
    )
    go_repository(
        name = "com_google_cloud_go_datacatalog",
        importpath = "cloud.google.com/go/datacatalog",
        sum = "h1:rbYNmHwvAOOwnW2FPXYkaK3Mf1MmGqRzK0mMiIEyLdo=",
        version = "v1.19.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataflow",
        importpath = "cloud.google.com/go/dataflow",
        sum = "h1:7VmCNWcPJBS/srN2QnStTB6nu4Eb5TMcpkmtaPVhRt4=",
        version = "v0.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_dataform",
        importpath = "cloud.google.com/go/dataform",
        sum = "h1:jV+EsDamGX6cE127+QAcCR/lergVeeZdEQ6DdrxW3sQ=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_google_cloud_go_datafusion",
        importpath = "cloud.google.com/go/datafusion",
        sum = "h1:Q90alBEYlMi66zL5gMSGQHfbZLB55mOAg03DhwTTfsk=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_datalabeling",
        importpath = "cloud.google.com/go/datalabeling",
        sum = "h1:zrq4uMmunf2KFDl/7dS6iCDBBAxBnKVDyw6+ajz3yu0=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_dataplex",
        importpath = "cloud.google.com/go/dataplex",
        sum = "h1:ACVOuxwe7gP0SqEso9SLyXbcZNk5l8hjcTX+XLntI5s=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataproc_v2",
        importpath = "cloud.google.com/go/dataproc/v2",
        sum = "h1:tTVP9tTxmc8fixxOd/8s6Q6Pz/+yzn7r7XdZHretQH0=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataqna",
        importpath = "cloud.google.com/go/dataqna",
        sum = "h1:NJnu1kAPamZDs/if3bJ3+Wb6tjADHKL83NUWsaIp2zg=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_datastore",
        importpath = "cloud.google.com/go/datastore",
        sum = "h1:0P9WcsQeTWjuD1H14JIY7XQscIPQ4Laje8ti96IC5vg=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_google_cloud_go_datastream",
        importpath = "cloud.google.com/go/datastream",
        sum = "h1:Z2sKPIB7bT2kMW5Uhxy44ZgdJzxzE5uKjavoW+EuHEE=",
        version = "v1.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_deploy",
        importpath = "cloud.google.com/go/deploy",
        sum = "h1:5OVjzm8MPC5kP+Ywbs0mdE0O7AXvAUXksSyHAyMFyMg=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_google_cloud_go_dialogflow",
        importpath = "cloud.google.com/go/dialogflow",
        sum = "h1:tLCWad8HZhlyUNfDzDP5m+oH6h/1Uvw/ei7B9AnsWMk=",
        version = "v1.47.0",
    )
    go_repository(
        name = "com_google_cloud_go_dlp",
        importpath = "cloud.google.com/go/dlp",
        sum = "h1:OFlXedmPP/5//X1hBEeq3D9kUVm9fb6ywYANlpv/EsQ=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_google_cloud_go_documentai",
        importpath = "cloud.google.com/go/documentai",
        sum = "h1:hlYieOXUwiJ7HpBR/vEPfr8nfSxveLVzbqbUkSK0c/4=",
        version = "v1.23.7",
    )
    go_repository(
        name = "com_google_cloud_go_domains",
        importpath = "cloud.google.com/go/domains",
        sum = "h1:ua4GvsDztZ5F3xqjeLKVRDeOvJshf5QFgWGg1CKti3A=",
        version = "v0.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_edgecontainer",
        importpath = "cloud.google.com/go/edgecontainer",
        sum = "h1:Szy3Q/N6bqgQGyxqjI+6xJZbmvPvnFHp3UZr95DKcQ0=",
        version = "v1.1.4",
    )
    go_repository(
        name = "com_google_cloud_go_errorreporting",
        importpath = "cloud.google.com/go/errorreporting",
        sum = "h1:kj1XEWMu8P0qlLhm3FwcaFsUvXChV/OraZwA70trRR0=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_essentialcontacts",
        importpath = "cloud.google.com/go/essentialcontacts",
        sum = "h1:S2if6wkjR4JCEAfDtIiYtD+sTz/oXjh2NUG4cgT1y/Q=",
        version = "v1.6.5",
    )
    go_repository(
        name = "com_google_cloud_go_eventarc",
        importpath = "cloud.google.com/go/eventarc",
        sum = "h1:+pFmO4eu4dOVipSaFBLkmqrRYG94Xl/TQZFOeohkuqU=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_google_cloud_go_filestore",
        importpath = "cloud.google.com/go/filestore",
        sum = "h1:/+wUEGwk3x3Kxomi2cP5dsR8+SIXxo7M0THDjreFSYo=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_firestore",
        importpath = "cloud.google.com/go/firestore",
        sum = "h1:8aLcKnMPoldYU3YHgu4t2exrKhLQkqaXAGqT0ljrFVw=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_google_cloud_go_functions",
        importpath = "cloud.google.com/go/functions",
        sum = "h1:ZjdiV3MyumRM6++1Ixu6N0VV9LAGlCX4AhW6Yjr1t+U=",
        version = "v1.15.4",
    )
    go_repository(
        name = "com_google_cloud_go_gkebackup",
        importpath = "cloud.google.com/go/gkebackup",
        sum = "h1:KhnOrr9A1tXYIYeXKqCKbCI8TL2ZNGiD3dm+d7BDUBg=",
        version = "v1.3.4",
    )
    go_repository(
        name = "com_google_cloud_go_gkeconnect",
        importpath = "cloud.google.com/go/gkeconnect",
        sum = "h1:1JLpZl31YhQDQeJ98tK6QiwTpgHFYRJwpntggpQQWis=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_gkehub",
        importpath = "cloud.google.com/go/gkehub",
        sum = "h1:J5tYUtb3r0cl2mM7+YHvV32eL+uZQ7lONyUZnPikCEo=",
        version = "v0.14.4",
    )
    go_repository(
        name = "com_google_cloud_go_gkemulticloud",
        importpath = "cloud.google.com/go/gkemulticloud",
        sum = "h1:NmJsNX9uQ2CT78957xnjXZb26TDIMvv+d5W2vVUt0Pg=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_google_cloud_go_gsuiteaddons",
        importpath = "cloud.google.com/go/gsuiteaddons",
        sum = "h1:uuw2Xd37yHftViSI8J2hUcCS8S7SH3ZWH09sUDLW30Q=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_iam",
        importpath = "cloud.google.com/go/iam",
        sum = "h1:1jTsCu4bcsNsE4iiqNT5SHwrDRCfRmIaaaVFhRveTJI=",
        version = "v1.1.5",
    )
    go_repository(
        name = "com_google_cloud_go_iap",
        importpath = "cloud.google.com/go/iap",
        sum = "h1:M4vDbQ4TLXdaljXVZSwW7XtxpwXUUarY2lIs66m0aCM=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_ids",
        importpath = "cloud.google.com/go/ids",
        sum = "h1:VuFqv2ctf/A7AyKlNxVvlHTzjrEvumWaZflUzBPz/M4=",
        version = "v1.4.4",
    )
    go_repository(
        name = "com_google_cloud_go_iot",
        importpath = "cloud.google.com/go/iot",
        sum = "h1:m1WljtkZnvLTIRYW1YTOv5A6H1yKgLHR6nU7O8yf27w=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_kms",
        importpath = "cloud.google.com/go/kms",
        sum = "h1:pj1sRfut2eRbD9pFRjNnPNg/CzJPuQAzUujMIM1vVeM=",
        version = "v1.15.5",
    )
    go_repository(
        name = "com_google_cloud_go_language",
        importpath = "cloud.google.com/go/language",
        sum = "h1:zg9uq2yS9PGIOdc0Kz/l+zMtOlxKWonZjjo5w5YPG2A=",
        version = "v1.12.2",
    )
    go_repository(
        name = "com_google_cloud_go_lifesciences",
        importpath = "cloud.google.com/go/lifesciences",
        sum = "h1:rZEI/UxcxVKEzyoRS/kdJ1VoolNItRWjNN0Uk9tfexg=",
        version = "v0.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_logging",
        importpath = "cloud.google.com/go/logging",
        sum = "h1:iEIOXFO9EmSiTjDmfpbRjOxECO7R8C7b8IXUGOj7xZw=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_longrunning",
        importpath = "cloud.google.com/go/longrunning",
        sum = "h1:w8xEcbZodnA2BbW6sVirkkoC+1gP8wS57EUUgGS0GVg=",
        version = "v0.5.4",
    )
    go_repository(
        name = "com_google_cloud_go_managedidentities",
        importpath = "cloud.google.com/go/managedidentities",
        sum = "h1:SF/u1IJduMqQQdJA4MDyivlIQ4SrV5qAawkr/ZEREkY=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_maps",
        importpath = "cloud.google.com/go/maps",
        sum = "h1:WxxLo//b60nNFESefLgaBQevu8QGUmRV3+noOjCfIHs=",
        version = "v1.6.2",
    )
    go_repository(
        name = "com_google_cloud_go_mediatranslation",
        importpath = "cloud.google.com/go/mediatranslation",
        sum = "h1:VRCQfZB4s6jN0CSy7+cO3m4ewNwgVnaePanVCQh/9Z4=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_memcache",
        importpath = "cloud.google.com/go/memcache",
        sum = "h1:cdex/ayDd294XBj2cGeMe6Y+H1JvhN8y78B9UW7pxuQ=",
        version = "v1.10.4",
    )
    go_repository(
        name = "com_google_cloud_go_metastore",
        importpath = "cloud.google.com/go/metastore",
        sum = "h1:94l/Yxg9oBZjin2bzI79oK05feYefieDq0o5fjLSkC8=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_google_cloud_go_monitoring",
        importpath = "cloud.google.com/go/monitoring",
        sum = "h1:blrdvF0MkPPivSO041ihul7rFMhXdVp8Uq7F59DKXTU=",
        version = "v1.17.0",
    )
    go_repository(
        name = "com_google_cloud_go_networkconnectivity",
        importpath = "cloud.google.com/go/networkconnectivity",
        sum = "h1:e9lUkCe2BexsqsUc2bjV8+gFBpQa54J+/F3qKVtW+wA=",
        version = "v1.14.3",
    )
    go_repository(
        name = "com_google_cloud_go_networkmanagement",
        importpath = "cloud.google.com/go/networkmanagement",
        sum = "h1:HsQk4FNKJUX04k3OI6gUsoveiHMGvDRqlaFM2xGyvqU=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_networksecurity",
        importpath = "cloud.google.com/go/networksecurity",
        sum = "h1:947tNIPnj1bMGTIEBo3fc4QrrFKS5hh0bFVsHmFm4Vo=",
        version = "v0.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_notebooks",
        importpath = "cloud.google.com/go/notebooks",
        sum = "h1:eTOTfNL1yM6L/PCtquJwjWg7ZZGR0URFaFgbs8kllbM=",
        version = "v1.11.2",
    )
    go_repository(
        name = "com_google_cloud_go_optimization",
        importpath = "cloud.google.com/go/optimization",
        sum = "h1:iFsoexcp13cGT3k/Hv8PA5aK+FP7FnbhwDO9llnruas=",
        version = "v1.6.2",
    )
    go_repository(
        name = "com_google_cloud_go_orchestration",
        importpath = "cloud.google.com/go/orchestration",
        sum = "h1:kgwZ2f6qMMYIVBtUGGoU8yjYWwMTHDanLwM/CQCFaoQ=",
        version = "v1.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_orgpolicy",
        importpath = "cloud.google.com/go/orgpolicy",
        sum = "h1:RWuXQDr9GDYhjmrredQJC7aY7cbyqP9ZuLbq5GJGves=",
        version = "v1.11.4",
    )
    go_repository(
        name = "com_google_cloud_go_osconfig",
        importpath = "cloud.google.com/go/osconfig",
        sum = "h1:OrRCIYEAbrbXdhm13/JINn9pQchvTTIzgmOCA7uJw8I=",
        version = "v1.12.4",
    )
    go_repository(
        name = "com_google_cloud_go_oslogin",
        importpath = "cloud.google.com/go/oslogin",
        sum = "h1:NP/KgsD9+0r9hmHC5wKye0vJXVwdciv219DtYKYjgqE=",
        version = "v1.12.2",
    )
    go_repository(
        name = "com_google_cloud_go_phishingprotection",
        importpath = "cloud.google.com/go/phishingprotection",
        sum = "h1:sPLUQkHq6b4AL0czSJZ0jd6vL55GSTHz2B3Md+TCZI0=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_policytroubleshooter",
        importpath = "cloud.google.com/go/policytroubleshooter",
        sum = "h1:sq+ScLP83d7GJy9+wpwYJVnY+q6xNTXwOdRIuYjvHT4=",
        version = "v1.10.2",
    )
    go_repository(
        name = "com_google_cloud_go_privatecatalog",
        importpath = "cloud.google.com/go/privatecatalog",
        sum = "h1:Vo10IpWKbNvc/z/QZPVXgCiwfjpWoZ/wbgful4Uh/4E=",
        version = "v0.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_pubsub",
        importpath = "cloud.google.com/go/pubsub",
        sum = "h1:6SPCPvWav64tj0sVX/+npCBKhUi/UjJehy9op/V3p2g=",
        version = "v1.33.0",
    )
    go_repository(
        name = "com_google_cloud_go_pubsublite",
        importpath = "cloud.google.com/go/pubsublite",
        sum = "h1:pX+idpWMIH30/K7c0epN6V703xpIcMXWRjKJsz0tYGY=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_recaptchaenterprise_v2",
        importpath = "cloud.google.com/go/recaptchaenterprise/v2",
        sum = "h1:Zrd4LvT9PaW91X/Z13H0i5RKEv9suCLuk8zp+bfOpN4=",
        version = "v2.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_recommendationengine",
        importpath = "cloud.google.com/go/recommendationengine",
        sum = "h1:JRiwe4hvu3auuh2hujiTc2qNgPPfVp+Q8KOpsXlEzKQ=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_recommender",
        importpath = "cloud.google.com/go/recommender",
        sum = "h1:tC+ljmCCbuZ/ybt43odTFlay91n/HLIhflvaOeb0Dh4=",
        version = "v1.12.0",
    )
    go_repository(
        name = "com_google_cloud_go_redis",
        importpath = "cloud.google.com/go/redis",
        sum = "h1:J9cEHxG9YLmA9o4jTSvWt/RuVEn6MTrPlYSCRHujxDQ=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_google_cloud_go_resourcemanager",
        importpath = "cloud.google.com/go/resourcemanager",
        sum = "h1:JwZ7Ggle54XQ/FVYSBrMLOQIKoIT/uer8mmNvNLK51k=",
        version = "v1.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_resourcesettings",
        importpath = "cloud.google.com/go/resourcesettings",
        sum = "h1:yTIL2CsZswmMfFyx2Ic77oLVzfBFoWBYgpkgiSPnC4Y=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_retail",
        importpath = "cloud.google.com/go/retail",
        sum = "h1:geqdX1FNqqL2p0ADXjPpw8lq986iv5GrVcieTYafuJQ=",
        version = "v1.14.4",
    )
    go_repository(
        name = "com_google_cloud_go_run",
        importpath = "cloud.google.com/go/run",
        sum = "h1:qdfZteAm+vgzN1iXzILo3nJFQbzziudkJrvd9wCf3FQ=",
        version = "v1.3.3",
    )
    go_repository(
        name = "com_google_cloud_go_scheduler",
        importpath = "cloud.google.com/go/scheduler",
        sum = "h1:eMEettHlFhG5pXsoHouIM5nRT+k+zU4+GUvRtnxhuVI=",
        version = "v1.10.5",
    )
    go_repository(
        name = "com_google_cloud_go_secretmanager",
        importpath = "cloud.google.com/go/secretmanager",
        sum = "h1:krnX9qpG2kR2fJ+u+uNyNo+ACVhplIAS4Pu7u+4gd+k=",
        version = "v1.11.4",
    )
    go_repository(
        name = "com_google_cloud_go_security",
        importpath = "cloud.google.com/go/security",
        sum = "h1:sdnh4Islb1ljaNhpIXlIPgb3eYj70QWgPVDKOUYvzJc=",
        version = "v1.15.4",
    )
    go_repository(
        name = "com_google_cloud_go_securitycenter",
        importpath = "cloud.google.com/go/securitycenter",
        sum = "h1:crdn2Z2rFIy8WffmmhdlX3CwZJusqCiShtnrGFRwpeE=",
        version = "v1.24.3",
    )
    go_repository(
        name = "com_google_cloud_go_servicedirectory",
        importpath = "cloud.google.com/go/servicedirectory",
        sum = "h1:5niCMfkw+jifmFtbBrtRedbXkJm3fubSR/KHbxSJZVM=",
        version = "v1.11.3",
    )
    go_repository(
        name = "com_google_cloud_go_shell",
        importpath = "cloud.google.com/go/shell",
        sum = "h1:nurhlJcSVFZneoRZgkBEHumTYf/kFJptCK2eBUq/88M=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_spanner",
        importpath = "cloud.google.com/go/spanner",
        sum = "h1:ttU+lhARPF/iZE3OkCpmfsemCz9mLaqBhGPd3Qub2sQ=",
        version = "v1.54.0",
    )
    go_repository(
        name = "com_google_cloud_go_speech",
        importpath = "cloud.google.com/go/speech",
        sum = "h1:qkxNao58oF8ghAHE1Eghen7XepawYEN5zuZXYWaUTA4=",
        version = "v1.21.0",
    )
    go_repository(
        name = "com_google_cloud_go_storage",
        importpath = "cloud.google.com/go/storage",
        sum = "h1:B59ahL//eDfx2IIKFBeT5Atm9wnNmj3+8xG/W4WB//w=",
        version = "v1.35.1",
    )

    go_repository(
        name = "com_google_cloud_go_storagetransfer",
        importpath = "cloud.google.com/go/storagetransfer",
        sum = "h1:YM1dnj5gLjfL6aDldO2s4GeU8JoAvH1xyIwXre63KmI=",
        version = "v1.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_talent",
        importpath = "cloud.google.com/go/talent",
        sum = "h1:LnRJhhYkODDBoTwf6BeYkiJHFw9k+1mAFNyArwZUZAs=",
        version = "v1.6.5",
    )
    go_repository(
        name = "com_google_cloud_go_texttospeech",
        importpath = "cloud.google.com/go/texttospeech",
        sum = "h1:ahrzTgr7uAbvebuhkBAAVU6kRwVD0HWsmDsvMhtad5Q=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_tpu",
        importpath = "cloud.google.com/go/tpu",
        sum = "h1:XIEH5c0WeYGaVy9H+UueiTaf3NI6XNdB4/v6TFQJxtE=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_trace",
        importpath = "cloud.google.com/go/trace",
        sum = "h1:2qOAuAzNezwW3QN+t41BtkDJOG42HywL73q8x/f6fnM=",
        version = "v1.10.4",
    )
    go_repository(
        name = "com_google_cloud_go_translate",
        importpath = "cloud.google.com/go/translate",
        sum = "h1:t5WXTqlrk8VVJu/i3WrYQACjzYJiff5szARHiyqqPzI=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_video",
        importpath = "cloud.google.com/go/video",
        sum = "h1:Xrpbm2S9UFQ1pZEeJt9Vqm5t2T/z9y/M3rNXhFoo8Is=",
        version = "v1.20.3",
    )
    go_repository(
        name = "com_google_cloud_go_videointelligence",
        importpath = "cloud.google.com/go/videointelligence",
        sum = "h1:YS4j7lY0zxYyneTFXjBJUj2r4CFe/UoIi/PJG0Zt/Rg=",
        version = "v1.11.4",
    )
    go_repository(
        name = "com_google_cloud_go_vision_v2",
        importpath = "cloud.google.com/go/vision/v2",
        sum = "h1:T/ujUghvEaTb+YnFY/jiYwVAkMbIC8EieK0CJo6B4vg=",
        version = "v2.7.5",
    )
    go_repository(
        name = "com_google_cloud_go_vmmigration",
        importpath = "cloud.google.com/go/vmmigration",
        sum = "h1:qPNdab4aGgtaRX+51jCOtJxlJp6P26qua4o1xxUDjpc=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_vmwareengine",
        importpath = "cloud.google.com/go/vmwareengine",
        sum = "h1:WY526PqM6QNmFHSqe2sRfK6gRpzWjmL98UFkql2+JDM=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_google_cloud_go_vpcaccess",
        importpath = "cloud.google.com/go/vpcaccess",
        sum = "h1:zbs3V+9ux45KYq8lxxn/wgXole6SlBHHKKyZhNJoS+8=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_webrisk",
        importpath = "cloud.google.com/go/webrisk",
        sum = "h1:iceR3k0BCRZgf2D/NiKviVMFfuNC9LmeNLtxUFRB/wI=",
        version = "v1.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_websecurityscanner",
        importpath = "cloud.google.com/go/websecurityscanner",
        sum = "h1:5Gp7h5j7jywxLUp6NTpjNPkgZb3ngl0tUSw6ICWvtJQ=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_workflows",
        importpath = "cloud.google.com/go/workflows",
        sum = "h1:qocsqETmLAl34mSa01hKZjcqAvt699gaoFbooGGMvaM=",
        version = "v1.12.3",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
        version = "v0.0.0-20161208181325-20d25e280405",
    )
    go_repository(
        name = "in_gopkg_fsnotify_v1",
        importpath = "gopkg.in/fsnotify.v1",
        sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
        version = "v1.4.7",
    )
    go_repository(
        name = "in_gopkg_ini_v1",
        importpath = "gopkg.in/ini.v1",
        sum = "h1:Dgnx+6+nfE+IfzjUEISNeydPJh9AXNNsWbGP9KzCsOA=",
        version = "v1.67.0",
    )
    go_repository(
        name = "in_gopkg_tomb_v1",
        importpath = "gopkg.in/tomb.v1",
        sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
        version = "v1.0.0-20141024135613-dd632973f1e7",
    )

    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:ZCJp+EgiOT7lHqUV2J862kp8Qj64Jo6az82+3Td9dZw=",
        version = "v2.2.2",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "io_etcd_go_etcd_api_v3",
        importpath = "go.etcd.io/etcd/api/v3",
        sum = "h1:szRajuUUbLyppkhs9K6BRtjY37l66XQQmw7oZRANE4k=",
        version = "v3.5.10",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_pkg_v3",
        importpath = "go.etcd.io/etcd/client/pkg/v3",
        sum = "h1:kfYIdQftBnbAq8pUWFXfpuuxFSKzlmM5cSn76JByiT0=",
        version = "v3.5.10",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v2",
        importpath = "go.etcd.io/etcd/client/v2",
        sum = "h1:MrmRktzv/XF8CvtQt+P6wLUlURaNpSDJHFZhe//2QE4=",
        version = "v2.305.10",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v3",
        importpath = "go.etcd.io/etcd/client/v3",
        sum = "h1:W9TXNZ+oB3MCd/8UjxHTWK5J9Nquw9fQBLJd5ne5/Ao=",
        version = "v3.5.10",
    )

    go_repository(
        name = "io_opencensus_go",
        importpath = "go.opencensus.io",
        sum = "h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=",
        version = "v0.24.0",
    )
    go_repository(
        name = "org_golang_google_api",
        importpath = "google.golang.org/api",
        sum = "h1:b2CqT6kG+zqJIVKRQ3ELJVLN1PwHZ6DJ3dW8yl82rgY=",
        version = "v0.149.0",
    )
    go_repository(
        name = "org_golang_google_appengine",
        importpath = "google.golang.org/appengine",
        sum = "h1:IhEN5q69dyKagZPYMSdIjS2HqprW324FRQZJcGqPAsM=",
        version = "v1.6.8",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:nz5NESFLZbJGPFxDT/HCn+V1mZ8JGNoY4nUpmW/Y2eg=",
        version = "v0.0.0-20240102182953-50ed04b92917",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_api",
        importpath = "google.golang.org/genproto/googleapis/api",
        sum = "h1:rcS6EyEaoCO52hQDupoSfrxI3R6C2Tq741is7X8OvnM=",
        version = "v0.0.0-20240102182953-50ed04b92917",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_bytestream",
        importpath = "google.golang.org/genproto/googleapis/bytestream",
        sum = "h1:kOLFPvCAOJaW878BYjdqPilFjH+Il4ILDAOsfAiEPYQ=",
        version = "v0.0.0-20240108191215-35c7eff3a6b1",
    )

    go_repository(
        name = "org_golang_google_genproto_googleapis_rpc",
        importpath = "google.golang.org/genproto/googleapis/rpc",
        sum = "h1:/jFB8jK5R3Sq3i/lmeZO0cATSzFfZaJq1J2Euan3XKU=",
        version = "v0.0.0-20231212172506-995d672761c0",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:26+wFr+cNqSGFcOXcabYC0lUVJVRa2Sb2ortSK7VrEU=",
        version = "v1.60.1",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        sum = "h1:pPC6BG5ex8PDFnkbrGU3EixyhKcQ2aDuBS36lqK/C7I=",
        version = "v1.32.0",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:mMMrFzRSCF0GvB7Ne27XVtVAaXLrPmgPC7/v0tkwHaY=",
        version = "v0.16.0",
    )
    go_repository(
        name = "org_golang_x_exp",
        importpath = "golang.org/x/exp",
        sum = "h1:c2HOrn5iMezYjSlGPncknSEr/8x5LELb/ilJbXi9DEA=",
        version = "v0.0.0-20190121172915-509febef88a4",
    )
    go_repository(
        name = "org_golang_x_lint",
        importpath = "golang.org/x/lint",
        sum = "h1:2M3HP5CCK1Si9FQhwnzYhXdG6DXeebvUHFpre8QvbyI=",
        version = "v0.0.0-20201208152925-83fdc39ff7b5",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:LUYupSeNrTNCGzR/hVBk2NHZO4hXcVaW1k4Qx7rjPx8=",
        version = "v0.8.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:zTwKpTd2XuCqf8huc7Fo2iSy+4RHPd10s4KzeTnVr1c=",
        version = "v0.19.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:jDDenyj+WgFtmV3zYVoi8aE2BwtXFLWOA67ZfNWftiY=",
        version = "v0.13.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:zxkM55ReGkDlKSM+Fu41A+zmbZuaPVbGMzvvdUPznYQ=",
        version = "v0.4.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:h48lPFYpsTvQJZF4EKyI4aLHaev3CxivZmv7yZig9pc=",
        version = "v0.15.0",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:y/Oo/a/q3IXu26lQgl04j/gjuBDOBlx7X6Om1j2CPW4=",
        version = "v0.15.0",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:ScX5w1eTa3QqT8oi6+ziP7dTV1S2+ALU0bI+0zXKWiQ=",
        version = "v0.14.0",
    )
    go_repository(
        name = "org_golang_x_time",
        importpath = "golang.org/x/time",
        sum = "h1:o7cqy6amK/52YcAKIPlM3a+Fpj35zvRj2TP+e1xFSfk=",
        version = "v0.5.0",
    )

    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        sum = "h1:BOw41kyTf3PuCW1pVQf8+Cyg8pMlkYB1oo9iJ6D/lKM=",
        version = "v0.6.0",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
        version = "v0.0.0-20200804184101-5ec99f83aff1",
    )
    go_repository(
        name = "org_uber_go_atomic",
        importpath = "go.uber.org/atomic",
        sum = "h1:ECmE8Bn/WFTYwEW/bpKD3M8VtR/zQVbavAoalC1PYyE=",
        version = "v1.9.0",
    )
    go_repository(
        name = "org_uber_go_multierr",
        importpath = "go.uber.org/multierr",
        sum = "h1:7fIwc/ZtS0q++VgcfqFDxSBZVv/Xo49/SYnDFupUwlI=",
        version = "v1.9.0",
    )
    go_repository(
        name = "org_uber_go_zap",
        importpath = "go.uber.org/zap",
        sum = "h1:WefMeulhovoZ2sYXz7st6K0sLj7bBhpiFaud4r4zST8=",
        version = "v1.21.0",
    )
