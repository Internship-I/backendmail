package mailApp_test

import (
	"fmt"
	"testing"

	module "github.com/internship1/backendmail/module"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertTransaction(t *testing.T) {
	sender := "Muthia"
	receiver := "Dara"
	addressReceiver := "Jl. Ambon No. 123"
	phone := "08123456789"
	item := "Dokumen Penting"
	status := "On Process"
	codValue := 50000.0

	// Call function
	insertedID, err := module.InsertTransaction(module.MongoConn, "MailApp", sender, receiver, addressReceiver, phone, item, status, codValue)
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

