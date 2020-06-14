# Disgord WebHook

A small library written in Go to post messages to Discord channels via web hooks.
![Image](https://i.imgur.com/NS7PJXA.png)

# Usage
## Send webhook directly
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
					IconURL: "https://avatars3.githubusercontent.com/u/32648600?s=460&u=c13e23efd4bca1aaca605b0874174fcaf580db9d&v=4",
					Name:    "LoudPacks",
					URL:     "https://github.com/LoudPacks",
				},
				Thumbnail: disgord.EmbedThumbnail{
					URL: "https://avatars3.githubusercontent.com/u/32648600?s=460&u=c13e23efd4bca1aaca605b0874174fcaf580db9d&v=4",
				},
				Footer: disgord.EmbedFooter{
					Text: "Footer",
				},
				Image: disgord.EmbedImage{
					URL: "https://avatars3.githubusercontent.com/u/32648600?s=460&u=c13e23efd4bca1aaca605b0874174fcaf580db9d&v=4",
				},
				URL: "https://github.com/LoudPacks",
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
		},
	}

    //Create an http.Request with the message paylaod
    httpRequest := disgord.CreateRequest(hook, jsonBody)
    ...
}
```