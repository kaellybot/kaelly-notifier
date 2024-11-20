package images

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetImageFromURL(ctx context.Context, rawURL string,
) (*bytes.Buffer, error) {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return nil, err
	}

	req, errReq := http.NewRequestWithContext(ctx, http.MethodGet, parsedURL.String(), nil)
	if errReq != nil {
		return nil, errReq
	}

	client := &http.Client{}
	resp, errDo := client.Do(req)
	if errDo != nil {
		return nil, errDo
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch image: status code %d", resp.StatusCode)
	}

	buffer := new(bytes.Buffer)
	_, errCopy := io.Copy(buffer, resp.Body)
	if errCopy != nil {
		return nil, errCopy
	}

	return buffer, nil
}
