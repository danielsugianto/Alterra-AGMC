package config

import (
	"context"
	"fmt"

	"github.com/danielsugianto/alterra-agmc-day6/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitialDB() *gorm.DB {

	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     "agmc",
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

	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
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
