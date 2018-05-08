package main

import (
	"errors"
	"fmt"
	"log"

	pb "github.com/inigofu/shippy-user-service/proto/auth"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

const topic = "user.created"

type serviceAuth struct {
	repo         Repository
	tokenService Authable
}

type serviceUser struct {
	repo         Repository
	tokenService Authable
}

func (srv *serviceAuth) Get(ctx context.Context, req *pb.User, res *pb.ResponseUser) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *serviceUser) GetUserMenus(ctx context.Context, req *pb.User, res *pb.ResponseMenu) error {
	menues, err := srv.repo.GetUserMenus(req.Email)
	if err != nil {
		return err
	}
	res.Menues = menues
	return nil
}

func (srv *serviceAuth) GetAll(ctx context.Context, req *pb.Request, res *pb.ResponseUser) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *serviceAuth) Auth(ctx context.Context, req *pb.User, res *pb.ResponseToken) error {
	log.Println("Logging in with:", req.Email, req.Password)
	user, err := srv.repo.GetByEmail(req.Email)
	log.Println(user, err)
	if err != nil {
		return err
	}

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = &pb.Token{Token: token}
	res.User = user
	return nil
}

func (srv *serviceAuth) Create(ctx context.Context, req *pb.User, res *pb.ResponseUser) error {

	log.Println("Creating user: ", req)

	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(fmt.Sprintf("error hashing password: %v", err))
	}

	req.Password = string(hashedPass)
	if err := srv.repo.Create(req); err != nil {
		return errors.New(fmt.Sprintf("error creating user: %v", err))
	}

	token, err := srv.tokenService.Encode(req)
	if err != nil {
		return err
	}

	res.User = req
	res.Token = &pb.Token{Token: token}

	/*
		if err := srv.Publisher.Publish(ctx, req); err != nil {
			return errors.New(fmt.Sprintf("error publishing event: %v", err))
		}*/

	return nil
}

func (srv *serviceAuth) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {

	// Decode token
	claims, err := srv.tokenService.Decode(req.Token)

	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Valid = true

	return nil
}
func (srv *serviceUser) CreateRole(ctx context.Context, req *pb.Role, res *pb.ResponseRole) error {
	log.Println("Creating role: ", req)

	if err := srv.repo.CreateRole(req); err != nil {
		return errors.New(fmt.Sprintf("error creating role: %v", err))
	}

	res.Role = req
	return nil
}
func (srv *serviceUser) GetRole(ctx context.Context, req *pb.Role, res *pb.ResponseRole) error {
	return nil
}
func (srv *serviceUser) GetAllRoles(ctx context.Context, req *pb.Request, res *pb.ResponseRole) error {
	return nil
}

func (srv *serviceUser) CreateMenu(ctx context.Context, req *pb.Menu, res *pb.ResponseMenu) error {
	log.Println("Creating menu: ", req)

	if err := srv.repo.CreateMenu(req); err != nil {
		return errors.New(fmt.Sprintf("error creating menu: %v", err))
	}

	res.Menu = req
	return nil
}
func (srv *serviceUser) GetMenu(ctx context.Context, req *pb.Menu, res *pb.ResponseMenu) error {
	return nil
}
func (srv *serviceUser) GetAllMenues(ctx context.Context, req *pb.Request, res *pb.ResponseMenu) error {
	return nil
}
