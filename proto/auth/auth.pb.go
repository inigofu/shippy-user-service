// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	proto/auth/auth.proto

It has these top-level messages:
	User
	Request
	ResponseUser
	ResponseRole
	ResponseMenu
	ResponseToken
	Token
	Error
	Role
	Menu
	Badge
	Wrapper
	Atributes
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Company  string `protobuf:"bytes,3,opt,name=company" json:"company,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	Token    string `protobuf:"bytes,6,opt,name=token" json:"token,omitempty"`
	// @inject_tag: gorm:"many2many:user_roles;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:role_id;jointable_foreignkey:user_id;"
	Roles []*Role `protobuf:"bytes,7,rep,name=roles" json:"roles,omitempty" gorm:"many2many:user_roles;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:role_id;jointable_foreignkey:user_id;"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *User) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type Request struct {
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ResponseUser struct {
	User   *User    `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Users  []*User  `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	Token  *Token   `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
}

func (m *ResponseUser) Reset()                    { *m = ResponseUser{} }
func (m *ResponseUser) String() string            { return proto.CompactTextString(m) }
func (*ResponseUser) ProtoMessage()               {}
func (*ResponseUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ResponseUser) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ResponseUser) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ResponseUser) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *ResponseUser) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type ResponseRole struct {
	Role   *Role    `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	Roles  []*Role  `protobuf:"bytes,2,rep,name=roles" json:"roles,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *ResponseRole) Reset()                    { *m = ResponseRole{} }
func (m *ResponseRole) String() string            { return proto.CompactTextString(m) }
func (*ResponseRole) ProtoMessage()               {}
func (*ResponseRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ResponseRole) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *ResponseRole) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *ResponseRole) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type ResponseMenu struct {
	Menu   *Menu    `protobuf:"bytes,1,opt,name=menu" json:"menu,omitempty"`
	Menues []*Menu  `protobuf:"bytes,2,rep,name=menues" json:"menues,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *ResponseMenu) Reset()                    { *m = ResponseMenu{} }
func (m *ResponseMenu) String() string            { return proto.CompactTextString(m) }
func (*ResponseMenu) ProtoMessage()               {}
func (*ResponseMenu) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ResponseMenu) GetMenu() *Menu {
	if m != nil {
		return m.Menu
	}
	return nil
}

func (m *ResponseMenu) GetMenues() []*Menu {
	if m != nil {
		return m.Menues
	}
	return nil
}

func (m *ResponseMenu) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type ResponseToken struct {
	Token  *Token   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	User   *User    `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *ResponseToken) Reset()                    { *m = ResponseToken{} }
func (m *ResponseToken) String() string            { return proto.CompactTextString(m) }
func (*ResponseToken) ProtoMessage()               {}
func (*ResponseToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ResponseToken) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *ResponseToken) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ResponseToken) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token  string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Valid  bool     `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code        int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Role struct {
	Id    string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Users []*User `protobuf:"bytes,3,rep,name=users" json:"users,omitempty"`
	// @inject_tag: gorm:"many2many:role_menus;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:menu_id;jointable_foreignkey:role_id;"
	Menues []*Menu `protobuf:"bytes,4,rep,name=menues" json:"menues,omitempty" gorm:"many2many:role_menus;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:menu_id;jointable_foreignkey:role_id;"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Role) GetMenues() []*Menu {
	if m != nil {
		return m.Menues
	}
	return nil
}

type Menu struct {
	Id      string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Url     string   `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Badge   *Badge   `protobuf:"bytes,4,opt,name=badge" json:"badge,omitempty"`
	Wrapper *Wrapper `protobuf:"bytes,5,opt,name=wrapper" json:"wrapper,omitempty"`
	Title   bool     `protobuf:"varint,6,opt,name=title" json:"title,omitempty"`
	// @inject_tag: gorm:"many2many:menu_childrens;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:children_id;jointable_foreignkey:parent_id;"
	Children  []*Menu `protobuf:"bytes,7,rep,name=children" json:"children,omitempty" gorm:"many2many:menu_childrens;foreignkey:id;association_foreignkey:id;association_jointable_foreignkey:children_id;jointable_foreignkey:parent_id;"`
	Icon      string  `protobuf:"bytes,8,opt,name=icon" json:"icon,omitempty"`
	Roles     []*Role `protobuf:"bytes,9,rep,name=roles" json:"roles,omitempty"`
	BadgeID   string  `protobuf:"bytes,10,opt,name=badgeID" json:"badgeID,omitempty"`
	WrapperID string  `protobuf:"bytes,11,opt,name=wrapperID" json:"wrapperID,omitempty"`
}

func (m *Menu) Reset()                    { *m = Menu{} }
func (m *Menu) String() string            { return proto.CompactTextString(m) }
func (*Menu) ProtoMessage()               {}
func (*Menu) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Menu) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Menu) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Menu) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Menu) GetBadge() *Badge {
	if m != nil {
		return m.Badge
	}
	return nil
}

