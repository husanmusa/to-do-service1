package postgres

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/husanmusa/to-do-service/config"
	pb "github.com/husanmusa/to-do-service/genproto"
	"github.com/husanmusa/to-do-service/pkg/db"
	"github.com/husanmusa/to-do-service/storage/repo"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	CleanupFunc func()
	Repository  repo.TaskStorageI
}

func (suite *TaskRepositoryTestSuite) SetupSuite() {
	pgPool, cleanup := db.ConnectToDBForSuite(config.Load())

	suite.Repository = NewTaskRepo(pgPool)
	suite.CleanupFunc = cleanup
}

func (suite *TaskRepositoryTestSuite) TestTaskCRUD() {
	id := "0d512776-60ed-4980-b8a3-6904a2234fd4"
	Assignee := "Assignee"
	Title := "Title"
	Summary := "Summary"
	Deadline := "2021-12-20T00:00:00Z"
	Status := "active"

	task := pb.Task{
		Id:       id,
		Assignee: "Assignee",
		Title:    "Title",
		Summary:  "Summary",
		Deadline: "2021-12-20",
		Status:   "active",
	}

	_ = suite.Repository.Delete(id)

	task, err := suite.Repository.Create(task)
	suite.Nil(err)

	getTask, err := suite.Repository.Get(task.Id)
	suite.Nil(err)
	suite.NotNil(getTask, "Task must not be nil")
	suite.Equal(Assignee, getTask.Assignee, "Assignee must match")
	suite.Equal(Title, getTask.Title, "Title must match")
	suite.Equal(Summary, getTask.Summary, "Summary must match")
	suite.Equal(Deadline, getTask.Deadline, "Deadline must match")
	suite.Equal(Status, getTask.Status, "Status must match")

	listTasks, _, err := suite.Repository.List(1, 2)
	suite.Nil(err)
	suite.NotEmpty(listTasks)
	suite.Equal(Assignee, listTasks[0].Assignee, "Assignee must match")
	suite.Equal(Title, listTasks[0].Title, "Title must match")
	suite.Equal(Summary, listTasks[0].Summary, "Summary must match")
	suite.Equal(Deadline, listTasks[0].Deadline, "Deadline must match")
	suite.Equal(Status, listTasks[0].Status, "Status must match")

	err = suite.Repository.Delete(id)
	suite.Nil(err)
}

func (suite *TaskRepositoryTestSuite) TearDownSuite() {
	suite.CleanupFunc()
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}
