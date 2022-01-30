package mongo_test

import (
	"context"
	"log"
	"testing"

	"github.com/codigician/profile/internal/about"
	aboutmongo "github.com/codigician/profile/internal/about/mongo"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (a *AboutMongoTestSuite) TestSave() {
	ctx := context.Background()

	expectedAbout := &about.About{
		Personal: about.Personal{
			Firstname:   "Kaan Yuksel",
			Lastname:    "Bilgin",
			Email:       "bilgin.yuksel96@gmail.com",
			PhoneNumber: "+905326541234",
			Country:     "Turkey",
		},
	}

	id, err := a.abmongo.Save(ctx, expectedAbout)
	log.Println(id)
	fetchedMongoAbout := a.getAboutWithOriginalMongoClient(ctx, id)

	a.Nil(err)
	a.EqualValues(expectedAbout.Personal, fetchedMongoAbout.Personal)
	a.Nil(fetchedMongoAbout.Education)
	a.Nil(fetchedMongoAbout.WorkHistory)
	a.Nil(fetchedMongoAbout.Websites)
	a.Empty(fetchedMongoAbout.Headline)
	a.Empty(fetchedMongoAbout.Me)
}

func (a *AboutMongoTestSuite) TestUpdate() {
	ctx := context.Background()

	email := "gin.yl9@gmail.com"

	emptyAbout := aboutmongo.About{
		ID:       primitive.NewObjectID(),
		Personal: aboutmongo.Personal{Email: email},
	}
	id := a.saveAboutWithMongoClient(ctx, email, emptyAbout)

	expectedAbout := &about.About{
		Headline: "Updated headline",
		Me:       "Updated me",
		Personal: about.Personal{
			Firstname:   "Haa kl",
			Lastname:    "Blg",
			Email:       email,
			PhoneNumber: "+900000000000",
			Country:     "China",
		},
		Education: []about.Education{{
			School:      "Updated school",
			Program:     "Updated program",
			Degree:      "Updated degree",
			Description: "Some description",
			Current:     true,
		},
		},
		WorkHistory: []about.WorkHistory{{
			Company: "Updated company",
			Role:    "Software Engineer",
			Current: true,
		},
		},
		Websites: []about.Website{{
			Title: "Google",
			URL:   "https://www.google.com",
		}},
	}

	err := a.abmongo.Update(ctx, id, expectedAbout)

	fetchedMongoAbout := a.getAboutWithOriginalMongoClient(ctx, id)

	a.Nil(err)
	a.Equal(expectedAbout.Headline, fetchedMongoAbout.Headline)
	a.Equal(expectedAbout.Me, fetchedMongoAbout.Me)
	a.EqualValues(expectedAbout.Personal, fetchedMongoAbout.Personal)
	a.EqualValues(expectedAbout.Websites[0], fetchedMongoAbout.Websites[0])
	a.EqualValues(expectedAbout.WorkHistory[0], fetchedMongoAbout.WorkHistory[0])
	a.EqualValues(expectedAbout.Education[0], fetchedMongoAbout.Education[0])
}

func (a *AboutMongoTestSuite) TestGet() {
	ctx := context.Background()

	expectedAboutMongo := aboutmongo.About{
		ID:          primitive.NewObjectID(),
		Headline:    "headline",
		Me:          "me",
		Personal:    aboutmongo.Personal{Email: "some@mail.com"},
		Education:   []aboutmongo.Education{{School: "school"}},
		WorkHistory: []aboutmongo.WorkHistory{{Company: "company"}},
		Websites:    []aboutmongo.Website{{Title: "title", URL: "url"}},
	}
	id := a.saveAboutWithMongoClient(ctx, "some@mail.com", expectedAboutMongo)

	actualAbout, err := a.abmongo.Get(ctx, id)

	a.Nil(err)
	a.Equal(expectedAboutMongo.Personal.Email, actualAbout.Personal.Email)
	a.Equal(expectedAboutMongo.Headline, actualAbout.Headline)
	a.Equal(expectedAboutMongo.Me, actualAbout.Me)
	a.Equal(expectedAboutMongo.Education[0].School, actualAbout.Education[0].School)
	a.Equal(expectedAboutMongo.WorkHistory[0].Company, actualAbout.WorkHistory[0].Company)
	a.Equal(expectedAboutMongo.Websites[0].Title, actualAbout.Websites[0].Title)
}

func (a *AboutMongoTestSuite) saveAboutWithMongoClient(ctx context.Context, email string, abmongo aboutmongo.About) string {
	res, err := a.client.Database("profile").Collection("about").InsertOne(ctx, abmongo)
	if err != nil {
		a.Fail("inserting about: %v\n", err)
	}

	oid, _ := res.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}

func (a *AboutMongoTestSuite) getAboutWithOriginalMongoClient(ctx context.Context, id string) *aboutmongo.About {
	var ma aboutmongo.About
	oid, _ := primitive.ObjectIDFromHex(id)
	if err := a.client.Database("profile").Collection("about").
		FindOne(ctx, primitive.M{"_id": oid}).Decode(&ma); err != nil {
		a.Fail("fetching about: %v\n", err)
	}

	return &ma
}

type AboutMongoTestSuite struct {
	suite.Suite

	container testcontainers.Container

	client  *mongo.Client
	abmongo *aboutmongo.Mongo
}

func TestIntegrationAboutMongo(t *testing.T) {
	if testing.Short() {
		t.Skip("about mongo test skipping")
	}

	suite.Run(t, new(AboutMongoTestSuite))
}

func (s *AboutMongoTestSuite) SetupSuite() {
	ctx := context.Background()
	uri := "mongodb://localhost:27017"

	s.container = s.createMongoDBContainer(ctx)
	s.client = s.createMongoDBClient(ctx, uri)
	s.abmongo = aboutmongo.New(uri)
	if err := s.abmongo.Connect(ctx); err != nil {
		s.Fail("mongo connect: %v", err)
	}
}

func (s *AboutMongoTestSuite) TearDownSuite() {
	ctx := context.Background()

	if err := s.abmongo.Disconnect(ctx); err != nil {
		s.Fail("mongo disconnect: %v", err)
	}
	log.Println("mongo disconnected")

	if err := s.container.Terminate(ctx); err != nil {
		s.Fail("mongo container terminate: %v", err)
	}
	log.Println("mongo container terminated")
}

func (s *AboutMongoTestSuite) createMongoDBClient(ctx context.Context, uri string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		s.Fail("mongo client creation: %v\n", err)
	}

	if err := client.Connect(ctx); err != nil {
		s.Fail("mongo client connect: %v\n", err)
	}

	return client
}

func (s *AboutMongoTestSuite) createMongoDBContainer(ctx context.Context) testcontainers.Container {
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "mongo",
			ExposedPorts: []string{"27017:27017"},
		},
		Started: true,
	})
	if err != nil {
		s.Fail("container create: %v\n", err)
	}

	return container
}
