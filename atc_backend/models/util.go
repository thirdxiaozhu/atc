package models

func GetCompaniesByRole(role string) *[]Company {
	companies := []Company{}
	err := DBH.QueryAllByField(&companies, "company", "role", role)
	if err != nil {
		return nil
	}

	return &companies
}

func GetPublishersByCompany(company_name string) *[]User {
	users := []User{}
	err := DBH.QueryAllByField(&users, "user", "company", company_name)
	if err != nil {
		return nil
	}

	return &users
}
