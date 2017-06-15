package model

// GetUser function for DB connection test
func GetUser() User {
	user := User{}
	db.Debug().Model(&User{}).First(&user)
	return user
}
