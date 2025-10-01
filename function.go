package mailApp

import (
	"context"
	"fmt"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	// "time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

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
	}
	return insertResult.InsertedID
}

// func InsertTransaction(db string, collection string, connote string, tracking_number string, phone_number string, item_content string, delivery_status string) (insertedID interface{}) {
// 	transaction := Transaction{
// 		ID:            primitive.NewObjectID(),
// 		ConsigmentNote: connote,
// 		TrackingNumber: tracking_number,
// 		PhoneNumber:    phone_number,
// 		Item_Content:   item_content,
// 		DeliveryStatus: delivery_status,
// 		CreatedAt:      time.Now().Format(time.RFC3339),
// 		UpdatedAt:      time.Now().Format(time.RFC3339),
// 	}
// 	return InsertOneDoc(db, collection, transaction)
// }

// func GetAllTransaction() (data []Transaction) {
// 	karyawan := MongoConnect("tesdb2024").Collection("presensi")
// 	filter := bson.M{}
// 	cursor, err := karyawan.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("GetALLData :", err)
// 	}
// 	err = cursor.All(context.TODO(), &data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }