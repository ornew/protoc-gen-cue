```bash
cp ../pkg/options/cue.proto vendor/cue.proto
buf generate proto
cue export ./gen/foo_gen.cue foo.cue | tee foo.json
```
