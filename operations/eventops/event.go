package eventops

import (
	"ekinoksgcc/model"
	"ekinoksgcc/operations/productops"
	"ekinoksgcc/repository"
	"log"
)

//Database operations for events

// Order of return values: eventID, orderPlacementEventID, error
func AddOrderPlacementEventToDB(placedOrderID int) (int, int, error) {
	var eventID int
	var orderPlacementEventID int
	err := repository.DB.QueryRow("INSERT INTO events(eventTimestamp) VALUES(NOW()) RETURNING eventID").Scan(&eventID)
	if err != nil {
		log.Fatal(err)
	}
	err1 := repository.DB.QueryRow("INSERT INTO orderPlacementEvent(eventID, placedOrderID) VALUES($1, $2) RETURNING orderPlacementEventID", eventID, placedOrderID).Scan(&orderPlacementEventID)
	if err1 != nil {
		log.Fatal(err1)
	}
	return eventID, orderPlacementEventID, err
}

// Order of return values: eventID, orderPlacementEventID, error
func RemoveOrderPlacementEventFromDB(orderPlacementEventID int) (int, int, error) {
	var eventID int
	err := repository.DB.QueryRow("SELECT eventID, placedOrderID FROM orderPlacementEvent WHERE orderPlacementEventID=$1", orderPlacementEventID).Scan(&eventID)
	if err != nil {
		log.Fatal(err)
	}
	_, err1 := repository.DB.Exec("DELETE FROM orderPlacementEvent WHERE orderPlacementEventID=$1", orderPlacementEventID)
	if err1 != nil {
		log.Fatal(err1)
	}
	_, err2 := repository.DB.Exec("DELETE FROM events WHERE eventID=$1", eventID)
	if err2 != nil {
		log.Fatal(err2)
	}
	return eventID, orderPlacementEventID, err
}

// Order of return values: eventID, userRegisterEventID, error
func AddUserRegisterEventToDB(userID int) (int, int, error) {
	var eventID int
	err := repository.DB.QueryRow("INSERT INTO events(eventTimestamp) VALUES(NOW()) RETURNING eventID").Scan(&eventID)
	if err != nil {
		log.Fatal(err)
	}
	err1 := repository.DB.QueryRow("INSERT INTO userRegisterEvent(eventID, userID) VALUES($1, $2) RETURNING userRegisterEventID", eventID, userID)
	if err1 != nil {
		log.Fatal(err1)
	}
	return eventID, userID, err
}

// Order of return values: eventID, userRegisterEventID, error
func RemoveUserRegisterEventFromDB(userRegisterEventID int) (int, error) {
	var eventID int
	err := repository.DB.QueryRow("SELECT eventID, userID FROM userRegisterEvent WHERE userRegisterEventID=$1", userRegisterEventID).Scan(&eventID)
	if err != nil {
		log.Fatal(err)
	}
	_, err1 := repository.DB.Exec("DELETE FROM userRegisterEvent WHERE userRegisterEventID=$1", userRegisterEventID)
	if err1 != nil {
		log.Fatal(err1)
	}
	_, err2 := repository.DB.Exec("DELETE FROM events WHERE eventID=$1", eventID)
	if err2 != nil {
		log.Fatal(err2)
	}
	return eventID, err
}

// Order of return values: eventID, userLoginEventID, error
func AddUserLoginEventToDB(userID int) (int, int, error) {
	var eventID int
	var userLoginEventID int
	err := repository.DB.QueryRow("INSERT INTO events(eventTimestamp) VALUES(NOW()) RETURNING eventID").Scan(&eventID)
	if err != nil {
		log.Fatal(err)
	}
	err1 := repository.DB.QueryRow("INSERT INTO userLoginEvent(eventID, userID) VALUES($1, $2) RETURNING userLoginEventID", eventID, userID).Scan(&userLoginEventID)
	if err1 != nil {
		log.Fatal(err1)
	}
	return eventID, userLoginEventID, err
}

// Order of return values: eventID, userLoginEventID, error
func RemoveUserLoginEventFromDB(userLoginEventID int) (int, int, error) {
	var eventID int
	err := repository.DB.QueryRow("SELECT eventID, userID FROM userLoginEvent WHERE userLoginEventID=$1", userLoginEventID).Scan(&eventID)
	if err != nil {
		log.Fatal(err)
	}
	_, err1 := repository.DB.Exec("DELETE FROM userLoginEvent WHERE userLoginEventID=$1", userLoginEventID)
	if err1 != nil {
		log.Fatal(err1)
	}
	_, err2 := repository.DB.Exec("DELETE FROM events WHERE eventID=$1", eventID)
	if err2 != nil {
		log.Fatal(err2)
	}
	return eventID, userLoginEventID, err
}

