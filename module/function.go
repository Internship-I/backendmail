package module

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	// "regexp"
	"time"

	"github.com/Internship-I/backendmail/model"
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

func InsertTransaction(db *mongo.Database, col string, sender string, sender_phone string, receiver string, addressReceiver string, receiver_phone string, item string, status string, codValue float64) (insertedID primitive.ObjectID, connote string, err error) {
	now := primitive.NewDateTimeFromTime(time.Now())
	connote = GenerateConnote() // assign ke variabel global agar bisa direturn

	transaction := bson.M{
		"consignment_note": connote,
		"sender_name":      sender,
		"sender_phone":     sender_phone,
		"receiver_name":    receiver,
		"address_receiver": addressReceiver,
		"receiver_phone":   receiver_phone,
		"item_content":     item,
		"delivery_status":  status,
		"cod_value":        codValue,
		"created_at":       now,
		"updated_at":       now,
	}

	result, err := db.Collection(col).InsertOne(context.Background(), transaction)
	if err != nil {
		fmt.Printf("❌ InsertTransaction error: %v\n", err)
		return insertedID, connote, err
	}

	insertedID = result.InsertedID.(primitive.ObjectID)
	fmt.Printf("✅ Data berhasil dimasukkan dengan ID: %v | Connote: %s\n", insertedID, connote)

	return insertedID, connote, nil
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

// // GetBySenderOrReceiver retrieves transactions by either sender or receiver name
// func GetBySenderOrReceiver(name string, db *mongo.Database, col string) ([]model.Transaction, error) {
// 	var transactions []model.Transaction
// 	collection := db.Collection(col)

// 	filter := bson.M{
// 		"$or": []bson.M{
// 			{"sender_name": name},
// 			{"receiver_name": name},
// 		},
// 	}

// 	cursor, err := collection.Find(context.TODO(), filter, options.Find())
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal mendapatkan transaction: %w", err)
// 	}
// 	defer cursor.Close(context.TODO())

// 	for cursor.Next(context.TODO()) {
// 		var t model.Transaction
// 		if err := cursor.Decode(&t); err != nil {
// 			continue
// 		}
// 		transactions = append(transactions, t)
// 	}

// 	if len(transactions) == 0 {
// 		return nil, fmt.Errorf("transaction dengan sender atau receiver %s tidak ditemukan", name)
// 	}

// 	return transactions, nil
// }

// FUNCTION USER
// GetUserByID retrieves a user from the database by its ID
func GetUserByID(_id primitive.ObjectID, db *mongo.Database, col string) (model.User, error) {
	var user model.User
	collection := db.Collection("User")
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, fmt.Errorf("GetUserByID: user dengan ID %s tidak ditemukan", _id.Hex())
		}
		return user, fmt.Errorf("GetUserByID: gagal mendapatkan data user: %w", err)
	}
	return user, nil
}

func GetRoleByAdmin(db *mongo.Database, collection string, role string) (*model.User, error) {
	var user model.User
	filter := bson.M{"role": role}
	opts := options.FindOne()

	err := db.Collection(collection).FindOne(context.Background(), filter, opts).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func InsertUsers(db *mongo.Database, col string, fullname string, phonenumber string, username string, password string, role string) (insertedID primitive.ObjectID, err error) {
	users := bson.M{
		"fullname": fullname,
		"phone":    phonenumber,
		"username": username,
		"password": password,
		"role":     role,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), users)
	if err != nil {
		fmt.Printf("InsertUser: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetByUsername(db *mongo.Database, col string, username string) (*model.User, error) {
	var admin model.User
	err := db.Collection(col).FindOne(context.Background(), bson.M{"username": username}).Decode(&admin)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func DeleteTokenFromMongoDB(db *mongo.Database, col string, token string) error {
	collection := db.Collection(col)
	filter := bson.M{"token": token}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

// GetAllUser retrieves all users from the database
func GetAllUser(db *mongo.Database, col string) ([]model.User, error) {
	var data []model.User
	user := db.Collection(col)

	cursor, err := user.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("GetAllUser error:", err)
		return nil, err
	}
	defer cursor.Close(context.TODO()) // Selalu tutup cursor

	if err := cursor.All(context.TODO(), &data); err != nil {
		fmt.Println("Error decoding users:", err)
		return nil, err
	}

	return data, nil
}

func SaveTokenToDatabase(db *mongo.Database, col string, adminID string, token string) error {
	collection := db.Collection(col)
	filter := bson.M{"admin_id": adminID}
	update := bson.M{
		"$set": bson.M{
			"token":      token,
			"updated_at": time.Now(),
		},
	}
	_, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	return nil
}

// InsertUser creates a new order in the database
func InsertUser(db *mongo.Database, col string, name string, phone string, username string, password string, role string) (insertedID primitive.ObjectID, err error) {
	user := bson.M{
		"name":                    name,
		"phone":                   phone,
		"username":                username,
		"password":                password,
		"role":                    role,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), user)
	if err != nil {
		fmt.Printf("InsertUser: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// UpdateUser updates an existing user in the database
func UpdateUser(ctx context.Context, db *mongo.Database, col string, _id primitive.ObjectID, name string, phone string, username string, password string, role string) (err error) {
	filter := bson.M{"_id": _id}
	update := bson.M{
		"$set": bson.M{
			"name":     name,
			"phone":    phone,
			"username": username,
			"password": password,
			"role":     role,
		},
	}
	result, err := db.Collection(col).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("UpdateUser: gagal memperbarui User: %w", err)
	}
	if result.MatchedCount == 0 {
		return errors.New("UpdateUser: tidak ada data yang diubah dengan ID yang ditentukan")
	}
	return nil
}

// DeleteUserByID deletes a menu item from the database by its ID
func DeleteUserByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	user := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := user.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

