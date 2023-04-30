package model

type RequestUserInfo struct {
	UserID int `json:"user_id"`
}

type RequestProductInfo struct {
	ProductID int `json:"product_id"`
}

type RequestRoleInfo struct {
	RoleID int `json:"role_id"`
}

// If UserID has specified, than OrderID is ignored
type RequestOrderInfo struct {
	OrderID int `json:"order_id"`
	UserID  int `json:"user_id"`
}

// If UserID has specified, than UserRoleID is ignored
type RequestRoleAssignmentEventInfo struct {
	UserRoleID int `json:"user_role_id"`
	UserID     int `json:"user_id"`
}

type RequestUserLoginEventInfo struct {
	UserID int `json:"user_id"`
}

// If UserID has specified, than PlacedOrderID is ignored
type RequestOrderPlacementEventInfo struct {
	PlacedOrderID int `json:"placed_order_id"`
	UserID        int `json:"user_id"`
}

// If UserID has specified, than ProductID is ignored
type RequestProductAddEventInfo struct {
	ProductID int `json:"product_id"`
	UserID    int `json:"user_id"`
}

type RequestAddUser struct {
	User User `json:"user"`
	Role Role `json:"role"`
}

type RequestAddProduct struct {
	Product Product `json:"product"`
}

type RequestUpdateProduct struct {
	Product Product `json:"product"`
}

type RequestPlaceOrder struct {
	Order  Order         `json:"order"`
	Items  []ProductItem `json:"items"`
	UserID int           `json:"user_id"`
}

type RequestAssignRole struct {
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

type RequestUserLogin struct {
	UserEmail string `json:"user_email"`
	Password  string `json:"password"`
}

type RequestUserRegister struct {
	User                 User   `json:"user"`
	Password             string `json:"password"`
	PasswordVerification string `json:"password_verification"`
}