// Order of return values: eventID, roleAssignmentEvent, error
func AddRoleAssignmentEventToDB(userRoleID int, actorUserID int) (int, int, error) {
	var eventID int
	var roleAssignmentEventID int
	err := repository.DB.QueryRow("INSERT INTO events(eventTimestamp) VALUES(NOW()) RETURNING eventID").Scan(&eventID)
	if err != nil {
		log.Fatal(err)
	}
	err1 := repository.DB.QueryRow("INSERT INTO roleAssignmentEvent(eventID, userRoleID, userID) VALUES($1, $2, $3) RETURNING roleAssignmentEventID", eventID, userRoleID, actorUserID).Scan(&roleAssignmentEventID)
	if err1 != nil {
		log.Fatal(err1)
	}
	return eventID, roleAssignmentEventID, err
}

// Order of return values: eventID, roleAssignmentEvent, error
func RemoveRoleAssignmentEventFromDB(roleAssignmentEventID int) (int, int, error) {
	var eventID int
	err := repository.DB.QueryRow("SELECT eventID, userRoleID, userID FROM roleAssignmentEvent WHERE roleAssignmentEventID=$1", roleAssignmentEventID).Scan(&eventID)
	if err != nil {
		log.Fatal(err)
	}
	_, err1 := repository.DB.Exec("DELETE FROM roleAssignmentEvent WHERE roleAssignmentEventID=$1", roleAssignmentEventID)
	if err1 != nil {
		log.Fatal(err1)
	}
	_, err2 := repository.DB.Exec("DELETE FROM events WHERE eventID=$1", eventID)
	if err2 != nil {
		log.Fatal(err2)
	}
	return eventID, roleAssignmentEventID, err
}

// Order of return values: eventID, productAddEventID, error
func AddProductAddEventToDB(addedProductID int) (int, int, error) {
	var eventID int
	var productAddEventID int
	err := repository.DB.QueryRow("INSERT INTO events(eventTimestamp) VALUES(NOW()) RETURNING eventID").Scan(&eventID)
	if err != nil {
		log.Fatal(err)
	}
	err1 := repository.DB.QueryRow("INSERT INTO productAddEvent(eventID, addedProductID) VALUES($1, $2) RETURNING productAddEventID", eventID, addedProductID).Scan(&productAddEventID)
	if err1 != nil {
		log.Fatal(err1)
	}
	return eventID, productAddEventID, err
}

// Order of return values: eventID, productAddEventID, error
func RemoveProductAddEventFromDB(productAddEventID int) (int, int, error) {
	var eventID int
	err := repository.DB.QueryRow("SELECT eventID, addedProductID FROM productAddEvent WHERE productAddEventID=$1", productAddEventID).Scan(&eventID)
	if err != nil {
		log.Fatal(err)
	}
	_, err1 := repository.DB.Exec("DELETE FROM productAddEvent WHERE productAddEventID=$1", productAddEventID)
	if err1 != nil {
		log.Fatal(err1)
	}
	_, err2 := repository.DB.Exec("DELETE FROM events WHERE eventID=$1", eventID)
	if err2 != nil {
		log.Fatal(err2)
	}
	return eventID, productAddEventID, err
}

//Get all

func GetAllProductAddEventsFromDB() ([]*model.ProductAddEvent, error) {
	rows, err := repository.DB.Query("SELECT productAddEventID, eventID, addedProductID FROM productAddEvent")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var productAddEvents []*model.ProductAddEvent
	for rows.Next() {
		var productAddEvent model.ProductAddEvent
		err := rows.Scan(&productAddEvent.EventID, &productAddEvent.EventID, &productAddEvent.AddedProductID)
		if err != nil {
			log.Fatal(err)
		}
		productAddEvents = append(productAddEvents, &productAddEvent)
	}
	return productAddEvents, err
}

func GetAllUserRegisterEventsFromDB() ([]*model.UserRegisterEvent, error) {
	rows, err := repository.DB.Query("SELECT userRegisterEventID, eventID, userID FROM userRegisterEvent")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var userRegisterEvents []*model.UserRegisterEvent
	for rows.Next() {
		var userRegisterEvent model.UserRegisterEvent
		err := rows.Scan(&userRegisterEvent.EventID, &userRegisterEvent.EventID, &userRegisterEvent.UserID)
		if err != nil {
			log.Fatal(err)
		}
		userRegisterEvents = append(userRegisterEvents, &userRegisterEvent)
	}
	return userRegisterEvents, err
}

