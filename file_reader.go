package ezconf

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Environment indicates what environment the application is running in.
type Environment string

// Environment types
const (
	DevelopmentEnv Environment = "development"
	ProductionEnv  Environment = "production"
)

// Read a json file at filename into an interface.
func Read(filename string, to interface{}) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(dat, to)
}

// CurrentEnvironment returns the current Environment.
// The current environment is determined by the "ENV" Environment variable.
func CurrentEnvironment() Environment {
	return Environment(os.Getenv("ENV"))
}

func (e Environment) String() string {
	switch e {
	case ProductionEnv:
		return "production"
	default:
		return "development"
	}
}
