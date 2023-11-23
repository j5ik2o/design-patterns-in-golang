package singleton

import "testing"
import "github.com/stretchr/testify/require"

func TestGetSingletonInstance(t *testing.T) {
	singleton := GetSingletonInstance()
	singleton2 := GetSingletonInstance()
	require.Equal(t, singleton, singleton2)
}
