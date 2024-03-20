# Go_Mean_Server

- 현재 공부중인 go 문법
```go
func main() {
	fmt.Println("Hello World!")

	// 변수 선언
	var a int = 10
	var f float32 = 11.
	var i, j, k int = 1, 2, 3
	var d = 5 // 타입 생략 가능 (타입 추론 가능)
	fmt.Println(a)
	fmt.Println(f)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(d)

	// 상수 선언
	const hi = "hi"
	const (
		Sky   = "Blue"
		Rose  = "Red"
		Gyuri = "Awesome"
	)
	fmt.Println(Sky)

	// iota - identifier (상수값을 순차적으로 0부터 부여 가능)
	const (
		apple = iota
		grape
		orange
	)

	// string - ``와 ""로만 선언 가능 (''이거 안쓴다)
	// ``로 쓰인 문자열은 문자 그대로 사용 (`\n`일 경우 줄바꿈 아님)
	// string은 한 번 생성되면 수정 불가

	// type conversion - 타입 자동 변환이 되지 않는다
	// go에서는 타입을 명시적으로 지정해주어야 한다.
	// 명시적 타입변환이 없을 때 runtime error가 발생한다.
	var l int = 100
	var b uint = uint(l)
	var c float32 = float32(l)
	fmt.Println(l, b, c)

	var str = "ABC"
	var bytes = []byte(str)
	var str2 = string(bytes)
	fmt.Println(str, bytes, str2)

	// 연산자
	var p *int
	var m int = 10

	p = &m          // k의 주소 할당
	fmt.Println(*p) // p가 가리키는 주소에 있는 값 출력

	// if/else
	// ()은 안 쓰지만 {}는 무조건 써 준다.
	// else if 혹은 else를 쓸 때는 반드시 전 조건의 {}와 같은 라인에 써준다.
	if m == 1 {
		print("One")
	} else if m == 2 {
		print("Two")
	} else {
		print("Other")
	}
}
```