package env

import (
	"fmt"
	"os"
)

func Get(key string) string {
	v, found := os.LookupEnv(key)
	if !found {
		panic(fmt.Errorf("the envioriment variable of key %v does not exist", key))
	}

	return v
}
