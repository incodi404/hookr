package utils

import "os"

func GetEnv(key string, def string) string {
	var val string
	val = os.Getenv(key)
	if val == "" {
		val = def
	}

	return val

}
