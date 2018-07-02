package main

import (
	"log"
	pb "github.com/quin47/shippy/consignment-service/proto/consignment"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/micro/go-micro/client"
	"context"
	"github.com/micro/go-micro/cmd"
	"errors"
	metadata "github.com/micro/go-micro/metadata"
)

const(
 defaultFilename = "consignment.json"
)

func main() {
	file := defaultFilename
	var token string
	log.Println(os.Args)

	if len(os.Args) < 3 {
		log.Fatal(errors.New("not enough arguments, expecting file and token"))
	}

	file = os.Args[1]
	token = os.Args[2]

	if len(os.Args) > 1 {
		file =os.Args[1]
	}


	consignment, err := parseFile(file)
	if err!=nil {
		log.Fatalf("parse consignment Failed: %v",err)
	}

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	serviceClient := pb.NewShippingService("go.micro.srv.consignment", client.DefaultClient)
	cmd.Init()
	response, e := serviceClient.CreateConsignment(ctx, consignment)
	if e != nil {
		log.Fatalf("could not create:%v ",err)
	}

	log.Printf("Created: %v",response)
	allConsignment, i := serviceClient.GetAllConsignment(ctx, &pb.GetRequest{})

	if i != nil {
		log.Fatalf("could not list consignment %v",allConsignment)
	}
	for _,v := range allConsignment.Consigments {
		log.Println(v)
	}
}

func parseFile(file string) (*pb.Consignment ,error){
	var consignment *pb.Consignment
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return  nil,err
	}
	json.Unmarshal(bytes,&consignment)
	return consignment, err


}