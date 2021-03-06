package entity

type Beautician struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Price int    `json:"price"`
}

type BeauticianInteractorInterface interface {
	GetBeauticians() ([]Beautician, error)
	GetBeautician(id string) (Beautician, error)
	AddBeautician(name string, sex string, price int) (string, error)
	UpdateBeautician(id string, name string, sex string, price int) error
	DeleteBeautician(id string) error
}
