package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-pg/pg/v10"
)

const (
	// DatabaseName is the name of the database
	DatabaseName = "nft"

	// UserName is the name of the user
	UserName = "nft"

	// Password is the password of the user
	Password = "nft"

	// Port is the port of the database
	Port = 5432

	// SSLMode is the ssl mode of the database
	SSLMode = "disable"
)

var (
	Host = setHost()

	// DatabaseURL is the url of the database
	DatabaseURL = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", UserName, Password, Host, Port, DatabaseName, SSLMode)
	db          = pg.Connect(&pg.Options{
		User:     UserName,
		Password: Password,
		Addr:     Host + ":" + strconv.Itoa(Port),

		Database: DatabaseName,
	})
)

// DB host for docker
func setHost() string {
	if os.Getenv("NFT_DB_HOST") != "" {
		return os.Getenv("NFT_DB_HOST")
	}
	return "localhost"

}
