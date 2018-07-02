package main

import (
	pb "github.com/quin47/shippy/user-service/proto/user"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro"
	"log"
	"os"
	"context"
)

func main() {
	cmd.Init()

	userServiceClient := pb.NewUserService("", client.DefaultClient)
	service := micro.NewService(micro.Name("go.micro.srv.user"),micro.Version("latest"))
	service.Init()

	name := "Ewan"
	email := "ewan@gmail.com"
	password := "test123"
	company := "BBC"

	r, err := userServiceClient.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})

	if err != nil {
		log.Fatalf("Could not create: %v ",err)
	}
	log.Printf("Created: %s",r.User.Id)

	getAll, err := userServiceClient.GetAll(context.TODO(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list user : %v",err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}
	token, err := userServiceClient.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("could not auth user : %s error %v",email,err)
	}
	log.Printf("Your access token is %s",token.Token)

	os.Exit(0)
}
