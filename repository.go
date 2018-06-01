package main

import (
	"log"

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
	GetForm(id string) (*pb.Form, error)
	UpdateForm(form *pb.Form) (*pb.Form, error)
	CreateForm(form *pb.Form) error
	GetAllForms() ([]*pb.Form, error)
	GetSchema(id string) (*pb.FormSchema, error)
	CreateSchema(schema *pb.FormSchema) error
	GetAllSchemas() ([]*pb.FormSchema, error)
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
	if err := repo.db.Set("gorm:association_autoupdate", false).Create(user).Error; err != nil {
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
	if err := repo.db.Set("gorm:association_autoupdate", false).Create(&role).Error; err != nil {
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
	if err := repo.db.Set("gorm:association_autoupdate", false).Create(menu).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetUserMenus(email string) ([]*pb.Menu, error) {
	user := &pb.User{}
	//var roles []*pb.Role
	var menues []*pb.Menu
	var rolmenuesall []*pb.Menu

	if err := repo.db.Preload("Roles.Menues").Select("id").Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	log.Println("Getting menues from:", user)
	for _, role := range user.Roles {
		rolmenuesall = append(rolmenuesall, role.Menues...)
	}
	var rolmenues []string
	for _, role := range rolmenuesall {
		rolmenues = append(rolmenues, role.Id)
	}
	type Result struct {
		Children_id string
	}

	var results []Result
	var childrenid []string
	if err := repo.db.Raw("SELECT children_id FROM menu_childrens").Scan(&results).Error; err != nil {
		return nil, err
	}
	for _, result := range results {
		childrenid = append(childrenid, result.Children_id)
	}
	// (*sql.Row)

	if err := repo.db.Not(childrenid).Where(rolmenues).Preload("Children", "id in (?)", rolmenues).Find(&menues).Error; err != nil {
		return nil, err
	}

	return menues, nil
}
func (repo *UserRepository) GetAllForms() ([]*pb.Form, error) {
	var forms []*pb.Form
	if err := repo.db.Preload("Fields", func(db *gorm.DB) *gorm.DB {
		return repo.db.Order("form_schemas.order ASC")
	}).Preload("Tabs").Preload("Tabs.Fields", func(db *gorm.DB) *gorm.DB {
		return repo.db.Order("form_schemas.order ASC")
	}).Find(&forms).Error; err != nil {
		return nil, err
	}
	return forms, nil
}

func (repo *UserRepository) GetForm(id string) (*pb.Form, error) {
	var form *pb.Form
	log.Println("Getting form with id:", id)
	form = &pb.Form{Id: id}
	if err := repo.db.Preload("Fields", func(db *gorm.DB) *gorm.DB {
		return repo.db.Order("form_schemas.order ASC")
	}).Preload("Tabs").Preload("Tabs.Fields", func(db *gorm.DB) *gorm.DB {
		return repo.db.Order("form_schemas.order ASC")
	}).First(&form).Error; err != nil {
		return nil, err
	}
	return form, nil
}
func (repo *UserRepository) UpdateForm(form *pb.Form) (*pb.Form, error) {
	log.Println("Updating form", form)
	if err := repo.db.Save(&form).Error; err != nil {
		return nil, err
	}
	return form, nil
}
func (repo *UserRepository) CreateForm(form *pb.Form) error {
	if err := repo.db.Create(&form).Error; err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) GetAllSchemas() ([]*pb.FormSchema, error) {
	var formschemas []*pb.FormSchema
	if err := repo.db.Find(&formschemas).Error; err != nil {
		return nil, err
	}
	return formschemas, nil
}

func (repo *UserRepository) GetSchema(id string) (*pb.FormSchema, error) {
	var formschema *pb.FormSchema
	formschema.Id = id
	if err := repo.db.First(&formschema).Error; err != nil {
		return nil, err
	}
	return formschema, nil
}
func (repo *UserRepository) CreateSchema(formschema *pb.FormSchema) error {
	if err := repo.db.Create(&formschema).Error; err != nil {
		return err
	}
	return nil
}
