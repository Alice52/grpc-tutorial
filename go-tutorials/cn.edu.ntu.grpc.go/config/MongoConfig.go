package config

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"time"
)

type MongoConfig struct {
	MongoHost string
	MongoPort string
	MongoDb   string
	Username  string
	Password  string
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

/**
init database connection with host, port,
[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
*/
func InitMongoSession(mongoConfig MongoConfig) *mgo.Database {
	info := &mgo.DialInfo{
		Addrs:    []string{mongoConfig.MongoHost},
		//Database: mongoConfig.MongoDb,
		Username: mongoConfig.Username,
		Password: mongoConfig.Password,
		Timeout:  60 * time.Second,
	}

	// connect to mongo
	s, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Printf("ERROR connecting mongo, %s ", err.Error())
		return nil
	}
	s.SetMode(mgo.Monotonic, true)
	s.Ping()

	return s.DB(mongoConfig.MongoDb)
}

func ClearSession(session *mgo.Session) {
	session.Close()
}

// usage
//var Session *mgo.Session
//
//var mongoConfig = MongoConfig{
//	MongoHost: GetEnv("MONGO_HOST", "101.132.45.28"),
//	MongoPort: GetEnv("MONGO_PORT", "27017"),
//	MongoDb:   GetEnv("MONGO_DB", "tutorials"),
//	Username:  GetEnv("MONGO_USER", "root"),
//	Password:  GetEnv("MONGO_PASS", "Yu125**8782?"),
//}
//
//func main() {
//	Session = InitMongoSession(mongoConfig)
//	Session.Ping()
//
//	sessionCopy := Session.Copy()
//	defer sessionCopy.Close()
//	var coll = sessionCopy.DB(mongoConfig.MongoDb).C("blog")
//	blog := &blogpb.Blog{}
//	coll.Find(bson.M{"uid": 1}).One(blog)
//}
