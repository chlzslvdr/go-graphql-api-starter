package scalar

import (
	"errors"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	uuid "github.com/satori/go.uuid"
)

// MarshalID scalar for ID
func MarshalID(id uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(id.String()))
	})
}

// UnmarshalID scalar for ID
func UnmarshalID(v interface{}) (uuid.UUID, error) {
	if tmpStr, ok := v.(string); ok {
		uid, err := uuid.FromString(tmpStr)

		if err != nil {
			return uid, err
		}

		return uid, nil
	}
	return uuid.NewV4(), errors.New("Unable to parse UUID")
}
