package storage

// import (
// 	"context"
// 	"fmt"
// 	"search_engine/model"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type Storage struct {
// 	db *mongo.Database
// }

// const (
// 	DbName   = "scraper"
// 	CollName = "products"
// )

// func ConnectStorage(storagePath string) (*Storage, error) {
// 	clientOptions := options.Client().ApplyURI(storagePath)

// 	client, err := mongo.Connect(context.Background(), clientOptions)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = client.Ping(context.Background(), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	names, err := client.ListDatabaseNames(context.TODO(), bson.D{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	index := -1
// 	for i, name := range names {
// 		if name == DbName {
// 			index = i
// 			break
// 		}
// 	}

// 	if index == -1 {
// 		return nil, fmt.Errorf("database  %s not found", DbName)
// 	}

// 	return &Storage{db: client.Database(DbName)}, nil
// }

// func (s *Storage) SelectItems() ([]model.Item, error) {
// 	coll := s.db.Collection(CollName)
// 	filter := bson.M{}
// 	cursor, err := coll.Find(context.Background(), filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(context.Background())

// 	var items []model.Item
// 	for cursor.Next(context.Background()) {
// 		var event model.Item
// 		if err := cursor.Decode(&event); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, event)
// 	}
// 	return items, nil
// }

// func (s *Storage) CloseStorage() error {
// 	err := s.db.Client().Disconnect(context.Background())
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
