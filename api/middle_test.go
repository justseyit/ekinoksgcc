package api

import (
	"database/sql"
	"ekinoksgcc/model"
	"ekinoksgcc/repository"
	"strings"
	"testing"
)

func TestGetOrderInfo(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := model.RequestOrderInfo{
		UserID: 1,
	}
	resp, err := getOrderInfo(req)
	if err != nil && err != sql.ErrNoRows || resp.Success != true || resp.Message != "" {
		if !strings.Contains(err.Error(), "sql: no rows in result set") {
			t.Errorf("Error: %v", err)
		}
	}

	req2 := model.RequestOrderInfo{
		OrderID: 1,
	}
	resp2, err2 := getOrderInfo(req2)
	if err2 != nil && err2 != sql.ErrNoRows || resp2.Success != true || resp2.Message != "" {
		if !strings.Contains(err2.Error(), "sql: no rows in result set") {
			t.Errorf("Error: %v", err2)
		}
	}

	req3 := model.RequestOrderInfo{
		UserID:  1,
		OrderID: 1,
	}
	resp3, err3 := getOrderInfo(req3)
	if err3 != nil && err != sql.ErrNoRows || resp3.Success != true || resp3.Message != "" {
		if !strings.Contains(err3.Error(), "sql: no rows in result set") {
			t.Errorf("Error: %v", err3)
		}
	}
}

func TestGetProductInfo(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := model.RequestProductInfo{
		ProductID: 1,
	}
	resp, err := getProductInfo(req)
	if err != nil && err != sql.ErrNoRows || resp.Success != true || resp.Message != "" {
		t.Errorf("Error: %v", err)
	}
}

func TestAddOrder(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := model.RequestPlaceOrder{
		Order: model.Order{
			OrderNote: "test",
		},
		Items: []model.ProductItem{
			{
				ProductItemID:       1,
				ProductID:           1,
				OrderItemID:         1,
				ProductItemQuantity: 2,
				ProductItemPrice:    10,
			},
		},
		UserID: 1,
	}

	resp, err := addOrder(req)
	if err != nil && err != sql.ErrNoRows || resp.Success != true || resp.Message != "" {
		t.Errorf("Error: %v", err)
	}
}

func TestAddProduct(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := model.RequestAddProduct{
		Product: model.Product{
			ProductName:        "test",
			ProductDescription: "test",
			ProductPrice:       10,
		},
	}

	resp, err := addProduct(req)
	if err != nil && err != sql.ErrNoRows || resp.Success != true || resp.Message != "" {
		t.Errorf("Error: %v", err)
	}
}

func TestUpdateProduct(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := model.RequestUpdateProduct{
		Product: model.Product{
			ProductID:          1,
			ProductName:        "test",
			ProductDescription: "test",
			ProductPrice:       10,
		},
	}

	resp, err := updateProduct(req)
	if err != nil && err != sql.ErrNoRows || resp.Success != true || resp.Message != "" {
		t.Errorf("Error: %v", err)
	}
}

func TestAddUser(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := model.RequestAddUser{
		User: model.User{
			UserFullName: "test",
			UserEmail:    "test@test.com",
		},
	}

	resp, err := addUser(req)
	if err != nil && err != sql.ErrNoRows || resp.Success != true || resp.Message != "" {
		t.Errorf("Error: %v", err)
	}

}

func TestGetAllEvents(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	resp, err := getAllEvents()
	if err != nil && err != sql.ErrNoRows || resp.Success != true || resp.Message != "" {
		t.Errorf("Error: %v", err)
	}
}

func TestFillOrderPlacementEvents(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	var responseEventsInfo model.ResponseEventsInfo
	err := fillOrderPlacementEvents(&responseEventsInfo)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestFillUserRegisterEvents(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	var responseEventsInfo model.ResponseEventsInfo
	err := fillUserRegisterEvents(&responseEventsInfo)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestFillProductAddEvents(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	var responseEventsInfo model.ResponseEventsInfo
	err := fillProductAddEvents(&responseEventsInfo)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestFillUserLoginEvents(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	var responseEventsInfo model.ResponseEventsInfo
	err := fillUserLoginEvents(&responseEventsInfo)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestFillRoleAssignmentEvents(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	var responseEventsInfo model.ResponseEventsInfo
	err := fillRoleAssignmentEvents(&responseEventsInfo)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}
