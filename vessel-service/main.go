package main

import (
	pb "github.com/quin47/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"fmt"
	"os"
	"github.com/labstack/gommon/log"
)
const (
	defaultHost = "localhost:27017"
)

func main()  {

	host := os.Getenv("DB_HOST")
	if host =="" {
		host=defaultHost
	}

	session, err := CreateSession(host)
	defer  session.Close()
	if err != nil {
		log.Fatalf("Error connect to mongo %v",err)
	}
	repository := &VesselRepository{session.Copy()}
	createVessels(repository)

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	srv.Init()
	pb.RegisterVesselServiceHandler(srv.Server(),&service{session})
	 if err := srv.Run();err!=nil{
	 	fmt.Println(err)
	 }
}
func createVessels(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret1", MaxWeight: 200000, Capacity: 500},
		{Id: "vessel002", Name: "Kane's Salty Secret2", MaxWeight: 200001, Capacity: 502},
	}
	for _, v := range vessels {
		repo.Create(v)
	}

}
