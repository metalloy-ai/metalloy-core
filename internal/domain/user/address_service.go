package user

import (
	"context"

	"metalloyCore/tools"
)

func (us Service) GetAddress(ctx context.Context, username string) (Address, error) {
	address, err := us.Repo.GetAddress(ctx, username)

	handledAddress, err := tools.HandleEmptyError(address, err)
	return handledAddress.(Address), err
}

func (us Service) UpdateAddress(ctx context.Context, address AddressBase, username string) (Address, error) {
	fieldMap := map[string]interface{}{
		"street_address": address.StreetAddress,
		"city":           address.City,
		"state":          address.State,
		"country":        address.Country,
		"postal_code":    address.PostalCode,
	}
	updateArr, args, argsCount := tools.BuildUpdateQueryArgs(fieldMap, username)
	newAddress, err := us.Repo.UpdateAddress(ctx, updateArr, args, argsCount, username)

	handledNewAddress, err := tools.HandleEmptyError(newAddress, err)
	return handledNewAddress.(Address), err
}
