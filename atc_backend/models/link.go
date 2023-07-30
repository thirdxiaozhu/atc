package models

type Flight struct {
	ID    uint `gorm:"column:id;autoIncrement"`
	Arn   uint `gorm:"column:arn;NOT NULL"`
	Route uint `gorm:"column:route;NOT NULL"`
}
type Link struct {
	ID     uint `gorm:"column:id;autoIncrement"`
	Flight uint `gorm:"column:flight;NOT NULL"`
	Auth   uint `gorm:"column:auth;NOT NULL"`
}
