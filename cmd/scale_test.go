package cmd

import (
	"testing"

	"github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/require"
)

func TestScale(t *testing.T) {

	k := types.ServiceConfig{
		Name:  "storagenode",
		Image: "foobar",
	}

	err := Scale(&k, "10")
	require.NoError(t, err)

	require.Equal(t, uint64(10), *k.Deploy.Replicas)
}
