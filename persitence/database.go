package persitence

import (
	. "doescher.ninja/twitter-service/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var conn *gorm.DB

type connectionConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func InitDatabase() {
	err := connect()
	FatalIfError(err)

	err = migrate()
	FatalIfError(err)
}

func generateConnectionString(c connectionConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin",
		c.Host,
		c.User,
		c.Password,
		c.Database,
		c.Port,
	)
}

func connect() error {
	c := Conf()
	var err error
	conn, err = gorm.Open(postgres.Open(generateConnectionString(connectionConfig{
		c.DbHost,
		c.DbPort,
		c.DbUser,
		c.DbPass,
		c.DbName,
	})), &gorm.Config{})
	return err
}

func migrate() error {
	return conn.AutoMigrate(&Profile{}, &Tweet{})
}

func getDb() *gorm.DB {
	return conn
}