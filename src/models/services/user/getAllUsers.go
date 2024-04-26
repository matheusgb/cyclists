package services

import (
	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
)

func (user *User) GetAllUsers(pagination *domainsP.Pagination, email string) (*domainsP.Pagination, error) {
	return user.repository.GetAllUsers(pagination, email)
}
