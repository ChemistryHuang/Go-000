package biz

import  "Week02/dao"

type AccountBiz struct{}

func (u *AccountBiz) GetUser(id int) (dao.Account, error) {
	ad := dao.AccountDao{}
	return ad.GetAccountById(id)
}
