import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

var gormConn *gorm.DB

func GetDatabaseConnection() *gorm.DB {
	if gormConn != nil && gormConn.DB() != nil && gormConn.DB().Ping() == nil {
		return gormConn
	}

	conn, err := gorm.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_CONNECTION"))
	if err != nil {
		logrus.Fatal("Could not connect to the database")
	}

	gormConn = conn

	return gormConn
}