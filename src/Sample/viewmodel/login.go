package viewmodel

type Login struct {
	Title    string
	Active   string
	Email    string
	Password string
}

func NewLogin() Login {
	result := Login{
		Active: "home",
		Title:  "SampleGo",
	}
	return result
}
