package db

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/gommon/log"
	mgo "gopkg.in/mgo.v2"
)

var (
	session      *mgo.Session
	databaseName = "hello-world"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
	databaseName = os.Getenv("DB_NAME")
	stgConnectionAtlas()
}

func stgConnection() {

	fmt.Println(os.Getenv("DB_HOST"))
	url := os.Getenv("DB_HOST")

	dialInfo, err := mgo.Dial(url)

	if err != nil {
		panic(err)
	}
	session = dialInfo

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

}
func stgConnectionAtlas() {

	fmt.Println(os.Getenv("DB_HOST"))
	url := os.Getenv("DB_HOST")
	tlsConfig := &tls.Config{}
	dialInfo, err := mgo.ParseURL(url)
	if err != nil {
		log.Fatal(err)
	}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		return tls.Dial("tcp", addr.String(), tlsConfig)
	}
	mgoSession, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err)
	}
	session = mgoSession
	session.SetMode(mgo.Monotonic, true)

}

func pullSession() *mgo.Session {
	return session.Copy()
}

// Ping connection
func Ping() error {
	sessionCopy := pullSession()
	defer sessionCopy.Close()
	return sessionCopy.Ping()
}
