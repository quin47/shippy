package main

import (
	pb "github.com/quin47/shippy/user-service/proto/user"
	"context"
	"golang.org/x/crypto/bcrypt"
	"log"
	"errors"
)

type service struct {
	repo Repository
	tokenService Authable
}

func (s *service) Create(ctx context.Context, user *pb.User, res *pb.Response) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password=string(password)
	if err := s.repo.Create(user);err!=nil{
		return err
	}
	encode, e := s.tokenService.Encode(user)
	if e != nil {
		return  e
	}
	res.Token=&pb.Token{Token:encode}
	res.User=user
	return nil
}

func (s *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, e := s.repo.Get(req.Id)
	if e!=nil {
		return e
	}
	res.User=user
	return nil
}

func ( s *service) GetAll(ctx context.Context,req *pb.Request, res *pb.Response) error {
	if users, err := s.repo.GetAll(); err != nil {
		return err
	}else {
		res.Users=users
	}
	return nil

}

func (s *service) Auth(ctx context.Context,req *pb.User,res *pb.Token) error {
	log.Println("logging in with",req.Email,req.Password)
	user, err := s.repo.getByEmail(req.Email)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return err
	}
	encode, e := s.tokenService.Encode(user)
	if e != nil {
		return e
	}
	res.Token=encode
	return nil
}

func (s *service) ValidateToken( ctx context.Context,req *pb.Token,res *pb.Token) error {
	decode, e := s.tokenService.Decode(req.Token)
	if e != nil {
		return e
	}
	if decode.User.Id == "" {
		return errors.New("invalid user")
	}
	res.Valid=true
	return nil
}