func GetAllUserLoginEventsFromDB() ([]*model.UserLoginEvent, error) {
	rows, err := repository.DB.Query("SELECT userLoginEventID, eventID, userID FROM userLoginEvent")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var userLoginEvents []*model.UserLoginEvent
	for rows.Next() {
		var userLoginEvent model.UserLoginEvent
		err := rows.Scan(&userLoginEvent.EventID, &userLoginEvent.EventID, &userLoginEvent.UserID)
		if err != nil {
			log.Fatal(err)
		}
		userLoginEvents = append(userLoginEvents, &userLoginEvent)
	}
	return userLoginEvents, err
}

func GetAllRoleAssignmentEventsFromDB() ([]*model.RoleAssignmentEvent, error) {
	rows, err := repository.DB.Query("SELECT roleAssignmentEventID, eventID, userRoleID, userID FROM roleAssignmentEvent")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var roleAssignmentEvents []*model.RoleAssignmentEvent
	for rows.Next() {
		var roleAssignmentEvent model.RoleAssignmentEvent
		err := rows.Scan(&roleAssignmentEvent.EventID, &roleAssignmentEvent.EventID, &roleAssignmentEvent.UserRoleID, &roleAssignmentEvent.UserID)
		if err != nil {
			log.Fatal(err)
		}
		roleAssignmentEvents = append(roleAssignmentEvents, &roleAssignmentEvent)
	}
	return roleAssignmentEvents, err
}

func GetAllOrderPlacementEventsFromDB() ([]*model.OrderPlacementEvent, error) {
	rows, err := repository.DB.Query("SELECT orderPlacementEventID, eventID, placedOrderID FROM orderPlacementEvent")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var orderPlacementEvents []*model.OrderPlacementEvent
	for rows.Next() {
		var orderPlacementEvent model.OrderPlacementEvent
		err := rows.Scan(&orderPlacementEvent.EventID, &orderPlacementEvent.EventID, &orderPlacementEvent.PlacedOrderID)
		if err != nil {
			log.Fatal(err)
		}
		orderPlacementEvents = append(orderPlacementEvents, &orderPlacementEvent)
	}
	return orderPlacementEvents, err
}

//Get by

func GetProductAddEventByEventIDFromDB(eventID int) (*model.Event, *model.ProductAddEvent, error) {
	var productAddEvent model.ProductAddEvent
	err := repository.DB.QueryRow("SELECT productAddEventID, eventID, addedProductID FROM productAddEvent WHERE eventID=$1", eventID).Scan(&productAddEvent.EventID, &productAddEvent.EventID, &productAddEvent.AddedProductID)
	if err != nil {
		log.Fatal(err)
	}

	var event model.Event
	err = repository.DB.QueryRow("SELECT eventID, eventTimestamp FROM events WHERE eventID=$1", productAddEvent.EventID).Scan(&event.EventID, &event.EventTimestamp)
	if err != nil {
		log.Fatal(err)
	}
	return &event, &productAddEvent, err
}

func GetUserRegisterEventByEventIDFromDB(eventID int) (*model.Event, *model.UserRegisterEvent, error) {
	var userRegisterEvent model.UserRegisterEvent
	err := repository.DB.QueryRow("SELECT userRegisterEventID, eventID, userID FROM userRegisterEvent WHERE eventID=$1", eventID).Scan(&userRegisterEvent.EventID, &userRegisterEvent.EventID, &userRegisterEvent.UserID)
	if err != nil {
		log.Fatal(err)
	}

	var event model.Event
	err = repository.DB.QueryRow("SELECT eventID, eventTimestamp FROM events WHERE eventID=$1", userRegisterEvent.EventID).Scan(&event.EventID, &event.EventTimestamp)
	if err != nil {
		log.Fatal(err)
	}
	return &event, &userRegisterEvent, err
}

func GetUserLoginEventByEventIDFromDB(eventID int) (*model.Event, *model.UserLoginEvent, error) {
	var userLoginEvent model.UserLoginEvent
	err := repository.DB.QueryRow("SELECT userLoginEventID, eventID, userID FROM userLoginEvent WHERE eventID=$1", eventID).Scan(&userLoginEvent.EventID, &userLoginEvent.EventID, &userLoginEvent.UserID)
	if err != nil {
		log.Fatal(err)
	}

	var event model.Event
	err = repository.DB.QueryRow("SELECT eventID, eventTimestamp FROM events WHERE eventID=$1", userLoginEvent.EventID).Scan(&event.EventID, &event.EventTimestamp)
	if err != nil {
		log.Fatal(err)
	}
	return &event, &userLoginEvent, err
}

