package balancer

import (
	"fmt"
	"net/url"
	"strings"
	"sync/atomic"
)

type Balancer struct {
	cdnHost      string
	requestCount uint64
}

func NewBalancer(cdnHost string) *Balancer {
	return &Balancer{
		cdnHost:      cdnHost,
		requestCount: 0,
	}
}

func (b *Balancer) GetRedirectURL(videoURL string) (string, error) {
	parsedURL, err := url.Parse(videoURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	hostParts := strings.Split(parsedURL.Host, ".")
	if len(hostParts) < 2 {
		return "", fmt.Errorf("invalid origin server format")
	}

	serverName := hostParts[0]
	path := parsedURL.Path

	count := atomic.AddUint64(&b.requestCount, 1)

	if count%10 == 0 {
		return videoURL, nil
	}

	return fmt.Sprintf("http://%s/%s%s", b.cdnHost, serverName, path), nil
}
