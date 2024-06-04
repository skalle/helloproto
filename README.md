# Hello world bazel play

Been trying to get a very simple Go application together with a proto to work with [bazel.io](Bazel).

Largely I've been following this [https://github.com/bazelbuild/rules_go/blob/master/docs/go/core/bzlmod.md](tutorial).

## Changes made

The one change made was to change the gazelle generated proto compiler from `@io_bazel_rules_go` -> `@rules_go`.

## Handy commands

- `bazel build //...` <- Builds everything
- `bazel run //cmd/client:client` <- Run the client
- `bazel run //:gazelle` <- Re run gazelle
- `bazel run  @rules_go//go -- mod tidy` <- Run the go command
