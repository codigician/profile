package about_test

import (
	"context"
	"testing"

	"github.com/codigician/profile/internal/about"
	mocks "github.com/codigician/profile/internal/mocks/about"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockRepository := mocks.NewMockRepository(gomock.NewController(t))
	service := about.NewService(mockRepository)

	personal := about.Personal{
		Firstname:   "firstname",
		Lastname:    "lastname",
		Email:       "hiko@gmail.com",
		PhoneNumber: "+905335555555",
		Country:     "Turkey",
	}

	expectedAbout := about.About{Personal: personal}
	mockRepository.EXPECT().Save(gomock.Any(), &expectedAbout).Return("", nil)

	_, err := service.Create(context.TODO(), personal)

	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	mockRepository := mocks.NewMockRepository(gomock.NewController(t))
	service := about.NewService(mockRepository)

	mockAbout := about.About{
		Headline: "headline",
		Me:       "me",
		Personal: about.Personal{
			Email: "kiki@gmail.com",
		},
	}
	mockRepository.EXPECT().Update(gomock.Any(), "kiki@gmail.com", &mockAbout).Return(nil)

	_ = service.Update(context.TODO(), "kiki@gmail.com", mockAbout)
}

func TestGet(t *testing.T) {
	mockRepository := mocks.NewMockRepository(gomock.NewController(t))
	service := about.NewService(mockRepository)

	mockAbout := &about.About{}
	mockRepository.EXPECT().Get(gomock.Any(), "5").Return(mockAbout, nil)

	actualAbout, _ := service.Get(context.TODO(), "5")

	assert.Equal(t, mockAbout, actualAbout)
}