func GetRoleAssignmentEventByEventIDFromDB(eventID int) (*model.Event, *model.RoleAssignmentEvent, error) {
	var roleAssignmentEvent model.RoleAssignmentEvent
	err := repository.DB.QueryRow("SELECT roleAssignmentEventID, eventID, userRoleID, userID FROM roleAssignmentEvent WHERE eventID=$1", eventID).Scan(&roleAssignmentEvent.EventID, &roleAssignmentEvent.EventID, &roleAssignmentEvent.UserRoleID, &roleAssignmentEvent.UserID)
	if err != nil {
		log.Fatal(err)
	}

	var event model.Event
	err = repository.DB.QueryRow("SELECT eventID, eventTimestamp FROM events WHERE eventID=$1", roleAssignmentEvent.EventID).Scan(&event.EventID, &event.EventTimestamp)
	if err != nil {
		log.Fatal(err)
	}
	return &event, &roleAssignmentEvent, err
}

func GetOrderPlacementEventByEventIDFromDB(eventID int) (*model.Event, *model.OrderPlacementEvent, error) {
	var orderPlacementEvent model.OrderPlacementEvent
	err := repository.DB.QueryRow("SELECT orderPlacementEventID, eventID, orderID FROM orderPlacementEvent WHERE eventID=$1", eventID).Scan(&orderPlacementEvent.EventID, &orderPlacementEvent.EventID, &orderPlacementEvent.PlacedOrderID)
	if err != nil {
		log.Fatal(err)
	}

	var event model.Event
	err = repository.DB.QueryRow("SELECT eventID, eventTimestamp FROM events WHERE eventID=$1", orderPlacementEvent.EventID).Scan(&event.EventID, &event.EventTimestamp)
	if err != nil {
		log.Fatal(err)
	}
	return &event, &orderPlacementEvent, err
}

