package main // 이 파일이 실행 가능한 프로그램(패키지의 시작점)임을 Go에게 알림

import (
	"fmt" // fmt는 출력(print)을 도와주는 기본 제공 기능 (package)
)

func main() { // main 함수는 Go 프로그램이 실행될 때 가장 먼저 호출되는 부분
	names := []string{"nico", "lynn", "dal"}
	// names라는 이름의 "슬라이스"를 선언함
	// 슬라이스는 문자열(string)들의 목록(리스트)임
	// 지금은 "nico", "lynn", "dal" 이라는 3개의 이름이 들어있음

	names = append(names, "flynn")
	// append 함수는 슬라이스에 새로운 요소를 추가할 때 사용됨
	// 여기서는 "flynn"이라는 이름을 기존 슬라이스에 추가하고
	// 그 결과를 다시 names에 저장함 (원래 슬라이스를 직접 수정하지 않기 때문)

	fmt.Println(names)
	// 최종적으로 names 슬라이스의 모든 요소를 출력함
	// 결과: [nico lynn dal flynn]
}
