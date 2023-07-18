```bash
cp ../pkg/options/cue.proto vendor/cue.proto
buf generate src
cue export ./gen/example/v1/example_gen.cue example_values.cue | tee example_values.json
cue export ./gen/example/v1/example_gen.cue example_default.cue | tee example_default.json
```
