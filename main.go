package main

import(
	"fmt"
	"strconv"
	"./boltdb"
	"./usertable"
)

var (
	//数据库名
	myDBName = "./DATA/myDBName.db"
	
	//数据表名
	myBucket = "myBucket"
	
	//错误信息
	errOut = "!!!!!!!!!!!! ------Command error------ !!!!!!!!!!!!!"
	
	//按键值
	myKey string
)

func main(){
	//创建数据库操作对象
	db := boltdb.NewBoltDB(myDBName, myBucket)


	//创建数据表
	db.CreateBucket()

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
		fmt.Println(cmdOut)

		fmt.Printf("Input your Operation Cmd：")
		fmt.Scanln(&myKey)
		key, err := strconv.Atoi(myKey)

		//输入有误，显示错误并跳过
		if err != nil{
			fmt.Println(errOut)
			continue
		}

		//创建用户表对象
		u := usertable.NewUser("","","")
		
		switch key{
		case 0:
			break
		case 1:
			db.UpdateBucket(u)	//插入一条记录
		case 2:
			var newUsername, newPassword, newAddress string
			//
			fmt.Printf("请输入，(输入格式：username, password, address)：")
			fmt.Scanln(&newUsername, &newPassword, &newAddress)
			//
			newUser := usertable.NewUser(newUsername, newPassword, newAddress)

			db.UpdateBucket(newUser) //插入一条记录
		case 3:
			var queryUserName string
			//
			fmt.Printf("Input your Query UserName：")
			fmt.Scanln(&queryUserName)

			db.GetUser(queryUserName)	//根据用户名，获取数据表记录
		case 4:
			db.GetUserofAll()			//获取数据表全部信息
		case 5:
			db.DeleteBucket()			//删除数据表
			db.CreateBucket()			//创建数据表
		default:
			fmt.Println(errOut)			//打印错误信息
		}

		if key == 0{
			break	//退出循环
		}
	}
	defer db.Close()	//关闭数据库
}