package config

import (
	"context"
	"fmt"
	"os"

	"github.com/danielsugianto/alterra-agmc-day-10/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitialDB() *gorm.DB {

	config := map[string]string{
		"DB_Username": os.Getenv("DB_USERNAME"),
		"DB_Password": os.Getenv("DB_PASSWORD"),
		"DB_Port":     os.Getenv("DB_PORT"),
		"DB_Host":     os.Getenv("DB_HOST"),
		"DB_Name":     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitialMigration(DB)
	return DB
}

func InitialMongoDB() *mongo.Database {
	var ctx = context.Background()

	configMongoDB := map[string]string{
		"MONGODB_HOST": os.Getenv("MONGODB_HOST"),
		"MONGODB_PORT": os.Getenv("MONGODB_PORT"),
	}
	clientOptions := options.Client()
	connectionMongoString := fmt.Sprintf("mongodb://%s:%s", configMongoDB["MONGODB_HOST"], configMongoDB["MONGODB_PORT"])
	clientOptions.ApplyURI(connectionMongoString)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	return client.Database("agmc")
}
func InitialMigration(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
}
