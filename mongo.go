package mongodb

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect initializes a new MongoDB client (and sends a Ping)
func Connect(config Config) (*mongo.Client, error) {
	if config.Host == "" || config.Port == "" {
		return nil, errors.New("empty config")
	}
	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		config.ToDataSourceURL()))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

// LoadEnvConfig loads config from environment variables:
// MONGO_HOST, MONGO_PORT, MONGO_ROOT_USERNAME, MONGO_ROOT_PASSWORD, MONGO_DATABASE
func LoadEnvConfig() Config {
	return Config{
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Database: os.Getenv("MONGO_DATABASE"),
	}
}

// Config can be loaded easily by calling func LoadEnvConfig
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string // schema (aka database) name
}

func (c Config) ToDataSourceURL() string {
	if c.Username == "" {
		return fmt.Sprintf("mongodb://%v:%v/%v", c.Host, c.Port, c.Database)
	} else {
		return fmt.Sprintf("mongodb://%v:%v@%v:%v/%v",
			c.Username, c.Password, c.Host, c.Port, c.Database)
	}
}
