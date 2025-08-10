# mockguard

`mockguard` checks if mockgen is used in conventional filename and options.

## Config

`mockguard` provides default convention. Please see [config.go](config.go)

You can overwrite via commandline option or golangci setting.

- FileName
- Content

## Install

```shell
$ go install github.com/cloverrose/mockguard/cmd/mockguard@latest
```

### Or Build from source

```shell
$ make build/mockguard
```

### Or Install via aqua

https://aquaproj.github.io/

## Usage

```shell
$ go vet -vettool=`which mockguard` ./...
```

When you specify config

```shell
go vet -vettool=`which mockguard` \
  -mockguard.FileName=types.go \
  -mockguard.Content='//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE --destination=mock_$GOFILE' \
   ./...
```

### Or golangci-lint custom plugin

https://golangci-lint.run/plugins/module-plugins/

Here are reference settings

`.custom-gcl.yml`

```yaml
version: v2.3.1
plugins:
  - module: 'github.com/cloverrose/mockguard'
    import: 'github.com/cloverrose/mockguard'
    version: v0.1.3
```

`.golangci.yml`

```yaml
linters-settings:
  custom:
    mockguard:
      type: "module"
      description: check mockgen usage.
      settings:
        FileName: types.go
        Content: "//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE --destination=mock_$GOFILE"
```
