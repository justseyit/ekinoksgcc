package model

type ResponseBase struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ResponseUserInfo struct {
	UserID       int    `json:"user_id"`
	UserFullName string `json:"user_full_name"`
	UserEmail    string `json:"user_email"`
	UserRole     string `json:"user_role"`
	ResponseBase
}

type ResponseProductInfo struct {
	Product Product `json:"product"`
	ResponseBase
}

type PlacedOrderInfo struct {
	OrderID   int    `json:"order_id"`
	OrderNote string `json:"order_note"`
	Items     []Item `json:"items"`
}

type Item struct {
	ProductItemID       int                 `json:"product_item_id"`
	ProductItemQuantity int                 `json:"product_item_quantity"`
	ProductItemPrice    float64             `json:"product_item_price"`
	ProductInfo         ResponseProductInfo `json:"product_info"`
}

type ResponseOrderInfo struct {
	Actor            ResponseUserInfo  `json:"actor"`
	PlacedOrdersInfo []PlacedOrderInfo `json:"placed_orders_info"`
	ResponseBase
}

type ResponseRoleAssignmentEventInfo struct {
	Timestamp    string           `json:"timestamp"`
	Actor        ResponseUserInfo `json:"actor"`
	AssignedUser ResponseUserInfo `json:"assigned_user"`
	ResponseBase
}

type ResponseUserLoginEventInfo struct {
	Timestamp string           `json:"timestamp"`
	Actor     ResponseUserInfo `json:"actor"`
	ResponseBase
}

type ResponseUserRegisterEventInfo struct {
	Timestamp string           `json:"timestamp"`
	Actor     ResponseUserInfo `json:"actor"`
	ResponseBase
}

type ResponseOrderPlacementEventInfo struct {
	Timestamp string            `json:"timestamp"`
	Order     ResponseOrderInfo `json:"order"`
	ResponseBase
}

type ResponseProductAddEventInfo struct {
	Timestamp string              `json:"timestamp"`
	Actor     ResponseUserInfo    `json:"actor"`
	Product   ResponseProductInfo `json:"product"`
	ResponseBase
}

type ResponseEventsInfo struct {
	RoleAssignmentEvents []ResponseRoleAssignmentEventInfo `json:"role_assignment_events,omitempty"`
	UserLoginEvents      []ResponseUserLoginEventInfo      `json:"user_login_events,omitempty"`
	UserRegisterEvents   []ResponseUserRegisterEventInfo   `json:"user_register_events,omitempty"`
	OrderPlacementEvents []ResponseOrderPlacementEventInfo `json:"order_placement_events,omitempty"`
	ProductAddEvents     []ResponseProductAddEventInfo     `json:"product_add_events,omitempty"`
	ResponseBase
}
