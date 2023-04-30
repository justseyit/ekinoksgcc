package orderops

import (
	"ekinoksgcc/model"
	"testing"
)

func TestAddOrderToDB(t *testing.T) {
	req := model.Order{
		OrderID:   1,
		OrderNote: "test",
	}

	_, err := AddOrderToDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}

func TestGetOrderByOrderIDFromDB(t *testing.T) {
	req := 1
	_, err := GetOrderByOrderIDFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}

func TestGetAllOrdersByOrderIDFromDB(t *testing.T) {
	req := []int{1, 2}
	_, err := GetAllOrdersByIDsFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestGetPlacedOrderByIDFromDB(t *testing.T) {
	req := 1
	_, err := GetPlacedOrderByIDFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestGetPlacedOrdersByUserIDFromDB(t *testing.T) {
	req := 1
	_, err := GetPlacedOrdersByUserIDFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestGetPlacedOrderByOrderIDFromDB(t *testing.T) {

	req := 1
	_, err := GetPlacedOrderByOrderIDFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestAddPlacedOrderToDB(t *testing.T) {

	req := model.PlacedOrder{
		OrderID: 1,
		UserID:  1,
	}

	_, err := AddPlacedOrderToDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestAddOrderItemToDB(t *testing.T) {

	_, err := AddOrderItemToDB()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestGetProductItemByIDFromDB(t *testing.T) {

	req := 1
	_, err := GetProductItemByIDFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestGetAllProductItemsByIDFromDB(t *testing.T) {

	req := []int{1, 2}
	_, err := GetAllProductItemsByIDFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestGetAllProductItemsByPlacedOrderIDFromDB(t *testing.T) {

	req := 1
	_, err := GetAllProductItemsByPlacedOrderIDFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestAddProductItemToDB(t *testing.T) {

	req := model.ProductItem{
		ProductID:           1,
		OrderItemID:         1,
		ProductItemQuantity: 2,
		ProductItemPrice:    10,
	}
	_, err := AddProductItemToDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}
