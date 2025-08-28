package main // 이 파일이 '실행 가능한 프로그램'의 시작점임을 알리는 패키지 이름. go run/build의 엔트리 포인트는 항상 main 패키지

import (      // 외부/표준 라이브러리를 가져오는 구문(여러 개를 묶어서 적을 수 있음)
	"fmt" // 표준 출력/서식 지정 출력에 쓰는 패키지. 콘솔에 글자를 찍을 때 사용
)

func main() { // 프로그램이 시작되면 가장 먼저 자동으로 호출되는 특별한 함수(진입점)
	// nico라는 이름의 변수를 선언하고 초기화:
	//  - 자료형: map[string]string (키도 문자열, 값도 문자열인 '사전' 구조)
	//  - 리터럴 { "키": "값", ... } 형태로 즉시 값 채움
	nico := map[string]string{"name": "nico", "age": "12"}

	// map을 순회(iteration)하는 for-range 문:
	//  - 매 반복에서 key에는 한 개의 '키', value에는 그 키에 해당하는 '값'이 들어옴
	//  - 주의: Go의 map은 '순서가 없음(unordered)'. 출력 순서는 실행할 때마다 달라질 수 있음
	for key, value := range nico {
		// fmt.Println: 전달된 값들을 공백으로 구분해 한 줄 출력하고 마지막에 줄바꿈(개행) 추가
		fmt.Println(key, value)
	} // for 블록 끝
} // main 함수 끝

// 출력 값 
// name nico
// age 12

