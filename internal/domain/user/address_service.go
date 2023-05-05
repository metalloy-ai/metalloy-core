package user

func (us Service) GetAddress(username string) (Address, error) {
	return us.Repo.GetAddress(username)
}

func (us Service) UpdateAddress(address AddressBase, username string) (Address, error) {
	return us.Repo.UpdateAddress(address, username)
}
