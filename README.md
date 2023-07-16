# `protoc-gen-cue`

This protoc plugin generates CUE files.

## Conversions

Complies with [protojson](https://protobuf.dev/programming-guides/proto3/#json).

### Basic Types

| Proto Type      | CUE Type                            | Comments                           |
|-----------------|-------------------------------------|------------------------------------|
| `map<K, V>`     | `{ [string]: V }`                   | All keys are converted to strings. |
| `repeated V`    | `[...V]`                            |                                    |
| `bool`          | `bool`                              |                                    |
| `string`        | `string`                            |                                    |
| `bytes`         | `bytes`                             |                                    |
| `int32`         | `int32`                             |                                    |
| `fixed32`       | `int32`                             |                                    |
| `uint32`        | `uint32`                            |                                    |
| `int64`         | `int64`                             |                                    |
| `fixed64`       | `fixed64`                           |                                    |
| `uint32`        | `uint64`                            |                                    |
| `float`         | `float32`                           |                                    |
| `double`        | `float64`                           |                                    |
| `Any`           | `null \| { "@type": string, ... }`  |                                    |
| `Struct`        | `null \| { [string]: _ }`           |                                    |
| `Value`         | `null \| _`                         |                                    |
| `ListValue`     | `null \| [...]`                     |                                    |
| `NullValue`     | `null \| null`                      |                                    |
| `BoolValue`     | `null \| bool`                      |                                    |
| `StringValue`   | `null \| string`                    |                                    |
| `Int32Value`    | `null \| int32`                     |                                    |
| `UInt32Value`   | `null \| uint32`                    |                                    |
| `Int64Value`    | `null \| int64`                     |                                    |
| `UInt64Value`   | `null \| uint64`                    |                                    |
| `FloatValue`    | `null \| float32`                   |                                    |
| `DoubleValue`   | `null \| double`                    |                                    |
| `Empty`         | `null \| close({})`                 |                                    |
| `Timestamp`     | `null \| string`                    | See the [`Timestamp`](#timestamp) section for more information. |
| `Duration`      | `null \| string`                    | See the [`Duration`](#duration) section for more information. |
| `FieldMask`     | `null \| { paths: [...string] }`    |                                    |

### Message

```proto
import "github.com/ornew/protoc-gen-cue/pkg/options/cue.proto";

message Foo {
  string name = 1 [(cue.field).expr = '!="xxx"'];
  int32 age = 2 [(cue.field).expr = '<100'];
  int32 age_next_year = 3 [(cue.field).expr = 'age+1'];
}
```

To:

```cue
#Foo: {
  name: string
  name: !="xxx"
  age: int32
  age: <100
  ageNextYear: age+1
}
```

### Enum

```proto
enum Bar {
  ZERO = 0;
  ONE = 1;
}
```

To:

```cue
#Bar: #Bar_ZERO | #Bar_ONE

#Bar_ZERO: "ZERO"
#Bar_ONE:  "ONE"
```

### Oneof

```proto
message Car {
  oneof id {
    string product_name = 1;
    int32 serial_number = 2;
  }
}
```

To:

```cue
#Car: {
  _oneof_id: productName & serialNumber
  productName?: string
  serialNumber?: int32
}
```

### Timestamp

Currently defined by an unconstrained `string`. This is due to the fact that CUE's built-in `time.Time` constraint is incompatible with the JSON format defined in the `timestamppb`. We plan to fix this issue in a future version to follow the original format. See for more details: [time.Time on pkg.go.dev](https://pkg.go.dev/cuelang.org/go@v0.5.0/pkg/time#Time) and [timestamppb.Timestamp](https://pkg.go.dev/google.golang.org/protobuf@v1.31.0/types/known/timestamppb#hdr-JSON_Mapping)

### Duration

Currently defined by an unconstrained `string`. This is due to the fact that CUE's built-in `time.Duration` constraint is incompatible with the JSON format defined in the `durationpb`. We plan to fix this issue in a future version to follow the original format. See for more details: [time.Duration in pkg.go.dev](https://pkg.go.dev/cuelang.org/go@v0.5.0/pkg/time#Duration) and [descriptorpb.Duration](https://pkg.go.dev/google.golang.org/protobuf/types/known/durationpb#hdr-JSON_Mapping)

### Optional (proto3)

```proto
message Dog {
  optional string nick_name = 1;
}
```

To:

```cue
#Dog: {
  nickName?: string
}
```
