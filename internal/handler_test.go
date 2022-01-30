package internal_test

import (
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
			scenario:           "given no id in path it should return 404",
			givenRawURLPath:    "/?start=2020-01-01&end=2021-01-01",
			expectedStatusCode: http.StatusNotFound,
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

func startTestServerWithProfileHandler(aboutService internal.AboutService, submissionService internal.SubmissionService) *httptest.Server {
	e := echo.New()

	profileHandler := internal.NewProfileHandler(aboutService, submissionService, nil)
	profileHandler.RegisterRoutes(e)

	return httptest.NewServer(e.Server.Handler)
}
