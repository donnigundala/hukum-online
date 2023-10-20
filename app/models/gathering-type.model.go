package models

type GatheringType struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`

	Gatherings []Gathering `json:"gatherings" gorm:"foreignKey:TypeID;references:ID"`
}
