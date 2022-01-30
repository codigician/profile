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
	mockRepository.EXPECT().Save(gomock.Any(), &expectedAbout).Return(nil)

	err := service.Create(context.TODO(), personal)

	assert.Nil(t, err)
}
