package layer

import (
	"fmt"
	"gorm.io/gorm"
)

type IDao interface {
	IFlow
	SetDB(*gorm.DB)
	GetDB() *gorm.DB
}

var (
	DefaultDBClient *gorm.DB
)

type Dao struct {
	Flow
	db           *gorm.DB
	defaultDB    *gorm.DB
	tableName    string
	partitionNum int
}

func (entity *Dao) SetTableName(tableName string) {
	entity.tableName = tableName
}

func (entity *Dao) GetTableName() string {
	return entity.tableName
}

func (entity *Dao) SetDB(db *gorm.DB) {
	entity.db = db
}

func (entity *Dao) GetDB() (db *gorm.DB) {
	if entity.db != nil {
		db = entity.db.WithContext(entity.GetContext())
	} else if entity.defaultDB != nil {
		db = entity.defaultDB.WithContext(entity.GetContext())
	} else if DefaultDBClient != nil {
		db = DefaultDBClient.WithContext(entity.GetContext())
	}

	if len(entity.tableName) > 0 {
		db = db.Table(entity.tableName)
	}

	return
}

func (entity *Dao) SetPartitionNum(num int) {
	entity.partitionNum = num
}

func (entity *Dao) GetPartitionNum() int {
	return entity.partitionNum
}

func (entity *Dao) GetPartitionTable(value int) string {
	return fmt.Sprintf("%s%d", entity.GetTableName(), value%entity.partitionNum)
}

func SetDefaultDB(db *gorm.DB) {
	DefaultDBClient = db
}
