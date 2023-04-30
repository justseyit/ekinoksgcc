package api

import (
	"ekinoksgcc/model"
	"ekinoksgcc/operations/eventops"
	"ekinoksgcc/operations/orderops"
	"ekinoksgcc/operations/productops"
	"ekinoksgcc/operations/userops"
	"log"
)

// If userID is specified, return order info for that user
// If userID is not specified, return order info for orderID
func getOrderInfo(req model.RequestOrderInfo) (model.ResponseOrderInfo, error) {
	var orderInfo model.ResponseOrderInfo
	var responseUserInfo model.ResponseUserInfo
	var placedOrders []model.PlacedOrder
	var products []model.Product
	var orderIDs []int
	var productItemIDs []int

	if req.UserID != 0 {
		// User ID is specified. Return order info for that user
		user, err5 := userops.GetUserFromDB(req.UserID)
		if err5 != nil {
			orderInfo.Success = false
			orderInfo.Message = "User not found"
			return orderInfo, err5
		}

		role, err4 := userops.GetRoleByUserIDFromDB(req.UserID)
		if err4 != nil {
			orderInfo.Success = false
			orderInfo.Message = "Role not found"
			return orderInfo, err4
		}

		responseUserInfo.UserID = user.UserID
		responseUserInfo.UserFullName = user.UserFullName
		responseUserInfo.UserEmail = user.UserEmail
		responseUserInfo.UserRole = role.RoleName

		placedOrders, err := orderops.GetPlacedOrdersByUserIDFromDB(req.UserID)
		if err != nil {
			orderInfo.Success = false
			orderInfo.Message = "Placed orders not found"
			return orderInfo, err
		}

		for _, placedOrder := range placedOrders {
			orderIDs = append(orderIDs, placedOrder.OrderID)
			productItemIDs = append(productItemIDs, placedOrder.ProductItemID)
		}

		orders, err1 := orderops.GetAllOrdersByIDsFromDB(orderIDs)
		if err1 != nil {
			orderInfo.Success = false
			orderInfo.Message = "Orders not found"
			return orderInfo, err1
		}

		productItems, err2 := orderops.GetAllProductItemsByIDFromDB(productItemIDs)
		if err2 != nil {
			orderInfo.Success = false
			orderInfo.Message = "Product items not found"
			return orderInfo, err2
		}

		for _, productItem := range productItems {
			product, err := productops.GetProductByProductIDFromDB(productItem.ProductID)
			if err != nil {
				orderInfo.Success = false
				orderInfo.Message = "Product not found"
				return orderInfo, err
			}
			products = append(products, product)
		}

		for _, order := range orders {
			for _, placedOrder := range placedOrders {
				if order.OrderID == placedOrder.OrderID {
					for _, productItem := range productItems {
						if productItem.ProductItemID == placedOrder.ProductItemID {
							orderInfo.Actor = responseUserInfo
							items := []model.Item{}
							for _, product := range products {
								if product.ProductID == productItem.ProductID {
									items = append(items, model.Item{
										ProductItemID:       productItem.ProductItemID,
										ProductItemQuantity: productItem.ProductItemQuantity,
										ProductItemPrice:    productItem.ProductItemPrice,
										ProductInfo: model.ResponseProductInfo{
											Product: product,
										},
									})
								}
							}
							orderInfo.PlacedOrdersInfo = append(orderInfo.PlacedOrdersInfo, model.PlacedOrderInfo{
								OrderID:   order.OrderID,
								OrderNote: order.OrderNote,
								Items:     items,
							})
						}
					}
				}
			}
		}

		orderInfo.Success = true
		orderInfo.Message = ""
		return orderInfo, nil
	} else {
		// User ID is not specified. Return order info for orderID

		placedOrder, err6 := orderops.GetPlacedOrderByOrderIDFromDB(req.OrderID)
		if err6 != nil {
			orderInfo.Success = false
			orderInfo.Message = "Placed order not found"
			return orderInfo, err6
		}

		user, err5 := userops.GetUserFromDB(placedOrder.UserID)
		if err5 != nil {
			orderInfo.Success = false
			orderInfo.Message = "User not found"
			return orderInfo, err5
		}

		role, err4 := userops.GetRoleByUserIDFromDB(placedOrder.UserID)
		if err4 != nil {
			orderInfo.Success = false
			orderInfo.Message = "Role not found"
			return orderInfo, err4
		}

		responseUserInfo.UserID = user.UserID
		responseUserInfo.UserFullName = user.UserFullName
		responseUserInfo.UserEmail = user.UserEmail
		responseUserInfo.UserRole = role.RoleName

		for _, placedOrder := range placedOrders {
			orderIDs = append(orderIDs, placedOrder.OrderID)
			productItemIDs = append(productItemIDs, placedOrder.ProductItemID)
		}

		orders, err1 := orderops.GetAllOrdersByIDsFromDB(orderIDs)
		if err1 != nil {
			orderInfo.Success = false
			orderInfo.Message = "Orders not found"
			return orderInfo, err1
		}

		productItems, err2 := orderops.GetAllProductItemsByIDFromDB(productItemIDs)
		if err2 != nil {
			orderInfo.Success = false
			orderInfo.Message = "Product items not found"
			return orderInfo, err2
		}

		for _, productItem := range productItems {
			productItemIDs = append(productItemIDs, productItem.ProductItemID)
		}

		for _, order := range orders {
			for _, placedOrder := range placedOrders {
				if order.OrderID == placedOrder.OrderID {
					for _, productItem := range productItems {
						if productItem.ProductItemID == placedOrder.ProductItemID {
							orderInfo.Actor = responseUserInfo
							items := []model.Item{}
							for _, product := range products {
								if product.ProductID == productItem.ProductID {
									items = append(items, model.Item{
										ProductItemID:       productItem.ProductItemID,
										ProductItemQuantity: productItem.ProductItemQuantity,
										ProductItemPrice:    productItem.ProductItemPrice,
										ProductInfo: model.ResponseProductInfo{
											Product: product,
										},
									})
								}
							}
							orderInfo.PlacedOrdersInfo = append(orderInfo.PlacedOrdersInfo, model.PlacedOrderInfo{
								OrderID:   order.OrderID,
								OrderNote: order.OrderNote,
								Items:     items,
							})
						}
					}
				}
			}
		}

		orderInfo.Success = true
		orderInfo.Message = ""
		return orderInfo, nil
	}
}

