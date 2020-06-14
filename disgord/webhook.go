package disgord

import (
	"fmt"
	"net/http"
)

//WebHook structure to hold webhook meta data
type WebHook struct {
	ID    uint64 //the webhook id
	Token string //the webhook token
}

const userAgent string = "disgord/1.0"

const baseURL string = "https://discordapp.com/api/"

type urlBuilder = func(WebHook) string

type endpoint struct {
	buildURL urlBuilder
	method   string
}

var (
	executeWebhook endpoint = endpoint{
		buildURL: func(hook WebHook) string {
			return baseURL + fmt.Sprintf("webhooks/%d/%s", hook.ID, hook.Token)
		},
		method: http.MethodPost,
	}
)

//CreateRequest creates an http request object that will execute the specified webhook with the provided body as it's payload.
func CreateRequest(hook WebHook, body RequestVisitor) (*http.Request, error) {
	request, err := http.NewRequest(executeWebhook.method, executeWebhook.buildURL(hook), nil)
	if err != nil {
		return nil, err
	}
	err = body.accept(request)
	request.Header.Add("User-Agent", userAgent)
	return request, err
}

//Send sends the specified webhook with the provided body as it's payload.
func Send(hook WebHook, body RequestVisitor) (int, error) {
	request, err := CreateRequest(hook, body)
	if err != nil {
		return 0, err
	}
	httpClient := http.Client{}
	resp, err := httpClient.Do(request)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	return resp.StatusCode, err
}
