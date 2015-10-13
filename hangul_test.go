package aheui

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func Test기능(t *testing.T) {
	스택 := 저장공간{}
	소리 := 스택.기능('희')
	assert.Equal(t, 초성_ㅎ, 소리.초)
	assert.Equal(t, false, 소리.중.움직임)
	assert.Equal(t, 0, 소리.종)
}

func Test한글(t *testing.T) {
	h := "한글"
	assert.Equal(t, 6, len(h))
	assert.Equal(t, 2, utf8.RuneCountInString(h))
	assert.Equal(t, []string{"한", "글"}, strings.Split(h, ""))
	assert.True(t, strings.Contains(h, "한"))
	assert.True(t, strings.Contains(h, "글"))
	assert.True(t, strings.ContainsRune(h, '한'))
	assert.True(t, strings.ContainsRune(h, '글'))
	runes := []rune(h)
	assert.Equal(t, []int32{54620, 44544}, runes)
	assert.Equal(t, "[]int32", reflect.TypeOf(runes).String())
	assert.Equal(t, "int32", reflect.TypeOf(runes[0]).String())
	r := reflect.ValueOf(runes[0])
	assert.Equal(t, "reflect.Value", reflect.TypeOf(r).String())
	assert.Equal(t, "int32", r.Type().String())
	assert.Equal(t, reflect.Int32, r.Kind())
	for i, c := range h {
		assert.True(t, strings.Contains(h, string(c)))
		assert.True(t, strings.ContainsRune(h, c))
		assert.Equal(t, "int32", reflect.TypeOf(c).String())
		assert.Equal(t, "int32", reflect.TypeOf(rune(c)).String())
		assert.Equal(t, "string", reflect.TypeOf(strconv.QuoteRune(c)).String())
		if 0 == i {
			assert.True(t, 54620 == c)
			assert.Equal(t, "한", string(c))
			assert.Equal(t, "'한'", strconv.QuoteRune(c))
		}
	}
}
