// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this info except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mailgun

import (
	"io"
	t "time"

	"github.com/abcum/orbit"

	"gopkg.in/mailgun/mailgun-go.v1"

	"github.com/abcum/cirrius/cpm/time"
)

type Email struct {
	orb  *orbit.Orbit
	api  mailgun.Mailgun
	Info email
}

type email struct {
	Atch    map[string]io.ReadCloser `validate:"-"`
	Bcc     []string                 `validate:"dive,required,email"`
	Cc      []string                 `validate:"dive,required,email"`
	From    string                   `validate:"required"`
	Hdrs    map[string]string        `validate:"dive,required"`
	Html    string                   `validate:"-"`
	Reply   string                   `validate:"omitempty,email"`
	Subject string                   `validate:"required,max=250"`
	Tags    []string                 `validate:"dive,required"`
	Test    bool                     `validate:"-"`
	Text    string                   `validate:"required"`
	Time    t.Time                   `validate:"-"`
	To      []string                 `validate:"required,gt=0,dive,required,email"`
}

func NewEmail(orb *orbit.Orbit, api mailgun.Mailgun) *Email {
	return &Email{
		orb: orb,
		api: api,
	}
}

func (this *Email) Attachment(name string, r io.ReadCloser) *Email {
	if this.Info.Atch == nil {
		this.Info.Atch = make(map[string]io.ReadCloser)
	}
	this.Info.Atch[name] = r
	return this
}

func (this *Email) BCC(email string) *Email {
	this.Info.Bcc = append(this.Info.Bcc, email)
	return this
}

func (this *Email) CC(email string) *Email {
	this.Info.Cc = append(this.Info.Cc, email)
	return this
}

func (this *Email) From(email string) *Email {
	this.Info.From = email
	return this
}

func (this *Email) Header(key, val string) *Email {
	if this.Info.Hdrs == nil {
		this.Info.Hdrs = make(map[string]string)
	}
	this.Info.Hdrs[key] = val
	return this
}

func (this *Email) Html(html string) *Email {
	this.Info.Html = html
	return this
}

func (this *Email) ReplyTo(email string) *Email {
	this.Info.Reply = email
	return this
}

func (this *Email) Subject(subject string) *Email {
	this.Info.Subject = subject
	return this
}

func (this *Email) Tag(tag string) *Email {
	this.Info.Tags = append(this.Info.Tags, tag)
	return this
}

func (this *Email) Testing(enabled bool) *Email {
	this.Info.Test = enabled
	return this
}

func (this *Email) Text(text string) *Email {
	this.Info.Text = text
	return this
}

func (this *Email) Time(when *time.Time) *Email {
	this.Info.Time = when.Val
	return this
}

func (this *Email) To(email string) *Email {
	this.Info.To = append(this.Info.To, email)
	return this
}

func (this *Email) Send() (id string) {

	if err := checker.Struct(this); err != nil {
		this.orb.Quit(mesgError)
	}

	m := this.api.NewMessage(this.Info.From, this.Info.Subject, this.Info.Text)

	m.SetDeliveryTime(this.Info.Time)

	if this.Info.Test {
		m.EnableTestMode()
	}

	if this.Info.Html != "" {
		m.SetHtml(this.Info.Html)
	}

	if this.Info.Reply != "" {
		m.SetReplyTo(this.Info.Reply)
	}

	for _, tag := range this.Info.Tags {
		m.AddTag(tag)
	}

	for hdr, val := range this.Info.Hdrs {
		m.AddHeader(hdr, val)
	}

	for txt, att := range this.Info.Atch {
		m.AddReaderAttachment(txt, att)
	}

	for _, recipient := range this.Info.Cc {
		m.AddCC(recipient)
	}

	for _, recipient := range this.Info.Bcc {
		m.AddBCC(recipient)
	}

	for _, recipient := range this.Info.To {
		m.AddRecipientAndVariables(recipient, nil)
	}

	_, id, err := this.api.Send(m)
	if err != nil {
		this.orb.Quit(authError)
	}

	this.Info.To = []string{}

	return id

}
