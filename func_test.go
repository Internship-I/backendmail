package mailApp

import (
	"fmt"
	"testing"
)

func TestInsertTransaction(t *testing.T) {
	SenderName := "Hanif"
	ReceiverName := "Jaka"
	PhoneNumber := "083174603834"
	CODValue := 125000.0
	Item_Content := "Sepatu"
	DeliveryStatus := "On Procces"

	insertedID, err := InsertTransaction(SenderName, ReceiverName, PhoneNumber, Item_Content, DeliveryStatus, CODValue)
	if err != nil {
		t.Fatal("InsertTransaction gagal:", err)
	}

	if insertedID == nil {
		t.Fatal("InsertTransaction gagal, InsertedID nil")
	}

	// if hasil.ConsigmentNote == "" {
	// 	t.Fatal("ConsigmentNote gagal tergenerate")
	// }
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
