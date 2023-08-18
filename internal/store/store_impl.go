package store

import (
	"sync"

	"github.com/nduyphuong/gorya/internal/models"
	"github.com/nduyphuong/gorya/internal/os"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

var modelStorage *Storage
var muModelStorage sync.Mutex

func GetSingleton() (Interface, error) {
	muModelStorage.Lock()
	defer func() {
		muModelStorage.Unlock()
	}()
	if modelStorage != nil {
		return modelStorage, nil
	}
	dbType := os.GetEnv("GORYA_DB_TYPE", "sqlite")
	host := os.GetEnv("GORYA_DB_HOST", "localhost:3306")
	user := os.GetEnv("GORYA_DB_USER", "root")
	password := os.GetEnv("GORYA_DB_PASSWORD", "root")
	dbName := os.GetEnv("GORYA_DB_NAME", "gorya")
	switch dbType {
	case "mysql":
		db, err := NewMySQLDB(host, user, password, dbName)
		if err != nil {
			return nil, err
		}
		store := NewStorage(db)
		return store, nil
	default:
		db, err := NewSqliteDB()
		if err != nil {
			return nil, err
		}
		store := NewStorage(db)
		return store, nil
	}
}

func NewStorage(db *gorm.DB) Interface {
	return &Storage{
		db,
	}
}

func (c *Storage) SavePolicy(m models.PolicyModel) error {
	if err := c.db.Save(&m).Error; err != nil {
		return err
	}
	return nil
}

func (c *Storage) SaveSchedule(m models.ScheduleModel) error {
	if err := c.db.Save(&m).Error; err != nil {
		return err
	}
	return nil
}
