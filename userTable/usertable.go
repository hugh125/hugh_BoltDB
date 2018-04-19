package usertable

import(
	"fmt"
)

type UserTable struct{
	Id uint64
	Username string
	Password string
	Address string
}

//构造函数
//传入用户属性
//返回结构体指针
func NewUser(uid uint64, username, password, address string) *UserTable{
	var u = new(UserTable)
	u.Id = uid
	u.Username = username
	u.Password = password
	u.Address = address

	return u
}

//打印用户信息
func (u *UserTable) Print(isOK bool) string{
	if !isOK {
		return ""
	}
	return fmt.Sprintf("{id:%d, username:%s, password:%s, address:%s}", u.Id, u.Username, u.Password, u.Address)
}
