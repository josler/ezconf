package ezconf

import (
	"fmt"
	"os"
)

type envStruct struct {
	Foo string
}

func ExampleEnvTransform() {
	// set up a parser and filter
	p := NewParser()
	filter := Filter{Type: String}
	p.RegisterTransformer(EnvTransform, filter)

	// artificially set an environment variable
	os.Setenv("EZCONF", "ez")

	result := p.Transform(envStruct{"ENV['EZCONF']"})
	fmt.Println(result.(envStruct).Foo)
	// Output: ez
}
