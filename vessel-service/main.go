package main

import (
	pb "github.com/quin47/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"errors"
	"context"
	"fmt"
)

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(specification *pb.Specification)(*pb.Vessel,error) {
	for _, v := range repo.vessels {
		if specification.Capacity<=v.Capacity &&specification.MaxWeight<=v.MaxWeight{
			return v,nil
		}
	}
	return  nil,errors.New("No vessel found by that spec")

}
func (s *service) FindAvailable(ctx context.Context,specification *pb.Specification,response *pb.Response) error {
	vessel, err := s.repo.FindAvailable(specification)
	if err != nil {
		return err
	}
	response.Vessel=vessel
	return  nil

}

type Repository interface {
	FindAvailable( *pb.Specification,)(*pb.Vessel,error)
}

type service struct {
	repo Repository
}

func main()  {
	vessels := []*pb.Vessel{
		{Id: "vessel00023", Name: "Boat McBoatface", MaxWeight: 980000, Capacity: 100232323},
	}
	repository := &VesselRepository{vessels}
	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	srv.Init()
	pb.RegisterVesselServiceHandler(srv.Server(),&service{repository})
	 if err := srv.Run();err!=nil{
	 	fmt.Println(err)
	 }
}
