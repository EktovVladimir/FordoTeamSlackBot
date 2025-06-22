package user_finder

import (
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/services/user_finder/mocks"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	rmocks "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v72/github"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	errTest = errors.New("test error")
)

func TestFindUserByEmail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ts := initTestStruct(ctrl)

	ts.ghClient.EXPECT().
		Users(gomock.Any(), "test@test.com in:email", gomock.Any()).
		Return(&github.UsersSearchResult{
			Total: github.Ptr(1),
			Users: []*github.User{
				{
					Login: github.Ptr("testGh"),
					Email: github.Ptr("test@test.com"),
				},
			},
		}, nil, nil)

	ts.slClient.EXPECT().
		GetUserByEmailContext(gomock.Any(), "test@test.com").
		Return(&slack.User{
			ID:   "U00A0A0A00A",
			Name: "testSl",
		}, nil)

	actual, err := ts.service.FindUserByEmail(t.Context(), "test@test.com")

	assert.NoError(t, err)
	require.NotNil(t, actual)

	assert.Equal(t, &models.User{
		Email:       "test@test.com",
		SlackId:     "U00A0A0A00A",
		SlackName:   "testSl",
		GithubLogin: "testGh",
	}, actual)
}

func TestFindUserByEmail_SLackNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ts := initTestStruct(ctrl)

	ts.slClient.EXPECT().
		GetUserByEmailContext(gomock.Any(), "test@test.com").
		Return(nil, errTest)
	ts.ghClient.EXPECT().
		Users(gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	_, err := ts.service.FindUserByEmail(t.Context(), "test@test.com")

	assert.Error(t, err)
}

func TestFindUserByEmail_GithubNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ts := initTestStruct(ctrl)

	ts.slClient.EXPECT().
		GetUserByEmailContext(gomock.Any(), "test@test.com").
		Return(&slack.User{
			ID:   "U00A0A0A00A",
			Name: "testSl",
		}, nil)
	ts.ghClient.EXPECT().
		Users(gomock.Any(), "test@test.com in:email", gomock.Any()).
		Return(nil, nil, errTest)

	actual, err := ts.service.FindUserByEmail(t.Context(), "test@test.com")

	assert.Error(t, err)
	require.NotNil(t, actual)
	assert.Equal(t, "U00A0A0A00A", actual.SlackId)
	assert.Equal(t, "testSl", actual.SlackName)
}

func TestFindUserByEmailWithDb_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ts := initTestStruct(ctrl)

	ts.repo.EXPECT().
		GetByEmail(gomock.Any(), "test@test.com").
		Return(&db.User{
			Id:         1,
			Email:      "test@test.com",
			GithubName: "testGh",
			SlackName:  "testSl",
			SlackId:    "U00A0A0A00A",
		}, nil)

	ts.slClient.EXPECT().
		GetUserByEmailContext(gomock.Any(), gomock.Any()).
		Times(0)

	ts.ghClient.EXPECT().
		Users(gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	actual, err := ts.service.FindUserByEmailWithDb(t.Context(), "test@test.com")

	assert.NoError(t, err)
	require.NotNil(t, actual)
	assert.Equal(t, &models.User{
		Id:          1,
		Email:       "test@test.com",
		SlackId:     "U00A0A0A00A",
		SlackName:   "testSl",
		GithubLogin: "testGh",
	}, actual)
}

func TestSaveUser_Update_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ts := initTestStruct(ctrl)

	ts.repo.EXPECT().
		GetByEmail(gomock.Any(), "test@test.com").
		Return(&db.User{
			Id:    1,
			Email: "test@test.com",
		}, nil)

	ts.repo.EXPECT().
		GetById(gomock.Any(), gomock.Any()).
		Times(0)

	ts.repo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Times(0)

	ts.repo.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil)

	err := ts.service.SaveUser(t.Context(), &models.User{
		Email:       "test@test.com",
		SlackId:     "U00A0A0A00A",
		SlackName:   "testSl",
		GithubLogin: "testGh",
	})

	assert.NoError(t, err)
}

func TestSaveUser_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ts := initTestStruct(ctrl)

	ts.repo.EXPECT().
		GetByEmail(gomock.Any(), "test@test.com").
		Return(nil, repository.ErrorNotFound)

	ts.repo.EXPECT().
		GetById(gomock.Any(), gomock.Any()).
		Times(0)

	ts.repo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	ts.repo.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Times(0)

	err := ts.service.SaveUser(t.Context(), &models.User{
		Email:       "test@test.com",
		SlackId:     "U00A0A0A00A",
		SlackName:   "testSl",
		GithubLogin: "testGh",
	})

	assert.NoError(t, err)
}

type testStruct struct {
	slClient *mocks.MockslackClient
	ghClient *mocks.MockgithubSearchService
	repo     *rmocks.MockUserRepository
	service  *UserFinder
}

func initTestStruct(ctrl *gomock.Controller) *testStruct {
	slClient := mocks.NewMockslackClient(ctrl)
	ghClient := mocks.NewMockgithubSearchService(ctrl)
	repo := rmocks.NewMockUserRepository(ctrl)

	return &testStruct{
		slClient: slClient,
		ghClient: ghClient,
		repo:     repo,
		service:  New(repo, slClient, ghClient),
	}
}
