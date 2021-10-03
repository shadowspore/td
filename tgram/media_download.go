package tgram

import (
	"context"
	"io"

	"github.com/gotd/td/telegram/downloader"
	"github.com/gotd/td/tg"
)

func (c *Client) downloadMedia(
	ctx context.Context,
	dcID int,
	location tg.InputFileLocationClass,
	w io.Writer,
) (err error) {
	conn, err := c.telegram.ConnectDC(ctx, dcID)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	_, err = downloader.NewDownloader().
		DownloadDirect(tg.NewClient(conn), location).
		Stream(ctx, w)
	return err
}

func (c *Client) DownloadMedia(
	ctx context.Context,
	dcID int,
	location tg.InputFileLocationClass,
	w io.Writer,
) error {
	return c.downloadMedia(ctx, dcID, location, w)
}
