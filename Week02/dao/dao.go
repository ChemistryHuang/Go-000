package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type Account struct {
	Id   int64
	Name string
}

var db *sql.DB

//func initDB() (err error) {
//	dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		return err
//	}
//	return db.Ping()
//}

type AccountDao struct {
}

func (ad *AccountDao) GetAccountById(id int) (Account, error) {

	var user Account
	//	查找数据库
	sqlStr := "select id, name from account where id=?"
	err := db.QueryRow(sqlStr, 1).Scan(&user.Id, &user.Name)
	if err != nil {
		return user, errors.Wrap(sql.ErrNoRows, "Not Found")
	}

	//user.Id=1
	//user.Name="noob"
	return user, nil
}
