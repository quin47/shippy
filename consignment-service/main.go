package main
 import(
	 pb "github.com/quin47/shippy/consignment-service/proto/consignment"
	 "context"
	 "net"
	 "log"
	 "google.golang.org/grpc"
	 "google.golang.org/grpc/reflection"
 )
const(
	port=":50051"
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
}

func (s *service) CreateConsignment(ctx context.Context, request *pb.Consignment)(*pb.Response,error)  {
	consignment, e := s.repo.Create(request)
	if e != nil {
		return nil,e
	}
	return &pb.Response{Created:true,Consignment:consignment},nil
}
func (s *service) GetAllConsignment(ctx context.Context,request *pb.GetRequest) (*pb.Response,error) {
	consignments, err := s.repo.GetAll()
	if err!=nil {
		return nil,err;
	}
	return  &pb.Response{Consigments:consignments},nil
}


func main() {
	repo :=&Repository{}
	listen, e := net.Listen("tcp", port)
	if e != nil {
		log.Fatalf("failed to listen : %s",e)
	}
	s:= grpc.NewServer()
	pb.RegisterShippingServiceServer(s, &service{repo})

	reflection.Register(s)
    if err :=s.Serve(listen);err!=nil{
    	log.Fatalf("failed to serve : %s",err)
	}

}

