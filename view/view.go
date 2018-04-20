package view

import(
	"fmt"
	"strconv"
	"../boltdb"
	"../usertable"
)
var (
	//错误信息
	errOut = "!!!!!!!!!!!! ------Command error------ !!!!!!!!!!!!!"	
	//按键值
	myKey string	//按键值

	newUsername, newPassword, newAddress string
)

//operation 操作对象
type Operation struct{
	Mydb *boltdb.BoltDB
	Myuser *usertable.UserTable
	isOK bool

	cmdOut string 
}

//操作对象构造函数
func NewOperation(db *boltdb.BoltDB) *Operation{
	op := new(Operation)
	op.Mydb = db
	op.cmdOut = 
	`
	>>>>>>-----------------------------------------------------------
	***  1: insert into DataDB (default format)
	***  2: insert into DataDB (custom  format : username, password, address)
	***  3: get one user information
	***  4: get all user information
	***  5: deleta DataDB
	***  0: Exit ! 
	>>>>>>-----------------------------------------------------------
	`
	return op
}

//创建用户表对象
func (op *Operation) createUser(){
	newUser := new(usertable.UserTable)
	newUser.Id =  op.Mydb.GetID() //获取自增列
	strID := fmt.Sprintf("%03d", newUser.Id)

	newUser.Username = "user_" + strID
	newUser.Password = "pswd_" + strID
	newUser.Address = "addr_" + strID
	op.Myuser = newUser
}

//case1，默认插入数据操作
func (op *Operation)insertDefaultInfor(){
		op.createUser()	//创建用户表对象
		op.isOK = op.Mydb.InsertBucket(op.Myuser)	//插入一条记录
		fmt.Println(op.Myuser.Print(op.isOK))
}

//case2，自定义插入数据操作
func (op *Operation)insertCustomInfor(){
		fmt.Printf("Please input you data, (eg.format: username, password, address): ")
		fmt.Scanln(&newUsername, &newPassword, &newAddress)
		if newUsername == ""{
			return
		}
		op.createUser()	//创建用户表对象
		//自定义用户信息
		op.Myuser.Username = newUsername
		op.Myuser.Password = newPassword
		op.Myuser.Address = newAddress

		op.isOK = op.Mydb.InsertBucket(op.Myuser)	//插入一条记录			
		fmt.Println(op.Myuser.Print(op.isOK))
}

//case3，按用户名 查找记录信息
func (op *Operation)getOneInforofUserName(){
		var queryUserName string
		//
		fmt.Printf("Input your Query UserName: ")
		fmt.Scanln(&queryUserName)

		value := op.Mydb.GetUser(queryUserName)	//根据用户名，获取数据表记录
		if string(value) != ""{
			fmt.Printf("key = %s, %s\n",queryUserName, value)
		}else{
			fmt.Printf("Not found (key = %s)\n", queryUserName)
		}
}

//case4，遍历数据表
func (op *Operation)getAllInfor(){
		allUser := op.Mydb.GetAllUser()			//获取数据表全部信息
		for k := range allUser{
			fmt.Printf("key = %s, %s\n",k, allUser[k])
		}
}

//功能区显示和命令行输入
func (op *Operation)cmdLoopOP() int{
	op.isOK = false
	fmt.Println(op.cmdOut)

	fmt.Printf("Input your Operation Cmd: ")
	fmt.Scanln(&myKey)
	key, err := strconv.Atoi(myKey)

	//输入有误，显示错误并跳过
	if err != nil{
		fmt.Println(errOut)
		return 0
	}
	return key
}

//命令行操作区
func (op *Operation)Run(){
	for{
		key := op.cmdLoopOP()	//功能区显示和命令行输入
		switch key{
		case 0:
			break
		case 1:
			op.insertDefaultInfor()		//默认插入数据操作
		case 2:
			op.insertCustomInfor()		//自定义插入数据操作
		case 3:
			op.getOneInforofUserName()		//按用户名 查找记录信息
		case 4:
			op.getAllInfor()		//遍历数据表
		case 5:
			op.Mydb.DeleteBucket()		//删除数据表
			op.Mydb.CreateBucket()		//创建数据表
		default:
			fmt.Println(errOut)			//打印错误信息
		}

		if key == 0{
			break	//退出循环
		}
	}
}