package module

import (
	"github.com/aiteung/atdb"
	"os"
)

var MongoString string = os.Getenv("MONGOINTERN")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "Internship1",
}

var MongoConn = atdb.MongoConnect(MongoInfo)
