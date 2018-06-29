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
)

const(
 defaultFilename = "consignment.json"
)




func main() {
	file := defaultFilename
	if len(os.Args) > 1 {
		file =os.Args[1]
	}
	consignment, err := parseFile(file)
	if err!=nil {

		log.Fatalf("parse consignment Failed: %v",err)
	}

	serviceClient := pb.NewShippingService("go.micro.srv.consignment", client.DefaultClient)
	cmd.Init()
	response, e := serviceClient.CreateConsignment(context.TODO(), consignment)
	if e != nil {
		log.Fatalf("could not create:%v ",err)
	}

	log.Printf("Created: %v",response)
	allConsignment, i := serviceClient.GetAllConsignment(context.Background(), &pb.GetRequest{})

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