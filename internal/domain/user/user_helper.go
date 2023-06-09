package user

import (
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/jackc/pgx/v5"

	"metalloyCore/tools"
)

func (u *User) ToReponse() *UserResponse {
	return &UserResponse{
		UserBase: u.UserBase,
	}
}

func (u *User) ScanFromRow(row pgx.Row) error {
	err := row.Scan(
		&u.UserID,
		&u.Username,
		&u.Email,
		&u.UserType,
		&u.FirstName,
		&u.LastName,
		&u.PhoneNumber,
		&u.AddressID,
		&u.RegistrationDate,
		&u.Password,
	)
	return err
}

func (u *UserResponse) ScanFromRow(row pgx.Row) error {
	err := row.Scan(
		&u.UserID,
		&u.Username,
		&u.Email,
		&u.UserType,
		&u.FirstName,
		&u.LastName,
		&u.PhoneNumber,
		&u.AddressID,
		&u.RegistrationDate,
	)
	return err
}

func (fr *FullUserResponse) ScanFromRow(row pgx.Row) error {
	err := row.Scan(
		&fr.UserID,
		&fr.Username,
		&fr.Email,
		&fr.UserType,
		&fr.FirstName,
		&fr.LastName,
		&fr.PhoneNumber,
		&fr.AddressID,
		&fr.RegistrationDate,
		&fr.StreetAddress,
		&fr.City,
		&fr.State,
		&fr.Country,
		&fr.PostalCode,
	)
	return err
}

func (u *UserCreate) DecodeBody(data io.ReadCloser) error {
	err := json.NewDecoder(data).Decode(u)
	if err != nil {
		return tools.NewBadRequestErr("Invalid JSON body: " + err.Error())
	}
	return nil
}

func (u *UserCreate) DecodeBase64(data string) error {
	decodedBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(decodedBytes, u)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserCreate) Validate() error {
	if u.Username == "" || u.Email == "" || u.UserType == "" ||
		u.FirstName == "" || u.LastName == "" || u.PhoneNumber == "" ||
		u.StreetAddress == "" || u.City == "" || u.State == "" ||
		u.Country == "" || u.PostalCode == "" || u.Password == "" {
		return tools.NewBadRequestErr("All fields are required")
	}
	return nil
}

func InitUserUpdate(username string) *UserUpdate {
	return &UserUpdate{
		UserBase: UserBase{Username: username},
	}
}

func (u *UserUpdate) DecodeBody(data io.ReadCloser) error {
	err := json.NewDecoder(data).Decode(u)
	if err != nil {
		return tools.NewBadRequestErr("Invalid JSON body: " + err.Error())
	}
	return nil
}
