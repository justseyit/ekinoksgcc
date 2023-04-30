package api

import (
	"database/sql"
	"ekinoksgcc/model"
	"encoding/json"
	"log"
	"net/http"
	"net/mail"

	"ekinoksgcc/auth"
)

// Endpoints

// Adds a new user to the system as a customer
func HandlerAddUser(w http.ResponseWriter, r *http.Request) {

	//TODO: Check if current user is admin
	var requestAddUser model.RequestAddUser
	var responseUserInfo model.ResponseUserInfo
	err := json.NewDecoder(r.Body).Decode(&requestAddUser)
	if err != nil {
		log.Print(err)
		responseUserInfo.Success = false
		responseUserInfo.Message = "Invalid JSON format"
		w.WriteHeader(http.StatusBadRequest)
	}

	_, mailErr := mail.ParseAddress(requestAddUser.User.UserEmail)
	if mailErr != nil {
		log.Print(mailErr)
		responseUserInfo.Success = false
		responseUserInfo.Message = "Invalid email address"
		w.WriteHeader(http.StatusBadRequest)
	}

	if requestAddUser.User.UserFullName == "" || requestAddUser.User.UserEmail == "" {
		log.Print("User full name or email cannot be empty")
		responseUserInfo.Success = false
		responseUserInfo.Message = "User full name or email cannot be empty"
		w.WriteHeader(http.StatusBadRequest)
	}

	responseUserInfo, erro := addUser(requestAddUser)
	if erro != nil {
		log.Print(erro)
		responseUserInfo.Success = false
		responseUserInfo.Message = erro.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseUserInfo)
}

// For login requests for both customers and admins
func HandlerUserLogin(w http.ResponseWriter, r *http.Request) {
	var requestLogin model.RequestUserLogin
	var responseUserInfo model.ResponseUserInfo
	err := json.NewDecoder(r.Body).Decode(&requestLogin)
	if err != nil {
		log.Print(err)
		responseUserInfo.Success = false
		responseUserInfo.Message = "Invalid JSON format"
		w.WriteHeader(http.StatusBadRequest)
	}

	token, ex, err1 := auth.UserLogin(requestLogin)
	if err1 != nil {
		log.Print(err1)
		responseUserInfo.Success = false
		responseUserInfo.Message = err1.Error()
	}

	responseUserInfo.Success = true
	responseUserInfo.Message = ""

	w.Header().Set("Content-Type", "application/json")
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: ex,
	})
	json.NewEncoder(w).Encode(responseUserInfo)
}

// For register requests for customers
func HandlerUserRegister(w http.ResponseWriter, r *http.Request) {
	var requestRegister model.RequestUserRegister
	var responseUserInfo model.ResponseUserInfo
	err := json.NewDecoder(r.Body).Decode(&requestRegister)
	if err != nil {
		log.Print(err)
		responseUserInfo.Success = false
		responseUserInfo.Message = "Invalid JSON format"
	}

	token, ex, err1 := auth.UserSignup(requestRegister)
	if err1 != nil {
		log.Print(err1)
		responseUserInfo.Success = false
		responseUserInfo.Message = err1.Error()
	}

	responseUserInfo.Success = true
	responseUserInfo.Message = ""

	w.Header().Set("Content-Type", "application/json")
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: ex,
	})
	json.NewEncoder(w).Encode(responseUserInfo)
}

// For logout requests for both customers and admins
func HandlerUserLogout(w http.ResponseWriter, r *http.Request) {
	var responseBase model.ResponseBase
	auth.UserLogout(w)

	responseBase.Success = true
	responseBase.Message = ""

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBase)
}

