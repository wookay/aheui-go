package aheui

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAheui(t *testing.T) {
	assert.Equal(t, "Hello, world!\n", Aheui(`
밤밣따빠밣밟따뿌
빠맣파빨받밤뚜뭏
돋밬탕빠맣붏두붇
볻뫃박발뚷투뭏붖
뫃도뫃희멓뭏뭏붘
뫃봌토범더벌뿌뚜
뽑뽀멓멓더벓뻐뚠
뽀덩벐멓뻐덕더벅
`))
}
