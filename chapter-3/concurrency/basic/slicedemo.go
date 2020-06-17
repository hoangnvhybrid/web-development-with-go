package main

import (
	"fmt"
)
func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}
func main() {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)
	fmt.Println("slice after function call", nos)
	return
	number := make([]int, 5, 5)
	for i := 0; i < 5; i++ {
		number[i] = i
	}
	fmt.Printf("0: length of slice %d capacity %d\n", len(number), cap(number))
	fmt.Println("number: ", number)
	fmt.Println("number: ", &number[0])
	number[0] = 1
	number = number[2:4]
	fmt.Printf("1: length of slice %d capacity %d\n", len(number), cap(number))
	fmt.Println("number: ", number)
	fmt.Println("number: ", &number[0])
	number = number[:cap(number)+1]
	fmt.Printf("2: length of slice %d capacity %d\n", len(number), cap(number))
	fmt.Println("number: ", number)
	fmt.Println("number: ", &number[0])
	//number = append(number, 2)
	//fmt.Println(&number[0])

	//Câu hỏi hay lắm bạn, bạn hiểu rằng cap (capacity) là con số giúp tối ưu việc reslice khi hàm append được gọi,
	//giả sử mình khai báo 1 slice có cap là 1 (tức là sức chứa của slice sẽ là 1 phần tử)
	//thì khi mình thêm 1 phần thử mới thì slice nó sẽ phải tạo ra 1 array mới có 2 phân tử rồi nó copy phần tử đẩu tiên qua rồi add phần tử mới bạn muốn thêm vào.
	//	Bạn hãy tưởng tượng nếu bạn gọi append cả ngàn phần tử thì chương trình sẽ cấp phát các array mới liên tục điều này ko tốt chút nào,
	//	do đó slice nó sinh ra cái cơ chế  grow slice và con số capacity nó thể hiện điều này. Khi bạn khai báo slice13 := make([]int, 2)
	//thì lúc nào cap=2, khi bạn  append(slice13, 100)
	//tức là lúc này len(slice) = 3 trong khi cap=2 nó ko đủ chỗ chứa phần tử 100,  thì lúc này slice nó tự động grow (mở rộng) lên là 4 (double cái con số 2 lên)
	//và khi số 100 nó được thêm vào tức là len(slice) = 3 như vậy slice nó ko cần tạo ra một array mới khi thêm 1 phần tử ở tương lai nữa vì nó có thể chứa tới 4 phần tử.
	//	khi mình tiếp tục gọi slice13 = append(slice13, 101) thì lúc này len(slice) = 4 và cái do capacity = 4 vẫn đủ chưa nên nó ko cần phải tạo ra một array mới nữa,
	//	tuy nhiên nếu bạn tiếp tục gọi gọi
	//slice13 = append(slice13, 102) thì lúc này len(slice)=5 vượt quá capacity=4 ban đầu thì lúc này cơ chế grow slice sẽ lại xảy ra nó lại tiếp tục double cái capacity ban đầu tức là lúc này new_capacity=4*2=8 cứ thế quá trình này sẽ tiếp diễn nếu bạn còn gọi hàm append . Trường hợp nếu cap(slice) lớn hơn 1024 thì lúc này con số grow slice được tính theo công thức khác bạn tham khảo cái hàm growSlice ở đây nhé https://golang.org/src/runtime/slice.go; ====>newcap += newcap / 4 = con số này có thể khác nhau giữa các máy tính của bạn vì nó bị phụ thuộc vào bộ nhớ .


	slice13 := make([]int, 2)
	fmt.Println(slice13)
	slice13 = append(slice13, 100)
	fmt.Println(slice13)
	fmt.Println(len(slice13))
	fmt.Println(cap(slice13))

	var myArr [4]int
	fmt.Println(&myArr)
	fmt.Println(&myArr[0])
	fmt.Println(&myArr[1])
	fmt.Println(&myArr[2])
	fmt.Println(&myArr[3])

	b := [2]string{"Penn", "Teller"}
	fmt.Println(len(b))
	c := [...]string{"Penn", "Teller"}
	fmt.Println(len(c))
	d := c
	fmt.Println(&c[0])
	fmt.Println(&d[0])

	letters := []string{"a", "b", "c", "d"}
	fmt.Printf("length of slice %d capacity %d\n", len(letters), cap(letters))


	var s []byte
	s = make([]byte, 5, 5)
	// s == []byte{0, 0, 0, 0, 0}
	fmt.Println(s)

	var s1 []byte
	fmt.Printf("s1: length of slice %d capacity %d\n", len(s1), cap(s1))

	e1 := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	var f []byte
	f = e1[1:4]  // == []byte{'o', 'l', 'a'}, sharing the same storage as b
	f[0] = 'g'
	fmt.Println("f: ", f)
	fmt.Println("e1: ", e1)

	fmt.Println(&e1[1])
	fmt.Println(&f[0])

	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s2 := x[:] // a slice referencing the storage of x
	s2[1] ="hoang"
	fmt.Println(s2)
	fmt.Println(x)
	fmt.Println(&x[0])
	fmt.Println(&s2[0])
	//fruitarray :=[...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	//fruitslice := fruitarray[1:3]
	//fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))

	//var intSlice = make([]int, 10)
	//fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	//fmt.Println(reflect.ValueOf(intSlice).Kind())
	//
	//var strSlice []string
	//fmt.Println(reflect.ValueOf(strSlice).Kind())
}

