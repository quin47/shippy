package main

import (
	vesselProto "github.com/quin47/shippy/vessel-service/proto/vessel"
	pb "github.com/quin47/shippy/consignment-service/proto/consignment"
	"github.com/globalsign/mgo"
	"context"
	"log"
)

type service struct {
	session *mgo.Session
	vesselService	vesselProto.VesselService
}

func (s *service) GetRepo() Repository {
	return  &ConsignmentRepository{s.session.Clone()}
}

func (s *service) CreateConsignment(ctx context.Context,req *pb.Consignment,res *pb.Response)  error{
	response, err := s.vesselService.FindAvailable(context.Background(), &vesselProto.Specification{MaxWeight: req.Weight, Capacity: int32(len(req.Containers))})
	log.Printf("Found  vessel: %s \n",response.Vessel.Name)
	if err != nil {
		return  err
	}
	req.VesselId = response.Vessel.Id
	s.GetRepo().Create(req)
	if err != nil {
		return  err
	}

	res.Created= true
	res.Consignment=req
	return  nil
}

func (s *service)GetAllConsignment(ctx context.Context, request *pb.GetRequest,res *pb.Response) error  {
	defer  s.GetRepo().Close()
	consignments, err := s.GetRepo().GetAll()
	if err !=nil {
		return err
	}
	res.Consigments = consignments
	return nil

}






