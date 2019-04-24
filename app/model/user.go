package model

// GetUserByName description
//
// creatTime:2019年04月23日 17:14:22
// author:hailaz
func GetUserByName() (*User, error) {
	u := User{}
	err := defDB.Table("user").Where("user_name", "admin").Struct(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
