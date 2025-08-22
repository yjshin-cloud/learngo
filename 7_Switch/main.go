package main

import "fmt"

func main() {
    a := 2
    b := &a // b는 a의 주소

    *b = 20 // b가 가리키는 메모리의 값을 변경
    fmt.Println(a) // 20 → a의 값이 바뀜
}
