package db

import (
	"bytes"
	"fmt"
	"os/exec"
	"sync"

	"devops.aishu.cn/AISHUDevOps/AnyFabric/_git/data-masking/common/log"
	"github.com/jinguoxing/af-go-frame/core/options"
	"gorm.io/gorm"
)

var (
	once sync.Once
)

// Data .
type Data struct {
	// TODO wrapped database client
	DB *gorm.DB
}

// NewData .
func NewData(database *Database) (*Data, func(), error) {
	log.Info("暂时不需要使用数据库")
	return nil, nil, nil
	// var err error
	// var client *gorm.DB
	// once.Do(func() {
	// 	client, err = database.Default.NewMySqlClient()

	// })
	// if err != nil {
	// 	log.Errorf("open mysql failed, err: %v", err)
	// 	return nil, nil, err
	// }

	// if os.Getenv("init_db") == "true" {
	// 	if err = initDB(database); err != nil {
	// 		log.Errorf("init db failed, err: %v\n", err.Error())
	// 		return nil, nil, err
	// 	}
	// 	os.Exit(0)
	// }
	// return &Data{
	// 		DB: client,
	// 	}, func() {
	// 		log.Info("closing the data resources")
	// 	}, nil
}

type Database struct {
	Default  options.DBOptions `json:"default"`
	Default1 options.DBOptions `json:"default1"`
}

func initDB(database *Database) error {

	fileSource := "file:/usr/local/bin/af/infrastructure/repository/db/gen/migration"
	dns := fmt.Sprintf("mysql://%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", database.Default.Username, database.Default.Password, database.Default.Host, database.Default.Database)

	cmd := exec.Command("/usr/local/bin/af/infrastructure/repository/db/gen/migration/migrate",
		[]string{
			"-source",
			fileSource,
			"-database",
			dns,
			"up",
		}...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	log.Infof("out: %v\n", outStr)
	if err != nil {
		log.Errorf("err: %v\n", errStr)
		return err
	}
	return nil
}
