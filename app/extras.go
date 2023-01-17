package app

import (
	"fmt"
	"log"
	"time"

	"dev.go/databases"
)

func loadusers() {
	xusers, err := databases.GetUsers()
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range xusers {
		err := databases.SetData(v.Username, v.Password, 120*time.Second)
		if err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Println("USerdata Loaded to memory")
}
