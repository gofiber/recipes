package config

import "os"

// Config func to get env value.
// godotenv.Load is called once in main(); this function just reads the already-set env vars.
func Config(key string) string {
	return os.Getenv(key)
}
