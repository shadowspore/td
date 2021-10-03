package tgram

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/gotd/td/tg"
)

// Document located in Telegram server.
type Document struct {
	doc       *tg.Document
	ttlPeriod int
	client    *Client
}

// ID returns document ID.
func (d *Document) ID() int64 { return d.doc.ID }

// Date of upload.
func (d *Document) Date() int { return d.doc.Date }

// Filename of the document.
func (d *Document) Filename() string {
	const dateLayout = "2006-01-02_15-04-05"

	var filename, ext string
	for _, attr := range d.doc.Attributes {
		switch v := attr.(type) {
		case *tg.DocumentAttributeImageSize:
			switch d.doc.MimeType {
			case "image/png":
				ext = ".png"
			case "image/webp":
				ext = ".webp"
			case "image/tiff":
				ext = ".tif"
			default:
				ext = ".jpg"
			}
		case *tg.DocumentAttributeAnimated:
			ext = ".gif"
		case *tg.DocumentAttributeSticker:
			ext = ".webp"
		case *tg.DocumentAttributeVideo:
			switch d.doc.MimeType {
			case "video/mpeg":
				ext = ".mpeg"
			case "video/webm":
				ext = ".webm"
			case "video/ogg":
				ext = ".ogg"
			default:
				ext = ".mp4"
			}
		case *tg.DocumentAttributeAudio:
			switch d.doc.MimeType {
			case "audio/webm":
				ext = ".webm"
			case "audio/aac":
				ext = ".aac"
			case "audio/ogg":
				ext = ".ogg"
			default:
				ext = ".mp3"
			}
		case *tg.DocumentAttributeFilename:
			filename = v.FileName
		}
	}

	if filename == "" {
		filename = fmt.Sprintf(
			"doc%d_%s%s", d.doc.ID,
			time.Unix(int64(d.doc.Date), 0).Format(dateLayout),
			ext,
		)
	}

	return filename
}

func (d *Document) Raw() *tg.Document { return d.doc }

// Size of the document (in bytes).
func (d *Document) Size() int { return d.doc.Size }

// MimeType returns the document mime type.
func (d *Document) MimeType() string { return d.doc.MimeType }

// Download downloads document to the provided writer.
func (d *Document) Download(ctx context.Context, w io.Writer) error {
	return d.client.downloadMedia(ctx, d.doc.DCID, d.asInputFileClass(), w)
}

func (d *Document) asInputFileClass() tg.InputFileLocationClass {
	return &tg.InputDocumentFileLocation{
		ID:            d.doc.ID,
		AccessHash:    d.doc.AccessHash,
		FileReference: d.doc.FileReference,
		ThumbSize:     "",
	}
}
