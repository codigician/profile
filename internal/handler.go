package internal

import (
	"context"
	"github.com/codigician/profile/internal/about"
	"github.com/codigician/profile/internal/submission"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type AboutService interface {
	Get(ctx context.Context, id string) (*about.About, error)
}

type SubmissionService interface {
	FindAllBetween(ctx context.Context, start, end time.Time) ([]submission.Submission, error)
}

type AnalyticsService interface {
}

type ProfileHandler struct {
	aboutService      AboutService
	submissionService SubmissionService
	analyticsService  AnalyticsService
}

func NewProfileHandler(aboutService AboutService, submissionService SubmissionService, analyticsService AnalyticsService) *ProfileHandler {
	return &ProfileHandler{
		aboutService:      aboutService,
		submissionService: submissionService,
		analyticsService:  analyticsService,
	}
}

// RegisterRoutes ...
func (p *ProfileHandler) RegisterRoutes(router *echo.Echo) {
	router.GET("/", p.GetComplete)
}

// GetComplete returns the detailed profile information of the current user
func (p *ProfileHandler) GetComplete(c echo.Context) error {
	//bearerAuthorization := c.Request().Header.Get("Authorization")
	//splitBearerToken := strings.Split(bearerAuthorization, " ")
	//if len(splitBearerToken) != 2 {
	//	return echo.ErrBadRequest
	//}

	id := "5"
	start := time.Now()
	end := time.Now().Add(5 * time.Hour)
	submissions, err := p.submissionService.FindAllBetween(c.Request().Context(), start, end)
	if err != nil {
		return err
	}

	aboutMe, err := p.aboutService.Get(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ProfileRes{
		About: AboutRes{
			Headline:    aboutMe.Headline,
			Me:          aboutMe.Me,
			Firstname:   aboutMe.Personal.Firstname,
			Lastname:    aboutMe.Personal.Lastname,
			Email:       aboutMe.Personal.Email,
			PhoneNumber: aboutMe.Personal.PhoneNumber,
			Country:     aboutMe.Personal.Country,
		},
		Websites:    FromWebsites(aboutMe.Websites),
		Education:   FromEducation(aboutMe.Education),
		WorkHistory: FromWorkHistory(aboutMe.WorkHistory),
		Submissions: FromSubmissions(submissions),
	})
}

// GetPublicInformation returns the detailed public information of a given user
func (p *ProfileHandler) GetPublicInformation() {
	// fetch public information using about service
}

// GetSubmitHistory of the profile in given timespan
func (p *ProfileHandler) GetSubmitHistory() {

}
