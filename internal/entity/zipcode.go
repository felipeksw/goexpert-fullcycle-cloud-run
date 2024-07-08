package entity

import (
	"errors"
	"regexp"
)

type Zipcode struct {
	Zipcode string
}

func (z *Zipcode) IsValid() error {

	var re = regexp.MustCompile(`^[0-9]{8}$`)

	if !re.MatchString(z.Zipcode) {
		return errors.New("zip Code must be 8 numeric digits")
	}
	return nil
}

func NewZipcode(zipcode string) (*Zipcode, error) {

	var zc = &Zipcode{
		Zipcode: zipcode,
	}

	err := zc.IsValid()
	if err != nil {
		return nil, err
	}

	return zc, nil
}
