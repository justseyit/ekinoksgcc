package productops

import (
	"ekinoksgcc/model"
	"ekinoksgcc/repository"
	"log"
)

// Database operations for product placement and management

// Product operations

// Adds a Product
func AddProductToDB(product model.Product) (int, error) {
	var id int
	err := repository.DB.QueryRow("INSERT INTO product(productName, productPrice, productDescription) VALUES($1, $2, $3) RETURNING productID", product.ProductName, product.ProductPrice, product.ProductDescription).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id, err
}

// Updates a Product by its ID. Do not change productID field of the argument
func UpdateProductInDB(product model.Product) (int, error) {
	_, err := repository.DB.Exec("UPDATE product SET productName=$1, productPrice=$2, productDescription=$3 WHERE productID=$4", product.ProductName, product.ProductPrice, product.ProductDescription, product.ProductID)
	if err != nil {
		log.Fatal(err)
	}
	return product.ProductID, err
}

// Deletes a Product by its ID
func DeleteProductFromDB(productID int) (int, error) {
	_, err := repository.DB.Exec("DELETE FROM product WHERE productID=$1", productID)
	if err != nil {
		log.Fatal(err)
	}
	return productID, err
}

// Returns a Product by its ID
func GetProductByProductIDFromDB(productID int) (model.Product, error) {
	var product model.Product
	err := repository.DB.QueryRow("SELECT productID, productName, productPrice, productDescription FROM product WHERE productID=$1", productID).Scan(&product.ProductID, &product.ProductName, &product.ProductPrice, &product.ProductDescription)
	if err != nil {
		log.Fatal(err)
	}
	return product, err
}

// Returns all Products by their IDs
func GetProductsByIDsFromDB(productIDs []int) ([]model.Product, error) {
	products := []model.Product{}
	for _, productID := range productIDs {
		product, err := GetProductByProductIDFromDB(productID)
		if err != nil {
			log.Fatal(err)
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
}

// Returns all Products
func GetAllProductsFromDB() ([]model.Product, error) {
	products := []model.Product{}
	rows, err := repository.DB.Query("SELECT productID, productName, productPrice, productDescription FROM product")
	if err != nil {
		log.Fatal(err)
		return products, err
	}
	defer rows.Close()
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ProductID, &product.ProductName, &product.ProductPrice, &product.ProductDescription)
		if err != nil {
			log.Fatal(err)
			return products, err
		}
		products = append(products, product)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return products, err
	}
	return products, nil
}

func GetAddedProductMappingsByUserIDFromDB(userID int) ([]model.AddedProduct, error) {
	addedProducts := []model.AddedProduct{}
	rows, err := repository.DB.Query("SELECT addedProductID, userID, productID FROM addedProduct WHERE userID=$1", userID)
	if err != nil {
		log.Fatal(err)
		return addedProducts, err
	}
	defer rows.Close()
	for rows.Next() {
		var addedProduct model.AddedProduct
		err := rows.Scan(&addedProduct.ProductID, &addedProduct.UserID, &addedProduct.AddedProductID)
		if err != nil {
			log.Fatal(err)
			return addedProducts, err
		}
		addedProducts = append(addedProducts, addedProduct)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return addedProducts, err
	}
	return addedProducts, nil
}

func GetAddedProductMappingByProductIDFromDB(productID int) (model.AddedProduct, error) {
	var addedProduct model.AddedProduct
	err := repository.DB.QueryRow("SELECT addedProductID, userID, productID FROM addedProduct WHERE productID=$1", productID).Scan(&addedProduct.ProductID, &addedProduct.UserID, &addedProduct.AddedProductID)
	if err != nil {
		log.Fatal(err)
	}
	return addedProduct, err
}

func GetAddedProductMappingByAddedProductMappingIDFromDB(addedProductID int) (model.AddedProduct, error) {
	var addedProduct model.AddedProduct
	err := repository.DB.QueryRow("SELECT addedProductID, userID, productID FROM addedProduct WHERE addedProductID=$1", addedProductID).Scan(&addedProduct.ProductID, &addedProduct.UserID, &addedProduct.AddedProductID)
	if err != nil {
		log.Fatal(err)
	}
	return addedProduct, err
}

func GetAllAddedProductMappingsFromDB() ([]model.AddedProduct, error) {
	addedProducts := []model.AddedProduct{}
	rows, err := repository.DB.Query("SELECT addedProductID, userID, productID FROM addedProduct")
	if err != nil {
		log.Fatal(err)
		return addedProducts, err
	}
	defer rows.Close()
	for rows.Next() {
		var addedProduct model.AddedProduct
		err := rows.Scan(&addedProduct.ProductID, &addedProduct.UserID, &addedProduct.AddedProductID)
		if err != nil {
			log.Fatal(err)
			return addedProducts, err
		}
		addedProducts = append(addedProducts, addedProduct)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return addedProducts, err
	}
	return addedProducts, nil
}
