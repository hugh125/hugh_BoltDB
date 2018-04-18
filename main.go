package main

import(
	"./boltdb"
)

var (
	myDBName = "./DATA/myDBName.db"	//数据库名
	myBucket = "myBucket"			//数据表名
)
func main(){
	// 1、创建数据库
	db := boltdb.NewBoltDB(myDBName, myBucket)

	// 2、打开数据表
	db.CreateBucket()

	// 3、开始循环等待命令行操作
	// 4、检测退出循环

	// 5、	关闭数据库
	defer db.Close()	//关闭数据库
}