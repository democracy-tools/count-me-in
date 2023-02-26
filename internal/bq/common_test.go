package bq_test

import (
	"testing"

	"github.com/democracy-tools/countmein/internal/bq"
	"github.com/stretchr/testify/require"
)

func TestToInterfaceSlice(t *testing.T) {

	type Person struct {
		Id   string
		Name string
	}

	items, err := bq.ToInterfaceSlice([]*Person{{
		Id:   "1",
		Name: "Israel",
	}, {
		Id:   "3",
		Name: "Moshe",
	}})

	require.Nil(t, err)
	require.Len(t, items, 2)
}
