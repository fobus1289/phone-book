package common

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	SECRET   = "SECRET"
	EXPIRED  = "EXPIRED"
	DATABASE = "DATABASE"
)

var (
	secret   string
	database string
	expired  int64
)

func init() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	//SECRET
	secret = os.Getenv(SECRET)
	{
		if secret == "" {
			panic("secret can be empty")
		}
	}

	//DATABASE
	database = os.Getenv(DATABASE)
	{
		if database == "" {
			panic("database can be empty")
		}
	}

	//EXPIRED
	exp, err := strconv.ParseInt(os.Getenv(EXPIRED), 10, 64)
	{
		if err != nil {
			panic(err)
		}
	}

	expired = exp

}

func Secret() []byte {
	return []byte(secret)
}

func Expired() int64 {
	return expired
}

func Database() string {
	return database
}
