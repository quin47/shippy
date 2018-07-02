package main

import (
	"github.com/globalsign/mgo"
	"context"
	pb "github.com/quin47/shippy/vessel-service/proto/vessel"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) FindAvailable(ctx context.Context, specification *pb.Specification, response *pb.Response) error {
	defer s.GetRepo().Close()
	vessel, err := s.GetRepo().FindAvailable(specification)
	if err != nil {
		return nil
	}
	response.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, vessel *pb.Vessel, response *pb.Response) error {
	defer s.GetRepo().Close()
	if err := s.GetRepo().Create(vessel); err != nil {
		return nil
	}
	response.Vessel = vessel
	response.Created = true
	return nil
}
