
# EZConf [![GoDoc](http://godoc.org/github.com/josler/ezconf?status.png)](http://godoc.org/github.com/josler/ezconf)

EZConf is a really dumb config parser.

### Reading Files

EZConf reads JSON files and parses them to Config structs.

```go
ezconf.Read(filename, &myStruct)
```

### Transformations

EZConf can transform struct values programatically; just register a `TransformerFunc` and a `Filter` and then ask it to transform the struct values.

For example, there's a built in `ezconf.EnvTransform` function, that parses string vales for ENV variables (marked by `"ENV['SOMEKEY']"`), and replaces the value with that from the environment.

```go
p := ezconf.NewParser()
filter := ezconf.Filter{Type: ezconf.String}
p.RegisterTransformer(ezconf.EnvTransform, filter)
outputStruct := p.Transform(myInputStruct)
```
