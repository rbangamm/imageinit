package user

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	UserName string `bson:"uname"`
	Password string `bson:"pwd"`
}