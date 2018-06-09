package db

import (
	"os"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/learning-microservice/event/mvc/models"
)

var (
	once    sync.Once
	cacheDB *gorm.DB
)

// DB is ...
func DB() *gorm.DB {
	if err := lazyInit(); err != nil {
		panic(err)
	}
	return cacheDB
}

// IsMySQLError is ...
func IsMySQLError(err error) bool {
	if _, ok := err.(*mysql.MySQLError); ok {
		//log.Printf("Number: %d", mysqlErr.Number)
		//log.Printf("Message: %s", mysqlErr.Message)
		//log.Printf("Error(): %s", mysqlErr.Error())
		return true
	}
	return false
}

func lazyInit() (err error) {
	once.Do(func() {
		if cacheDB, err = gorm.Open("mysql", os.Getenv("DB_ADDRESS")); err != nil {
			return
		}
		//if os.Getenv("DB_DEBUG") != "" {
		cacheDB.LogMode(true)
		//}
		//db.SingularTable(true)
		err = cacheDB.DB().Ping()

		// auto migration
		tables := []interface{}{
			&models.Unassignment{},
			&models.Assignment{},
			&models.Cancel{},
			&models.Booking{},
			&models.Event{},
		}
		cacheDB.DropTableIfExists(tables...)
		cacheDB.AutoMigrate(tables...)
		cacheDB.Exec(`
			CREATE OR REPLACE VIEW event_assignment_view AS 
			  SELECT ea.* FROM event_assignments ea
			  WHERE NOT EXISTS (
				SELECT 'e' FROM event_unassignments eu
				WHERE ea.id = eu.assignment_id
			  ) 
		`)
		cacheDB.Exec(`
			CREATE OR REPLACE VIEW event_booking_view AS 
			  SELECT eb.* FROM event_bookings eb
			  WHERE NOT EXISTS (
				SELECT 'e' FROM event_cancels ec
				WHERE eb.id = ec.booking_id
			  ) 
		`)
	})
	return
}
