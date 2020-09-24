package utils

import (
	"fmt"
	"time"
)

func TimeFormat(s string) (string, error) {
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		fmt.Printf("TimeFormat time parse err: %v", err)
		return "", err
	}
	ret := t.Format("2006-01-02T15:04:05")
	return ret, nil
}
