package usertable

type UserTable struct{
	Id uint64
	Username string
	Password string
	Address string
}

//构造函数
//传入用户属性
//返回结构体指针
func NewUser(username, password, address string) *UserTable{
	var u = new(UserTable)
	u.Id = 0
	u.Username = username
	u.Password = password
	u.Address = address

	return u
}