package models

type Member struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`

	// ManyToMany
	Gatherings []Gathering `json:"gatherings" gorm:"many2many:Invitation"`
}
