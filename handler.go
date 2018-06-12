package main

import (
	"errors"
	"fmt"
	"log"

	pb "github.com/inigofu/shippy-user-service/proto/auth"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

const topic = "user.created"

type service struct {
	repo         Repository
	tokenService Authable
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.ResponseUser) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *service) GetUserMenus(ctx context.Context, req *pb.User, res *pb.ResponseMenu) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.New("no auth meta-data found in request")
	}

	// Note this is now uppercase (not entirely sure why this is...)
	token := meta["Authorization"]
	if token == "" {
		return errors.New("no auth meta-data found in request")
	}
	log.Println("Authenticating with token: ", token)
	tokin := &pb.Token{
		Token: token,
	}
	tokout := &pb.ResponseToken{}
	err := srv.ValidateToken(ctx, tokin, tokout)
	if err != nil {
		return err
	}
	menues, err := srv.repo.GetUserMenus(req.Email)
	if err != nil {
		return err
	}
	res.Menues = menues
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.ResponseUser) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}
func (srv *service) GetAllUsersRole(ctx context.Context, req *pb.Request, res *pb.ResponseUser) error {
	users, err := srv.repo.GetAllUsersRole()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}
func (srv *service) UpdateUser(ctx context.Context, req *pb.User, res *pb.ResponseUser) error {
	err := srv.repo.UpdateUser(req)
	if err != nil {
		return err
	}
	res.User = req
	return nil
}
func (srv *service) DeleteUser(ctx context.Context, req *pb.User, res *pb.ResponseUser) error {
	err := srv.repo.DeleteUser(req)
	if err != nil {
		return err
	}
	res.User = nil
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.ResponseToken) error {
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

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.ResponseUser) error {

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

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.ResponseToken) error {
	log.Println("validating token:", req.Token)
	log.Println("validating token complete:", req)
	// Decode token
	claims, err := srv.tokenService.Decode(req.Token)

	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Token = &pb.Token{Valid: true}
	res.User = &pb.User{Email: claims.User.Email}

	return nil
}
func (srv *service) CreateRole(ctx context.Context, req *pb.Role, res *pb.ResponseRole) error {
	log.Println("Creating role: ", req)

	if err := srv.repo.CreateRole(req); err != nil {
		return errors.New(fmt.Sprintf("error creating role: %v", err))
	}

	res.Role = req
	return nil
}
func (srv *service) UpdateRole(ctx context.Context, req *pb.Role, res *pb.ResponseRole) error {
	log.Println("Creating role: ", req)

	if err := srv.repo.UpdateRole(req); err != nil {
		return errors.New(fmt.Sprintf("error creating role: %v", err))
	}

	res.Role = req
	return nil
}
func (srv *service) GetRole(ctx context.Context, req *pb.Role, res *pb.ResponseRole) error {
	role, err := srv.repo.GetRole(req.Id)
	if err != nil {
		return err
	}
	res.Role = role
	return nil
}
func (srv *service) DeleteRole(ctx context.Context, req *pb.Role, res *pb.ResponseRole) error {
	err := srv.repo.DeleteRole(req)
	if err != nil {
		return err
	}
	res.Role = nil
	return nil
}
func (srv *service) GetAllRoles(ctx context.Context, req *pb.Request, res *pb.ResponseRole) error {
	roles, err := srv.repo.GetAllRoles()
	if err != nil {
		return err
	}
	res.Roles = roles
	return nil
}

func (srv *service) CreateMenu(ctx context.Context, req *pb.Menu, res *pb.ResponseMenu) error {
	log.Println("Creating menu: ", req)

	if err := srv.repo.CreateMenu(req); err != nil {
		return errors.New(fmt.Sprintf("error creating menu: %v", err))
	}

	res.Menu = req
	return nil
}
func (srv *service) GetMenu(ctx context.Context, req *pb.Menu, res *pb.ResponseMenu) error {
	return nil
}
func (srv *service) GetAllMenues(ctx context.Context, req *pb.Request, res *pb.ResponseMenu) error {
	return nil
}

func (srv *service) CreateForm(ctx context.Context, req *pb.Form, res *pb.ResponseForm) error {
	log.Println("Creating form: ", req)

	if err := srv.repo.CreateForm(req); err != nil {
		return errors.New(fmt.Sprintf("error creating form: %v", err))
	}

	res.Form = req
	return nil
}
func (srv *service) GetForm(ctx context.Context, req *pb.Form, res *pb.ResponseForm) error {
	log.Println("Getting form: ", req, "with id:", req.Id)
	form, err := srv.repo.GetForm(req.Id)
	if err != nil {
		return err
	}
	res.Form = form
	return nil
}
func (srv *service) DeleteForm(ctx context.Context, req *pb.Form, res *pb.ResponseForm) error {
	log.Println("Getting form: ", req, "with id:", req.Id)
	err := srv.repo.DeleteForm(req)
	if err != nil {
		return err
	}
	res.Form = nil
	return nil
}
func (srv *service) UpdateForm(ctx context.Context, req *pb.Form, res *pb.ResponseForm) error {
	log.Println("Updating form: ", req, "with id:", req.Id)
	form, err := srv.repo.UpdateForm(req)
	if err != nil {
		return err
	}
	res.Form = form
	return nil
}
func (srv *service) GetAllForms(ctx context.Context, req *pb.Request, res *pb.ResponseForm) error {
	forms, err := srv.repo.GetAllForms()
	if err != nil {
		return err
	}
	res.Forms = forms
	return nil
}
func (srv *service) DeleteFields(ctx context.Context, req *pb.Form, res *pb.Error) error {
	log.Println("Deleting fields")
	err := srv.repo.DeleteFields(req)
	if err != nil {
		return err
	}
	res = nil
	return nil
}
func (srv *service) DeleteTabs(ctx context.Context, req *pb.Form, res *pb.Error) error {
	log.Println("Deleting tabs")
	err := srv.repo.DeleteTabs(req)
	if err != nil {
		return err
	}
	res = nil
	return nil
}

func (srv *service) CreateSchema(ctx context.Context, req *pb.FormSchema, res *pb.ResponseFormSchema) error {
	log.Println("Creating schema: ", req)

	if err := srv.repo.CreateSchema(req); err != nil {
		return errors.New(fmt.Sprintf("error creating schema: %v", err))
	}

	res.Formschema = req
	return nil
}
func (srv *service) GetSchema(ctx context.Context, req *pb.FormSchema, res *pb.ResponseFormSchema) error {
	return nil
}
func (srv *service) GetAllSchemas(ctx context.Context, req *pb.Request, res *pb.ResponseFormSchema) error {
	return nil
}
