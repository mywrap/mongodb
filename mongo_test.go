package mongodb

import (
	"context"
	"math"
	"testing"
)

// this test need a local running MongoDB server
func TestConnect(t *testing.T) {
	//c := Config{Host: "127.0.0.1", Port: "27017", Database: "database0"}
	c := LoadEnvConfig()
	client, err := Connect(c)
	if err != nil {
		t.Fatalf("error Connect: %v", err)
	}

	qr, err := client.Database(c.Database).Collection("testColl0").
		InsertOne(context.Background(),
			map[string]interface{}{
				"FullName": "Đào Thị Lán",
				"Phone":    "09xxx28543",
				"NetWorth": 10 * math.Pow10(9),
			},
		)
	if err != nil {
		t.Errorf("error InsertOne: %v", err)
	}
	t.Logf("insertedId: %v", qr.InsertedID)
}

func TestConfig_ToDataSourceURL(t *testing.T) {
	c := Config{
		Host:     "127.0.0.1",
		Port:     "27017",
		Username: "tungdt",
		Password: "MNrZ_3RbckGM",
		Database: "database1",
	}
	e := "mongodb://tungdt:MNrZ_3RbckGM@127.0.0.1:27017/database1"
	r := c.ToDataSourceURL()
	if r != e {
		t.Errorf("error ToDataSourceURL: real: %v, expected: %v", r, e)
	}
	if false {
		_, err := Connect(c)
		t.Log(err)
	}
}

func Test_BasicQuery(t *testing.T) {
	// TODO: examples options find limit, offset, decode, aggregate "join"
}