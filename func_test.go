package mailApp

import (
	"fmt"
	"testing"
)

func TestInsertTransaction(t *testing.T) {
	SenderName := "Nida"
	ReceiverName := "Qinthar"
	PhoneNumber := "083174603834"
	Item_Content := "Makanan"
	DeliveryStatus := "On Process"

	insertedID, err := InsertTransaction(SenderName, ReceiverName, PhoneNumber, Item_Content, DeliveryStatus)
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
	phone := "083174603834" // ganti dengan nomor hp yang ada di database
	results := GetByPhoneNumber(phone)

	if len(results) == 0 {
		t.Fatalf("GetByPhoneNumber gagal, data dengan phone %s tidak ditemukan", phone)
	}
	fmt.Println("Hasil pencarian berdasarkan phone number:", results)
}

func TestGetByName(t *testing.T) {
	name := "nid" // cukup sebagian aja, regex akan tetap match misal "Dewi"
	results := GetByName(name)

	if len(results) == 0 {
		t.Fatalf("GetByName gagal, data dengan nama %s tidak ditemukan", name)
	}
	fmt.Println("Hasil pencarian berdasarkan nama:", results)
}
