package mongo

import (
	"context"
	"fmt"

	"github.com/codigician/profile/internal/about"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	uri    string
	client *mongo.Client
}

func New(uri string) *Mongo {
	return &Mongo{uri: uri}
}

func (m *Mongo) Save(ctx context.Context, a *about.About) error {
	res, err := m.pa().InsertOne(ctx, FromAbout(a))
	id, _ := res.InsertedID.(primitive.ObjectID)
	fmt.Println(id) // TODO:
	return err
}

func (m *Mongo) Connect(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.uri))
	m.client = client
	return err
}

func (m *Mongo) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *Mongo) pa() *mongo.Collection {
	return m.client.Database("profile").Collection("about")
}
