package scalar

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

const timeLayout = "2006-01-02"

// MarshalDate for scalar date and datetime
func MarshalDate(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format(timeLayout)))
	})
}

// UnmarshalDate for scalar date and datetime
func UnmarshalDate(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse(timeLayout, tmpStr)
	}
	return time.Time{}, errors.New("Date should follow the format yyyy-mm-dd")
}
