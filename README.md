# AAA Test linter for Golang

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
