package models

type Arn struct {
	ID         uint   `gorm:"column:id;autoIncrement"`
	Flight_Arn string `gorm:"column:flight_arn;unique;size:20"`
	Company    string `gorm:"column:company;size:20"`
}

func getArns() []Arn {
	arns := []Arn{
		{
			Flight_Arn: ".B-1234",
			Company:    "中国南方航空",
		},
		{
			Flight_Arn: ".B-5678",
			Company:    "中国国际航空",
		},
		{
			Flight_Arn: ".C-919A",
			Company:    "中国东方航空",
		},
		{
			Flight_Arn: ".C-919B",
			Company:    "中国东方航空",
		},
	}

	return arns
}
