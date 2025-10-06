package module

import (
	"context"
	"fmt"
	"math/rand"
	// "regexp"
	"time"

	"github.com/internship1/backendmail/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect gagal: %v\n", err)
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

func InsertTransaction(db *mongo.Database, col string, sender string, receiver string, addressReceiver string, phone string, item string, status string, codValue float64) (insertedID primitive.ObjectID, err error) {
	now := primitive.NewDateTimeFromTime(time.Now())
	connote := GenerateConnote()

	transaction := bson.M{
		"consignment_note": connote,
		"sender_name":      sender,
		"receiver_name":    receiver,
		"address_receiver": addressReceiver,
		"phone_number":     phone,
		"item_content":     item,
		"delivery_status":  status,
		"cod_value":        codValue,
		"created_at":       now,
		"updated_at":       now,
	}

	result, err := db.Collection(col).InsertOne(context.Background(), transaction)
	if err != nil {
		fmt.Printf("InsertTransaction error: %v\n", err)
		return
	}

	insertedID = result.InsertedID.(primitive.ObjectID)
	fmt.Printf("✅ Data berhasil dimasukkan dengan ID: %v\n", insertedID)

	return insertedID, nil
}

// GetAllTransaction retrieves all transaction from the database
func GetAllTransaction(db *mongo.Database, col string) (data []model.Transaction) {
	transaction := db.Collection(col)
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

// GetTransactionByConnote retrieves transactions by consignment note
func GetTransactionByConnote(connote string, db *mongo.Database, col string) ([]model.Transaction, error) {
	var transactions []model.Transaction
	collection := db.Collection(col)
	filter := bson.M{"consignment_note": connote}

	cursor, err := collection.Find(context.TODO(), filter, options.Find())
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan transaction: %w", err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var t model.Transaction
		if err := cursor.Decode(&t); err != nil {
			continue
		}
		transactions = append(transactions, t)
	}

	if len(transactions) == 0 {
		return nil, fmt.Errorf("transaction dengan connote %s tidak ditemukan", connote)
	}

	return transactions, nil
}

// // Get transactions by Phone Number
// func GetByPhoneNumber(phone string) []Transaction {
// 	transaction := MongoConnect("Internship1").Collection("MailApp")
// 	filter := bson.M{"phone_number": phone}

// 	var results []Transaction
// 	cursor, err := transaction.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("GetByPhoneNumber error:", err)
// 		return nil
// 	}
// 	err = cursor.All(context.TODO(), &results)
// 	if err != nil {
// 		fmt.Println("Cursor decode error:", err)
// 		return nil
// 	}
// 	return results
// }

// // Get transactions by either Sender Name or Receiver Name (with regex search)
// func GetByName(name string) []Transaction {
// 	transaction := MongoConnect("Internship1").Collection("MailApp")

// 	// filter dengan regex (case-insensitive)
// 	filter := bson.M{
// 		"$or": []bson.M{
// 			{"sender_name": bson.M{"$regex": name, "$options": "i"}},
// 			{"receiver_name": bson.M{"$regex": name, "$options": "i"}},
// 		},
// 	}

// 	var results []Transaction
// 	cursor, err := transaction.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("GetByName error:", err)
// 		return nil
// 	}
// 	err = cursor.All(context.TODO(), &results)
// 	if err != nil {
// 		fmt.Println("Cursor decode error:", err)
// 		return nil
// 	}
// 	return results
// }

// // GetByAddress mencari transaksi berdasarkan alamat penerima
// func GetByAddress(address string) ([]Transaction, error) {
// 	var results []Transaction

// 	// Koneksi ke collection
// 	collection := MongoConnect("Internship1").Collection("MailApp")

// 	regexPattern := fmt.Sprintf(".*%s.*", regexp.QuoteMeta(address))
// 	filter := bson.M{
// 		"address_receiver": bson.M{
// 			"$regex":   regexPattern,
// 			"$options": "i", // ignore case
// 		},
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	cursor, err := collection.Find(ctx, filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	if err := cursor.All(ctx, &results); err != nil {
// 		return nil, err
// 	}

// 	return results, nil
// }
