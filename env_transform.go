package ezconf

import (
	"os"
	"regexp"
)

// EnvTransform is a TransformFunc that transforms string values with an
// ENV variable inside into the corresponding value from the environment.
func EnvTransform(val interface{}) interface{} {
	stringVal := val.(string)
	if key, ok := isEnv(stringVal); ok {
		return os.Getenv(key)
	}
	return val
}

func isEnv(val string) (string, bool) {
	matched, err := regexp.MatchString("ENV\\['.*'\\]", val)
	if err != nil {
		panic(err)
	}
	if !matched {
		return "", false
	}
	envKey := val[5 : len(val)-2]
	return envKey, true
}
