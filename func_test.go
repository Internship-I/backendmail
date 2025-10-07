package mailApp_test

import (
	"fmt"
	"testing"

	module "github.com/internship1/backendmail/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertTransaction(t *testing.T) {
	sender := "Muthia"
	sender_phone := "083174603834"
	receiver := "Dara"
	addressReceiver := "Jl. Ambon No. 123"
	receiver_phone := "08123456789"
	item := "Dokumen Penting"
	status := "On Process"
	codValue := 50000.0

	// Call function
	insertedID, err := module.InsertTransaction(module.MongoConn, "MailApp", sender, sender_phone, receiver, addressReceiver, receiver_phone, item, status, codValue)
	if err != nil {
		t.Errorf("InsertTransaction gagal: %v", err)
	}

	// Assertion: cek apakah ObjectID valid
	if insertedID.IsZero() {
		t.Errorf("InsertTransaction gagal, ObjectID kosong")
	}

	fmt.Printf("✅ TestInsertTransaction berhasil dengan ID: %v\n", insertedID)
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
	phone := "083867818081" // ganti sesuai data nyata atau hasil InsertTransaction

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
	col := "MailApp" // ganti sesuai nama collection

	// Gunakan consignment note yang ada di database
	addressReceiver := "Jakarta" // ganti sesuai data nyata atau hasil InsertTransaction

	transactions, err := module.GetByAddress(addressReceiver, db, col)
	if err != nil {
		t.Fatalf("error calling GetByAddress: %v", err)
	}

	fmt.Printf("Ditemukan %d transaksi:\n", len(transactions))
	for _, tx := range transactions {
		fmt.Printf("%+v\n", tx)
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