package entity

import (
	"errors"
	"log/slog"
	"regexp"
)

type zipcodeEntity struct {
	zipcode string
}

func (z *zipcodeEntity) Zipcode() string {
	return z.zipcode
}

func NewZipcode(zipcode string) (*zipcodeEntity, error) {

	var zc = &zipcodeEntity{
		zipcode: zipcode,
	}

	err := zc.IsValid()
	if err != nil {
		slog.Error("[invalid zipcode]", "error", err.Error())
		return nil, err
	}

	return &zipcodeEntity{
		zipcode: zc.zipcode,
	}, nil
}

func (z *zipcodeEntity) IsValid() error {

	var re = regexp.MustCompile(`^[0-9]{8}$`)

	if !re.MatchString(z.zipcode) {
		return errors.New("zip code must be 8 numeric digits")
	}
	return nil
}
