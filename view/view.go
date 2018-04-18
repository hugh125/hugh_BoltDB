package view

import(
	"fmt"
	"strconv"
	"../boltdb"
	"../usertable"
)
var (
	//数据库名
	myDBName = "./DATA/myDBName.db"	
	//数据表名
	myBucket = "myBucket"	
	//错误信息
	errOut = "!!!!!!!!!!!! ------Command error------ !!!!!!!!!!!!!"	
	//按键值
	myKey string	//
	newUid uint64
	newUsername, newPassword, newAddress string

)

//operation 操作对象
type Operation struct{
	Mydb *boltdb.BoltDB
	Myuser *usertable.UserTable
	isOK bool
}

//操作对象构造函数
func NewOperation(db *boltdb.BoltDB) *Operation{
	op := new(Operation)
	op.Mydb = db
	return op
}

//创建用户表对象
func (op *Operation) CreateUser(){
	newUser := new(usertable.UserTable)
	newUid :=  op.Mydb.GetID() //获取自增列
	strID := fmt.Sprintf("%03d", newUid)

	newUser.Id = newUid
	newUser.Username = "user_" + strID
	newUser.Password = "pswd_" + strID
	newUser.Address = "addr_" + strID
	op.Myuser = newUser
}

//命令行操作区
func (op *Operation)CmdLoop(){
	cmdOut := 
	`
	>>>>>>-----------------------------------------------------------
	***  1：写入一条记录(默认格式)
	***  2：写入一条记录(输入格式：username, password, address)
	***  3：获取一条记录
	***  4：获取全部记录
	***  5：删除数据表
	***  0：退出！
	>>>>>>-----------------------------------------------------------
	`
	for{
		op.isOK = false
		fmt.Println(cmdOut)

		fmt.Printf("Input your Operation Cmd：")
		fmt.Scanln(&myKey)
		key, err := strconv.Atoi(myKey)

		//输入有误，显示错误并跳过
		if err != nil{
			fmt.Println(errOut)
			continue
		}
		
		switch key{
		case 0:
			break
		case 1:
			op.CreateUser()
			op.isOK = op.Mydb.InsertBucket(op.Myuser)	//插入一条记录
			fmt.Println(op.Myuser.Print(op.isOK))
		case 2:
			//
			fmt.Printf("请输入，(输入格式：username, password, address)：")
			fmt.Scanln(&newUsername, &newPassword, &newAddress)
			if newUsername == ""{
				break
			}
			op.CreateUser()
			//
			op.Myuser.Username = newUsername
			op.Myuser.Password = newPassword
			op.Myuser.Address = newAddress

			op.isOK = op.Mydb.InsertBucket(op.Myuser)	//插入一条记录			
			fmt.Println(op.Myuser.Print(op.isOK))
		case 3:
			var queryUserName string
			//
			fmt.Printf("Input your Query UserName：")
			fmt.Scanln(&queryUserName)

			value := op.Mydb.GetUser(queryUserName)	//根据用户名，获取数据表记录
			if string(value) != ""{
				fmt.Printf("key = %s, %s\n",queryUserName, value)
			}else{
				fmt.Printf("Not found (key = %s)\n", queryUserName)
			}
		case 4:
			allUser := op.Mydb.GetAllUser()			//获取数据表全部信息
			for k := range allUser{
				fmt.Printf("key = %s, %s\n",k, allUser[k])
			}
		case 5:
			op.Mydb.DeleteBucket()			//删除数据表
			op.Mydb.CreateBucket()			//创建数据表
		default:
			fmt.Println(errOut)			//打印错误信息
		}

		if key == 0{
			break	//退出循环
		}
	}
}