func getProductInfo(req model.RequestProductInfo) (model.ResponseProductInfo, error) {
	var productInfo model.ResponseProductInfo
	product, err := productops.GetProductByProductIDFromDB(req.ProductID)
	if err != nil {
		productInfo.Success = false
		productInfo.Message = "Product not found"
		return productInfo, err
	}
	productInfo.Success = true
	productInfo.Message = ""
	productInfo.Product = product
	return productInfo, nil
}

func addOrder(req model.RequestPlaceOrder) (model.ResponseOrderInfo, error) {
	var responseOrderInfo model.ResponseOrderInfo
	var responseUserInfo model.ResponseUserInfo

	responseOrderInfo.PlacedOrdersInfo = []model.PlacedOrderInfo{}

	user, err_ := userops.GetUserFromDB(req.UserID)
	if err_ != nil {
		responseOrderInfo.Success = false
		responseOrderInfo.Message = "User not found"
		return responseOrderInfo, err_
	}

	role, err0 := userops.GetRoleByUserIDFromDB(req.UserID)
	if err0 != nil {
		responseOrderInfo.Success = false
		responseOrderInfo.Message = "Role not found"
		return responseOrderInfo, err0
	}

	responseUserInfo = model.ResponseUserInfo{
		UserID:       user.UserID,
		UserFullName: user.UserFullName,
		UserEmail:    user.UserEmail,
		UserRole:     role.RoleName,
	}

	orderID, err := orderops.AddOrderToDB(req.Order)
	if err != nil {
		responseOrderInfo.Success = false
		responseOrderInfo.Message = "Order not added"
		return responseOrderInfo, err
	}

	_, err1 := orderops.AddOrderItemToDB()
	if err1 != nil {
		responseOrderInfo.Success = false
		responseOrderInfo.Message = "Order item not added"
		return responseOrderInfo, err1
	}

	var items []model.Item
	var placedOrderID int

	for _, item := range req.Items {
		_, err2 := orderops.AddProductItemToDB(item)
		if err2 != nil {
			responseOrderInfo.Success = false
			responseOrderInfo.Message = "Product item not added"
			return responseOrderInfo, err2
		}

		var err3 error

		placedOrderID, err3 = orderops.AddPlacedOrderToDB(model.PlacedOrder{
			OrderID:       orderID,
			UserID:        req.UserID,
			ProductItemID: item.ProductItemID,
		})
		if err3 != nil {
			responseOrderInfo.Success = false
			responseOrderInfo.Message = "Placed order not added"
			return responseOrderInfo, err3
		}

		product, err4 := productops.GetProductByProductIDFromDB(item.ProductID)
		if err4 != nil {
			responseOrderInfo.Success = false
			responseOrderInfo.Message = "Product not found"
			return responseOrderInfo, err4
		}

		items = append(items, model.Item{
			ProductItemID:       item.ProductItemID,
			ProductItemQuantity: item.ProductItemQuantity,
			ProductItemPrice:    item.ProductItemPrice,
			ProductInfo: model.ResponseProductInfo{
				Product: product,
			},
		})
	}

	eventops.AddOrderPlacementEventToDB(placedOrderID)

	responseOrderInfo.Success = true
	responseOrderInfo.Message = ""
	responseOrderInfo.Actor = responseUserInfo
	responseOrderInfo.PlacedOrdersInfo = append(responseOrderInfo.PlacedOrdersInfo, model.PlacedOrderInfo{
		OrderID:   req.Order.OrderID,
		OrderNote: req.Order.OrderNote,
		Items:     items,
	})
	return responseOrderInfo, nil
}

