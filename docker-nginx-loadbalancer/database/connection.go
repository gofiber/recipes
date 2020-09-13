package database

import (
	"context"
	"encoding/json"
	"fiber-docker-nginx/models"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN : content a mongoDB connection*/
var MongoCN = ConnectDB()

/*Config : all config file content*/
var Config = LoadConfiguration()

var clientOptions = options.Client().ApplyURI(
	fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		Config.DB.User,
		Config.DB.Password,
		Config.DB.Server,
		Config.DB.Cluster,
	),
)

/*LoadConfiguration : load all configuration file*/
func LoadConfiguration() models.Configuration {
	var configuration models.Configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal("Error reading the file")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}
	return configuration
}

/*ConnectDB : Create a connection to mongoDB and return the connection*/
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Connect successfully")
	return client
}

/*ConnectionOK : Check the connection and return true or false */
func ConnectionOK() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}
