package storage

import (
	"labix.org/v2/mgo"
)

var (
	dbName = "avengers"
)

func collection(col string, xx func(*mgo.Collection)) {
	session, err := mgo.Dial("server1.example.com,server2.example.com")
    if err != nil {
        panic(err)
    }
	defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)
	
	c := session.DB(dbName).C(col)

	xx(c);
}