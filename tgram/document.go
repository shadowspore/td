package tgram

import (
	"context"
	"io"

	"github.com/gotd/td/telegram/downloader"
	"github.com/gotd/td/tg"
)

type Document struct {
	doc       *tg.Document
	ttlPeriod int
	client    *Client
}

func (d *Document) ID() int64 { return d.doc.ID }

func (d *Document) Filename() (string, bool) {
	for _, attr := range d.doc.Attributes {
		if nattr, ok := attr.(*tg.DocumentAttributeFilename); ok {
			return nattr.FileName, true
		}
	}

	return "", false
}

func (d *Document) Size() int { return d.doc.Size }

func (d *Document) MimeType() string { return d.doc.MimeType }

func (d *Document) Download(ctx context.Context, w io.Writer) error {
	conn, err := d.client.telegram.DC(ctx, d.doc.DCID, 1)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	_, err = downloader.NewDownloader().
		Download(tg.NewClient(conn), d.asInputFileClass()).
		Stream(ctx, w)
	return err
}

func (d *Document) asInputFileClass() tg.InputFileLocationClass {
	return &tg.InputDocumentFileLocation{
		ID:            d.doc.ID,
		AccessHash:    d.doc.AccessHash,
		FileReference: d.doc.FileReference,
		ThumbSize:     "",
	}
}
