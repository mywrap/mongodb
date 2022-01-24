package mongodb

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// this test need a local MongoDB server
func TestConnect(t *testing.T) {
	cfg := Config{Host: "127.0.0.1", Port: "27017", Database: "database0"}
	client, err := Connect(cfg)
	if err != nil {
		t.Fatalf("error Connect: %v", err)
	}
	ret, err := client.Database(cfg.Database).Collection("TestStruct").
		UpdateOne(context.TODO(),
			bson.M{"_id": "TestKey0"},
			bson.M{"$set": bson.M{"Value": "test value", "UpdatedAt": time.Now()}},
			options.Update().SetUpsert(true),
		)
	if err != nil {
		t.Errorf("error InsertOne: %#v", err)
	} else {
		t.Logf("upsertedID: %v", ret.UpsertedID)
	}
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
	// TODO: examples options find limit, offset, decode, aggregate (join)
}
