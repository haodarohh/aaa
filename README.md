# AAA Test linter for Golang

## Installation

Requires Go 1.24.5 or later. Install `golangci-lint` separately if you want to run this linter through a golangci-lint plugin.

```shell
git clone https://github.com/haodarohh/AAA.git
cd AAA
go mod download
```

To install the standalone analyzer binary:

```shell
go install github.com/haodarohh/aaa@latest
```

For local development, install it from the repository checkout:

```shell
go install .
```

## Usage

#### Go Plugin Syetem

- Check the golangci-lint package version and modify the `go.mod` to exact version.

```shell
golangci-lint version --debug | grep "golang.org/x/tools"
```

```shell
# build
go build -buildmode=plugin main.go

# run
golangci-lint run ./...
```

- Reference
  - https://golangci-lint.run/plugins/go-plugins/
  - https://github.com/golangci/example-plugin-linter
  - https://blog.csdn.net/Dusong_/article/details/144088947

#### Module Plugin System

- Add `.custom-gcl.yml`, the `module` in the file must contain `register.Plugin(xxx)`
- build new custom golangci-lint

```shell
golangci-lint custom -v
```

```shell
./custom-gcl -c .golangci.custom.yml run ./...
```

- Reference
  - https://golangci-lint.run/plugins/module-plugins/
  - https://github.com/golangci/example-plugin-module-linter
