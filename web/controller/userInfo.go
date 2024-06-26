
package controller

import "education/service"

type Application struct {
	Setup *service.ServiceSetup
}

type User struct {
	LoginName	string
	Password	string
	IsAdmin	string
}


var users []User

func init() {

	admin := User{LoginName:"admin", Password:"123456", IsAdmin:"T"}
	jack := User{LoginName:"jack", Password:"123456", IsAdmin:"F"}

	users = append(users, admin)
	users = append(users, jack)

}