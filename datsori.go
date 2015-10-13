package aheui

import ()

func (저장 *저장공간) 없음() {
}

func (저장 *저장공간) 나눗셈() {
	갑 := 저장.뽑().(int)
	을 := 저장.뽑().(int)
	저장.집어넣기(을 / 갑)
}

func (저장 *저장공간) 덧셈() {
	저장.집어넣기(저장.뽑().(int) + 저장.뽑().(int))
}

func (저장 *저장공간) 곱셈() {
	저장.집어넣기(저장.뽑().(int) * 저장.뽑().(int))
}

func (저장 *저장공간) 나머지() {
	갑 := 저장.뽑().(int)
	을 := 저장.뽑().(int)
	저장.집어넣기(을 % 갑)
}

func (저장 *저장공간) 뽑() interface{} {
	값 := 저장.공간[len(저장.공간)-1]
	저장.공간 = 저장.공간[:len(저장.공간)-1]
	return 값
}

func (저장 *저장공간) 뽑기(받침 int) {
	if isempty(저장.공간) {
	} else {
		switch 받침 {
		case 종성_ㅎ:
			저장.출력 = append(저장.출력, string(저장.뽑().(int)))
		default:
			저장.출력 = append(저장.출력, 저장.뽑())
		}
	}
}

func (저장 *저장공간) 집어넣기(값 int) {
	switch 값 {
	case 종성_ㅇ:
	case 종성_ㅎ:
	default:
		저장.공간 = append(저장.공간, 값)
	}
}

func (저장 *저장공간) 중복() {
	값 := 저장.공간[len(저장.공간)-1]
	저장.공간 = append(저장.공간, 값)
}

func (저장 *저장공간) 선택() {
}

func (저장 *저장공간) 이동() {
}

func (저장 *저장공간) 비교() {
	if 저장.뽑().(int) >= 저장.뽑().(int) {
		저장.집어넣기(1)
	} else {
		저장.집어넣기(0)
	}
}

func (저장 *저장공간) 조건() {
}

func (저장 *저장공간) 뺄셈() {
	저장.집어넣기(-저장.뽑().(int) + 저장.뽑().(int))
}

func (저장 *저장공간) 바꿔치기() {
	end := len(저장.공간) - 1
	저장.공간[end], 저장.공간[end-1] = 저장.공간[end-1], 저장.공간[end]
}

func (저장 *저장공간) 끝냄() {
	if isempty(저장.공간) {
	} else {
		저장.뽑()
	}
}

func isempty(a []int) bool {
	return 0 == len(a)
}

func 닿소리들() []interface{} {
	없음 := func(저장 *저장공간, 값 int) { 저장.없음() }
	나눗셈 := func(저장 *저장공간, 값 int) { 저장.나눗셈() }
	덧셈 := func(저장 *저장공간, 값 int) { 저장.덧셈() }
	곱셈 := func(저장 *저장공간, 값 int) { 저장.곱셈() }
	나머지 := func(저장 *저장공간, 값 int) { 저장.나머지() }
	뽑기 := func(저장 *저장공간, 값 int) { 저장.뽑기(값) }
	집어넣기 := func(저장 *저장공간, 값 int) { 저장.집어넣기(값) }
	중복 := func(저장 *저장공간, 값 int) { 저장.중복() }
	선택 := func(저장 *저장공간, 값 int) { 저장.선택() }
	이동 := func(저장 *저장공간, 값 int) { 저장.이동() }
	비교 := func(저장 *저장공간, 값 int) { 저장.비교() }
	조건 := func(저장 *저장공간, 값 int) { 저장.조건() }
	뺄셈 := func(저장 *저장공간, 값 int) { 저장.뺄셈() }
	바꿔치기 := func(저장 *저장공간, 값 int) { 저장.바꿔치기() }
	끝냄 := func(저장 *저장공간, 값 int) { 저장.끝냄() }
	return []interface{}{없음, 없음, 나눗셈, 덧셈, 곱셈, 나머지, 뽑기, 집어넣기, 중복, 선택, 이동, 없음, 비교, 없음, 조건, 없음, 뺄셈, 바꿔치기, 끝냄}
}

var 닿소리표 = 닿소리들()