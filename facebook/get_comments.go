package facebook

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/trongcuong1710/facebook_tools/config"
	"github.com/trongcuong1710/facebook_tools/constant"
)

const (
	path              = "/" + postIDPlaceHolder + "/comments?access_token="
	postIDPlaceHolder = "{post_id}"
)

// GetCommentsResponse represents response from facebook graph api post's comments
type GetCommentsResponse struct {
	Comments []*Comment `json:"data"`
}

// Comment represents facebook's comment object
type Comment struct {
	CreatedAt string `json:"created_time"`
	Message   string `json:"message"`
	ID        string `json:"id"`
}

// GetComments return comment for a specific post
func (c *clientImpl) GetComments(postID string) ([]*Comment, error) {
	url := c.buildGetCommentsURL(postID)

	logger := logrus.WithField("url", url)
	logger.Info("params")

	client := http.DefaultClient
	resp, err := client.Get(url)
	if err != nil {
		logger.Warn(constant.ErrorLogTag, err, "client.Get")
		return nil, err
	}

	var getCommentResp *GetCommentsResponse
	decoder := json.NewDecoder(resp.Body)

	if err = decoder.Decode(&getCommentResp); err != nil {
		logger.Warn(constant.ErrorLogTag, err, "decoder.Decode")
		return nil, err
	}

	return getCommentResp.Comments, nil
}

func (c *clientImpl) buildGetCommentsURL(postID string) string {
	accessToken := config.GetFBAppToken()
	pageID := config.GetPageID()
	url := c.buildRequestURL(strings.ReplaceAll(path, postIDPlaceHolder, pageID+"_"+postID))
	return url + accessToken + "&filter=stream&limit=1000"
}
