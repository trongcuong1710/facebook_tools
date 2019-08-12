package facebook

const (
	statsTag                  = "facebook.client"
	fbGraphAPIURL             = "https://graph.facebook.com/v4.0"
	accessTokenQueryStringKey = "access_token"
)

// DefaultClient represents facebook default client
var DefaultClient Client = &clientImpl{}

// Client represents facebook's client
type Client interface {
	GetComments(postID string) ([]*Comment, error)
}

type clientImpl struct{}

func (c *clientImpl) buildRequestURL(path string) string {
	return fbGraphAPIURL + path
}
