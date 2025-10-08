package mailApp_test

import (
	"fmt"
	"testing"

	module "github.com/Internship-I/backendmail/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertTransaction(t *testing.T) {
	sender := "Qinthar"
	sender_phone := "082127854156"
	receiver := "Hilwa"
	addressReceiver := "Jl. Banda No.5"
	receiver_phone := "085624064624"
	item := "Makanan "
	status := "Delivered"
	codValue := 105000.0

	// Call function
	insertedID, connote, err := module.InsertTransaction(module.MongoConn, "MailApp", sender, sender_phone, receiver, addressReceiver, receiver_phone, item, status, codValue)
	if err != nil {
		t.Errorf("InsertTransaction gagal: %v", err)
	}

	// Assertion: cek apakah ObjectID valid
	if insertedID.IsZero() {
		t.Errorf("InsertTransaction gagal, ObjectID kosong")
	}

	fmt.Printf("✅ TestInsertTransaction berhasil dengan ID: %v | Connote: %s\n", insertedID, connote)
}

func TestGetAllTransaction(t *testing.T) {
	data := module.GetAllTransaction(module.MongoConn, "MailApp")
	fmt.Println(data)
}

func TestGetByConsignmentNote(t *testing.T) {
	db := module.MongoConn
	col := "MailApp" // ganti sesuai nama collection

	// Gunakan consignment note yang ada di database
	connote := "P0610253848524" // ganti sesuai data nyata atau hasil InsertTransaction

	transactions, err := module.GetTransactionByConnote(connote, db, col)
	if err != nil {
		t.Fatalf("error calling GetByConsignmentNote: %v", err)
	}

	fmt.Printf("Ditemukan %d transaksi:\n", len(transactions))
	for _, tx := range transactions {
		fmt.Printf("%+v\n", tx)
	}
}

// TestGetByPhoneNumber
func TestGetByPhoneNumber(t *testing.T) {
	db := module.MongoConn
	col := "MailApp" // ganti sesuai nama collection

	// Gunakan consignment note yang ada di database
	phone := "082127854156" // ganti sesuai data nyata atau hasil InsertTransaction

	transactions, err := module.GetByPhoneNumber(phone, db, col)
	if err != nil {
		t.Fatalf("error calling GetByPhoneNumber: %v", err)
	}

	fmt.Printf("Ditemukan %d transaksi:\n", len(transactions))
	for _, tx := range transactions {
		fmt.Printf("%+v\n", tx)
	}
}

// TestGetByAddress
func TestGetByAddress(t *testing.T) {
	db := module.MongoConn
	col := "MailApp"
	testAddress := "Banda" // bisa match "Jl. Banda No. 5", "Banda Aceh", dll

	transactions, err := module.GetByAddress(testAddress, db, col)
	if err != nil {
		t.Errorf("Error memanggil GetByAddress: %v", err)
		return
	}

	if len(transactions) == 0 {
		t.Errorf("Tidak ditemukan transaksi dengan alamat mengandung: %s", testAddress)
		return
	}

	fmt.Printf("✅ Ditemukan %d transaksi yang mengandung '%s':\n", len(transactions), testAddress)
	for i, tx := range transactions {
		fmt.Printf("%d. Connote: %s | Sender: %s | Receiver: %s | Address: %s\n",
			i+1, tx.ConsigmentNote, tx.SenderName, tx.ReceiverName, tx.AddressReceiver)
	}
}

// // TestGetBySenderOrReceiver
// func TestGetBySenderOrReceiver(t *testing.T) {
// 	db := module.MongoConn
// 	col := "transactions"

// 	name := "Alice"

// 	results, err := module.GetBySenderOrReceiver(name, db, col)
// 	if err != nil {
// 		t.Fatalf("GetBySenderOrReceiver error: %v", err)
// 	}

// 	fmt.Printf("Ditemukan %d transaksi dengan sender/receiver %s:\n", len(results), name)
// 	for _, tx := range results {
// 		fmt.Printf("%+v\n", tx)
// 	}
// }

//FUNCTION USER
func TestInsertUser(t *testing.T) {
    // Test data
	name := "Muhammad Qinthar"
    phone_number := "081234567890"
    username := "Qintharalmaliki"
    password := "Qinthar123"
    role := "Admin"
 
	 // Call the function
	 insertedID, err := module.InsertUser(module.MongoConn, "User", name, phone_number, username, password, role)
	 if err != nil {
		 t.Fatalf("Error inserting user: %v", err)
	 }
 
	 // Print the result
	 fmt.Printf("Data berhasil disimpan dengan id %s\n", insertedID.Hex())
}

//GetUserByID retrieves a user from the database by its ID
func TestGetUserByID(t *testing.T) {
	_id := "68e473a524e138fe0dee8caf"
	objectID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	menu, err := module.GetUserByID(objectID, module.MongoConn, "User")
	if err != nil {
		t.Fatalf("error calling GetMenuItemByID: %v", err)
	}
	fmt.Println(menu)
}

func TestGetAllUsers(t *testing.T) {
	data, err := module.GetAllUser(module.MongoConn, "User")
	if err != nil {
		t.Fatalf("error calling GetAllUsers: %v", err)
	}
	fmt.Println(data)
}	

func TestDeleteUserByID(t *testing.T) {
    id := "68e473a524e138fe0dee8caf"
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        t.Fatalf("error converting id to ObjectID: %v", err)
    }

    err = module.DeleteUserByID(objectID, module.MongoConn, "User")
    if err != nil {
        t.Fatalf("error calling DeleteUserByID: %v", err)
    }

    _, err = module.GetUserByID(objectID, module.MongoConn, "User")
    if err == nil {
        t.Fatalf("expected data to be deleted, but it still exists")
    }
}