package mailApp

import (
	"context"
	"fmt"
	"math/rand"
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

func GenerateConnote() string {
	timestamp := time.Now().Format("020106") // format: ddmmyy
	randomNum := rand.Intn(9999999) //random 7 digit
	return fmt.Sprintf("P%s%07d", timestamp, randomNum)
}

func InsertTransaction(sender, receiver, phone, item, status string, codValue float64) (interface{}, error) {
    now := primitive.NewDateTimeFromTime(time.Now())
    connote := GenerateConnote()

    newTx := &Transaction{
        ID:             primitive.NewObjectID(),
        ConsigmentNote: connote,
        SenderName:     sender,
        ReceiverName:   receiver,
        PhoneNumber:    phone,
        ItemContent:    item,
        DeliveryStatus: status,
		CODValue: 		codValue,
        CreatedAt:      now,
        UpdatedAt:      now,
    }

    insertedID := InsertOneDoc("Internship1", "MailApp", newTx)
	if insertedID == nil {
		return nil, fmt.Errorf("gagal insert data")
	}

	// Print konfirmasi sekali aja
	fmt.Printf("✅ Data berhasil dimasukkan dengan ID: %v\n", insertedID)

	return insertedID, nil
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

// Get transaction by Consignment Note (resi)
func GetByConsignmentNote(connote string) *Transaction {
	transaction := MongoConnect("Internship1").Collection("MailApp")
	filter := bson.M{"connote": connote}

	var result Transaction
	err := transaction.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("GetByConsignmentNote error:", err)
		return nil
	}
	return &result
}

// Get transactions by Phone Number
func GetByPhoneNumber(phone string) []Transaction {
	transaction := MongoConnect("Internship1").Collection("MailApp")
	filter := bson.M{"phone_number": phone}

	var results []Transaction
	cursor, err := transaction.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetByPhoneNumber error:", err)
		return nil
	}
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		fmt.Println("Cursor decode error:", err)
		return nil
	}
	return results
}

// Get transactions by either Sender Name or Receiver Name (with regex search)
func GetByName(name string) []Transaction {
	transaction := MongoConnect("Internship1").Collection("MailApp")

	// filter dengan regex (case-insensitive)
	filter := bson.M{
		"$or": []bson.M{
			{"sender_name": bson.M{"$regex": name, "$options": "i"}},
			{"receiver_name": bson.M{"$regex": name, "$options": "i"}},
		},
	}

	var results []Transaction
	cursor, err := transaction.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetByName error:", err)
		return nil
	}
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		fmt.Println("Cursor decode error:", err)
		return nil
	}
	return results
}