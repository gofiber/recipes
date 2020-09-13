package models

type Configuration struct {
	DB struct {
		User string
		Password string
		Server string
		Cluster string
	}
	ENV string
}
