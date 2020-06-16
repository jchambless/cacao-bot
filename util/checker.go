package util

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
	"unicode/utf16"
	"unicode/utf8"

	"github.com/jchambless/cacao/framework"
)

const NUM_FIELDS int = 6
const DEFAULT_TIMEOUT int = 5

type OnlineStatus struct {
	Online         bool
	Version        string
	Motd           string
	CurrentPlayers string
	MaxPlayers     string
	Latency        time.Duration
}

// Based on code from https://github.com/ldilley/minestat
func MCServerChecker(ctx framework.Context) *OnlineStatus {
	log.Printf("MCServerChecker called for %s", ctx.Conf.ServerIP)

	status := OnlineStatus{}
	timeout := DEFAULT_TIMEOUT
	start_time := time.Now()

	// Todo: Change hardcoded port and move back to ctx.Conf
	conn, err := net.DialTimeout("tcp", ctx.Conf.ServerIP+":25565", time.Duration(timeout)*time.Second)
	status.Latency = time.Since(start_time)
	status.Latency = status.Latency.Round(time.Millisecond)
	if err != nil {
		status.Online = false
		return &status
	}

	_, err = conn.Write([]byte("\xFE\x01"))
	if err != nil {
		status.Online = false
		return &status
	}

	raw_data := make([]byte, 512)
	_, err = conn.Read(raw_data)
	if err != nil {
		status.Online = false
		return &status
	}
	conn.Close()

	if raw_data == nil || len(raw_data) == 0 {
		status.Online = false
		return &status
	}

	raw_string, err := decodeUtf16(raw_data)
	if err != nil {
		status.Online = false
		return &status
	}

	data := strings.Split(raw_string, "\x00")
	if data != nil && len(data) >= NUM_FIELDS {
		status.Online = true
		status.Version = data[2]
		status.Motd = data[3]
		status.CurrentPlayers = data[4]
		status.MaxPlayers = data[5]
	} else {
		status.Online = false
	}

	return &status
}

// https://gist.github.com/bradleypeabody/185b1d7ed6c0c2ab6cec
func decodeUtf16(b []byte) (string, error) {
	if len(b)%2 != 0 {
		return "", fmt.Errorf("Must have even length byte slice")
	}

	u16s := make([]uint16, 1)

	ret := &bytes.Buffer{}

	b8buf := make([]byte, 4)

	lb := len(b)
	for i := 0; i < lb; i += 2 {
		u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}

	return ret.String(), nil
}
