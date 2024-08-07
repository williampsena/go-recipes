# About

This repository offers valuable test examples that demonstrate how to create "crashable" tests using Panic / Log.Errors, some of which I made while writing articles.

More information about this application can be found in this article: [Golang: How to Test Code That Exits or Crashes?](https://willsena.dev/golang-how-to-test-code-that-exits-or-crashes/).


# Installing dependencies

```shell
go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@latest
go install golang.org/x/vuln/cmd/govulncheck@latest
```

# Running tests

```shell
make test
```