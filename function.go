package mailApp

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOINTERN")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
		return nil
	}
	return insertResult.InsertedID

}

func InsertTransaction(connote string, tracking_number string, phone_number string, item_content string, delivery_status string) (insertedID interface{}) {
	var transaction Transaction
	transaction.ConsigmentNote = connote
	transaction.TrackingNumber = tracking_number
	transaction.PhoneNumber = phone_number
	transaction.Item_Content = item_content
	transaction.DeliveryStatus = delivery_status
	transaction.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	transaction.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	return InsertOneDoc("Internship1", "MailApp", transaction)
}

func GetAllTransaction() (data []Transaction) {
	transaction := MongoConnect("Internship1").Collection("MailApp")
	filter := bson.M{}
	cursor, err := transaction.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllTransaction :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
