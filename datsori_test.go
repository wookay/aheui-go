package aheui

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test저장공간(t *testing.T) {
	스택 := 저장공간{}
	스택.집어넣기(5)
	스택.집어넣기(70)
	assert.Equal(t, []int{5, 70}, 스택.공간)
	스택.덧셈()
	스택.뽑기(0)
	assert.Equal(t, []interface{}{75}, 스택.출력)
	assert.Equal(t, []int{}, 스택.공간)
}
