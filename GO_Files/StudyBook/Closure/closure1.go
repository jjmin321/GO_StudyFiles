package main

import "fmt"

func main() {
	//클로저(closure)
	//익명함수 사용할 경우 함수를 변수에 할당해서 사용 가능
	//함수 안에서 함수를 선언 및 정의 가능 -> 이 때 외부 함수에 선언된 변수에 접근 가능한 함수
	//함수가 선언되는 순간에 함수가 실행될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)
	//함수를 호출할 때 이전에 존재했던 값을 유지하기 위해서 -> 비동기, 누적카운트(캡쳐 후 1,2,3,4,5 또 캡쳐해서 6,7,8,9,10 ....) -> 무분별한 전역변수 남발하지 않을 수 있음
	//클로저 남발 -> 객체들이 메모리에 존재하므로 -> 메모리 부족, 오버플로우 현상 발생
	//클로저 정확하게 이해하고 사용해야 효율적으로 프로그램 작성 가능.

	//ex1	//일반함수
	multiply := func(x, y int) int {
		return x * y
	}

	r1 := multiply(5, 10)
	fmt.Println("ex1 : ", r1)

	//ex2	//클로저
	m, n := 5, 10            //변수 m, n이 캡쳐됨
	sum := func(c int) int { //익명함수 변수 할당
		return m + n + c //지역 변수 소멸되지 않는다. (함수 호출 시 마다 사용 가능) //함수의 일회용성을 방지할 수 있음.
	}

	r2 := sum(10)
	+fmt.Println("ex2 : ", r2)
}