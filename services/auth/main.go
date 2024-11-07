package main

import (
	auth "gomall/services/auth/kitex_gen/auth/authservice"
	"log"
)

func main() {
	svr := auth.NewServer(new(AuthServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
