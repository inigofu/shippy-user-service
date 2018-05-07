package main

import (
	pb "github.com/inigofu/shippy-user-service/proto/auth"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmail(email string) (*pb.User, error)
	GetAllRoles() ([]*pb.Role, error)
	GetRole(id string) (*pb.Role, error)
	CreateRole(role *pb.Role) error
	GetAllMenues() ([]*pb.Menu, error)
	GetMenu(id string) (*pb.Menu, error)
	CreateMenu(menu *pb.Menu) error
	GetUserMenus(userid string) ([]*pb.Menu, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
	user := &pb.User{}
	if err := repo.db.Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetAllRoles() ([]*pb.Role, error) {
	var roles []*pb.Role
	if err := repo.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (repo *UserRepository) GetRole(id string) (*pb.Role, error) {
	var role *pb.Role
	role.Id = id
	if err := repo.db.First(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}
func (repo *UserRepository) CreateRole(role *pb.Role) error {
	if err := repo.db.Create(&role).Error; err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) GetAllMenues() ([]*pb.Menu, error) {
	var menues []*pb.Menu
	if err := repo.db.Find(&menues).Error; err != nil {
		return nil, err
	}
	return menues, nil
}

func (repo *UserRepository) GetMenu(id string) (*pb.Menu, error) {
	var menu *pb.Menu
	menu.Id = id
	if err := repo.db.First(&menu).Error; err != nil {
		return nil, err
	}
	return menu, nil
}
func (repo *UserRepository) CreateMenu(menu *pb.Menu) error {
	if err := repo.db.Create(menu).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetUserMenus(userid string) ([]*pb.Menu, error) {
	var user *pb.User
	var roles []*pb.Role
	var menues []*pb.Menu
	user.Id = userid
	if err := repo.db.Model(&user).Related(&roles, "Roles"); err != nil {
		return nil, err
	}
	if err := repo.db.Find(&menues).Error; err != nil {
		return nil, err
	}
	return menues, nil
}
