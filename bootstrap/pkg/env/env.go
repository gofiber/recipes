package env

import "github.com/joho/godotenv"

var Env map[string]string

func GetEnv(key, def string) string {
	if val, ok := Env[key]; ok {
		return val
	}
	return def
}

func SetupEnvFile() {
	envFile := ".env"
	var err error
	Env, err = godotenv.Read(envFile)
	if err != nil {
		panic(err)
	}

}
