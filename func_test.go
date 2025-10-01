package mailApp

import (
	"fmt"
	"testing"
)

func TestInsertTransaction(t *testing.T) {
	ConsigmentNote := "CN-20251001-001"
	TrackingNumber := "TRK123456789"
	PhoneNumber := "6281234567890"
	Item_Content := "Dokumen Penting"
	DeliveryStatus := "On Process"
	hasil := InsertTransaction(ConsigmentNote, TrackingNumber, PhoneNumber, Item_Content, DeliveryStatus)
	fmt.Println(hasil)

	if hasil == nil {
    t.Fatal("InsertTransaction gagal, hasil nil")
}

}

func TestGetAllTransaction(t *testing.T) {
	data := GetAllTransaction()
	fmt.Println(data)
}