func addProduct(req model.RequestAddProduct) (model.ResponseProductInfo, error) {
	var responseProductInfo model.ResponseProductInfo
	productID, err := productops.AddProductToDB(req.Product)
	if err != nil {
		responseProductInfo.Success = false
		responseProductInfo.Message = "Product not added"
		return responseProductInfo, err
	}
	responseProductInfo.Success = true
	responseProductInfo.Message = ""
	responseProductInfo.Product = model.Product{
		ProductID:          productID,
		ProductName:        req.Product.ProductName,
		ProductDescription: req.Product.ProductDescription,
		ProductPrice:       req.Product.ProductPrice,
	}
	return responseProductInfo, nil
}

func updateProduct(req model.RequestUpdateProduct) (model.ResponseProductInfo, error) {
	var responseProductInfo model.ResponseProductInfo
	productID, err := productops.UpdateProductInDB(req.Product)
	if err != nil {
		responseProductInfo.Success = false
		responseProductInfo.Message = "Product not updated"
		return responseProductInfo, err
	}
	responseProductInfo.Success = true
	responseProductInfo.Message = ""
	responseProductInfo.Product = model.Product{
		ProductID:          productID,
		ProductName:        req.Product.ProductName,
		ProductDescription: req.Product.ProductDescription,
		ProductPrice:       req.Product.ProductPrice,
	}
	return responseProductInfo, nil
}

func addUser(req model.RequestAddUser) (model.ResponseUserInfo, error) {
	var responseUserInfo model.ResponseUserInfo
	userID, err := userops.CreateUserInDB(req.User)
	if err != nil {
		responseUserInfo.Success = false
		responseUserInfo.Message = "User not added"
		return responseUserInfo, err
	}
	responseUserInfo.Success = true
	responseUserInfo.Message = ""
	responseUserInfo = model.ResponseUserInfo{
		UserID:       userID,
		UserFullName: req.User.UserFullName,
		UserEmail:    req.User.UserEmail,
		UserRole:     req.Role.RoleName,
	}
	return responseUserInfo, nil
}

func getAllEvents() (model.ResponseEventsInfo, error) {
	var responseEventsInfo model.ResponseEventsInfo
	err0 := fillOrderPlacementEvents(&responseEventsInfo)
	if err0 != nil {
		responseEventsInfo.Success = false
		responseEventsInfo.Message = err0.Error()
		return responseEventsInfo, err0
	}
	err1 := fillUserRegisterEvents(&responseEventsInfo)
	if err1 != nil {
		responseEventsInfo.Success = false
		responseEventsInfo.Message = err1.Error()
		return responseEventsInfo, err1
	}
	err2 := fillUserLoginEvents(&responseEventsInfo)
	if err2 != nil {
		responseEventsInfo.Success = false
		responseEventsInfo.Message = err2.Error()
		return responseEventsInfo, err2
	}
	err3 := fillProductAddEvents(&responseEventsInfo)
	if err3 != nil {
		responseEventsInfo.Success = false
		responseEventsInfo.Message = err3.Error()
		return responseEventsInfo, err3
	}
	err4 := fillRoleAssignmentEvents(&responseEventsInfo)
	if err4 != nil {
		responseEventsInfo.Success = false
		responseEventsInfo.Message = err4.Error()
		return responseEventsInfo, err4
	}

	return responseEventsInfo, nil
}

