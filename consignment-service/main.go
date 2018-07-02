package main

import (
	pb "github.com/quin47/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/quin47/shippy/vessel-service/proto/vessel"
	"os"
	"github.com/micro/go-micro"
	"fmt"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/metadata"
	"log"
	"github.com/micro/go-micro/client"
	 userService "github.com/quin47/shippy/user-service/proto/user"
	 "context"
	 "errors"
 )

const defaultHost  = "localhost:27017"

func main() {
	host:= os.Getenv("DB_HOST")

	if host == "" {
		host= defaultHost
	}

	session, err := CreateSession(host)
	defer  session.Close()
	if err !=nil {
		log.Panicf("could not connect to your database %s - %v",host,err)
	}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)
	srv.Init()
	vesselClient := vesselProto.NewVesselService("go.micro.srv.vessel", srv.Client())
	pb.RegisterShippingServiceHandler(srv.Server(),&service{session.Clone(),vesselClient})

	if err := srv.Run();err!=nil{
		fmt.Println(err)
	}

}


func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// Auth here
		authClient := userService.NewUserService("go.micro.srv.user", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userService.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}

