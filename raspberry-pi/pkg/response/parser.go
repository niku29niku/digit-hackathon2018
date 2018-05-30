package response

import (
	"fmt"
	"strings"
)

// Parser parse response string
type Parser interface {
	ParseStatus(resp string) (Status, error)
}

type parserImpl struct{}

func (p *parserImpl) ParseStatus(resp string) (st Status, er error) {
	r := strings.ToLower(strings.Trim(resp, "\r\n"))
	if r == "ok" {
		st = Ok
	} else if r == "ng" {
		st = Ng
	} else {
		er = fmt.Errorf("response is invalid %s", resp)
	}
	return
}

// NewParser create new parser instance
func NewParser() Parser {
	return &parserImpl{}
}