func GetProductAddEventsByUserIDFromDB(userID int) ([]*model.ProductAddEvent, error) {
	addedProds, err := productops.GetAddedProductMappingsByUserIDFromDB(userID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var productAddEvents []*model.ProductAddEvent
	for _, addedProd := range addedProds {
		var productAddEvent model.ProductAddEvent
		err := repository.DB.QueryRow("SELECT productAddEventID, eventID, addedProductID FROM productAddEvent WHERE addedProductID=$1", addedProd.AddedProductID).Scan(&productAddEvent.EventID, &productAddEvent.EventID, &productAddEvent.AddedProductID)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		productAddEvents = append(productAddEvents, &productAddEvent)
	}
	return productAddEvents, err
}

func GetUserRegisterEventByUserIDFromDB(userID int) (*model.UserRegisterEvent, error) {
	var userRegisterEvent model.UserRegisterEvent
	err := repository.DB.QueryRow("SELECT userRegisterEventID, eventID, userID FROM userRegisterEvent WHERE userID=$1", userID).Scan(&userRegisterEvent.EventID, &userRegisterEvent.EventID, &userRegisterEvent.UserID)
	if err != nil {
		log.Fatal(err)
	}
	return &userRegisterEvent, err
}

func GetUserLoginEventsByUserIDFromDB(userID int) ([]*model.UserLoginEvent, error) {
	rows, err := repository.DB.Query("SELECT userLoginEventID, eventID, userID FROM userLoginEvent WHERE userID=$1", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var userLoginEvents []*model.UserLoginEvent
	for rows.Next() {
		var userLoginEvent model.UserLoginEvent
		err := rows.Scan(&userLoginEvent.EventID, &userLoginEvent.EventID, &userLoginEvent.UserID)
		if err != nil {
			log.Fatal(err)
		}
		userLoginEvents = append(userLoginEvents, &userLoginEvent)
	}
	return userLoginEvents, err
}

func GetRoleAssignmentEventsByActorIDFromDB(actorUserID int) ([]*model.RoleAssignmentEvent, error) {
	rows, err := repository.DB.Query("SELECT roleAssignmentEventID, eventID, userRoleID, userID FROM roleAssignmentEvent WHERE userID=$1", actorUserID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var roleAssignmentEvents []*model.RoleAssignmentEvent
	for rows.Next() {
		var roleAssignmentEvent model.RoleAssignmentEvent
		err := rows.Scan(&roleAssignmentEvent.EventID, &roleAssignmentEvent.EventID, &roleAssignmentEvent.UserRoleID, &roleAssignmentEvent.UserID)
		if err != nil {
			log.Fatal(err)
		}
		roleAssignmentEvents = append(roleAssignmentEvents, &roleAssignmentEvent)
	}
	return roleAssignmentEvents, err
}

func GetRoleAssignmentEventByAssignedUserIDFromDB(assignedUserID int) (*model.RoleAssignmentEvent, error) {
	var roleAssignmentEvent model.RoleAssignmentEvent
	var userRole model.UserRole
	err := repository.DB.QueryRow("SELECT userRoleID, userID, roleID FROM userRole WHERE userID=$1", assignedUserID).Scan(&userRole.UserRoleID, &userRole.UserID, &userRole.RoleID)
	if err != nil {
		log.Fatal(err)
	}
	err = repository.DB.QueryRow("SELECT roleAssignmentEventID, eventID, userRoleID, userID FROM roleAssignmentEvent WHERE userRoleID=$1", userRole.UserRoleID).Scan(&roleAssignmentEvent.EventID, &roleAssignmentEvent.EventID, &roleAssignmentEvent.UserRoleID, &roleAssignmentEvent.UserID)
	if err != nil {
		log.Fatal(err)
	}
	return &roleAssignmentEvent, err
}

//Get by ID

func GetEventByEventID(eventID int) (*model.Event, error) {
	var event model.Event
	err := repository.DB.QueryRow("SELECT eventID, eventTimestamp FROM events WHERE eventID=$1", eventID).Scan(&event.EventID, &event.EventTimestamp)
	if err != nil {
		log.Fatal(err)
	}
	return &event, err
}

func GetAddProductEventByIDFromDB(eventID int) (*model.ProductAddEvent, error) {
	var productAddEvent model.ProductAddEvent
	err := repository.DB.QueryRow("SELECT productAddEventID, eventID, addedProductID FROM productAddEvent WHERE productAddEventID=$1", eventID).Scan(&productAddEvent.EventID, &productAddEvent.EventID, &productAddEvent.AddedProductID)
	if err != nil {
		log.Fatal(err)
	}
	return &productAddEvent, err
}

func GetUserRegisterEventByIDFromDB(eventID int) (*model.UserRegisterEvent, error) {
	var userRegisterEvent model.UserRegisterEvent
	err := repository.DB.QueryRow("SELECT userRegisterEventID, eventID, userID FROM userRegisterEvent WHERE userRegisterEventID=$1", eventID).Scan(&userRegisterEvent.EventID, &userRegisterEvent.EventID, &userRegisterEvent.UserID)
	if err != nil {
		log.Fatal(err)
	}
	return &userRegisterEvent, err
}

func GetUserLoginEventByIDFromDB(eventID int) (*model.UserLoginEvent, error) {
	var userLoginEvent model.UserLoginEvent
	err := repository.DB.QueryRow("SELECT userLoginEventID, eventID, userID FROM userLoginEvent WHERE userLoginEventID=$1", eventID).Scan(&userLoginEvent.EventID, &userLoginEvent.EventID, &userLoginEvent.UserID)
	if err != nil {
		log.Fatal(err)
	}
	return &userLoginEvent, err
}

func GetRoleAssignmentEventByIDFromDB(eventID int) (*model.RoleAssignmentEvent, error) {
	var roleAssignmentEvent model.RoleAssignmentEvent
	err := repository.DB.QueryRow("SELECT roleAssignmentEventID, eventID, userRoleID, userID FROM roleAssignmentEvent WHERE roleAssignmentEventID=$1", eventID).Scan(&roleAssignmentEvent.EventID, &roleAssignmentEvent.EventID, &roleAssignmentEvent.UserRoleID, &roleAssignmentEvent.UserID)
	if err != nil {
		log.Fatal(err)
	}
	return &roleAssignmentEvent, err
}

func GetOrderPlacementEventByIDFromDB(eventID int) (*model.OrderPlacementEvent, error) {
	var orderPlacementEvent model.OrderPlacementEvent
	err := repository.DB.QueryRow("SELECT orderPlacementEventID, eventID, orderID FROM orderPlacementEvent WHERE orderPlacementEventID=$1", eventID).Scan(&orderPlacementEvent.EventID, &orderPlacementEvent.EventID, &orderPlacementEvent.PlacedOrderID)
	if err != nil {
		log.Fatal(err)
	}
	return &orderPlacementEvent, err
}
