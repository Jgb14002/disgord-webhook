# Usage
## Send webhook directly
```go
package main

import (
	"time"

	"github.com/jgb14002/disgord-webhook/disgord"
)

func main() {
	hook := disgord.WebHook{
		ID:    12345,
		Token: "token",
	}

	jsonBody := disgord.WebHookJSONBody{
		Username:  "TestUser",
		AvatarURL: "https://st3.depositphotos.com/15648834/17930/v/600/depositphotos_179308454-stock-illustration-unknown-person-silhouette-glasses-profile.jpg",
		Embeds: []disgord.EmbedObject{
			{
				Title:       "Sample Webhook",
				Description: "A discord webhook sent with Go",
				Color:       (255 << 16) | (0 << 8) | 100,
				Timestamp:   time.Now().Format(time.RFC3339),
				Fields: []disgord.EmbedField{
					{
						Name:   "Inline Field A",
						Value:  "100",
						Inline: true,
					},
					{
						Name:   "Inline Field B",
						Value:  "Hello World",
						Inline: true,
					},
				},
				Author: disgord.EmbedAuthor{
					IconURL: "https://st3.depositphotos.com/15648834/17930/v/600/depositphotos_179308454-stock-illustration-unknown-person-silhouette-glasses-profile.jpg",
					Name:    "TestUser",
					URL:     "https://github.com/jgb14002",
				},
				Thumbnail: disgord.EmbedThumbnail{
					URL: "https://st3.depositphotos.com/15648834/17930/v/600/depositphotos_179308454-stock-illustration-unknown-person-silhouette-glasses-profile.jpg",
				},
				Footer: disgord.EmbedFooter{
					Text: "Footer",
				},
				Image: disgord.EmbedImage{
					URL: "https://st3.depositphotos.com/15648834/17930/v/600/depositphotos_179308454-stock-illustration-unknown-person-silhouette-glasses-profile.jpg",
				},
				URL: "https://github.com/jgb14002",
			},
		},
	}

    //Creates and sends an http.Request
	disgord.Send(hook, jsonBody)
}
```
## Send webhook request manually
```go
package main

import (
	"time"

	"github.com/LoudPacks/disgord-webhook/disgord"
)

func main() {
	hook := disgord.WebHook{
		ID:    12345,
		Token: "token",
	}

	jsonBody := disgord.WebHookJSONBody{
		Username:  "LoudPacks",
		AvatarURL: "https://avatars3.githubusercontent.com/u/32648600?s=460&u=c13e23efd4bca1aaca605b0874174fcaf580db9d&v=4",
		Embeds: []disgord.EmbedObject{
			{
				Title:       "Sample Webhook",
				Description: "A discord webhook sent with Go",
				Color:       (255 << 16) | (0 << 8) | 100,
				Timestamp:   time.Now().Format(time.RFC3339)
			}
		},
	}

    //Create an http.Request with the message paylaod
    httpRequest := disgord.CreateRequest(hook, jsonBody)
    ...
}
```