func fillOrderPlacementEvents(responseEventsInfo *model.ResponseEventsInfo) error {
	orderPlacementEvents, err := eventops.GetAllOrderPlacementEventsFromDB()
	if err != nil {
		responseEventsInfo.Success = false
		responseEventsInfo.Message = "Events not found"
		return err
	}

	for _, event := range orderPlacementEvents {
		log.Println(event)
		eventForOrder, err1 := eventops.GetEventByEventID(event.EventID)
		if err1 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "Event not found"
			return err1
		}
		placedOrder, err2 := orderops.GetPlacedOrderByIDFromDB(event.PlacedOrderID)
		if err2 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "Placed order not found"
			return err2
		}
		order, err3 := orderops.GetOrderByOrderIDFromDB(placedOrder.OrderID)
		if err3 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "Order not found"
			return err3
		}
		user, err4 := userops.GetUserFromDB(placedOrder.UserID)
		if err4 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "User not found"
			return err4
		}

		info := model.ResponseOrderPlacementEventInfo{
			Timestamp: eventForOrder.EventTimestamp,
			Order: model.ResponseOrderInfo{
				Actor: model.ResponseUserInfo{
					UserID:       user.UserID,
					UserFullName: user.UserFullName,
					UserEmail:    user.UserEmail,
				},
				PlacedOrdersInfo: make([]model.PlacedOrderInfo, 0),
			},
		}

		info.Order.PlacedOrdersInfo = append(info.Order.PlacedOrdersInfo, model.PlacedOrderInfo{
			OrderID:   order.OrderID,
			OrderNote: order.OrderNote,
			Items:     make([]model.Item, 0),
		})

		productItems, err5 := orderops.GetAllProductItemsByPlacedOrderIDFromDB(placedOrder.PlacedOrderID)
		if err5 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "Product items not found"
			return err5
		}

		for _, productItem := range productItems {
			product, err6 := productops.GetProductByProductIDFromDB(productItem.ProductID)
			if err6 != nil {
				responseEventsInfo.Success = false
				responseEventsInfo.Message = "Product not found"
				return err6
			}
			info.Order.PlacedOrdersInfo[0].Items = append(info.Order.PlacedOrdersInfo[0].Items, model.Item{
				ProductItemID:       productItem.ProductItemID,
				ProductItemQuantity: productItem.ProductItemQuantity,
				ProductItemPrice:    productItem.ProductItemPrice,
				ProductInfo: model.ResponseProductInfo{
					Product: model.Product{
						ProductID:          product.ProductID,
						ProductName:        product.ProductName,
						ProductDescription: product.ProductDescription,
						ProductPrice:       product.ProductPrice,
					},
				},
			})
		}

		responseEventsInfo.OrderPlacementEvents = append(responseEventsInfo.OrderPlacementEvents, info)
	}

	return nil
}

func fillUserRegisterEvents(responseEventsInfo *model.ResponseEventsInfo) error {
	userRegisterEvents, err := eventops.GetAllUserRegisterEventsFromDB()
	if err != nil {
		responseEventsInfo.Success = false
		responseEventsInfo.Message = "Events not found"
		return err
	}

	for _, event := range userRegisterEvents {
		eventForUser, err1 := eventops.GetEventByEventID(event.EventID)
		if err1 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "Event not found"
			return err1
		}
		user, err2 := userops.GetUserFromDB(event.UserID)
		if err2 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "User not found"
			return err2
		}

		info := model.ResponseUserRegisterEventInfo{
			Timestamp: eventForUser.EventTimestamp,
			Actor: model.ResponseUserInfo{
				UserID:       user.UserID,
				UserFullName: user.UserFullName,
				UserEmail:    user.UserEmail,
			},
		}

		responseEventsInfo.UserRegisterEvents = append(responseEventsInfo.UserRegisterEvents, info)
	}

	return nil
}

