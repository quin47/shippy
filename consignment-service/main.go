package main
 import(
	 pb "github.com/quin47/shippy/consignment-service/proto/consignment"
	 vesselProto "github.com/quin47/shippy/vessel-service/proto/vessel"
	 "context"
	 "github.com/micro/go-micro"
	 "fmt"
	 "log"
 )
type IRepository interface{
 Create(*pb.Consignment)(*pb.Consignment, error)
 GetAll()([]*pb.Consignment,error)
}

type Repository struct {
	consignments []*pb.Consignment
}

func(repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment,error){
	update := append(repo.consignments,consignment)
	repo.consignments=update
	return consignment,nil

}
func(repo *Repository) GetAll()([]*pb.Consignment,error){
	return  repo.consignments,nil
}

type service struct {
	repo IRepository
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, request *pb.Consignment ,response *pb.Response)(error)  {

	available, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: request.Weight,
		Capacity:  int32(len(request.Containers)),
	})
	if err != nil {
		return  err
	}
	log.Printf("Found vessel: %s \n",available.Vessel.Name)
	request.VesselId=available.Vessel.Id
	consignment, e := s.repo.Create(request)
	if e != nil {
		return e
	}
	response.Created= true
	response.Consignment=consignment
	return  nil
}
func (s *service) GetAllConsignment(ctx context.Context,request *pb.GetRequest,response *pb.Response) (error) {
	consignments, err := s.repo.GetAll()
	if err!=nil {
		return nil
	}
	response.Consigments=consignments
	return  nil
}


func main() {
	repo :=&Repository{}
	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	srv.Init()
	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())
	pb.RegisterShippingServiceHandler(srv.Server(),&service{repo,vesselClient})

	if err := srv.Run();err!=nil{
		fmt.Println(err)
	}

}

