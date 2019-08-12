package facebook

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
)

type GetCommentsSuite struct {
	suite.Suite

	postID string
	client *clientImpl
}

func TestGetCommentsSuite(t *testing.T) {
	suite.Run(t, new(GetCommentsSuite))
}

func (s *GetCommentsSuite) SetupSuite() {
	s.postID = "abc"
	s.client = &clientImpl{}
}

func (s *GetCommentsSuite) TestGetComments_Success() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		s.client.buildGetCommentsURL(s.postID),
		httpmock.NewStringResponder(200, `{"data": [{"created_time": "2019-08-05T02:08:13+0000","message":"17, Huỳnh Tuấn Kiệt, Orangy Rita","id": "2391535717726871_2391551047725338"}]}`),
	)

	comments, err := s.client.GetComments(s.postID)

	s.NoError(err)
	s.NotEmpty(comments)
	s.Equal("17, Huỳnh Tuấn Kiệt, Orangy Rita", comments[0].Message)
	s.Equal("2391535717726871_2391551047725338", comments[0].ID)
	s.Equal("2019-08-05T02:08:13+0000", comments[0].CreatedAt)
}

func (s *GetCommentsSuite) TestGetComments_Failed() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		s.client.buildGetCommentsURL(s.postID),
		httpmock.NewStringResponder(200, `{"data": [{"created_time": "2019-08-05T02:08:13+0000","message":"17, Huỳnh Tuấn Kiệt, Orangy Rita","id": "2391535717726871_2391551047725338"}]}`),
	)

	comments, err := s.client.GetComments(s.postID)

	s.NoError(err)
	s.NotEmpty(comments)
	s.Equal("17, Huỳnh Tuấn Kiệt, Orangy Rita", comments[0].Message)
	s.Equal("2391535717726871_2391551047725338", comments[0].ID)
	s.Equal("2019-08-05T02:08:13+0000", comments[0].CreatedAt)
}
