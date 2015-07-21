package strfmt

import (
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var _ sql.Scanner = &Date{}
var _ driver.Valuer = Date{}

func TestDate(t *testing.T) {
	pp := Date{}
	err := pp.UnmarshalText([]byte{})
	assert.NoError(t, err)
	err = pp.UnmarshalText([]byte("yada"))
	assert.Error(t, err)
	orig := "2014-12-15"
	err = pp.UnmarshalText([]byte(orig))
	assert.NoError(t, err)
	txt, err := pp.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, orig, string(txt))
}

func TestDate_Scan(t *testing.T) {
	ref := time.Now().Truncate(24 * time.Hour).UTC()
	date, str := Date{ref}, ref.Format(RFC3339FullDate)

	values := []interface{}{str, []byte(str), ref}
	for _, value := range values {
		result := Date{}
		(&result).Scan(value)
		assert.Equal(t, date, result, "value: %#v", value)
	}
}

func TestDate_Value(t *testing.T) {
	ref := time.Now().Truncate(24 * time.Hour).UTC()
	date := Date{ref}
	dbv, err := date.Value()
	assert.NoError(t, err)
	assert.Equal(t, dbv, ref)
}
