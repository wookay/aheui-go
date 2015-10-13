# 아희고

Linux, OSX: [![Build Status](https://api.travis-ci.org/wookay/aheui-go.svg?branch=master)](https://travis-ci.org/wookay/aheui-go)
Windows: [![Build status](https://ci.appveyor.com/api/projects/status/jccn0dum0ab8totr?svg=true)](https://ci.appveyor.com/project/wookay/aheui-go)
[![Coverage Status](https://coveralls.io/repos/wookay/aheui-go/badge.svg?branch=master&service=github)](https://coveralls.io/github/wookay/aheui-go?branch=master)

```
$ go get github.com/wookay/aheui-go

$ cat a.go
package main

import (
	"fmt"
	"github.com/wookay/aheui-go"
)

func main() {
	fmt.Print(aheui.Aheui(`
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

$ go run a.go
Hello, world!
```

Aheui(아희) is the esoteric programming language ever to be designed for the Hangul (Hangeul) . The aim of the language is to reflect the graphical design of Hangul.
- https://github.com/aheui/aheui.github.io/blob/master/_posts/2006-10-27-specification.en.markdown
