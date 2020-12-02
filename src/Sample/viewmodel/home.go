package viewmodel

import "SampleGo/src/Sample/model"

type Home struct {
	Title  string
	Active string
	User   model.User
}

func NewHome(user model.User) Home {
	result := Home{
		Active: "home",
		Title:  "Malzenstwa",
		User:   user,
	}
	return result
}

func NewHome2() Home {
	result := Home{
		Active: "home",
		Title:  "Malzenstwa",
	}
	return result
}
