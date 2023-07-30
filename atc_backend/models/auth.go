package models

type Auth struct {
	ID        uint   `gorm:"column:id;autoIncrement"`
	Auth_name string `gorm:"column:auth_name;unique;size:20"`
}

func getAuths() []Auth {
	auths := []Auth{
		{
			Auth_name: "气象",
		},
		{
			Auth_name: "VDL-2",
		},
		{
			Auth_name: "ADS-A/C",
		},
		{
			Auth_name: "ACARS",
		},
		{
			Auth_name: "satcom",
		},
	}

	return auths
}
