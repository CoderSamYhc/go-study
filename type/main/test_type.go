package main

func main() {
	// 整型
	//var i int // 根据系统走 32位系统就是 int32  64位系统就是 int64
	//var i2 int8  = -1 // 有符号 int8 int16 int32 int64
	//var ui uint // 无符号 根据系统走 32位系统就是 int32  64位系统就是 int64
	//fmt.Println(i, i2, ui)
	//fmt.Printf("%T\n", i)
	//fmt.Printf("%T\n", ui)
	// 浮点型
	//var f1 float64 = 1.443253624 // 浮点数 没有 符号之分
	//var f2 float64 = 1.443253629 // 精度更高
	//rs := isEqual(f1, f2, 0.1)
	//fmt.Println(rs) // true
	//
	//fmt.Println(f1)
	//fmt.Println(f2)
	//
	//// 字符
	//var b1 byte = 'a' // ASCII 值 utf-8编码 中文3字节
	//fmt.Println(b1)
	//
	//var b2 int = '山' // 中文用int存储
	//fmt.Println(b2)

	// 字符串
	//var s1 string = "sam"
	//fmt.Println(s1)
	//fmt.Println(s1[0]) // 获取字符对应的ascii码

	//str := "hello"
	//str2 := "中国"
	//strTemp := []byte(str) // 切片
	//strTemp[0] = 'c'
	//str = "c" + str[1:]
	//fmt.Println(str)
	//fmt.Println(len(str)) // 5
	//fmt.Println(len(str2)) // 6
	//fmt.Println(utf8.RuneCountInString(str2)) // 2

	//for i := 0; i < len(str); i++ {
	//	fmt.Println(str[i])
	//}

	// i 索引 ch 值
	//for i, ch := range str {
	//	fmt.Println(i, ch)
	//}

	// 其他类型转字符串 string  fmt.Sprintf
	//num := 12
	//num2 := 12.3
	//bool := true
	//mChar := 's'
	//fmt.Printf("%T \n", string(num))
	//str1 := fmt.Sprintf("%d \n", num) // 转整型
	//str2 := fmt.Sprintf("%f \n", num2) // 转浮点型
	//str3 := fmt.Sprintf("%t \n", bool) // 转布尔型
	//str4 := fmt.Sprintf("%c \n", mChar) // 转byte
	//fmt.Println(str1, str2, str3, str4)
	//fmt.Printf("%T \n", str1) // string
	//fmt.Printf("%T \n", str2) // string
	//fmt.Printf("%T \n", str3) // string
	//fmt.Printf("%T \n", str4) // string

	// strconv 转换系列
	//a := strconv.FormatBool(false) // 转布尔
	//b := strconv.FormatFloat(32.32, 'f', 1, 64) // 转浮点 值，fmt表示格式 f 浮点， g 科学计数，精度，进制
	//c := strconv.FormatInt(20, 10) // 转整型 值， 进制
	//d := strconv.Itoa(123) // 转int 值
	//fmt.Println(a, b)
	//fmt.Printf("%T \n", a)
	//fmt.Printf("%T \n", b)
	//fmt.Println(c, d)
	//fmt.Printf("%T \n", c)
	//fmt.Printf("%T \n", d)


	// 字符串转其他类型

	//var str string = "123"
	//var str2 string = "123.321421"
	//i, _ := strconv.ParseInt(str, 10, 64) // 转int
	//j, _ := strconv.ParseFloat(str2, 64) // 转浮点
	//t, _ := strconv.ParseBool("false") // 转布尔
	//l, _ := strconv.Atoi("231")  // 转int
	//
	//fmt.Println(i, j, t, l)
	//fmt.Printf("%T \n", i)
	//fmt.Printf("%T \n", j)
	//fmt.Printf("%T \n", t)
	//fmt.Printf("%T \n", l)

	// 字符串拼接
	//str1 := "hello"
	//str2 := " world"
	//
	//var stringBuilder bytes.Buffer
	//stringBuilder.WriteString(str1)
	//stringBuilder.WriteString(str2)
	//fmt.Println(stringBuilder.String())

	//str := make([]byte, 0) // 类型， 大小
	//str = strconv.AppendInt(str, 100, 10) // 切片， 数值， 进制
	//str = strconv.AppendBool(str, false)
	//str = strconv.AppendFloat(str, 5.5, 'f',3,64) // 切片，数值， 格式化类型， 小数位， 进制
	//fmt.Println(string(str)) // 切片转字符串

	// 复数
	//var t complex128
	//t = 2.1 + 3.14i
	//t := complex(2.1,3.14) // 结果同上
	//fmt.Println(real(t))   // 实部：2.1
	//fmt.Println(imag(t))   // 虚部：3.14
	//fmt.Println(t)   // 虚部：3.14
}

//func isEqual(f1,f2,p float64) bool {
//	// p为用户自定义精度，如：0.00001
//	return math.Abs(f1-f2) < p
//}