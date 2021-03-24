package services

import(
	"context"
	s "github.com/user/models"
	"github.com/user/database"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
)

type User struct{}

var (
	db = database.DB
)

type Args struct {}

type ReplyGetAllAcl struct {
	DataMainAcl []s.MainAcl `json:data_main_acl`
}

func (u *User) GetAllAcl(ctx context.Context, args *Args, reply *ReplyGetAllAcl) error {
	var main_acl []s.MainAcl
	// "SELECT * FROM `main_acl` ORDER BY parent_id,sort ASC"
	db.Table("main_acl").Find(&main_acl).Order("parent_id,sort ASC")
	fmt.Println("GetAllAcl",main_acl)
	reply.DataMainAcl = main_acl
	return nil
}


type ResGetAllRole struct{
	Id string
	RoleName string
	RoleDesc string
	IsSystem string
	RoleTag string
	AclId string
}

type ArgsGetAllRole struct {
	Id int
}

type ReplyGetAllRole struct{
	DataAllRole []ResGetAllRole
}

func (u *User) GetAllRole(ctx context.Context, args *ArgsGetAllRole, reply *ReplyGetAllRole) error {
	fmt.Println("52-args id:",args.Id)
	var results []ResGetAllRole
	db.Table("main_role").Select("main_role.id, main_role.role_name, main_role.role_desc, main_role.is_system, main_role.role_tag, main_role_acl.acl_id").Joins("left join main_role_acl on main_role_acl.role_id = main_role.id").Where("is_delete != ? AND user_parent_id = ? ","1","0").Or("user_parent_id = ?",args.Id).Order("main_role.id ASC").Scan(&results)
	fmt.Println("getAllRole",results);
	reply.DataAllRole = results
	return nil
}


type ArgsAddRoleAcl struct {
	RoleName string
	RoleDesc string
	CreateUserId int
	UserParentId int
	AclId int
}

type ReplyAddRoleAcl struct{
	Id int
}

func (u *User) AddRoleAcl(ctx context.Context, args *ArgsAddRoleAcl, reply *ReplyAddRoleAcl) error {
	main_role := s.MainRole{RoleName:args.RoleName, RoleDesc:args.RoleDesc, CreateUserId:args.CreateUserId, UserParentId:args.UserParentId, CreateTime:time.Now()}
	tx := db.Begin()
	result := db.Create(&main_role)
	if result.Error != nil {
		tx.Rollback()
	}
	main_role_acl := s.MainRoleAcl{RoleId:main_role.Id, AclId:args.AclId}
	result = db.Create(&main_role_acl)
	if result.Error != nil {
		tx.Rollback()
	}else{
		tx.Commit()
	}
	reply.Id = main_role_acl.Id
	return nil
}