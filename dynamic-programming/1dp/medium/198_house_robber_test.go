package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rob_case1(t *testing.T) {
	got := rob([]int{2, 7, 9, 3, 1})
	require.Equal(t, 12, got)
}
