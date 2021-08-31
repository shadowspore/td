package tgram

import (
	"context"
	"io"

	"github.com/gotd/td/telegram/downloader"
	"github.com/gotd/td/tg"
)

type Photo struct {
	photo     *tg.Photo
	ttlPeriod int
	client    *Client
}

func (p *Photo) ID() int64 { return p.photo.ID }

func (p *Photo) Sizes() []tg.PhotoSizeClass { return p.photo.Sizes }

func (p *Photo) Download(ctx context.Context, size tg.PhotoSizeClass, w io.Writer) error {
	conn, err := p.client.telegram.DC(ctx, p.photo.DCID, 1)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	_, err = downloader.NewDownloader().
		Download(
			tg.NewClient(conn),
			p.asInputFileLocation(size.GetType()),
		).Stream(ctx, w)
	return err
}

func (p *Photo) asInputFileLocation(size string) tg.InputFileLocationClass {
	return &tg.InputPhotoFileLocation{
		ID:            p.photo.ID,
		AccessHash:    p.photo.AccessHash,
		FileReference: p.photo.FileReference,
		ThumbSize:     size,
	}
}
