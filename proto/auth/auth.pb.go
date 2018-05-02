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
	Id       string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Company  string  `protobuf:"bytes,3,opt,name=company" json:"company,omitempty"`
	Email    string  `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Password string  `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	Token    string  `protobuf:"bytes,6,opt,name=token" json:"token,omitempty"`
	Roles    []*Role `protobuf:"bytes,7,rep,name=roles" json:"roles,omitempty"`
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

type Token struct {
	Token  string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Valid  bool     `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
	Id     string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name   string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Users  []*User `protobuf:"bytes,3,rep,name=users" json:"users,omitempty"`
	Menues []*Menu `protobuf:"bytes,4,rep,name=menues" json:"menues,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
	Id       string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Url      string   `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Badge    *Badge   `protobuf:"bytes,4,opt,name=badge" json:"badge,omitempty"`
	Wrapper  *Wrapper `protobuf:"bytes,5,opt,name=wrapper" json:"wrapper,omitempty"`
	Title    bool     `protobuf:"varint,6,opt,name=title" json:"title,omitempty"`
	Children []*Menu  `protobuf:"bytes,7,rep,name=children" json:"children,omitempty"`
	Icon     string   `protobuf:"bytes,8,opt,name=icon" json:"icon,omitempty"`
	Roles    []*Role  `protobuf:"bytes,9,rep,name=roles" json:"roles,omitempty"`
}

func (m *Menu) Reset()                    { *m = Menu{} }
func (m *Menu) String() string            { return proto.CompactTextString(m) }
func (*Menu) ProtoMessage()               {}
func (*Menu) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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

type Badge struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Variant string `protobuf:"bytes,2,opt,name=variant" json:"variant,omitempty"`
	Text    string `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
}

func (m *Badge) Reset()                    { *m = Badge{} }
func (m *Badge) String() string            { return proto.CompactTextString(m) }
func (*Badge) ProtoMessage()               {}
func (*Badge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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
	Id        string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Element   string     `protobuf:"bytes,2,opt,name=element" json:"element,omitempty"`
	Atributes *Atributes `protobuf:"bytes,3,opt,name=atributes" json:"atributes,omitempty"`
}

func (m *Wrapper) Reset()                    { *m = Wrapper{} }
func (m *Wrapper) String() string            { return proto.CompactTextString(m) }
func (*Wrapper) ProtoMessage()               {}
func (*Wrapper) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

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

type Atributes struct {
}

