package disgord

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	//MentionTypeRoles MentionType for @roles.
	MentionTypeRoles string = "roles"
	//MentionTypeUsers MentionType for @users.
	MentionTypeUsers string = "users"
	//MentionTypeEveryone MentionType for @everyone.
	MentionTypeEveryone string = "everyone"
)

//EmbedObject holds rich data attributes that can be embded into web hook messages.
type EmbedObject struct {
	Title       string         `json:"title,omitempty"`       //title of embed
	EmbedType   string         `json:"type,omitempty"`        //type of embed (always "rich" for webhook embeds)
	Description string         `json:"description,omitempty"` //description of embed
	URL         string         `json:"url,omitempty"`         //url of embed
	Timestamp   string         `json:"timestamp,omitempty"`   //timestamp of embed content (ISO8601 format)
	Color       int32          `json:"color,omitempty"`       //color code of the embed
	Footer      EmbedFooter    `json:"footer,omitempty"`      //foot information
	Image       EmbedImage     `json:"image,omitempty"`       //image information
	Thumbnail   EmbedThumbnail `json:"thumbnail,omitempty"`   //thumbnail information
	Video       EmbedVideo     `json:"video,omitempty"`       //video information
	Provider    EmbedProvider  `json:"provider,omitempty"`    //provider information
	Author      EmbedAuthor    `json:"author,omitempty"`      //author information
	Fields      []EmbedField   `json:"fields,omitempty"`      //fields information
}

//EmbedField field attribute for EmbedObjects
type EmbedField struct {
	Name   string `json:"name,omitempty"`   //name of the field
	Value  string `json:"value,omitempty"`  //value of the field
	Inline bool   `json:"inline,omitempty"` //whether or not this field should display inline
}

//EmbedFooter footer attribute for EmbedObjects
type EmbedFooter struct {
	Text         string `json:"text,omitempty"`           //footer text
	IconURL      string `json:"icon_url,omitempty"`       //url of footer icon (http(s))
	ProxyIconURL string `json:"proxy_icon_url,omitempty"` //a proxied url of footer icon
}

//EmbedImage image attribute for EmbedObjects
type EmbedImage struct {
	URL      string `json:"url,omitempty"`       //source url of image (http(s))
	ProxyURL string `json:"proxy_url,omitempty"` //a proxied url of the image
	Height   int32  `json:"height,omitempty"`    //height of image
	Width    int32  `json:"width,omitempty"`     //width of image
}

//EmbedThumbnail thumbnail attribute for EmbedObjects
type EmbedThumbnail struct {
	URL      string `json:"url,omitempty"`       //source of thumbnail (http(s))
	ProxyURL string `json:"proxy_url,omitempty"` //a proxied url of the thumbnail
	Height   int32  `json:"height,omitempty"`    //height of thumbnail
	Width    int32  `json:"width,omitempty"`     //width of thumbnail
}

//EmbedVideo video attribute for EmbedObjects
type EmbedVideo struct {
	URL    string `json:"url,omitempty"`    //srouce url of video
	Height int32  `json:"height,omitempty"` //height of video
	Width  int32  `json:"width,omitempty"`  //width of video
}

//EmbedProvider provider attribute for EmbedObjects
type EmbedProvider struct {
	Name string `json:"name,omitempty"` //name of provider
	URL  string `json:"url,omitempty"`  //url of provider
}

//EmbedAuthor author attribute for EmbedObjects
type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`           //name of author
	URL          string `json:"url,omitempty"`            //url of author
	IconURL      string `json:"icon_url,omitempty"`       //url of author icon (http(s))
	ProxyIconURL string `json:"proxy_icon_url,omitempty"` //a proxied url of author icon
}

//AllowedMention allows for more granular control over mentions without having to alter message content directly
type AllowedMention struct {
	Parse []string `json:"parse"` //An array of AllowedMentionTypes
	Roles []uint64 `json:"roles"` //An array of role IDs to mention (max len 100)
	Users []uint64 `json:"users"` //An array of user IDs to mention (max len 100)
}

//WebHookJSONBody contains the data for a webhook json only request body
type WebHookJSONBody struct {
	Content         string         `json:"content"`          //the message contents (max len 2000 chars)
	Username        string         `json:"username"`         //override the default username of the webhook
	AvatarURL       string         `json:"avatar_url"`       //override the default avatar of the webhook
	Tts             bool           `json:"tts"`              //true if this is a TTS message
	Embeds          []EmbedObject  `json:"embeds"`           //up to 10 embed objects
	AllowedMentions AllowedMention `json:"allowed_mentions"` //allowed mentions for the message
}

//RequestVisitor an interface for modifying http.Request objects
type RequestVisitor interface {
	accept(request *http.Request) error
}

func (hook WebHookJSONBody) accept(request *http.Request) error {
	request.Header.Add("Content-Type", "application/json")
	json, err := json.Marshal(hook)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(json)
	request.Body = ioutil.NopCloser(buffer)
	request.ContentLength = int64(buffer.Len())
	return nil
}
