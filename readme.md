# MongoDB client
Quick config to connect MongoDB from environment vars.  
Wrapped [mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver.git).

## Usage

````go
cfg := Config{Host: "127.0.0.1", Port: "27017", Database: "paave_news"}
client, err := Connect(cfg)
if err != nil {
    log.Fatalf("error Connect: %v", err)
}
ret, err := client.Database(cfg.Database).Collection("TestStruct").
    UpdateOne(context.TODO(),
        bson.M{"_id": "TestKey0"},
        bson.M{"$set": bson.M{"Value": "test value", "UpdatedAt": time.Now()}},
        options.Update().SetUpsert(true),
    )
````
