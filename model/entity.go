package model

type User struct {
	UserID       int    `json:"userID"`
	UserFullName string `json:"userFullName"`
	UserEmail    string `json:"userEmail"`
}

type Order struct {
	OrderID   int    `json:"orderID"`
	OrderNote string `json:"orderNote"`
}

type OrderItem struct {
	OrderItemID int `json:"orderItemID"`
}

type Product struct {
	ProductID          int     `json:"productID"`
	ProductName        string  `json:"productName"`
	ProductPrice       float64 `json:"productPrice"`
	ProductDescription string  `json:"productDescription"`
}

type Password struct {
	PasswordID    int    `json:"passwordID"`
	EncryptedData string `json:"encryptedData"`
}

type Event struct {
	EventID        int    `json:"eventID"`
	EventTimestamp string `json:"eventTimestamp"`
}

type Role struct {
	RoleID   int    `json:"roleID"`
	RoleName string `json:"roleName"`
}
