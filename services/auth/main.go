package main

import (
	auth "gomall/kitex_gen/auth/authservice"
	"gomall/services/auth/handler"
	"log"
)

func main() {
	svr := auth.NewServer(new(handler.AuthServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
