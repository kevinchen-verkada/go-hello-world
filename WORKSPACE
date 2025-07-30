workspace(name = "myws")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "c6cf9da6668ac84c470c43cbfccb8fdc844ead2b5a8b918e2816d44f2986f644",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.55.0/rules_go-v0.55.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.55.0/rules_go-v0.55.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "d76bf7a60fd8b050444090dfa2837a4eaf9829e1165618ee35dceca5cbdf58d5",
    urls = [
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.37.0/bazel-gazelle-v0.37.0.tar.gz",
    ],
)

load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@io_bazel_rules_go//go:deps.bzl", "go_download_sdk", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.24.4")

go_download_sdk(
    name = "go_sdk_linux",
    goarch = "amd64",
    goos = "linux",
    version = "1.24.4",
)

go_download_sdk(
    name = "go_sdk_linux_arm64",
    goarch = "arm64",
    goos = "linux",
    version = "1.24.4",
)

go_download_sdk(
    name = "go_sdk_darwin",
    goarch = "amd64",
    goos = "darwin",
    version = "1.24.4",
)

go_download_sdk(
    name = "go_sdk_darwin_arm64",
    goarch = "arm64",
    goos = "darwin",
    version = "1.24.4",
)

# gazelle:repo bazel_gazelle
gazelle_dependencies(go_sdk = "go_sdk")
