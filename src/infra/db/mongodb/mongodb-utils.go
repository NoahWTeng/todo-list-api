package mongodb

import (
	"bytes"
	"fmt"
	"strings"
)

func formatUrl(config *Config) string {
	if config.URI != "" {
		return config.URI
	}

	var b bytes.Buffer
	b.WriteString("mongodb://")
	if config.Username != "" {
		b.WriteString(config.Username)
		b.WriteString(":")
	}
	if config.Password != "" {
		b.WriteString(config.Password)
		b.WriteString("@")
	}
	b.WriteString(config.DatabaseHost)
	b.WriteString("/")
	b.WriteString(config.DatabaseName)

	var urlQueryString []string

	if config.PoolSize != 0 {
		urlQueryString = append(urlQueryString, fmt.Sprintf("maxPoolSize=%d", config.PoolSize))
	}

	if config.ReplicaSet != "" {
		urlQueryString = append(urlQueryString, fmt.Sprintf("replicaSet=%s", config.ReplicaSet))
	}

	if config.AuthSource != "" {
		urlQueryString = append(urlQueryString, fmt.Sprintf("authSource=%s", config.AuthSource))
	}

	if len(urlQueryString) > 0 {
		s := strings.Join(urlQueryString, "&")
		s = "?" + s
		b.WriteString(s)
	}

	return b.String()
}
