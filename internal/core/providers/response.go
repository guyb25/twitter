package providers

type Response struct {
	UserResponse    *userResponse
	GeneralResponse *generalResponse
	TweetResponse   *tweetResponse
}

func NewResponse() *Response {
	return &Response{
		UserResponse:    &userResponse{},
		GeneralResponse: &generalResponse{},
		TweetResponse:   &tweetResponse{},
	}
}
