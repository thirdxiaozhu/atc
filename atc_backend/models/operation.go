package models

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	shell "github.com/ipfs/go-ipfs-api"

	"gorm.io/driver/mysql" // gorm mysql 驱动包
	"gorm.io/gorm"         // gorm
)

type DBHandle struct {
	orm orm.Ormer
}

var (
	DBH    DBHandle
	config Config
	sh     *shell.Shell

	G_db *gorm.DB
	err  error
)

func init() {
	logger = logs.GetLogger()
	// read config
	config = ReadConfig("/root/go/src/atc/atc_backend/models/config.json")
	initDBH()
	initIPFS()
	initGorm()

	for _, com := range getCompanies() {
		DBH.Insert(&com)
	}

}

func initDBH() {
	// set to debug mode
	orm.Debug = true
	// register model and database
	databaseUrl := config.SqlUser + ":" + config.SqlPassword + "@tcp(172.17.0.2:" +
		strconv.Itoa(config.SqlPort) + ")/" + config.SqlDatabase + "?charset=utf8"

	orm.RegisterDataBase("default", "mysql", databaseUrl)
	orm.RegisterModel(new(User), new(Company))
	// automatically build table
	orm.RunSyncdb("default", false, true)

	// create orm object with specified database as recommended in the document
	DBH.orm = orm.NewOrmUsingDB("default")
}

func initIPFS() {
	sh = shell.NewShell("localhost:5001")
}

func initGorm() {
	// MySQL 配置信息
	username := "root"  // 账号
	password := "root"  // 密码
	host := "127.0.0.1" // 地址
	port := 3306        // 端口
	DBname := "atc"     // 数据库名称
	timeout := "10s"    // 连接超时，10秒
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, DBname, timeout)
	// Open 连接
	G_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect mysql.")
	}
	G_db.AutoMigrate(&Arn{})
	G_db.AutoMigrate(&Route{})
	G_db.AutoMigrate(&Auth{})
	G_db.AutoMigrate(&Flight{})
	G_db.AutoMigrate(&Link{})
	G_db.Create(getArns())
	G_db.Create(getRoutes())
	G_db.Create(getAuths())
}

// @Title Insert
// @Description insert new data
// @Param mode: must be a pointer
func (this *DBHandle) Insert(mode interface{}) bool {
	_, err := this.orm.Insert(mode)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// @Title Delete
// @Description delete by primary key
// @Param mode: must be a pointer
func (this *DBHandle) Delete(mode interface{}) {
	_, err := this.orm.Delete(mode)
	if err != nil {
		fmt.Println(err)
	}
}

// @Title Update
// @Description update by primary key
// @Param mode: must be a pointer
func (this *DBHandle) Update(mode interface{}) {
	_, err := this.orm.Update(mode)
	if err != nil {
		fmt.Println(err)
	}
}

// @Title QueryByPK
// @Description query by primary key
//
//	remenber to specify the primary key for mode before pass it
//
// @Param mode: must be a pointer
func (this *DBHandle) QueryByPK(mode interface{}) {
	err := this.orm.Read(mode)
	if err != nil {
		fmt.Println(err)
	}
}

// @Title QueryByField
// @Description query by field
// @Param mode: must be a pointer
func (this *DBHandle) QueryAllByField(mode interface{}, table string, field string,
	value interface{}) error {
	_, err := this.orm.QueryTable(table).Filter(field, value).All(mode)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (this *DBHandle) QueryOneByField(mode interface{}, table string, field string,
	value interface{}) error {
	err := this.orm.QueryTable(table).Filter(field, value).One(mode)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (this *DBHandle) QueryAll(mode interface{}, table string) error {
	_, err := this.orm.QueryTable(table).All(mode)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UploadIPFS(str string) (string, error) {

	hash, err := sh.Add(bytes.NewBufferString(str))
	if err != nil {
		fmt.Println("上传ipfs时错误：", err)
		return "", err
	}
	return hash, nil
}

// 从ipfs下载数据
func CatIPFS(hash string) string {
	//sh = shell.NewShell("localhost:5001")

	read, err := sh.Cat(hash)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(read)

	return string(body)
}
