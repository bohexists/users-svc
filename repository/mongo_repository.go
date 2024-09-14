package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"

	"github.com/bohexists/users-svc/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(uri, dbName, collectionName string) *MongoRepository {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	collection := client.Database(dbName).Collection(collectionName)
	return &MongoRepository{collection: collection}
}

func (r *MongoRepository) CreateUser(user models.User) (string, error) {
	user.ID = primitive.NewObjectID().Hex()
	_, err := r.collection.InsertOne(context.TODO(), user)
	return user.ID, err
}

func (r *MongoRepository) GetUser(id string) (*models.User, error) {
	var user models.User
	filter := bson.M{"id": id}
	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	return &user, err
}

func (r *MongoRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	cursor, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}

func (r *MongoRepository) UpdateUser(id string, user models.User) error {
	filter := bson.M{"id": id}
	update := bson.M{"$set": user}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *MongoRepository) DeleteUser(id string) error {
	filter := bson.M{"id": id}
	_, err := r.collection.DeleteOne(context.TODO(), filter)
	return err
}
