package productops

import (
	"ekinoksgcc/model"
	"testing"
)

func TestAddProductToDB(t *testing.T) {
	req := model.Product{
		ProductID:          1,
		ProductName:        "test",
		ProductPrice:       1,
		ProductDescription: "test",
	}
	_, err := AddProductToDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestUpdateProductInDB(t *testing.T) {
	req := model.Product{
		ProductID:          1,
		ProductName:        "test",
		ProductPrice:       1,
		ProductDescription: "test",
	}
	_, err := UpdateProductInDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestGetProductByProductIDFromDB(t *testing.T) {
	req := 1
	_, err := GetProductByProductIDFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
