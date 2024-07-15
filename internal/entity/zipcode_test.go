package entity_test

import (
	"testing"

	"github.com/felipeksw/goexpert-fullcycle-cloud-run/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewZipcode(t *testing.T) {

	type zipcodeLote struct {
		zipcode string
		err     string
		status  bool
	}

	table := []zipcodeLote{
		{"", "zip code must be 8 numeric digits", false},
		{"1300000z", "zip code must be 8 numeric digits", false},
		{"130000010", "zip code must be 8 numeric digits", false},
		{"13000001012345678", "zip code must be 8 numeric digits", false},
		{"#30000010", "zip code must be 8 numeric digits", false},
		{"13000001$", "zip code must be 8 numeric digits", false},
		{"^#3000001", "zip code must be 8 numeric digits", false},
		{"'3000001", "zip code must be 8 numeric digits", false},
		{"\"3000001", "zip code must be 8 numeric digits", false},
		{"130000-010", "zip code must be 8 numeric digits", false},
		{"ABCDEFGH", "zip code must be 8 numeric digits", false},
		{"13000001", "", true},
		{"00000000", "", true},
		{"99999999", "", true},
	}
	for _, item := range table {
		zpEntity, err := entity.NewZipcode(item.zipcode)
		if item.status {
			assert.Nil(t, err)
			assert.Equal(t, item.zipcode, zpEntity.Zipcode())
		} else {
			assert.Error(t, err, item.err)
		}
	}
}
