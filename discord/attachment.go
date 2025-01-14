package discord

import (
	"time"

	"github.com/snekROmonoro/snowflake"
)

// Attachment is used for files sent in a Message
type Attachment struct {
	ID           snowflake.ID    `json:"id,omitempty"`
	Filename     string          `json:"filename,omitempty"`
	Description  *string         `json:"description,omitempty"`
	ContentType  *string         `json:"content_type,omitempty"`
	Size         int             `json:"size,omitempty"`
	URL          string          `json:"url,omitempty"`
	ProxyURL     string          `json:"proxy_url,omitempty"`
	Height       *int            `json:"height,omitempty"`
	Width        *int            `json:"width,omitempty"`
	Ephemeral    bool            `json:"ephemeral,omitempty"`
	DurationSecs *float64        `json:"duration_secs,omitempty"`
	Waveform     *string         `json:"waveform,omitempty"`
	Flags        AttachmentFlags `json:"flags"`
}

func (a Attachment) CreatedAt() time.Time {
	return a.ID.Time()
}

type AttachmentFlags int

const (
	AttachmentFlagIsClip AttachmentFlags = 1 << iota
	AttachmentFlagIsThumbnail
	AttachmentFlagIsRemix
	AttachmentFlagsNone AttachmentFlags = 0
)

type AttachmentUpdate interface {
	attachmentUpdate()
}

type AttachmentKeep struct {
	ID snowflake.ID `json:"id,omitempty"`
}

func (AttachmentKeep) attachmentUpdate() {}

type AttachmentCreate struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func (AttachmentCreate) attachmentUpdate() {}
