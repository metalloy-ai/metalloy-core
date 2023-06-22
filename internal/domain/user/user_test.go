package user

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"metalloyCore/internal/config"
)

func ConvertFromString(s string) []UserResponse {
	var users []UserResponse
	json.Unmarshal([]byte(s), &users)
	return users
}

func SetupAndGetService() (UserService, context.Context) {
	config.LoadEnv("../../../.env")
	cfg, ctx := config.LoadBaseConfig(), context.Background()

	repo := InitRepository(*cfg)
	service := InitUserService(repo)

	return service, ctx
}

func TestGetAllUsers(t *testing.T) {
	service, ctx := SetupAndGetService()

	cases := []struct {
		name     string
		pageIdx  string
		pageSize string
		expected interface{}
	}{
		{"Get all users limit 3", "", "3", `[
            {
                "user_id": "86255d58-6f0a-4c64-a4f9-c931d4846b0e",
                "username": "admin_user",
                "email": "adminuser@email.com",
                "user_type": "admin",
                "first_name": "Admin",
                "last_name": "User",
                "phone_number": "214-555-4321",
                "address_id": 859762567944601601,
                "registration_date": "2023-04-25T14:02:05.339805-05:00"
            },
            {
                "user_id": "e489e469-47ae-4cca-a0aa-127d43166745",
                "username": "jane_smith",
                "email": "janesmith@email.com",
                "user_type": "supplier",
                "first_name": "Jane",
                "last_name": "Smith",
                "phone_number": "817-555-5678",
                "address_id": 859762567942275073,
                "registration_date": "2023-04-25T14:02:05.339805-05:00"
            },
            {
                "user_id": "c4bc6330-006a-4d09-b1b4-f959df2abc79",
                "username": "john_doe",
                "email": "johndoe@email.com",
                "user_type": "customer",
                "first_name": "John",
                "last_name": "Doe",
                "phone_number": "682-555-1234",
                "address_id": 859762567931002881,
                "registration_date": "2023-04-25T14:02:05.339805-05:00"
            }
        ]`},
		{"Get all users with page index", "john_doe", "3", `[
            {
                "user_id": "4a337baf-acf7-4048-8fa7-31e5ca8e6b8f",
                "username": "test",
                "email": "test@email.com",
                "user_type": "admin",
                "first_name": "test",
                "last_name": "test",
                "phone_number": "123-456-7890",
                "address_id": 873710857358508033,
                "registration_date": "2023-06-13T20:23:13.38865-05:00"
            },
            {
                "user_id": "3ffaca96-41a4-4207-a1a3-71664cf30c99",
                "username": "test2",
                "email": "test2@email.com",
                "user_type": "admin",
                "first_name": "test",
                "last_name": "test",
                "phone_number": "123-456-7891",
                "address_id": 863209351863861249,
                "registration_date": "2023-05-07T18:09:48.24748-05:00"
            },
            {
                "user_id": "1ecdc616-3353-4526-b8f9-a20514699964",
                "username": "test3",
                "email": "test3@email.com",
                "user_type": "admin",
                "first_name": "test",
                "last_name": "test",
                "phone_number": "123-456-7892",
                "address_id": 863209384751923201,
                "registration_date": "2023-05-07T18:09:58.328637-05:00"
            }
        ]`},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			users, err := service.GetAllUser(ctx, c.pageIdx, c.pageSize)
			if err != nil {
				t.Errorf("Error getting all users: %v", err)
			}

			if len(users) != len(ConvertFromString(c.expected.(string))) {
				t.Errorf("Expected %v users but got %v", len(ConvertFromString(c.expected.(string))), len(users))
			}
		})
	}
}

func TestGetSingleUser(t *testing.T) {
	service, ctx := SetupAndGetService()

	raw := `{
		"user_id": "4a337baf-acf7-4048-8fa7-31e5ca8e6b8f",
		"username": "test",
		"email": "test@email.com",
		"user_type": "admin",
		"first_name": "test",
		"last_name": "test",
		"phone_number": "123-456-7890",
		"address_id": 873710857358508033,
		"registration_date": "2023-06-13T20:23:13.38865-05:00",
		"street_address": "123 Test St",
		"city": "Test",
		"state": "TS",
		"country": "TEST",
		"postal_code": "12345"
	}`
	expected := &FullUserResponse{}
	json.Unmarshal([]byte(raw), expected)

	user, err := service.GetFullUser(ctx, expected.Username)

	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("Expected %v but got %v", expected, user)
	}

}

func TestGetSingleUserAddres(t *testing.T) {
	service, ctx := SetupAndGetService()

	rawUser := `{
		"user_id": "4a337baf-acf7-4048-8fa7-31e5ca8e6b8f",
		"username": "test",
		"email": "test@email.com",
		"user_type": "admin",
		"first_name": "test",
		"last_name": "test",
		"phone_number": "123-456-7890",
		"address_id": 873710857358508033,
		"registration_date": "2023-06-13T20:23:13.38865-05:00",
		"street_address": "123 Test St",
		"city": "Test",
		"state": "TS",
		"country": "TEST",
		"postal_code": "12345"
	}`
	raw := `{
		"address_id": 873710857358508033,
        "street_address": "123 Test St",
        "city": "Test",
        "state": "TS",
        "country": "TEST",
        "postal_code": "12345"
	}`
		
	expectedUser := &FullUserResponse{}
	json.Unmarshal([]byte(rawUser), expectedUser)

	expected := &Address{}
	json.Unmarshal([]byte(raw), expected)

	address, err := service.GetAddress(ctx, "test")

	if err != nil {
		t.Errorf("Error getting address: %v", err)
	}

	if address.AddressID != expectedUser.AddressID && address.AddressID != expected.AddressID {
		t.Errorf("Expected %v but got %v", expected.AddressID, expectedUser.AddressID)
	}

}
