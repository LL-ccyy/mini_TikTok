package model

func migrant() {
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{})
}
