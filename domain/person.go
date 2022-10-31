package domain

type Person struct {
	ID      uint64 `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (Person) TableName() string {
	return "person"
}
