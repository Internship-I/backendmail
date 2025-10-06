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

// GetTransactionByPhoneNumber retrieves transactions by phone number
func GetByPhoneNumber(phone string, db *mongo.Database, col string) ([]model.Transaction, error) {
	var transactions []model.Transaction
	collection := db.Collection(col)
	filter := bson.M{"phone_number": phone}

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
		return nil, fmt.Errorf("transaction dengan nomor hp %s tidak ditemukan", phone)
	}

	return transactions, nil
}

// GetTransactionByAddress retrieves transactions by address
func GetByAddress(addressReceiver string, db *mongo.Database, col string) ([]model.Transaction, error) {
	var transactions []model.Transaction
	collection := db.Collection(col)
	filter := bson.M{"address_receiver": addressReceiver}

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
		return nil, fmt.Errorf("transaction dengan alamat %s tidak ditemukan", addressReceiver)
	}

	return transactions, nil
}

// GetBySenderOrReceiver retrieves transactions by either sender or receiver name
func GetBySenderOrReceiver(name string, db *mongo.Database, col string) ([]model.Transaction, error) {
	var transactions []model.Transaction
	collection := db.Collection(col)

	filter := bson.M{
		"$or": []bson.M{
			{"sender_name": name},
			{"receiver_name": name},
		},
	}

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
		return nil, fmt.Errorf("transaction dengan sender atau receiver %s tidak ditemukan", name)
	}

	return transactions, nil
}