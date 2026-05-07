package repository

import (
	"context"
	"time"

	"catalog/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	client  *mongo.Client
	db      string
	timeout time.Duration
}

func (r *mongoRepository) Find(code string) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	product := &domain.Product{}
	collection := r.client.Database(r.db).Collection("products")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Error Finding a catalogue item")
		}
		return nil, errors.Wrap(err, "repository research")
	}
	return product, nil
}

func (r *mongoRepository) Store(product *domain.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.db).Collection("products")
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"code":  product.Code,
			"name":  product.Name,
			"price": product.Price,
		},
	)
	if err != nil {
		return errors.Wrap(err, "Error writing to repository")
	}
	return nil
}

func (r *mongoRepository) Update(product *domain.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.db).Collection("products")
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"code": product.Code},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "name", Value: product.Name}, {Key: "price", Value: product.Price}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoRepository) FindAll() ([]*domain.Product, error) {
	var items []*domain.Product

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.db).Collection("products")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, "repository find all")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var item domain.Product
		if err := cur.Decode(&item); err != nil {
			return nil, errors.Wrap(err, "repository decode")
		}
		items = append(items, &item)
	}
	return items, nil
}

func (r *mongoRepository) Delete(code string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"code": code}

	collection := r.client.Database(r.db).Collection("products")
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func newMongoClient(mongoServerURL string, timeout int) (*mongo.Client, error) {
	clientOptions := options.Client().
		ApplyURI(mongoServerURL)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// NewMongoRepository  mongoDB adapter is created by constructing a mongo repository based on the mongoRepository struct
func NewMongoRepository(mongoServerURL, mongoDb string, timeout int) (domain.Repository, error) {
	mongoClient, err := newMongoClient(mongoServerURL, timeout)
	repo := &mongoRepository{
		client:  mongoClient,
		db:      mongoDb,
		timeout: time.Duration(timeout) * time.Second,
	}
	if err != nil {
		return nil, errors.Wrap(err, "client error")
	}

	return repo, nil
}
