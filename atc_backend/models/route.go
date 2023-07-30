package models

type Route struct {
	ID          uint   `gorm:"column:id;autoIncrement"`
	Flight_num  string `gorm:"column:flight_num;unique;size:20"`
	Origin      string `gorm:"column:origin;size:20"`
	Destination string `gorm:"column:destination;size:20"`
}

func getRoutes() []Route {
	routes := []Route{
		{
			Flight_num:  "CA1234",
			Origin:      "北京",
			Destination: "上海",
		},
		{
			Flight_num:  "CA5678",
			Origin:      "兰州",
			Destination: "庆阳",
		},
		{
			Flight_num:  "CZ7890",
			Origin:      "郑州",
			Destination: "周口",
		},
		{
			Flight_num:  "MU9197",
			Origin:      "上海",
			Destination: "成都",
		},
	}

	return routes
}
