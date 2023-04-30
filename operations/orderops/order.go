package orderops

import (
	"ekinoksgcc/model"
	"ekinoksgcc/repository"
)

// Database operations for order placement and management

// Order operations

// Adds an Order
func AddOrderToDB(order model.Order) (int, error) {
	var id int
	err := repository.DB.QueryRow("INSERT INTO orders(orderNote) VALUES($1) RETURNING orderID", order.OrderNote).Scan(&id)

	return id, err
}

// Updates an Order by its ID. Do not change orderID field of the argument
func UpdateOrderInDB(order model.Order) (int, error) {
	_, err := repository.DB.Exec("UPDATE orders SET orderNote=$1 WHERE orderID=$2", order.OrderNote, order.OrderID)

	return order.OrderID, err
}

// Deletes an Order by its ID
func DeleteOrderFromDB(orderID int) (int, error) {
	_, err := repository.DB.Exec("DELETE FROM orders WHERE orderID=$1", orderID)

	return orderID, err
}

// Returns an Order by its ID
func GetOrderByOrderIDFromDB(orderID int) (model.Order, error) {
	var order model.Order
	err := repository.DB.QueryRow("SELECT orderID, orderNote FROM orders WHERE orderID=$1", order.OrderID).Scan(&order.OrderID, &order.OrderNote)
	return order, err
}

// Returns all Orders by their IDs
func GetAllOrdersByIDsFromDB(orderIDs []int) ([]model.Order, error) {
	orders := []model.Order{}
	for _, orderID := range orderIDs {
		order, err := GetOrderByOrderIDFromDB(orderID)
		if err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// Returns all Orders
func GetAllOrdersFromDB() ([]model.Order, error) {
	rows, err := repository.DB.Query("SELECT orderID, orderNote FROM orders")
	if err != nil {
		return []model.Order{}, err
	}
	defer rows.Close()

	orders := []model.Order{}
	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.OrderID, &order.OrderNote)
		if err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}
	return orders, err
}

// PlacedOrder operations

// Returns a PlacedOrder by its ID
func GetPlacedOrderByIDFromDB(placedOrderID int) (model.PlacedOrder, error) {
	var placedOrder model.PlacedOrder
	err := repository.DB.QueryRow("SELECT placedOrderID, orderID, productItemID, userID FROM placedOrder WHERE orderID=$1", placedOrderID).Scan(&placedOrder.PlacedOrderID, &placedOrder.ProductItemID, &placedOrder.UserID, &placedOrder.OrderID)
	if err != nil {
		return model.PlacedOrder{}, err
	}
	return placedOrder, err
}

// Returns all PlacedOrders by UserID
func GetPlacedOrdersByUserIDFromDB(userID int) ([]model.PlacedOrder, error) {
	rows, err := repository.DB.Query("SELECT placedOrderID, orderID, productItemID, userID FROM placedOrder WHERE userID=$1", userID)
	if err != nil {
		return []model.PlacedOrder{}, err
	}
	defer rows.Close()

	placedOrders := []model.PlacedOrder{}
	for rows.Next() {
		var placedOrder model.PlacedOrder
		err := rows.Scan(&placedOrder.PlacedOrderID, &placedOrder.OrderID, &placedOrder.ProductItemID, &placedOrder.UserID)
		if err != nil {
			return []model.PlacedOrder{}, err
		}
		placedOrders = append(placedOrders, placedOrder)
	}
	return placedOrders, err
}

// Returns PlacedOrder by OrderID
func GetPlacedOrderByOrderIDFromDB(orderID int) (model.PlacedOrder, error) {
	var placedOrder model.PlacedOrder
	err := repository.DB.QueryRow("SELECT placedOrderID, orderID, productItemID, userID FROM placedOrder WHERE orderID=$1", orderID).Scan(&placedOrder.PlacedOrderID, &placedOrder.OrderID, &placedOrder.ProductItemID, &placedOrder.UserID)
	if err != nil {
		return model.PlacedOrder{}, err
	}
	return placedOrder, err
}

// Returns all PlacedOrders
func GetAllPlacedOrdersFromDB() ([]model.PlacedOrder, error) {
	rows, err := repository.DB.Query("SELECT placedOrderID, orderID, productItemID, userID FROM placedOrder")
	if err != nil {
		return []model.PlacedOrder{}, err
	}
	defer rows.Close()

	placedOrders := []model.PlacedOrder{}
	for rows.Next() {
		var placedOrder model.PlacedOrder
		err := rows.Scan(&placedOrder.PlacedOrderID, &placedOrder.OrderID, &placedOrder.ProductItemID, &placedOrder.UserID)
		if err != nil {
			return placedOrders, err
		}
		placedOrders = append(placedOrders, placedOrder)
	}
	return placedOrders, err
}

// Updates a PlacedOrder by its ID. Do not change placedOrderID field of the argument
func UpdatePlacedOrderInDB(placedOrder model.PlacedOrder) (int, error) {
	_, err := repository.DB.Exec("UPDATE placedOrder SET orderID=$1, productItemID=$2, userID=$3 WHERE placedOrderID=$4", placedOrder.OrderID, placedOrder.ProductItemID, placedOrder.UserID, placedOrder.PlacedOrderID)
	if err != nil {
		return 0, err
	}
	return placedOrder.PlacedOrderID, err
}

