package db

import (
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/learning-microservice/event/mvc/models"
)

var (
	once    sync.Once
	cacheDB *gorm.DB
)

func DB() *gorm.DB {
	if err := lazyInit(); err != nil {
		panic(err)
	}
	return cacheDB
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
