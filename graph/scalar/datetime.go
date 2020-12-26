package scalar

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalDateTime for scalar date and datetime
func MarshalDateTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(t.Unix(), 10))
	})
}

// UnmarshalDateTime for scalar date and datetime
func UnmarshalDateTime(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int64); ok {
		return time.Unix(tmpStr, 0), nil
	}
	return time.Time{}, errors.New("time should be a unix timestamp")
}