// Deletes a PlacedOrder by its ID
func DeletePlacedOrderByIDFromDB(placedOrderID int) (int, error) {
	_, err := repository.DB.Exec("DELETE FROM placedOrder WHERE placedOrderID=$1", placedOrderID)
	if err != nil {
		return 0, err
	}
	return placedOrderID, err
}

// Adds a PlacedOrder
func AddPlacedOrderToDB(placedOrder model.PlacedOrder) (int, error) {
	var id int
	err := repository.DB.QueryRow("INSERT INTO placedOrder(orderID, productItemID, userID) VALUES($1, $2, $3) RETURNING placedOrderID", placedOrder.OrderID, placedOrder.ProductItemID, placedOrder.UserID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

// OrderItem operations

// Adds an OrderItem
func AddOrderItemToDB() (int, error) {
	var id int
	err := repository.DB.QueryRow("INSERT INTO orderItem() VALUES() RETURNING orderItemID").Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

// Returns an OrderItem by its ID
func GetOrderItemByIDFromDB(orderItemID int) (model.OrderItem, error) {
	var orderItem model.OrderItem
	err := repository.DB.QueryRow("SELECT orderItemID FROM orderItem WHERE orderItemID=$1", orderItemID).Scan(&orderItem.OrderItemID)
	if err != nil {
		return orderItem, err
	}
	return orderItem, err
}

// Returns all OrderItems
func GetAllOrderItemsFromDB() ([]model.OrderItem, error) {
	rows, err := repository.DB.Query("SELECT orderItemID FROM orderItem")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orderItems := []model.OrderItem{}
	for rows.Next() {
		var orderItem model.OrderItem
		err := rows.Scan(&orderItem.OrderItemID)
		if err != nil {
			return orderItems, err
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, err
}

// Deletes an OrderItem by its ID
func DeleteOrderItemByIDFromDB(orderItemID int) (int, error) {
	_, err := repository.DB.Exec("DELETE FROM orderItem WHERE orderItemID=$1", orderItemID)
	if err != nil {
		return orderItemID, err
	}
	return orderItemID, err
}

// ProductItem operations

// Returns a ProductItem by its ID
func GetProductItemByIDFromDB(productItemID int) (model.ProductItem, error) {
	var productItem model.ProductItem
	err := repository.DB.QueryRow("SELECT productItemID FROM productItem WHERE productItemID=$1", productItemID).Scan(&productItem.ProductItemID)
	if err != nil {
		return productItem, err
	}
	return productItem, err
}

// Returns all ProductItems by their ID
func GetAllProductItemsByIDFromDB(productItemIDs []int) ([]model.ProductItem, error) {
	productItems := []model.ProductItem{}
	for _, productItemID := range productItemIDs {
		productItem, err := GetProductItemByIDFromDB(productItemID)
		if err != nil {
			return productItems, err
		}
		productItems = append(productItems, productItem)
	}
	return productItems, nil
}

// Returns all ProductItems by their placedOrderID
func GetAllProductItemsByPlacedOrderIDFromDB(placedOrderID int) ([]model.ProductItem, error) {
	rows, err := repository.DB.Query("SELECT productItemID FROM productItem WHERE placedOrderID=$1", placedOrderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	productItems := []model.ProductItem{}
	for rows.Next() {
		var productItem model.ProductItem
		err := rows.Scan(&productItem.ProductItemID)
		if err != nil {
			return productItems, err
		}
		productItems = append(productItems, productItem)
	}
	return productItems, err
}

// Returns all ProductItems
func GetAllProductItemsFromDB() ([]model.ProductItem, error) {
	rows, err := repository.DB.Query("SELECT productItemID FROM productItem")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	productItems := []model.ProductItem{}
	for rows.Next() {
		var productItem model.ProductItem
		err := rows.Scan(&productItem.ProductItemID)
		if err != nil {
			return productItems, err
		}
		productItems = append(productItems, productItem)
	}
	return productItems, err
}

// Deletes a ProductItem by its ID
func DeleteProductItemByIDFromDB(productItemID int) (int, error) {
	_, err := repository.DB.Exec("DELETE FROM productItem WHERE productItemID=$1", productItemID)

	return productItemID, err
}

// Adds a ProductItem
func AddProductItemToDB(productItem model.ProductItem) (int, error) {
	var id int
	err := repository.DB.QueryRow("INSERT INTO productItem(productID, orderItemID, productItemPrice, productItemquantity) VALUES($1, $2, $3, $4) RETURNING productItemID", productItem.ProductID, productItem.OrderItemID, productItem.ProductItemPrice, productItem.ProductItemQuantity).Scan(&id)

	return id, err
}

// Updates a ProductItem by its ID. Do not change productItemID field of the argument
func UpdateProductItemInDB(productItem model.ProductItem) (int, error) {
	_, err := repository.DB.Exec("UPDATE productItem SET productID=$1, orderItemID=$2, productItemPrice=$3, productItemquantity=$4 WHERE productItemID=$5", productItem.ProductID, productItem.OrderItemID, productItem.ProductItemPrice, productItem.ProductItemQuantity, productItem.ProductItemID)

	return productItem.ProductItemID, err
}
