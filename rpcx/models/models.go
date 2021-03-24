package models

type CustUser struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Mobile string `json:"mobile"`
	Avatar string `json:"avatar"`
	RegSource string `json:"reg_source"`
	IsBlock int `json:"is_block"`
	IsDelete int `json:"is_delete"`
	UpdateTime int `json:"update_time"`
	CreateTime int `json:"create_time"`
}

type CustUserBind struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	BindType int `json:"bind_type"`
	BindOpenid string `json:"bind_openid"`
	BindMobile string `json:"bind_mobile"`
	BindNo string `json:"bind_no"`
	BindName string `json:"bind_name"`
	BindAvatar string `json:"bind_avatar"`
	BindPlateNo string `json:"bind_plate_no"`
	IsDelete int `json:"is_delete"`
	UpdateTime int `json:"update_time"`
	CreateTime int `json:"create_time"`
}

type MainAcl struct {
	Id int `json:"id"`
	ParentId int `json:"parent_id"`
	AclTag string `json:"acl_tag"`
	AclName string `json:"acl_name"`
	AclDesc string `json:"acl_desc"`
	IsSuper int `json:"is_super"`
	IsBlock int `json:"is_block"`
	Sort int `json:"sort"`
	UpdateTime int `json:"update_time"`
	CreateTime int `json:"create_time"`
}

type MainRole struct {
	Id int `json:"id"`
	RoleTag string `json:"role_tag"`
	RoleName string `json:"role_name"`
	RoleDesc string `json:"role_desc"`
	UserParentId int `json:"user_parent_id"`
	CreateUserId int `json:"create_user_id"`
	IsDelete int `json:"is_delete"`
	IsSystem int `json:"is_system"`
	UpdateTime int `json:"update_time"`
	CreateTime int `json:"create_time"`
}

type MainRoleAcl struct {
	Id int `json:"id"`
	RoleId int `json:"role_id"`
	AclId string `json:"acl_id"`
	UpdateTime int `json:"update_time"`
	CreateTime int `json:"create_time"`
}