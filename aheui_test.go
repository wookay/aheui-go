package aheui

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test아희(t *testing.T) {
	assert.Equal(t, "", 아희(""))
	assert.Equal(t, "", 아희("밯망희"))
	assert.Equal(t, "45", 아희("발밞따망희"))
	assert.Equal(t, "5", 아희("반받다망희"))
}

func Test헬로우월드(t *testing.T) {
	assert.Equal(t, "Hello, world!\n", 아희(`밤밣따빠밣밟따뿌
빠맣파빨받밤뚜뭏
돋밬탕빠맣붏두붇
볻뫃박발뚷투뭏붖
뫃도뫃희멓뭏뭏붘
뫃봌토범더벌뿌뚜
뽑뽀멓멓더벓뻐뚠
뽀덩벐멓뻐덕더벅`))
}
