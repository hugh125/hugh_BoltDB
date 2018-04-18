package boltdb

import(
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
		return nil
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
			return err
		}
		return nil
	})
	return err
}

func (b* BoltDB) GetID() uint64{
	var iret uint64 //返回ID
	b.MyBoltDB.Update(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte(b.myBucket))		
		//自增
		id , _ := b.NextSequence()
		iret = id
		return nil
		})
	return iret
}

//插入一条记录
//传入用户信息表
//返回插入结果
func (b *BoltDB) InsertBucket(newUser *usertable.UserTable)bool{
	err := b.MyBoltDB.Update(func(tx *bolt.Tx) error{
		//创建数据表
		b := tx.Bucket([]byte(b.myBucket))
		//整理记录条
		buf ,err := json.Marshal(newUser)
		if err != nil{
			return err
		}
		//根据用户名，插入记录
		b.Put([]byte(newUser.Username), buf)
		return nil
	})
	if err != nil{
		return false
	}
	return true
}

//根据用户名，获取数据表记录
func (b *BoltDB) GetUser(key string) []uint8{
	var ret []uint8
	b.MyBoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(b.myBucket))
		//根据用户名，提取记录
		ret = b.Get([]byte(key))	//key = username
		return nil
	})
	return ret
}

//获取数据表全部信息
func (b *BoltDB) GetAllUser() (map[string] string){
	retMap := make(map[string] string)
	b.MyBoltDB.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte(b.myBucket))
		b.ForEach(func(k, v []byte) error{
			retMap[string(k)] = string(v)	//赋值给 map 并返回
			return nil
		})
		return nil
	})
	return retMap
}

//删除数据表
func (b *BoltDB) DeleteBucket() bool {
	err := b.MyBoltDB.Update(func(tx *bolt.Tx) error{
		err := tx.DeleteBucket([]byte(b.myBucket))
		if err != nil{
			return err
		}
		return nil
	})
	if err != nil{
		return false
	}
	return true
}