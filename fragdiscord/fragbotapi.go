package main

type CreateBotResponse struct {
	MsAuthInfo *AuthUserData `json:"msAuthInfo,omitempty"`
}

type ErrorResponse struct {
	Err string `json:"error,omitempty"`
}

type AuthUserData struct {
	UserCode        string `json:"userCode"`
	VerificationUrl string `json:"VerificationUrl"`
	Email           string `json:"email"`
	Password        string `json:"password"`
}

type PostBotRequest struct {
	Stage    int    `json:"stage" form:"stage"`
	Email    string `json:"email,omitempty" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
	UserCode string `json:"userCode,omitempty" form:"userCode"`
}

type DeleteBotRequest struct {
	Delete bool `json:"delete" form:"delete"`
}

func CreateBot2(botId string, request PostBotRequest) *ErrorResponse {
	errRes := ErrorResponse{}
	_, err := ReqClient.R().
		SetHeader("access-token", AccessToken).
		SetBodyJsonMarshal(request).
		SetError(&errRes).
		Post(BackendUrl + "/bots/" + botId)
	if err != nil {
		return &ErrorResponse{"Request timed out!"}
	}
	if errRes.Err != "" {
		return &errRes
	}
	return nil
}

func CreateBot(botId string, request PostBotRequest) (*CreateBotResponse, *ErrorResponse) {
	result := CreateBotResponse{}
	errRes := ErrorResponse{}
	_, err := ReqClient.R().
		SetHeader("access-token", AccessToken).
		SetBodyJsonMarshal(request).
		SetResult(&result).
		SetError(&errRes).
		Post(BackendUrl + "/bots/" + botId)
	if err != nil {
		return nil, &ErrorResponse{"Backend Offline"}
	}
	if errRes.Err != "" {
		return nil, &errRes
	}
	return &result, nil
}

func StartBot(botId string) *ErrorResponse {
	errRes := ErrorResponse{}
	_, err := ReqClient.R().
		SetHeader("access-token", AccessToken).
		SetError(&errRes).
		Put(BackendUrl + "/bots/" + botId)
	if err != nil {
		return &ErrorResponse{"Request timed out!"}
	}
	if errRes.Err != "" {
		return &errRes
	}
	return nil
}

func StopBot(botId string) *ErrorResponse {
	errRes := ErrorResponse{}
	_, err := ReqClient.R().
		SetHeader("access-token", AccessToken).
		SetError(&errRes).
		Delete(BackendUrl + "/bots/" + botId)
	if err != nil {
		return &ErrorResponse{"Request timed out!"}
	}
	if errRes.Err != "" {
		return &errRes
	}
	return nil
}

func DeleteBot(botId string) *ErrorResponse {
	errRes := ErrorResponse{}
	_, err := ReqClient.R().
		SetHeader("access-token", AccessToken).
		SetBody(DeleteBotRequest{Delete: true}).
		SetError(&errRes).
		Delete(BackendUrl + "/bots/" + botId)
	if err != nil {
		return &ErrorResponse{"Request timed out!"}
	}
	if errRes.Err != "" {
		return &errRes
	}
	return nil
}
