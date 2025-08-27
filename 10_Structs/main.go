package main // 이 파일이 속한 패키지 이름을 선언. 'main'은 "실행 가능한 프로그램의 시작점"이라는 뜻.

import (       // 다른 코드 묶음(패키지)을 가져오는 구역 시작
	"fmt" // 표준 라이브러리 'fmt' 패키지: Println 등 화면 출력(프린트) 기능 제공
) // import 구역 끝

type person struct { // 'person'이라는 새 자료형(구조체) 정의 시작: 사람을 표현하는 틀
	name    string   // name 필드: 글자(문자열)로 이름을 담음. 예) "nico"
	age     int      // age 필드: 숫자(정수)로 나이를 담음. 예) 18
	favFood []string // favFood 필드: 글자 목록(슬라이스). 예) ["kimchi", "ramen"]
} // 구조체 정의 끝

func main() { // 프로그램이 실행될 때 가장 먼저 호출되는 함수(진입점)
	favFood := []string{"kimchi", "ramen"}                 // 문자열 슬라이스(가변 길이 목록) 만들기
	nico := person{name: "nico", age: 18, favFood: favFood} // person 구조체 값 생성(필드명:값 형태로 명확하게 초기화)
	fmt.Println(nico)                                      // 구조체 전체 출력(기본 형식). 예: {nico 18 [kimchi ramen]}
	fmt.Println(nico.name)                                 // 점(.)으로 특정 필드 접근 → name만 출력. 예: nico
} // main 함수 끝


// 결과값
// {nico 18 [kimchi ramen]}
// nico
