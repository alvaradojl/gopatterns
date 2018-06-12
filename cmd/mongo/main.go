package main

import (
	"fmt"
	"log"
	"strconv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	ApiKey  string        `bson:"apiKey"`
	IsValid bool          `bson:"isValid"`
}

const (
	Host         = "51.144.255.72:27017"
	Username     = "admin"
	Password     = "login123"
	AuthDatabase = "admin"
	DocDatabase  = "security"
	Collection   = "securityv1"
)

func main() {
	fmt.Println("starting...")

	session, err := GetSession()

	if err != nil {
		log.Fatal(err)
	}

	ReadDoc(session)

	defer session.Close()

}

func GetSession() (*mgo.Session, error) {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{Host},
		Username: Username,
		Password: Password,
		Database: AuthDatabase,
		Direct:   true,
	})

	fmt.Println("reading documents from collection...")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nConnected to %v!\n", session.LiveServers())

	session.SetMode(mgo.Monotonic, true)

	return session, err
}

func ReadDoc(session *mgo.Session) {
	c := *session.DB(DocDatabase).C(Collection)

	item := &Item{}

	err := c.Find(bson.M{"apiKey": "00001"}).One(&item)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Id: " + item.Id.Hex() + " ApiKey: " + item.ApiKey + " isValid: " + strconv.FormatBool(item.IsValid))
}
