package rdb

import (
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/learning-microservice/event/ddd/infrastructure/rdb/records"
)

var (
	once    sync.Once
	cacheDB *mySQL
)

func lazyInit() (err error) {
	once.Do(func() {
		var db *gorm.DB
		if db, err = gorm.Open("mysql", os.Getenv("DB_ADDRESS")); err != nil {
			return
		}
		if os.Getenv("DB_DEBUG") != "" {
			db.LogMode(true)
		}
		//db.SingularTable(true)
		if err = db.DB().Ping(); err != nil {
			return
		}

		cacheDB = &mySQL{DB: db}

		//if os.Getenv("DB_AUTO_MIGRATE") != "" {
		migration()
		//}

	})
	return
}

func migration() {
	tables := []interface{}{&records.Event{},
		&records.Assignment{},
		&records.Unassignment{},
		&records.Booking{},
		&records.Cancel{},
		&records.Control{},
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
}

////////////////////////////////////////////
// mySQL
////////////////////////////////////////////

type mySQL struct {
	*gorm.DB
}
