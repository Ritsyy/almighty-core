package models_test

import (
	"testing"

	"github.com/almighty/almighty-core/application"
	"github.com/almighty/almighty-core/errors"
	"github.com/almighty/almighty-core/gormsupport"
	"github.com/almighty/almighty-core/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
)

type workItemRepoBlackBoxTest struct {
	gormsupport.DBTestSuite
	repo application.WorkItemRepository
}

func TestRunWorkTypeRepoBlackBoxTest(t *testing.T) {
	suite.Run(t, &workItemRepoBlackBoxTest{DBTestSuite: gormsupport.NewDBTestSuite("../config.yaml")})
}

func (s *workItemRepoBlackBoxTest) SetupTest() {
	s.repo = models.NewWorkItemRepository(s.DB)
}

func (s *workItemRepoBlackBoxTest) TestFailDeleteZeroID() {
	defer gormsupport.DeleteCreatedEntities(s.DB)()

	// Create at least 1 item to avoid RowsEffectedCheck
	_, err := s.repo.Create(
		context.Background(), "system.bug",
		map[string]interface{}{
			models.SystemTitle: "Title",
			models.SystemState: models.SystemStateNew,
		}, "xx")

	if err != nil {
		s.T().Error("Could not create workitem", err)
	}

	err = s.repo.Delete(context.Background(), "0")
	require.IsType(s.T(), errors.NotFoundError{}, err)
}

func (s *workItemRepoBlackBoxTest) TestFailSaveZeroID() {
	defer gormsupport.DeleteCreatedEntities(s.DB)()

	// Create at least 1 item to avoid RowsEffectedCheck
	wi, err := s.repo.Create(
		context.Background(), "system.bug",
		map[string]interface{}{
			models.SystemTitle: "Title",
			models.SystemState: models.SystemStateNew,
		}, "xx")

	if err != nil {
		s.T().Error("Could not create workitem", err)
	}
	wi.ID = "0"

	_, err = s.repo.Save(context.Background(), *wi)
	require.IsType(s.T(), errors.NotFoundError{}, err)
}

func (s *workItemRepoBlackBoxTest) TestFaiLoadZeroID() {
	defer gormsupport.DeleteCreatedEntities(s.DB)()

	// Create at least 1 item to avoid RowsEffectedCheck
	_, err := s.repo.Create(
		context.Background(), "system.bug",
		map[string]interface{}{
			models.SystemTitle: "Title",
			models.SystemState: models.SystemStateNew,
		}, "xx")

	if err != nil {
		s.T().Error("Could not create workitem", err)
	}

	_, err = s.repo.Load(context.Background(), "0")
	require.IsType(s.T(), errors.NotFoundError{}, err)
}
