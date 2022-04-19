package entity

type Menu struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Price int    `json:"price"`
	Time  int    `json:"time"`
}

type MenuInteractorInterface interface {
	GetMenus() ([]Menu, error)
	GetMenu(id string) (Menu, error)
	AddMenu(name string, sex string, price int, time int) (string, error)
	UpdateMenu(id string, name string, sex string, price int, time int) error
	DeleteMenu(id string) error
}
