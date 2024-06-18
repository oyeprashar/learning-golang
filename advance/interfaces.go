package main

import "fmt"

type Service interface {
	GET(firstName string, lastName string) string
	POST(firstName string, lastName string) string
}

type DatabaseService struct {
	UserName string
	Password string
}

func (*DatabaseService) GET(firstName string, lastName string) string {
	return "GET_API_Called"

}

func (*DatabaseService) POST(firstName string, lastName string) string {
	return "POST_API_Called"
}

func main() {
	dbService := DatabaseService{
		UserName: "admin",
		Password: "123",
	}

	fmt.Println(dbService.GET("", ""))
	fmt.Println(dbService.POST("", ""))
}
