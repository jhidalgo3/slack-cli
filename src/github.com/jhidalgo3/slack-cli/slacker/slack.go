package slacker

import (
	"fmt"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
)

// Slack is struct
type Slack struct {
	C *Config
}

// Message is post payload
type Message struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

var cli = &http.Client{Timeout: time.Duration(5) * time.Second}

// Post is post message to slack
func (s *Slack) Post(msg Message) (err []error) {
	attachment1 := Attachment{}
	attachment1.AddField(Field{Title: "Author", Value: "jhidalgo3"}).AddField(Field{Title: "Status", Value: "Completed"})
	payload := Payload{
		Text:        msg.Text,
		Username:    "Jenkins",
		Channel:     msg.Channel,
		IconEmoji:   ":jenkins:",
		Attachments: []Attachment{attachment1},
	}

	if msg.Channel == "#" {
		payload = Payload{
			Text:        msg.Text,
			Username:    "Jenkins",
			IconEmoji:   ":jenkins:",
			Attachments: []Attachment{attachment1},
		}
	}

	fmt.Println(payload.Channel)

	err = Send(s.C.URL, "", payload)
	if len(err) > 0 {
		return err
	}

	return
}

// ----

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Attachment struct {
	Fallback   *string   `json:"fallback"`
	Color      *string   `json:"color"`
	PreText    *string   `json:"pretext"`
	AuthorName *string   `json:"author_name"`
	AuthorLink *string   `json:"author_link"`
	AuthorIcon *string   `json:"author_icon"`
	Title      *string   `json:"title"`
	TitleLink  *string   `json:"title_link"`
	Text       *string   `json:"text"`
	ImageUrl   *string   `json:"image_url"`
	Fields     []*Field  `json:"fields"`
	Footer     *string   `json:"footer"`
	FooterIcon *string   `json:"footer_icon"`
	Timestamp  *int64    `json:"ts"`
	MarkdownIn *[]string `json:"mrkdwn_in"`
}

type Payload struct {
	Parse       string       `json:"parse,omitempty"`
	Username    string       `json:"username,omitempty"`
	IconUrl     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
	LinkNames   string       `json:"link_names,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links,omitempty"`
	UnfurlMedia bool         `json:"unfurl_media,omitempty"`
	Markdown    bool         `json:"mrkdwn,omitempty"`
}

func (attachment *Attachment) AddField(field Field) *Attachment {
	attachment.Fields = append(attachment.Fields, &field)
	return attachment
}

func redirectPolicyFunc(req gorequest.Request, via []gorequest.Request) error {
	return fmt.Errorf("Incorrect token (redirection)")
}

func Send(webhookUrl string, proxy string, payload Payload) []error {
	request := gorequest.New().Proxy(proxy)
	resp, _, err := request.
		Post(webhookUrl).
		RedirectPolicy(redirectPolicyFunc).
		Send(payload).
		End()

	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return []error{fmt.Errorf("Error sending msg. Status: %v", resp.Status)}
	}

	return nil
}
