Build the target with remote execution from an M1 MacBook using
```shell
bazel build --config=remote :all
```
Will get the following error
```
{HOME}/go-hello-world/BUILD.bazel:3:11: While resolving toolchains for target //:go_lib: no matching toolchains found for types @io_bazel_rules_go//go:toolchain
ERROR: Analysis of target '//:go_lib' failed; build aborted: 
```

*note: Please remember to supply your own x-buildbuddy-api-key in .bazelrc
