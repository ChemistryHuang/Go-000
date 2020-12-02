package service

import (
	"Week02/biz"
	"Week02/dao"
)

type AccountService struct{}

func (u *AccountService) GetUser(id int) (dao.Account, error) {
	b := biz.AccountBiz{}
	user, err := b.GetUser(id)
	return user, err
}
