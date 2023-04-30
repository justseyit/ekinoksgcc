package eventops

import (
	"database/sql"
	"ekinoksgcc/repository"
	"testing"
)

func TestAddOrderPlacementEventToDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, _, err := AddOrderPlacementEventToDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestRemoveOrderPlacementEventFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, _, err := RemoveOrderPlacementEventFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestAddUserRegisterEventToDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, _, err := AddUserRegisterEventToDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestRemoveUserRegisterEventFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, err := RemoveUserRegisterEventFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestAddUserLoginEventToDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, _, err := AddUserLoginEventToDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestRemoveUserLoginEventFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, _, err := RemoveUserLoginEventFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestAddRoleAssignmentEventToDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	userRoleID := 1
	actorID := 1
	_, _, err := AddRoleAssignmentEventToDB(actorID, userRoleID)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestRemoveRoleAssignmentEventFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, _, err := RemoveRoleAssignmentEventFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestAddProductAddEventToDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, _, err := AddProductAddEventToDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestRemoveProductAddEventFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, _, err := RemoveProductAddEventFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetAllProductAddEventsFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	_, err := GetAllProductAddEventsFromDB()
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetAllUserRegisterEventsFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	_, err := GetAllUserRegisterEventsFromDB()
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetAllUserLoginEventsFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	_, err := GetAllUserLoginEventsFromDB()
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetAllRoleAssignmentEventsFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	_, err := GetAllRoleAssignmentEventsFromDB()
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetAllOrderPlacementEventsFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	_, err := GetAllOrderPlacementEventsFromDB()
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetProductAddEventByEventIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, _, err := GetProductAddEventByEventIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetUserRegisterEventByEventIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, _, err := GetUserRegisterEventByEventIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetUserLoginEventByEventIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, _, err := GetUserLoginEventByEventIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetRoleAssignmentEventByEventIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, _, err := GetRoleAssignmentEventByEventIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetOrderPlacementEventByEventIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, _, err := GetOrderPlacementEventByEventIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetProductAddEventsByUserIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, err := GetProductAddEventsByUserIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetUserRegisterEventByUserIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, err := GetUserRegisterEventByUserIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetUserLoginEventsByUserIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, err := GetUserLoginEventsByUserIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetRoleAssignmentEventsByActorIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, err := GetRoleAssignmentEventsByActorIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetRoleAssignmentEventByAssignedUserIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, err := GetRoleAssignmentEventByAssignedUserIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetEventByEventID(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, err := GetEventByEventID(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetAddProductEventByIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, err := GetAddProductEventByIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetUserRegisterEventByIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, err := GetUserRegisterEventByIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetUserLoginEventByIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, err := GetUserLoginEventByIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetRoleAssignmentEventByIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, err := GetRoleAssignmentEventByIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}

func TestGetOrderPlacementEventByIDFromDB(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := 1
	_, err := GetOrderPlacementEventByIDFromDB(req)
	if err != nil && err != sql.ErrNoRows {
		t.Errorf("Error: %v", err)
	}
}
