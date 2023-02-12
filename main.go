package main

import (
	"lazycodecampaign/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/lazycode_campaign?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}

	userInput.Name = "hari user"
	userInput.Occupation = "technician"
	userInput.Email = "user.hari@gmail.com"
	userInput.Password = "123123"

	userService.RegisterUser(userInput)
}
