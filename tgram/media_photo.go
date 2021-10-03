package tgram

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/gotd/td/tg"
)

// Photo located in Telegram server.
type Photo struct {
	photo     *tg.Photo
	ttlPeriod int
	client    *Client
}

// ID of a photo.
func (p *Photo) ID() int64 { return p.photo.ID }

// Date of upload.
func (p *Photo) Date() int { return p.photo.Date }

// Sizes returns available photo sizes to download.
func (p *Photo) Sizes() []tg.PhotoSizeClass { return p.photo.Sizes }

// Filename of the photo.
func (p *Photo) Filename() string {
	const dateLayout = "2006-01-02_15-04-05"
	return fmt.Sprintf(
		"photo%d_%s.jpg", p.photo.ID,
		time.Unix(int64(p.photo.Date), 0).Format(dateLayout),
	)
}

func (p *Photo) Raw() *tg.Photo { return p.photo }

// Download downloads photo with specified size to writer.
func (p *Photo) Download(ctx context.Context, size tg.PhotoSizeClass, w io.Writer) error {
	return p.client.downloadMedia(ctx, p.photo.DCID, p.asInputFileLocation(size.GetType()), w)
}

func (p *Photo) asInputFileLocation(size string) tg.InputFileLocationClass {
	return &tg.InputPhotoFileLocation{
		ID:            p.photo.ID,
		AccessHash:    p.photo.AccessHash,
		FileReference: p.photo.FileReference,
		ThumbSize:     size,
	}
}