func (m *Menu) GetWrapper() *Wrapper {
	if m != nil {
		return m.Wrapper
	}
	return nil
}

func (m *Menu) GetTitle() bool {
	if m != nil {
		return m.Title
	}
	return false
}

func (m *Menu) GetChildren() []*Menu {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Menu) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Menu) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Menu) GetBadgeID() string {
	if m != nil {
		return m.BadgeID
	}
	return ""
}

func (m *Menu) GetWrapperID() string {
	if m != nil {
		return m.WrapperID
	}
	return ""
}

type Badge struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Variant string `protobuf:"bytes,2,opt,name=variant" json:"variant,omitempty"`
	Text    string `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
}

func (m *Badge) Reset()                    { *m = Badge{} }
func (m *Badge) String() string            { return proto.CompactTextString(m) }
func (*Badge) ProtoMessage()               {}
func (*Badge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Badge) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Badge) GetVariant() string {
	if m != nil {
		return m.Variant
	}
	return ""
}

func (m *Badge) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Wrapper struct {
	Id          string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Element     string     `protobuf:"bytes,2,opt,name=element" json:"element,omitempty"`
	Atributes   *Atributes `protobuf:"bytes,3,opt,name=atributes" json:"atributes,omitempty"`
	AtributesID string     `protobuf:"bytes,4,opt,name=atributesID" json:"atributesID,omitempty"`
}

func (m *Wrapper) Reset()                    { *m = Wrapper{} }
func (m *Wrapper) String() string            { return proto.CompactTextString(m) }
func (*Wrapper) ProtoMessage()               {}
func (*Wrapper) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Wrapper) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Wrapper) GetElement() string {
	if m != nil {
		return m.Element
	}
	return ""
}

func (m *Wrapper) GetAtributes() *Atributes {
	if m != nil {
		return m.Atributes
	}
	return nil
}

func (m *Wrapper) GetAtributesID() string {
	if m != nil {
		return m.AtributesID
	}
	return ""
}

type Atributes struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Atributes) Reset()                    { *m = Atributes{} }
func (m *Atributes) String() string            { return proto.CompactTextString(m) }
func (*Atributes) ProtoMessage()               {}
func (*Atributes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Atributes) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "auth.User")
	proto.RegisterType((*Request)(nil), "auth.Request")
	proto.RegisterType((*ResponseUser)(nil), "auth.ResponseUser")
	proto.RegisterType((*ResponseRole)(nil), "auth.ResponseRole")
	proto.RegisterType((*ResponseMenu)(nil), "auth.ResponseMenu")
	proto.RegisterType((*ResponseToken)(nil), "auth.ResponseToken")
	proto.RegisterType((*Token)(nil), "auth.Token")
	proto.RegisterType((*Error)(nil), "auth.Error")
	proto.RegisterType((*Role)(nil), "auth.Role")
	proto.RegisterType((*Menu)(nil), "auth.Menu")
	proto.RegisterType((*Badge)(nil), "auth.Badge")
	proto.RegisterType((*Wrapper)(nil), "auth.Wrapper")
	proto.RegisterType((*Atributes)(nil), "auth.Atributes")
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 753 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xcd, 0x6e, 0xdb, 0x38,
	0x10, 0xf6, 0x8f, 0x64, 0xd9, 0xe3, 0x78, 0x77, 0xc1, 0xcd, 0x02, 0x44, 0x76, 0xb1, 0xf0, 0x72,
	0x81, 0x26, 0x6d, 0x91, 0xb4, 0x70, 0xce, 0x3d, 0xb8, 0x4d, 0x60, 0xe4, 0x90, 0x0b, 0xd1, 0xbf,
	0xab, 0x62, 0x0d, 0x1a, 0xa1, 0xb2, 0xa4, 0x52, 0x54, 0xd2, 0xde, 0xfb, 0x10, 0x7d, 0x8a, 0x9e,
	0x7a, 0xef, 0xab, 0x15, 0x1c, 0xd2, 0x92, 0x9c, 0x56, 0x81, 0x2e, 0x16, 0xe7, 0x9b, 0x21, 0xe7,
	0x9b, 0x5f, 0x18, 0xfe, 0xca, 0x55, 0xa6, 0xb3, 0x27, 0x61, 0xa9, 0xaf, 0xe9, 0xe7, 0x84, 0x64,
	0xe6, 0x99, 0xb3, 0xf8, 0xda, 0x07, 0xef, 0x55, 0x81, 0x8a, 0xfd, 0x06, 0x83, 0x38, 0xe2, 0xfd,
	0x79, 0xff, 0x68, 0x22, 0x07, 0x71, 0xc4, 0x18, 0x78, 0x69, 0xb8, 0x41, 0x3e, 0x20, 0x84, 0xce,
	0x8c, 0x43, 0xb0, 0xce, 0x36, 0x79, 0x98, 0x7e, 0xe2, 0x43, 0x82, 0xb7, 0x22, 0xdb, 0x07, 0x1f,
	0x37, 0x61, 0x9c, 0x70, 0x8f, 0x70, 0x2b, 0xb0, 0x03, 0x18, 0xe7, 0x61, 0x51, 0xdc, 0x66, 0x2a,
	0xe2, 0x3e, 0x29, 0x2a, 0xd9, 0xdc, 0xd0, 0xd9, 0x7b, 0x4c, 0xf9, 0xc8, 0xde, 0x20, 0x81, 0xcd,
	0xc1, 0x57, 0x59, 0x82, 0x05, 0x0f, 0xe6, 0xc3, 0xa3, 0xe9, 0x02, 0x4e, 0x88, 0xb0, 0xcc, 0x12,
	0x94, 0x56, 0x21, 0x26, 0x10, 0x48, 0xfc, 0x50, 0x62, 0xa1, 0xc5, 0x97, 0x3e, 0xec, 0x49, 0x2c,
	0xf2, 0x2c, 0x2d, 0x90, 0x62, 0xf8, 0x17, 0xbc, 0xb2, 0x40, 0x45, 0x51, 0x54, 0x97, 0x8d, 0x46,
	0x12, 0x6e, 0x5e, 0x37, 0xdf, 0x82, 0x0f, 0x9a, 0xaf, 0x93, 0x81, 0x55, 0xb0, 0xff, 0x61, 0x84,
	0x4a, 0x65, 0xaa, 0xe0, 0x43, 0x32, 0x99, 0x5a, 0x93, 0x73, 0x83, 0x49, 0xa7, 0x62, 0xff, 0x6d,
	0xa9, 0x7b, 0xe4, 0xc7, 0xd9, 0xbc, 0x34, 0x90, 0x8b, 0x43, 0x94, 0x35, 0x33, 0x43, 0xde, 0x30,
	0x33, 0xf4, 0x77, 0x99, 0x51, 0x58, 0x84, 0xd7, 0x71, 0x0f, 0x5a, 0xe2, 0xee, 0xc4, 0x4c, 0xdc,
	0xd6, 0x6e, 0x2f, 0x31, 0x2d, 0x8d, 0xdb, 0x0d, 0xa6, 0xe5, 0xae, 0x5b, 0xa3, 0x91, 0x84, 0x33,
	0x01, 0x23, 0xf3, 0xbd, 0xeb, 0x97, 0x2c, 0x9c, 0xa6, 0xab, 0xe3, 0xd9, 0xd6, 0x31, 0xe5, 0xa1,
	0xce, 0x51, 0xbf, 0x2d, 0x47, 0x55, 0xb5, 0x06, 0x2d, 0xd5, 0xea, 0xe4, 0xf8, 0x2d, 0xf8, 0xd6,
	0xe1, 0x7e, 0xd3, 0x61, 0xd5, 0x4f, 0xfb, 0xe0, 0xdf, 0x84, 0x49, 0x1c, 0x91, 0x93, 0xb1, 0xb4,
	0x42, 0xb7, 0x97, 0x9f, 0x81, 0x4f, 0x80, 0x99, 0x84, 0x75, 0x16, 0xd9, 0xda, 0xf9, 0x92, 0xce,
	0x6c, 0x0e, 0xd3, 0x08, 0x8b, 0xb5, 0x8a, 0x73, 0x1d, 0x67, 0xa9, 0x1b, 0x92, 0x26, 0x24, 0x72,
	0xf0, 0xa8, 0xf2, 0x5d, 0xe6, 0xaa, 0xea, 0xcb, 0x61, 0x5b, 0x5f, 0xd6, 0x85, 0xf2, 0xda, 0x0a,
	0x25, 0xbe, 0x0d, 0xc0, 0xa3, 0xaa, 0x77, 0x71, 0xf9, 0x07, 0x0c, 0x4b, 0x95, 0xb8, 0x31, 0x36,
	0x47, 0x53, 0xb1, 0xab, 0x30, 0x7a, 0x87, 0xbb, 0x5d, 0xfd, 0xdc, 0x40, 0xd2, 0x6a, 0xd8, 0x21,
	0x04, 0xb7, 0x2a, 0xcc, 0x73, 0x54, 0x34, 0xce, 0xd3, 0xc5, 0xcc, 0x1a, 0xbd, 0xb1, 0xa0, 0xdc,
	0x6a, 0xa9, 0x18, 0xb1, 0x4e, 0x90, 0x86, 0x7b, 0x2c, 0xad, 0xc0, 0x1e, 0xc0, 0x78, 0x7d, 0x1d,
	0x27, 0x91, 0xc2, 0x74, 0x77, 0xbe, 0x29, 0x8c, 0x4a, 0x67, 0xf8, 0xc6, 0xeb, 0x2c, 0xe5, 0x63,
	0xcb, 0xd7, 0x9c, 0xeb, 0x01, 0x99, 0xb4, 0x0d, 0x08, 0x87, 0x80, 0x58, 0x5e, 0x9c, 0x71, 0xb0,
	0xcb, 0xc9, 0x89, 0xec, 0x1f, 0x98, 0x38, 0x62, 0x17, 0x67, 0x7c, 0x4a, 0xba, 0x1a, 0x10, 0xe7,
	0xe0, 0x53, 0x90, 0x3f, 0xa5, 0x8d, 0x43, 0x70, 0x13, 0xaa, 0x38, 0x4c, 0xb5, 0xcb, 0xdc, 0x56,
	0x34, 0x04, 0x35, 0x7e, 0xd4, 0x2e, 0x7b, 0x74, 0x16, 0x9f, 0xfb, 0x10, 0xb8, 0x3c, 0xfc, 0xea,
	0x25, 0x4c, 0x70, 0x83, 0xf5, 0x4b, 0x4e, 0x64, 0xc7, 0x30, 0x09, 0xb5, 0x8a, 0xaf, 0x4a, 0x8d,
	0x05, 0x3d, 0x37, 0x5d, 0xfc, 0x6e, 0x43, 0x5b, 0x6e, 0x61, 0x59, 0x5b, 0x98, 0xb6, 0xab, 0x84,
	0x8b, 0x33, 0xb7, 0x6c, 0x9b, 0x90, 0xf8, 0x1b, 0x26, 0xd5, 0xcd, 0xbb, 0x3c, 0x16, 0xdf, 0x3d,
	0xf0, 0x96, 0xa5, 0xbe, 0x66, 0x8f, 0x60, 0xf4, 0x42, 0x61, 0xa8, 0x91, 0x35, 0x7a, 0xed, 0x80,
	0xb9, 0xa4, 0x36, 0x56, 0xaa, 0xe8, 0xb1, 0x43, 0x18, 0xae, 0x50, 0x77, 0x30, 0x3c, 0x86, 0xd1,
	0x0a, 0xf5, 0x32, 0x49, 0xd8, 0x6c, 0xab, 0xa7, 0x3d, 0xdd, 0x62, 0xfe, 0xd0, 0x71, 0x69, 0x3e,
	0xfc, 0xe7, 0xae, 0x25, 0x4d, 0xb6, 0xe8, 0xb1, 0xa7, 0xb0, 0xb7, 0x42, 0x6d, 0x2c, 0x4c, 0xa7,
	0x14, 0xf7, 0x71, 0x31, 0x06, 0xa2, 0xc7, 0x4e, 0x61, 0xf6, 0xda, 0x8c, 0x7a, 0xa8, 0xdd, 0x3e,
	0x6a, 0x2e, 0xa0, 0x36, 0x37, 0x27, 0x00, 0x36, 0x2b, 0x34, 0xb8, 0x8d, 0x16, 0xbb, 0xeb, 0xc4,
	0x60, 0xa2, 0xc7, 0x1e, 0x43, 0xb0, 0x42, 0xdd, 0xd1, 0x78, 0x01, 0x53, 0x9b, 0x1d, 0x49, 0xdd,
	0x7a, 0x7f, 0x8a, 0xdc, 0x9d, 0x8a, 0x10, 0x8d, 0x75, 0x63, 0x58, 0x5a, 0xa2, 0xb6, 0x84, 0x3a,
	0x1a, 0x9f, 0x52, 0x52, 0x97, 0x49, 0x72, 0x69, 0xf7, 0xfc, 0xfd, 0x8c, 0xec, 0xa5, 0xab, 0x11,
	0xfd, 0x77, 0x38, 0xfd, 0x11, 0x00, 0x00, 0xff, 0xff, 0xea, 0xd9, 0xe0, 0x2c, 0x54, 0x08, 0x00,
	0x00,
}
