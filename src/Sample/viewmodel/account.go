package viewmodel

type Account struct {
	Title    string
	Active   string
	Email    string
	Password string
}

func NewAccount() Account {
	result := Account{
		Active: "home",
		Title:  "Konto",
	}
	return result
}
