# protoc-gen-cue

This protoc plugin generates CUE files.

## Options

### `root`

Defaults to `.`.

See for more details: [`(protobuf.Config).Root` on pkg.go.dev](https://pkg.go.dev/cuelang.org/go@v0.5.0/encoding/protobuf?utm_source=gopls#Config.Root)

### `module`

Defaults to `.`.

See for more details: [`(protobuf.Config).Module` on pkg.go.dev](https://pkg.go.dev/cuelang.org/go@v0.5.0/encoding/protobuf?utm_source=gopls#Config.Module)

### `imports`

Defaults to empty. `;`-separated. e.g, `.;include;vendor`

See for more details: [`(protobuf.Config).Paths` on pkg.go.dev](https://pkg.go.dev/cuelang.org/go@v0.5.0/encoding/protobuf?utm_source=gopls#Config.Paths)

### `enum_mode`

Defaults to `json`.

See for more details: [`(protobuf.Config).EnumMode` on pkg.go.dev](https://pkg.go.dev/cuelang.org/go@v0.5.0/encoding/protobuf?utm_source=gopls#Config.EnumMode)
