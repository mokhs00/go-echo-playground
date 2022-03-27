package data

import userpb "go-grpc-playground/protos/v1/users"

var Users = []*userpb.UserInfo{
	{
		UserId:      "1",
		Name:        "mokhs",
		PhoneNumber: "01012341234",
		Age:         22,
	},
	{
		UserId:      "2",
		Name:        "mokhs2",
		PhoneNumber: "01043214321",
		Age:         23,
	},
}
