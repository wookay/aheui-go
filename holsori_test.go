package aheui

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test좌표(t *testing.T) {
	위치 := 좌표{1, 2}
	assert.Equal(t, 1, 위치.행)
	assert.Equal(t, 2, 위치.렬)
}
