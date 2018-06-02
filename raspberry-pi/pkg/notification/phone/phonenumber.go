package phone

import "github.com/nyaruka/phonenumbers"

// NumberParser is interface to operato phonenumber
type NumberParser interface {
	ParseToE164(origin string) (string, error)
}

// NewParser creates new parser instance
func NewParser() NumberParser {
	return &libNumberParser{}
}

type libNumberParser struct{}

func (l *libNumberParser) ParseToE164(origin string) (string, error) {
	number, err := phonenumbers.Parse(origin, "JP")
	if err != nil {
		return "", err
	}
	return phonenumbers.Format(number, phonenumbers.E164), nil
}