// For getting orders of a customer
func HandlerGetUserOrders(w http.ResponseWriter, r *http.Request) {
	var requestGetUserOrders model.RequestOrderInfo
	var responseGetUserOrders model.ResponseOrderInfo
	err := json.NewDecoder(r.Body).Decode(&requestGetUserOrders)
	if err != nil {
		log.Print(err)
		responseGetUserOrders.Success = false
		responseGetUserOrders.Message = "Invalid JSON format"
	}

	responseGetUserOrders, err = getOrderInfo(requestGetUserOrders)

	if err != nil {
		log.Printf("Error while getting order info: %v", err)
		responseGetUserOrders.Success = false
		if err.Error() == sql.ErrNoRows.Error() {
			responseGetUserOrders.Message = "No orders found"
		} else {
			responseGetUserOrders.Message = err.Error()
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {

		responseGetUserOrders.Success = true
		responseGetUserOrders.Message = ""
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseGetUserOrders)
}

// For placing a new order
func HandlerPlaceOrder(w http.ResponseWriter, r *http.Request) {
	var requestPlaceOrder model.RequestPlaceOrder
	var responsePlaceOrder model.ResponseOrderInfo
	err := json.NewDecoder(r.Body).Decode(&requestPlaceOrder)
	if err != nil {
		log.Print(err)
		responsePlaceOrder.Success = false
		responsePlaceOrder.Message = "Invalid JSON format"
	}

	responsePlaceOrder, err = addOrder(requestPlaceOrder)

	if err != nil {
		log.Print(err)
		responsePlaceOrder.Success = false
		if err == sql.ErrNoRows {
			responsePlaceOrder.Message = "No orders found"
		} else {
			responsePlaceOrder.Message = err.Error()
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {

		responsePlaceOrder.Success = true
		responsePlaceOrder.Message = ""
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responsePlaceOrder)
}

// For adding a new product
func HandlerAddProduct(w http.ResponseWriter, r *http.Request) {
	var requestAddProduct model.RequestAddProduct
	var responseAddProduct model.ResponseProductInfo
	err := json.NewDecoder(r.Body).Decode(&requestAddProduct)
	if err != nil {
		log.Print(err)
		responseAddProduct.Success = false
		responseAddProduct.Message = "Invalid JSON format"
	}

	responseAddProduct, err = addProduct(requestAddProduct)

	if err != nil {
		log.Print(err)
		responseAddProduct.Success = false
		if err == sql.ErrNoRows {
			responseAddProduct.Message = "No products found"
		} else {
			responseAddProduct.Message = err.Error()
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		responseAddProduct.Success = true
		responseAddProduct.Message = ""
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseAddProduct)
}

// For updating a product
func HandlerUpdateProduct(w http.ResponseWriter, r *http.Request) {
	var requestUpdateProduct model.RequestUpdateProduct
	var responseUpdateProduct model.ResponseProductInfo
	err := json.NewDecoder(r.Body).Decode(&requestUpdateProduct)
	if err != nil {
		log.Print(err)
		responseUpdateProduct.Success = false
		responseUpdateProduct.Message = "Invalid JSON format"
	}

	responseUpdateProduct, err = updateProduct(requestUpdateProduct)

	if err != nil {
		log.Print(err)
		responseUpdateProduct.Success = false
		if err == sql.ErrNoRows {
			responseUpdateProduct.Message = "No products found"
		} else {
			responseUpdateProduct.Message = err.Error()
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {

		responseUpdateProduct.Success = true
		responseUpdateProduct.Message = ""
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseUpdateProduct)
}

// For getting product details
func HandlerProductInfo(w http.ResponseWriter, r *http.Request) {
	var requestProductInfo model.RequestProductInfo
	var responseProductInfo model.ResponseProductInfo
	err := json.NewDecoder(r.Body).Decode(&requestProductInfo)
	if err != nil {
		log.Print(err)
		responseProductInfo.Success = false
		responseProductInfo.Message = "Invalid JSON format"
	}

	responseProductInfo, err1 := getProductInfo(requestProductInfo)

	if err1 != nil {
		log.Println("Error in get prod info")
		log.Print(err1)
		responseProductInfo.Success = false
		if err == sql.ErrNoRows {
			responseProductInfo.Message = "No products found"
		} else {
			responseProductInfo.Message = err1.Error()
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		responseProductInfo.Success = true
		responseProductInfo.Message = ""
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseProductInfo)
}

// For Getting all events
func HandlerGetAllEvents(w http.ResponseWriter, r *http.Request) {
	var responseEventsInfo model.ResponseEventsInfo
	responseEventsInfo, err := getAllEvents()

	if err != nil {
		log.Printf("")
		responseEventsInfo.Success = false
		if err == sql.ErrNoRows {
			responseEventsInfo.Message = "No events found"
		} else {
			responseEventsInfo.Message = err.Error()
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	responseEventsInfo.Success = true
	responseEventsInfo.Message = ""

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseEventsInfo)
}
