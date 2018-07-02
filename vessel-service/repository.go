package main

import (
	pb "github.com/quin47/shippy/vessel-service/proto/vessel"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
)

const (
	vesselCollection = "vessels"
	dbName           = "shippy"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
	Close()
}
type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) FindAvailable( specification *pb.Specification) (*pb.Vessel,error) {
	var vessel *pb.Vessel
	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": specification.Capacity},
		"maxweight": bson.M{"$gte": specification.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil,err
	}
	return vessel,nil
}
func (repo *VesselRepository)Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}
func ( repo *VesselRepository)Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
