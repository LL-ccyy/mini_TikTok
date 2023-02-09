package model

func migrant() {
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}, &Video{}, &Favourite{}, &Comment{}, &Follow{})
	//DB.Model(&Video{}).AddForeignKey("uid","User(id)","CASCADE","CASCADE")

}