func (m *Atributes) Reset()                    { *m = Atributes{} }
func (m *Atributes) String() string            { return proto.CompactTextString(m) }
func (*Atributes) ProtoMessage()               {}
func (*Atributes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*User)(nil), "auth.User")
	proto.RegisterType((*Request)(nil), "auth.Request")
	proto.RegisterType((*ResponseUser)(nil), "auth.ResponseUser")
	proto.RegisterType((*ResponseRole)(nil), "auth.ResponseRole")
	proto.RegisterType((*ResponseMenu)(nil), "auth.ResponseMenu")
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
	// 671 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x6e, 0x62, 0x3b, 0x4e, 0xc6, 0x2d, 0xa0, 0x55, 0x91, 0xac, 0x1e, 0x50, 0x6a, 0x24, 0x5a,
	0x81, 0x5a, 0xa4, 0xf4, 0xcc, 0xa1, 0xa0, 0xaa, 0xa7, 0x5e, 0x56, 0xfc, 0x5d, 0x37, 0xf1, 0x88,
	0x5a, 0x38, 0xb6, 0xd9, 0x5d, 0xb7, 0xf0, 0x26, 0x3c, 0x05, 0x4f, 0x08, 0x12, 0xda, 0xd9, 0xf5,
	0x4f, 0x02, 0xa9, 0x7c, 0x89, 0x77, 0xbe, 0x99, 0x9d, 0xf9, 0xbe, 0xdd, 0x99, 0x0d, 0x3c, 0xad,
	0x64, 0xa9, 0xcb, 0xd7, 0xa2, 0xd6, 0xb7, 0xf4, 0x73, 0x4e, 0x36, 0xf3, 0xcd, 0x3a, 0xf9, 0x35,
	0x02, 0xff, 0x83, 0x42, 0xc9, 0x1e, 0xc1, 0x38, 0x4b, 0xe3, 0xd1, 0x7c, 0x74, 0x3a, 0xe3, 0xe3,
	0x2c, 0x65, 0x0c, 0xfc, 0x42, 0xac, 0x31, 0x1e, 0x13, 0x42, 0x6b, 0x16, 0x43, 0xb8, 0x2a, 0xd7,
	0x95, 0x28, 0x7e, 0xc4, 0x1e, 0xc1, 0x8d, 0xc9, 0x0e, 0x21, 0xc0, 0xb5, 0xc8, 0xf2, 0xd8, 0x27,
	0xdc, 0x1a, 0xec, 0x08, 0xa6, 0x95, 0x50, 0xea, 0xbe, 0x94, 0x69, 0x1c, 0x90, 0xa3, 0xb5, 0xcd,
	0x0e, 0x5d, 0x7e, 0xc5, 0x22, 0x9e, 0xd8, 0x1d, 0x64, 0xb0, 0x39, 0x04, 0xb2, 0xcc, 0x51, 0xc5,
	0xe1, 0xdc, 0x3b, 0x8d, 0x16, 0x70, 0x4e, 0x84, 0x79, 0x99, 0x23, 0xb7, 0x8e, 0x64, 0x06, 0x21,
	0xc7, 0x6f, 0x35, 0x2a, 0x9d, 0xfc, 0x1c, 0xc1, 0x3e, 0x47, 0x55, 0x95, 0x85, 0x42, 0xd2, 0xf0,
	0x0c, 0xfc, 0x5a, 0xa1, 0x24, 0x15, 0xed, 0x66, 0xe3, 0xe1, 0x84, 0x9b, 0xec, 0xe6, 0xab, 0xe2,
	0x71, 0x3f, 0x3b, 0x05, 0x58, 0x07, 0x7b, 0x0e, 0x13, 0x94, 0xb2, 0x94, 0x2a, 0xf6, 0x28, 0x24,
	0xb2, 0x21, 0x57, 0x06, 0xe3, 0xce, 0xc5, 0x8e, 0x1b, 0xea, 0x3e, 0xd5, 0x71, 0x31, 0xef, 0x0d,
	0xe4, 0x74, 0x24, 0x75, 0xc7, 0xcc, 0x90, 0x37, 0xcc, 0x0c, 0xfd, 0x4d, 0x66, 0x24, 0x8b, 0xf0,
	0x4e, 0xf7, 0x78, 0x87, 0xee, 0x41, 0xcc, 0x92, 0xfb, 0xae, 0xec, 0x0d, 0x16, 0xb5, 0x29, 0xbb,
	0xc6, 0xa2, 0xde, 0x2c, 0x6b, 0x3c, 0x9c, 0x70, 0x96, 0xc0, 0xc4, 0x7c, 0xb7, 0xeb, 0x52, 0x84,
	0xf3, 0x0c, 0x2b, 0xfc, 0x19, 0x02, 0xd2, 0xdf, 0x5d, 0xeb, 0xa8, 0x7f, 0xad, 0x87, 0x10, 0xdc,
	0x89, 0x3c, 0x4b, 0xa9, 0x9b, 0xa6, 0xdc, 0x1a, 0xc3, 0x32, 0xbf, 0x81, 0x80, 0x00, 0xd3, 0x90,
	0xab, 0x32, 0xb5, 0x47, 0x18, 0x70, 0x5a, 0xb3, 0x39, 0x44, 0x29, 0xaa, 0x95, 0xcc, 0x2a, 0x9d,
	0x95, 0x85, 0xeb, 0xd5, 0x3e, 0x94, 0x54, 0xe0, 0xd3, 0x05, 0x0c, 0x69, 0xef, 0xb6, 0x3d, 0xbc,
	0x5d, 0xed, 0xd1, 0x9d, 0x97, 0xbf, 0xeb, 0xbc, 0x92, 0xdf, 0x23, 0xf0, 0xe9, 0xf0, 0x87, 0x94,
	0x7c, 0x02, 0x5e, 0x2d, 0x73, 0x37, 0x4d, 0x66, 0x69, 0x9a, 0x6b, 0x29, 0xd2, 0x2f, 0xb8, 0xd9,
	0x5c, 0x6f, 0x0d, 0xc4, 0xad, 0x87, 0x9d, 0x40, 0x78, 0x2f, 0x45, 0x55, 0xa1, 0xa4, 0xa9, 0x8a,
	0x16, 0x07, 0x36, 0xe8, 0x93, 0x05, 0x79, 0xe3, 0xa5, 0xcb, 0xc8, 0x74, 0x8e, 0x34, 0x63, 0x53,
	0x6e, 0x0d, 0xf6, 0x02, 0xa6, 0xab, 0xdb, 0x2c, 0x4f, 0x25, 0x16, 0x9b, 0x63, 0x46, 0x32, 0x5a,
	0x9f, 0xe1, 0x9b, 0xad, 0xca, 0x22, 0x9e, 0x5a, 0xbe, 0x66, 0xdd, 0xf5, 0xe9, 0x6c, 0xd7, 0x7c,
	0x5e, 0x41, 0x40, 0x64, 0xff, 0x91, 0x1f, 0x43, 0x78, 0x27, 0x64, 0x26, 0x0a, 0xed, 0x4e, 0xa0,
	0x31, 0x4d, 0x21, 0x8d, 0xdf, 0xb5, 0x3b, 0x05, 0x5a, 0x27, 0x4b, 0x08, 0x9d, 0x9c, 0xff, 0x25,
	0xc2, 0x1c, 0xd7, 0xd8, 0x25, 0x72, 0x26, 0x3b, 0x83, 0x99, 0xd0, 0x32, 0x5b, 0xd6, 0x1a, 0x15,
	0x65, 0x8b, 0x16, 0x8f, 0x2d, 0xc3, 0xcb, 0x06, 0xe6, 0x5d, 0x44, 0x12, 0xc1, 0xac, 0xc5, 0x17,
	0x7f, 0x3c, 0xf0, 0x2f, 0x6b, 0x7d, 0xcb, 0x5e, 0xc2, 0xe4, 0x9d, 0x44, 0xa1, 0x91, 0xf5, 0x1a,
	0xe0, 0x88, 0x39, 0xa5, 0xbd, 0xe7, 0x26, 0xd9, 0x63, 0x27, 0xe0, 0x5d, 0xa3, 0x1e, 0x10, 0x78,
	0x06, 0x93, 0x6b, 0xd4, 0x97, 0x79, 0xce, 0x0e, 0x1a, 0x3f, 0xbd, 0x61, 0x3b, 0xc2, 0x8f, 0x1d,
	0x97, 0x7e, 0xe2, 0xfe, 0x33, 0x93, 0xec, 0xb1, 0x57, 0x70, 0xf0, 0xd1, 0x4c, 0x91, 0xd0, 0x68,
	0x27, 0xaf, 0xef, 0xdf, 0x0e, 0x3e, 0x07, 0xb0, 0x9a, 0x68, 0x16, 0x7a, 0xb7, 0xb6, 0x5d, 0xdf,
	0x60, 0x94, 0x3c, 0xbc, 0x46, 0x3d, 0x30, 0x78, 0x01, 0x91, 0xd5, 0xc6, 0xe9, 0xa1, 0x7a, 0x58,
	0xa0, 0xdb, 0xd3, 0x12, 0xa2, 0x49, 0xe9, 0xf5, 0xdf, 0x76, 0xbc, 0xc1, 0x5a, 0x42, 0x03, 0x83,
	0x2f, 0x60, 0xdf, 0x12, 0xba, 0xb1, 0x2f, 0xd8, 0xc3, 0x8c, 0xec, 0xa6, 0xe5, 0x84, 0xfe, 0x15,
	0x2f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xab, 0xcf, 0xdb, 0xaf, 0x2e, 0x07, 0x00, 0x00,
}
