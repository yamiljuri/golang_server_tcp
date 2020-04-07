package database

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"

	"github.com/yamiljuri/server_tcp/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once              sync.Once
	mongoDBConnection MongoDBConnection
	DB                MongoDBConnection
)

type MongoDBConnection struct {
	Host     string
	Port     string
	Driver   string
	Database string
	client   *mongo.Client
}

func Default() DBConnection {
	once.Do(func() {
		mongoDBConnection = MongoDBConnection{
			Host:     config.Getenv("MONGO_DB_HOST"),
			Port:     config.Getenv("MONGO_DB_PORT"),
			Database: config.Getenv("MONGO_DB_DATABASE"),
			Driver:   config.Getenv("MONGO_DB_DRIVER"),
		}
	})
	DB = mongoDBConnection
	return &mongoDBConnection
}

func (m *MongoDBConnection) Connect() DBConnection {
	clientsOpts := options.Client().ApplyURI(fmt.Sprintf("%s://%s:%s", m.Driver, m.Host, m.Port))
	client, err := mongo.Connect(context.TODO(), clientsOpts)
	if err != nil {
		log.Fatalf("Error %s DB Connect %v", m.Driver, err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Error %s DB Ping %v", m.Driver, err)
	}
	m.client = client
	return &mongoDBConnection
}

func (m *MongoDBConnection) Disconnect() {
	err := m.client.Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("Error %s DB Disconnect %v", m.Driver, err)
	}
}

func (m *MongoDBConnection) Insert(objects ...interface{}) {
	if objects != nil {
		m.Connect()
		defer m.Disconnect()
		for _, object := range objects {
			collectionName := reflect.ValueOf(object).Type().Name()
			m.client.Database(m.Database).Collection(strings.ToLower(collectionName)).InsertOne(context.TODO(), object)
		}
	}
}
