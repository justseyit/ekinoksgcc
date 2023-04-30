package model

type UserPassword struct {
	UserPasswordID int `json:"userPasswordID"`
	UserID         int `json:"userID"`
	PasswordID     int `json:"passwordID"`
}

type PlacedOrder struct {
	PlacedOrderID int `json:"placedOrderID"`
	OrderID       int `json:"orderID"`
	UserID        int `json:"userID"`
	ProductItemID int `json:"productItemID"`
}

type ProductItem struct {
	ProductItemID       int     `json:"productItemID"`
	ProductID           int     `json:"productID"`
	OrderItemID         int     `json:"orderItemID"`
	ProductItemQuantity int     `json:"productItemQuantity"`
	ProductItemPrice    float64 `json:"productItemPrice"`
}

type AddedProduct struct {
	AddedProductID int `json:"addedProductID"`
	ProductID      int `json:"productID"`
	UserID         int `json:"userID"`
}

type UserRole struct {
	UserRoleID int `json:"userRoleID"`
	UserID     int `json:"userID"`
	RoleID     int `json:"roleID"`
}

type OrderPlacementEvent struct {
	OrderPlacementEventID int `json:"orderPlacementEventID"`
	EventID               int `json:"eventID"`
	PlacedOrderID         int `json:"placedOrderID"`
}

type UserRegisterEvent struct {
	UserRegisterEventID int `json:"userRegisterEventID"`
	EventID             int `json:"eventID"`
	UserID              int `json:"userID"`
}

type UserLoginEvent struct {
	UserLoginEventID int `json:"userLoginEventID"`
	EventID          int `json:"eventID"`
	UserID           int `json:"userID"`
}

type RoleAssignmentEvent struct {
	RoleAssignmentEventID int `json:"roleAssignmentEventID"`
	EventID               int `json:"eventID"`
	UserID                int `json:"userID"`
	UserRoleID            int `json:"userRoleID"`
}

type ProductAddEvent struct {
	ProductAddEventID int `json:"productAddEventID"`
	EventID           int `json:"eventID"`
	AddedProductID    int `json:"addedProductID"`
}
