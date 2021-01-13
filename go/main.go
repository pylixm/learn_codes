package main

import "fmt"

func main() {

	//name := "DeanWu"
	//// 相加
	//fmt.Println("我叫" + name)
	//
	//// 下标取值
	//fmt.Println(name[0])  // 直接取，是对应的ascii码，需要传下
	//fmt.Printf("%c\n", name[0])
	//fmt.Println(name[:3])
	//
	//// 使用内建函数len获取字符串长度
	//fmt.Println(len(name))
	//
	//
	//// 字符串包含
	//fmt.Println(strings.Contains(name, "a"))
	//
	//// 字符串开头，结尾
	//fmt.Println(strings.HasPrefix(name, "D"))
	//fmt.Println(strings.HasSuffix(name, "u"))
	//
	//// 字符串分割组合
	//arr := strings.Split(name, "e")
	//fmt.Println(strings.Join(arr, "e"))
	//
	//// 字符串格式化
	////fmt.Printf("%s, %.2f \n", a, f64)
	//fmt.Println(fmt.Sprintf("我叫，%s", name))


	// 数组
	//// 先声明，后赋值
	//var a [2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//
	//// 声明时，直接赋值
	//b := [5]int{10,20,30,40,50}
	//
	//// 可直接通过下标来访问数组
	//fmt.Println(a)
	//fmt.Println(a[0], a[1])
	//fmt.Println(b)
	//fmt.Println(b[1], b[2])
	//
	//// 通过len()函数可获取数组长度
	//fmt.Println(len(a))
	//fmt.Println(len(b))
	//
	//// 数组元素赋值
	//b[1] = 25
	//fmt.Println(b)
	//fmt.Println(b[1])

	// 多维数组
	//// 声明一个二维整型数组，两个维度分别存储 4 个元素和 2 个元素
	//var arr [4][2]int
	//arr[0] = [2]int{10, 11}
	//arr[0][1] = 15
	//
	//// 使用数组字面量来声明并初始化一个二维整型数组
	//arr1 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	//
	//// 声明并初始化外层数组中索引为 1 和 3 的元素
	//arr2 := [4][2]int{1: {20, 21}, 3: {40, 41}}
	//// 声明并初始化外层数组和内层数组的单个元素
	//arr3 := [4][2]int{1: {0: 20}, 3: {1: 41}}
	//
	//// 将 arr1 的索引为 1 的维度复制到一个同类型的新数组里
	//var arr4 [2]int = arr1[1]
	//// 将外层数组的索引为 1、内层数组的索引为 0 的整型值复制到新的整型变量里
	//var value int = arr1[1][0]
	//
	//// 使用索引获取数组值
	//fmt.Println(arr)
	//fmt.Println(arr1)
	//fmt.Println(arr1[0][1])
	//fmt.Println(arr2)
	//fmt.Println(arr3)
	//fmt.Println(arr4)
	//fmt.Println(value)

	// 切片
	//// 切片使用make(slice, len, cap) 声明, cap可省略，省略时，等于len
	//sli := make([]int, 2, 3)
	//fmt.Println(sli)  // 未赋值时，各元素为对应的零值
	//fmt.Printf("%p %v %v \n", sli, len(sli), cap(sli))
	//
	//// 声明时，初始化
	//nums := []int{10, 20, 30, 40}
	//fmt.Println(nums)
	//fmt.Printf("%p %v %v \n", nums, len(nums), cap(nums))
	//
	//// nil 切片，指针为空，长度和容量为0
	//var nums1 []int
	//fmt.Println(nums1)
	//
	//// 空切片，指针为指向一个空数组，长度和容量为0
	//var nums2 = make([]int, 0)
	//nums3 := []int{}
	//fmt.Println(nums2)
	//fmt.Println(nums3)
	//
	//nums[0] = 12
	//fmt.Println(nums)
	//
	//fmt.Println(nums)
	//// 从下标1，到下标2，新切片为容量4
	//fmt.Println(nums[1:2])  // 长度为2-1 =1，容量1
	//fmt.Println(nums[1:2:4]) // 长度为2-1=1，容量为4-1=3
	////slice[i:]  // 从 i 切到最尾部
	////slice[:j]  // 从最开头切到 j(不包含 j)
	////slice[:]   // 从头切到尾，等价于复制整个 slice
	//
	//
	//// 切片的追加, 使用内建函数 append(src, item) 返回 新的切片
	//nums = append(nums, 10) // 添加一个
	//nums = append(nums, 10, 20) // 同时添加多个
	//newNums := append(nums, nums1...)  // 合并两个切片
	//fmt.Println(newNums)
	//fmt.Println(nums)
	//
	// 切片的复制, 使用内建函数 copy(dst, src) 返回复制的元素个数
	// 赋值时，接收的切片容量需要大于原切片，否则复制失败
	//copyNums := make([]int, 5)
	//count := copy(copyNums, nums)
	//fmt.Println(count)
	//fmt.Println(copyNums)
	//fmt.Printf("slice: %p ,len: %v ,cap: %v \n", copyNums, len(copyNums), cap(copyNums))
	//fmt.Println(nums)
	//fmt.Printf("slice: %p ,len: %v ,cap: %v \n", nums, len(nums), cap(nums))

	//
	//
	//// 数组共享的切片
	//sli := make([]int, 3)  // 定义一个长度，大小都为3的切片
	//sli1 := sli[:2]  // 由切片再创建切片，sli 和sli1 底层引用同一个数组
	//// slice: 0xc0000160c0 ,len: 3 ,cap: 3
	//fmt.Printf("slice: %p ,len: %v ,cap: %v \n", sli, len(sli), cap(sli))
	//// slice1: 0xc0000160c0 ,len: 2 ,cap: 3
	//fmt.Printf("slice1: %p ,len: %v ,cap: %v \n", sli1, len(sli1), cap(sli1))
	//
	//sli[0] = 10
	//fmt.Println(sli)   // [10 0 0]
	//fmt.Println(sli1)  // [10 0]
	//
	//// 切片扩容
	//fmt.Printf("%p %v %v \n", sli, len(sli), cap(sli)) // 0xc0000160c0 3 3
	//sli = append(sli, 1)
	//sli = append(sli, 2)
	//fmt.Println(sli)
	//fmt.Printf("%p %v %v \n", sli, len(sli), cap(sli)) // 0xc00001a0c0 5 6


	//sli := make([]int, 3)  // 定义一个长度，大小都为3的切片
	//fmt.Println(sli)
	//sli1 := sli[:3:3] // 长度
	//fmt.Println(sli1)
	//fmt.Printf("%p %v %v \n", sli, len(sli), cap(sli))
	//fmt.Printf("%p %v %v \n", sli1, len(sli1), cap(sli1))
	//sli1 = append(sli1, 2) // sli1 容量为1， append时，
	//fmt.Printf("%p %v %v \n", sli1, len(sli1), cap(sli1))

	// 初始化
	//// 使用make函数
	//myMap := make(map[string]int)
	//fmt.Println(myMap) // map[]
	//
	//// 使用字面量
	//myResume := map[string]string{"name": "DeanWu", "job": "SRE"}
	//fmt.Println(myResume)  // map[job:SRE name:DeanWu]

	// 声明一个空map
	//var myResume1 map[string]string
	//myResume1["name"] = "DeanWu"  //panic: assignment to entry in nil map
	// 空的map，系统并没有分配内存，并能赋值。

	// 键值的类型可以是内置的类型，也可以是结构类型，只要这个值可以使用 == 运算符做比较
	// 切片、函数以及包含切片的结构类型，这些类型由于具有引用语义，可被其他引用改变原值，不能作为映射的键。
	//myMap1 := map[[]string]int{}
	//fmt.Println(myMap1)  // invalid map key type []string

	//// 通过key赋值
	//myResume["job"] = "web deployment"
	//fmt.Println(myResume)  // map[job:web deployment name:DeanWu]
	//
	//// 获取某个key的值
	//value, exists := myResume["name"]
	//if exists {
	//	fmt.Println(value) // DeanWu
	//}
	//value1 := myResume["name"]
	//if value1 != ""{
	//	fmt.Println(value1) // DeanWu
	//	// 推荐上边的写法，因为即使map无此key也会返回对应的零值。需要根据数据类型，做相应的判断，不如上边的统一，方便。
	//}
	//
	//// 删除键值对
	//delete(myResume, "job")
	//delete(myResume, "year")  // 当map中没有这个key时，什么都不执行。
	//fmt.Println(myResume)  // map[name:DeanWu]
	//
	//
	//// map嵌套
	//myNewResume := map[string]map[string]string{
	//	"name": {
	//		"first": "Dean",
	//		"last":"Wu",
	//	},
	//}
	//fmt.Println(myNewResume) // map[name:map[first:Dean last:Wu]]

	//// 定义结构体，自动名必须大写开头，作为公共变量
	//type Vertex struct {
	//	X int
	//	Y int
	//}
	//// 初始化
	//var ver Vertex
	//ver.X = 4  // 可使用. 来赋值和访问结构体
	//fmt.Println(ver.X)  // 4
	//
	//// 可使用指针来访问
	//v := Vertex{1, 2}
	//p := &v
	//p.X = 1e9
	//fmt.Println(v)  // {1000000000 2}
	//
	//type NewVertex struct {
	//	Vertex
	//	Z int
	//}
	//var v1 NewVertex
	//v1.X = 12
	//v1.Z = 13
	//fmt.Println(v1.X) // 12
	//fmt.Println(v1)  // {{12 0} 13}

	//x:= 5.5
	//n:= 6.1
	//lim := 20.0
	//v := math.Pow(x, n)
	//// if 判断
	//if v < lim {
	//	fmt.Println(fmt.Sprint("v --->:",v))
	//}
	//// if 简约形式，可赋值加判断
	//if v := math.Pow(x, n); v < lim {
	//	fmt.Println(fmt.Sprint("v --->:",v))
	//}
	//fmt.Println(fmt.Sprint("lim --->:",lim))
	//
	//// if else
	//if v:=math.Pow(x,n);v<lim{
	//	fmt.Println(fmt.Sprint("v --->:",v))
	//}else{
	//	fmt.Println(fmt.Sprint("lim --->:",lim))
	//}

	//// switch 后边可以直接跟值，也可以赋值加值
	//os := runtime.GOOS
	//switch os {
	//case "darwin":
	//	fmt.Println("OS X.")
	//case "linux":
	//	fmt.Println("Linux.")
	//default:
	//	fmt.Printf("%s.", os)
	//}
	//
	//switch os := runtime.GOOS; os {
	//default:
	//	fmt.Printf("%s.", os)
	//case "darwin":
	//	fmt.Println("OS X.")
	//case "linux":
	//	fmt.Println("Linux.")
	//}
	//
	//// switch也可以没有条件，在case处再判断
	//t := time.Now()
	//switch {
	//case t.Hour() < 12:
	//	fmt.Println("Good morning!")
	//case t.Hour() < 17:
	//	fmt.Println("Good afternoon.")
	//default:
	//	fmt.Println("Good evening.")
	//}

	//m := 2
	//if n:=2; n !=0{
	//	n = 1
	//	fmt.Println(n)  // 1
	//	fmt.Println(m)
	//	m = 4
	//}
	//fmt.Println(m)  // 4
	////fmt.Println(n)  // undefined: n


	//sum := 0
	//// 标准的for 写法
	//for i := 0; i < 10; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)
	//
	//// 省略初始化语句
	//for ; sum < 1000; {
	//	sum += sum
	//}
	//fmt.Println(sum)
	//
	//// 省略初始化和后置语句，用法类似python中的while
	//for sum < 1000 {
	//	sum += sum
	//}
	//
	//// 无限循环
	//for {
	//}

	// 遍历列表
	ids := []int{1,2,3,4,5,6}
	for i, v := range ids{
		fmt.Println(i, v)
	}

	// 遍历map
	own := map[string]string{
		"name": "DeanWu",
		"age": "30",
	}
	for k, v := range own{
		fmt.Println(k, v )
	}

	//
	//sum := 0
	//// 可使用break 来退出信息，或使用continue 跳过本次循环的之后的语句
	//for i := 0; i < 10; i++ {
	//	sum += i
	//	if sum > 5{
	//		break
	//		//continue
	//	}
	//	fmt.Println(sum)
	//}
	//fmt.Println(sum)

}