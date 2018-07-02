package main

import (
	pb "github.com/quin47/shippy/user-service/proto/user"
	"github.com/labstack/gommon/log"
	"github.com/micro/go-micro"
	"fmt"
)

func main()  {
	db, e := CreateConnection()
	defer db.Close()
	if e != nil {
		log.Fatalf("Could not connect to postgres : %v",e)
	}
	db.AutoMigrate(&pb.User{})
	repo := &UserRepository{db}
	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)
	srv.Init()
	pb.RegisterUserServiceHandler(srv.Server(),&service{repo,tokenService})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
