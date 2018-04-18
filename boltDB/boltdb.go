package boltdb

import(
	"fmt"
	"encoding/json"
	"github.com/boltdb/bolt"
	"../usertable"
)

type BoltDB struct{
	myDBName string
	myBucket string
	MyBoltDB *bolt.DB
}

//构造函数
//传入数据库名和数据表名
//返回结构体指针
func NewBoltDB(myDBName, myBucket string) *BoltDB{
	var b = new(BoltDB)
	b.myDBName = myDBName
	b.myBucket = myBucket

	//打开数据库
	db, err := bolt.Open(myDBName, 0600, nil)
	if err != nil{
		fmt.Println(err)
	}

	b.MyBoltDB = db

	return b
}

//关闭数据库
func (b *BoltDB)Close(){
	b.MyBoltDB.Close()
}

//创建数据表
func (b *BoltDB) CreateBucket() error {
	err := b.MyBoltDB.Update(func(tx *bolt.Tx) error{
		//如果不存在则创建数据表
		_, err := tx.CreateBucketIfNotExists([]byte(b.myBucket))
		if err != nil{
			fmt.Println(err.Error)
			return err
		}
		return nil
	})
	return err
}

//插入一条记录
//传入用户信息表
//返回插入结果
func (b *BoltDB) UpdateBucket(newUser *usertable.UserTable) error{
	err := b.MyBoltDB.Update(func(tx *bolt.Tx) error{
		//创建数据表
		b := tx.Bucket([]byte(b.myBucket))
		
		//自增
		id , _ := b.NextSequence()
		newUser.Id = id

		strID := fmt.Sprintf("%03d", id)
		fmt.Println(strID)
		if newUser.Username == ""{
			newUser.Username = "user_" + strID
		}
		if newUser.Password == ""{
			newUser.Password = "pswd_" + strID			
		}
		if newUser.Address == ""{
			newUser.Address = "addr_" + strID			
		}

		//整理记录条
		buf ,err := json.Marshal(newUser)
		if err != nil{
			return err
		}
		//根据用户名，插入记录
		b.Put([]byte(newUser.Username), buf)
		fmt.Println(string(buf))
		return nil
	})
	return err
}

//根据用户名，获取数据表记录
func (b *BoltDB) GetUser(key string) error{
	err := b.MyBoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(b.myBucket))
		//根据用户名，提取记录
		v := b.Get([]byte(key))	//key = username
		if string(v) != ""{
			fmt.Printf("key = %s, %s\n",key, v)
		}else{
			fmt.Printf("Not found (key = %s)\n", key)
		}
		//fmt.Println(v)
		return nil
	})
	return err
}

//获取数据表全部信息
func (b *BoltDB) GetUserofAll() {
	b.MyBoltDB.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte(b.myBucket))
		b.ForEach(func(k, v []byte) error{
			fmt.Printf("Key = %s, %s\n", k, v)
			return nil
		})
		return nil
	})
}

//删除数据表
func (b *BoltDB) DeleteBucket() error {
	err := b.MyBoltDB.Update(func(tx *bolt.Tx) error{
		err := tx.DeleteBucket([]byte(b.myBucket))
		if err != nil{
			fmt.Println(err)
			return err
		}
		return nil
	})
	return err
}