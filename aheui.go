package aheui

import (
	"fmt"
	"strings"
)

type 저장공간 struct {
	공간 []int
	출력 []interface{}
}

func 아희격자(입력 map[좌표]rune) string {
	if 0 == len(입력) {
		return ""
	}
	스택 := 저장공간{}
	커서 := 좌표{0, 0}
	for true {
		글자 := 입력[커서]
		소리 := 스택.기능(글자)
		if 초성_ㅎ == 소리.초 {
			break
		} else {
			스택.가자(소리.기능, 소리.종)
			커서 = 스택.진행(소리.중, 커서)
		}
	}
	출력 := ""
	for _, 아이템 := range 스택.출력 {
		출력 += fmt.Sprint(아이템)
	}
	return 출력
}

func 아희(입력 string) string {
	격자 := map[좌표]rune{}
	여러줄 := strings.Split(strings.TrimSpace(입력), "\n")
	for 행값, 줄 := range 여러줄 {
		for 렬값, 글자 := range []rune(줄) {
			격자[좌표{행값, 렬값}] = 글자
		}
	}
	return 아희격자(격자)
}

func (저장 *저장공간) 가자(기능 interface{}, 종 int) {
	기능.(func(*저장공간, int))(저장, 종)
}
