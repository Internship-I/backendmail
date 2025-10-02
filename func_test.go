package mailApp

import (
	"fmt"
	"testing"
)

func TestInsertTransaction(t *testing.T) {
	SenderName := "Hani"
	ReceiverName := "Kibo"
	AddressReceiver := "Jakarta"
	PhoneNumber := "09876643"
	CODValue := 75000.0
	Item_Content := "Makanan"
	DeliveryStatus := "Delivered"

	insertedID, err := InsertTransaction(SenderName, ReceiverName, AddressReceiver, PhoneNumber, Item_Content, DeliveryStatus, CODValue)
	if err != nil {
		t.Fatal("InsertTransaction gagal:", err)
	}

	if insertedID == nil {
		t.Fatal("InsertTransaction gagal, InsertedID nil")
	}
}

func TestGetAllTransaction(t *testing.T) {
	data := GetAllTransaction()
	fmt.Println(data)
}

func TestGetByConsignmentNote(t *testing.T) {
	// pastikan dulu ada data dengan connote tertentu
	connote := "P0210256827510" // ganti dengan resi yang sudah pernah ke-insert
	result := GetByConsignmentNote(connote)

	if result == nil {
		t.Fatalf("GetByConsignmentNote gagal, data dengan resi %s tidak ditemukan", connote)
	}
	fmt.Println("Hasil pencarian berdasarkan resi:", result)
}

func TestGetByPhoneNumber(t *testing.T) {
	phone := "083174603834"
	results := GetByPhoneNumber(phone)

	if len(results) == 0 {
		t.Fatalf("GetByPhoneNumber gagal, data dengan phone %s tidak ditemukan", phone)
	}
	fmt.Println("Hasil pencarian berdasarkan phone number:", results)
}

func TestGetByName(t *testing.T) {
	name := "nid"
	results := GetByName(name)

	if len(results) == 0 {
		t.Fatalf("GetByName gagal, data dengan nama %s tidak ditemukan", name)
	}
	fmt.Println("Hasil pencarian berdasarkan nama:", results)
}

func TestGetByAddress(t *testing.T) {
	results, err := GetByAddress("Jakarta")
	if err != nil {
		t.Fatal("Error GetByAddress:", err)
	}

	if len(results) == 0 {
		t.Fatal("Tidak ada transaksi dengan alamat tersebut")
	}

	for _, r := range results {
		fmt.Println("Hasil:", r.ReceiverName, "-", r.AddressReceiver)
	}
}
