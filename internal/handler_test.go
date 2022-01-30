package internal_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/codigician/profile/internal"
	"github.com/codigician/profile/internal/about"
	"github.com/codigician/profile/internal/mocks"
	"github.com/codigician/profile/internal/submission"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetComplete(t *testing.T) {
	mockAboutService := mocks.NewMockAboutService(gomock.NewController(t))
	mockSubmissionService := mocks.NewMockSubmissionService(gomock.NewController(t))

	srv := startTestServerWithProfileHandler(mockAboutService, mockSubmissionService)
	defer srv.Close()

	testCases := []struct {
		scenario           string
		givenRawURLPath    string
		mockAboutErr       error
		mockSubmissionErr  error
		expectedID         string
		expectedStatusCode int
		expectedStartTime  time.Time
		expectedEndTime    time.Time
	}{
		{
			scenario:           "given valid query string, call service with correct parameters and return 200",
			givenRawURLPath:    "/5uid?start=2020-01-01&end=2021-01-01",
			expectedID:         "5uid",
			expectedStartTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedEndTime:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedStatusCode: http.StatusOK,
		},
		{
			// it will return 405 because when you don't provide id
			// it will be the same with create path
			scenario:           "given no id in path it should return 405",
			givenRawURLPath:    "/?start=2020-01-01&end=2021-01-01",
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
		{
			scenario:           "given invalid start time, return 400",
			givenRawURLPath:    "/5uid?start=2020-31-01&end=2021-01-01",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			scenario:           "given invalid end time, return 400",
			givenRawURLPath:    "/5uid?start=2020-01-01&end=2021-3s-01",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			scenario:           "given valid query string, submission service fails return 500",
			givenRawURLPath:    "/839?start=2021-01-01&end=2022-01-01",
			mockSubmissionErr:  assert.AnError,
			expectedID:         "839",
			expectedStartTime:  time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedEndTime:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			scenario:           "given valid query string, about service fails return 500",
			givenRawURLPath:    "/sample-id?start=2019-01-01&end=2020-01-01",
			mockAboutErr:       assert.AnError,
			expectedID:         "sample-id",
			expectedStartTime:  time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedEndTime:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.scenario, func(t *testing.T) {
			mockAboutService.EXPECT().
				Get(gomock.Any(), tC.expectedID).
				Return(&about.About{}, tC.mockAboutErr).
				AnyTimes()

			mockSubmissionService.EXPECT().
				FindAllBetween(gomock.Any(), tC.expectedStartTime, tC.expectedEndTime).
				Return([]submission.Submission{}, tC.mockSubmissionErr).
				AnyTimes()

			res, _ := http.Get(srv.URL + tC.givenRawURLPath)

			assert.Equal(t, tC.expectedStatusCode, res.StatusCode)
		})
	}
}

func TestCreate(t *testing.T) {
	mockAboutService := mocks.NewMockAboutService(gomock.NewController(t))

	srv := startTestServerWithProfileHandler(mockAboutService, nil)
	defer srv.Close()

	testCases := []struct {
		scenario           string
		givenRequest       interface{}
		mockErr            error
		expectedPersonal   about.Personal
		expectedStatusCode int
	}{
		{
			scenario:           "given invalid request body return 400",
			givenRequest:       "invalid body",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			scenario:           "given valid request body service returns error return 500",
			givenRequest:       internal.CreateProfileReq{Firstname: "kaan", Email: "gigi@mail.com"},
			mockErr:            assert.AnError,
			expectedPersonal:   about.Personal{Firstname: "kaan", Email: "gigi@mail.com"},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			scenario:           "given valid request body service returns success return 201",
			givenRequest:       internal.CreateProfileReq{Firstname: "yuksel", Email: "bobo@gmail.com"},
			expectedPersonal:   about.Personal{Firstname: "yuksel", Email: "bobo@gmail.com"},
			expectedStatusCode: http.StatusCreated,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.scenario, func(t *testing.T) {
			mockAboutService.EXPECT().
				Create(gomock.Any(), tC.expectedPersonal).
				Return("", tC.mockErr).
				AnyTimes()

			bytesReq, _ := json.Marshal(tC.givenRequest)
			res, _ := http.Post(srv.URL, "application/json", bytes.NewBuffer(bytesReq))

			assert.Equal(t, tC.expectedStatusCode, res.StatusCode)
		})
	}
}

func startTestServerWithProfileHandler(aboutService internal.AboutService, submissionService internal.SubmissionService) *httptest.Server {
	e := echo.New()

	profileHandler := internal.NewProfileHandler(aboutService, submissionService, nil)
	profileHandler.RegisterRoutes(e)

	return httptest.NewServer(e.Server.Handler)
}