func fillUserLoginEvents(responseEventsInfo *model.ResponseEventsInfo) error {
	userLoginEvents, err := eventops.GetAllUserLoginEventsFromDB()
	if err != nil {
		responseEventsInfo.Success = false
		responseEventsInfo.Message = "Events not found"
		return err
	}

	for _, event := range userLoginEvents {
		eventForUser, err1 := eventops.GetEventByEventID(event.EventID)
		if err1 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "Event not found"
			return err1
		}
		user, err2 := userops.GetUserFromDB(event.UserID)
		if err2 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "User not found"
			return err2
		}

		info := model.ResponseUserLoginEventInfo{
			Timestamp: eventForUser.EventTimestamp,
			Actor: model.ResponseUserInfo{
				UserID:       user.UserID,
				UserFullName: user.UserFullName,
				UserEmail:    user.UserEmail,
			},
		}

		responseEventsInfo.UserLoginEvents = append(responseEventsInfo.UserLoginEvents, info)
	}

	return nil
}

func fillProductAddEvents(responseEventsInfo *model.ResponseEventsInfo) error {
	productAddEvents, err := eventops.GetAllProductAddEventsFromDB()
	if err != nil {
		responseEventsInfo.Success = false
		responseEventsInfo.Message = "Events not found"
		return err
	}

	for _, event := range productAddEvents {
		eventForProduct, err1 := eventops.GetEventByEventID(event.EventID)
		if err1 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "Event not found"
			return err1
		}
		addedProduct, err2 := productops.GetAddedProductMappingByAddedProductMappingIDFromDB(event.AddedProductID)
		if err2 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "Added product not found"
			return err2
		}

		user, err3 := userops.GetUserFromDB(addedProduct.UserID)
		if err3 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "User not found"
			return err3
		}

		role, err4 := userops.GetRoleByUserIDFromDB(user.UserID)
		if err4 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "Role not found"
			return err4
		}

		product, err5 := productops.GetProductByProductIDFromDB(addedProduct.ProductID)
		if err5 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "Product not found"
			return err5
		}

		info := model.ResponseProductAddEventInfo{
			Timestamp: eventForProduct.EventTimestamp,
			Actor: model.ResponseUserInfo{
				UserID:       user.UserID,
				UserFullName: user.UserFullName,
				UserEmail:    user.UserEmail,
				UserRole:     role.RoleName,
			},
			Product: model.ResponseProductInfo{
				Product: model.Product{
					ProductID:          product.ProductID,
					ProductName:        product.ProductName,
					ProductDescription: product.ProductDescription,
					ProductPrice:       product.ProductPrice,
				},
			},
		}

		responseEventsInfo.ProductAddEvents = append(responseEventsInfo.ProductAddEvents, info)
	}

	return nil
}

func fillRoleAssignmentEvents(responseEventsInfo *model.ResponseEventsInfo) error {
	roleAssignmentEvents, err := eventops.GetAllRoleAssignmentEventsFromDB()
	if err != nil {
		responseEventsInfo.Success = false
		responseEventsInfo.Message = "Events not found"
		return err
	}

	for _, event := range roleAssignmentEvents {
		eventForRoleAssignment, err1 := eventops.GetEventByEventID(event.EventID)
		if err1 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "Event not found"
			return err1
		}
		user, err2 := userops.GetUserFromDB(event.UserID)
		if err2 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "User not found"
			return err2
		}

		assignedUserRole, err3 := userops.GetUserRoleFromDB(event.UserRoleID)
		if err3 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "User role not found"
			return err3
		}

		assignedUser, err4 := userops.GetUserFromDB(assignedUserRole.UserID)
		if err4 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "Assigned user not found"
			return err4
		}

		assignedRole, err5 := userops.GetRoleByUserIDFromDB(assignedUser.UserID)
		if err5 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "Role not found"
			return err5
		}

		role, err6 := userops.GetRoleByUserIDFromDB(event.UserID)
		if err6 != nil {
			responseEventsInfo.Success = false
			responseEventsInfo.Message = "Role not found"
			return err6
		}

		info := model.ResponseRoleAssignmentEventInfo{
			Timestamp: eventForRoleAssignment.EventTimestamp,
			Actor: model.ResponseUserInfo{
				UserID:       user.UserID,
				UserFullName: user.UserFullName,
				UserEmail:    user.UserEmail,
				UserRole:     role.RoleName,
			},
			AssignedUser: model.ResponseUserInfo{
				UserID:       assignedUser.UserID,
				UserFullName: assignedUser.UserFullName,
				UserEmail:    assignedUser.UserEmail,
				UserRole:     assignedRole.RoleName,
			},
		}

		responseEventsInfo.RoleAssignmentEvents = append(responseEventsInfo.RoleAssignmentEvents, info)
	}

	return nil
}
