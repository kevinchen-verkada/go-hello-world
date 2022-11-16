The unused dependency 
```
"@com_github_bazelbuild_rules_go//go/tools/bazel:go_default_library",
```
in 
```
:go_lib
```
is used to reproduce the following error
```
ERROR: /private/var/tmp/_bazel_yue.chen/444218253d55c31a8c81bb3d3b5d6919/external/com_github_bazelbuild_rules_go/go/tools/coverdata/BUILD.bazel:3:16: in go_tool_library rule @com_github_bazelbuild_rules_go//go/tools/coverdata:coverdata: 
Traceback (most recent call last):
	File "/private/var/tmp/_bazel_yue.chen/444218253d55c31a8c81bb3d3b5d6919/external/com_github_bazelbuild_rules_go/go/private/rules/library.bzl", line 41, column 20, in _go_library_impl
		go = go_context(ctx)
	File "/private/var/tmp/_bazel_yue.chen/444218253d55c31a8c81bb3d3b5d6919/external/com_github_bazelbuild_rules_go/go/private/context.bzl", line 409, column 67, in go_context
		stdlib = _flatten_possibly_transitioned_attr(attr._stdlib)[GoStdLib]
Error: <target @com_github_bazelbuild_rules_go//:stdlib> (rule 'stdlib') doesn't contain declared provider 'GoStdLib'
ERROR: /private/var/tmp/_bazel_yue.chen/444218253d55c31a8c81bb3d3b5d6919/external/com_github_bazelbuild_rules_go/go/tools/coverdata/BUILD.bazel:3:16: Analysis of target '@com_github_bazelbuild_rules_go//go/tools/coverdata:coverdata' failed
ERROR: Analysis of target '//:go_lib' failed; build aborted: 
```

Tried this on M1 Macbook.
