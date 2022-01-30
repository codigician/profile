package internal

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/codigician/profile/internal/about"
	"github.com/codigician/profile/internal/submission"
	"github.com/labstack/echo/v4"
)

type AboutService interface {
	Get(ctx context.Context, id string) (*about.About, error)
	Create(ctx context.Context, personal about.Personal) (string, error)
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
	router.GET(":id", p.GetComplete)
	router.POST("", p.Create)
}

func (p *ProfileHandler) Create(c echo.Context) error {
	var req CreateProfileReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	id, err := p.aboutService.Create(c.Request().Context(), req.toPersonal())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, CreateProfileRes{id})
}

// GetComplete returns the detailed profile information of the current user
func (p *ProfileHandler) GetComplete(c echo.Context) error {
	id := c.Param("id")

	startTimeString := c.QueryParam("start")
	endTimeString := c.QueryParam("end")

	start, err := ParseTime(startTimeString)
	if err != nil {
		log.Printf("start time %s is not valid: %v\n", startTimeString, err)
		return echo.NewHTTPError(http.StatusBadRequest, "start time is not valid")
	}

	end, err := ParseTime(endTimeString)
	if err != nil {
		log.Printf("end time %s is not valid: %v\n", endTimeString, err)
		return echo.NewHTTPError(http.StatusBadRequest, "end time is not valid")
	}

	submissions, err := p.submissionService.FindAllBetween(c.Request().Context(), start, end)
	if err != nil {
		log.Printf("find all between: %v\n", err)
		return err
	}

	aboutMe, err := p.aboutService.Get(c.Request().Context(), id)
	if err != nil {
		log.Printf("get about me: %v\n", err)
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
		Websites:    fromWebsites(aboutMe.Websites),
		Education:   fromEducation(aboutMe.Education),
		WorkHistory: fromWorkHistory(aboutMe.WorkHistory),
		Submissions: fromSubmissions(submissions),
	})
}

// GetPublicInformation returns the detailed public information of a given user
// fetch public information using about service
func (p *ProfileHandler) GetPublicInformation() {
}

// GetSubmitHistory of the profile in given timespan
func (p *ProfileHandler) GetSubmitHistory() {

}
