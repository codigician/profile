package mongo

import (
	"context"

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

func (m *Mongo) Save(ctx context.Context, a *about.About) (string, error) {
	res, err := m.pa().InsertOne(ctx, fromAbout(a, primitive.NewObjectID()))
	id, _ := res.InsertedID.(primitive.ObjectID)
	return id.Hex(), err
}

func (m *Mongo) Get(ctx context.Context, id string) (*about.About, error) {
	oid, _ := primitive.ObjectIDFromHex(id)

	res := m.pa().FindOne(ctx, primitive.M{"_id": oid})

	var a About
	err := res.Decode(&a)
	return a.to(), err
}

// TOOD: maybe put email outside of the nested struct
// to make it easier to query and index
func (m *Mongo) Update(ctx context.Context, id string, a *about.About) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := primitive.M{"_id": oid}
	_, err := m.pa().UpdateOne(ctx, filter, primitive.M{"$set": fromAbout(a, oid)})
	return err
}

func (m *Mongo) pa() *mongo.Collection {
	return m.client.Database("profile").Collection("about")
}

func (m *Mongo) Connect(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.uri))
	m.client = client
	return err
}

func (m *Mongo) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
