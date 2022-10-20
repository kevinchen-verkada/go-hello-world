Build the target with remote execution from an M1 MacBook using
```shell
bazel build --config=remote :all
```
Will get the following error
```
/external/io_bazel_rules_go/BUILD.bazel:86:17: While resolving toolchains for target @io_bazel_rules_go//:cgo_context_data: No matching toolchains found for types @bazel_tools//tools/cpp:toolchain_type. Maybe --incompatible_use_cc_configure_from_rules_cc has been flipped and there is no default C++ toolchain added in the WORKSPACE file? See https://github.com/bazelbuild/bazel/issues/10134 for details and migration instructions.
ERROR: Analysis of target '//:go_lib' failed; build aborted:
```

*note: Please remember to supply your own x-buildbuddy-api-key in .bazelrc
