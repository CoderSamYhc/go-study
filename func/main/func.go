package main

import "fmt"

//func makeSuffix(suffix string) func (string) string {
//	return func (str string) string {
//		if strings.HasSuffix(str, suffix) {
//			return str
//		}
//		str += ".jpg"
//		return str
//
//	}
//
//}

func main() {
	//var a = "a.jpg"
	//var b = "b"
	//f := makeSuffix(".jpg")
	//fmt.Println(f(a))
	//fmt.Println(f(b))

	//var str = "!!!hello,中国,中文!!!"
	//fmt.Println(len(str)) // len() 获取字符串长度

	//r := []rune(str)
	//
	//for _, v := range r {
	//	fmt.Printf("%c \n", v)
	//}

	//fmt.Println(strings.Count(str, "l")) // strings.Count(str, substr) 统计子字符串出现的次数
	//fmt.Println(strings.EqualFold("abc", "ABC")) // strings.EqualFold(str1, str2) 不区分大小写，对比字符是否相等
	//
	//fmt.Println(strings.Index(str, "中")) //strings.Index(str, substr) 返回子字符串第一次出现的索引位置
	//fmt.Println(strings.LastIndex(str, "中文")) // strings.LastIndex(str, substr) 返回子字符串最后一次出现的索引位置
	//
	//fmt.Println(strings.Replace(str, "中文", "你好", 15))  // strings.Replace(str, old, new, index) 子字符串替换 index = -1 表示全部替换
	//
	//fmt.Println(strings.Split(str, ",")) // strings.Split(str, substr) 根据子字符串切割，返回数组
	//fmt.Println(strings.ToUpper(str)) // strings.ToUpper(str) 英文字母转大写
	//str = "HELLO,中国"
	//fmt.Println(strings.ToLower(str)) // strings.ToLower(str) 英文字母转小写
	//
	//fmt.Println(strings.TrimSpace(str)) // strings.TrimSpace(str) 剔除字符串两边的空格
	//
	//fmt.Println(strings.Trim(str, "!")) // strings.Trim(str, cutstr) 剔除两边指定字符串
	//fmt.Println(strings.TrimLeft(str, "!")) // strings.TrimLeft(str, cutstr) 剔除左边指定字符串
	//fmt.Println(strings.TrimRight(str, "!")) // strings.TrimRight(str, cutstr) 剔除右边指定字符串
	//
	//fmt.Println(strings.HasPrefix(str, "!")) // strings.HasPrefix(str, prefix) 验证字符串是否以指定字符串开头
	//fmt.Println(strings.HasSuffix(str, "!")) // strings.HasSuffix(str, suffix) 验证字符串是否以指定字符串结尾

	//now := time.Now() // 返回Time构造体

	//fmt.Printf("%T,  %v \n", now, now)

	//fmt.Println(now.Year()) // 年
	//fmt.Printf("%T, %v \n", now.Month(), now.Month()) // 月
	//fmt.Println(int(now.Month())) // 月
	//fmt.Println(now.Day()) // 日
	//fmt.Println(now.Hour()) // 时
	//fmt.Println(now.Minute()) // 分
	//fmt.Println(now.Second()) // 秒
	//y, m, d := now.Date() // 同时获取年月日
	//fmt.Println(y,int(m),d)

	//date := fmt.Sprintf("%d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	//fmt.Println(date)

	//fmt.Printf("%v \n", now.Format("2020/09/15 13:42:00"))

	//fmt.Println(now.Unix()) // 时间戳
	//fmt.Println(now.UnixNano()) // 毫秒

	//var num = getNum(7)
	//fmt.Println(num)

	n1 := 1
	n2 := 3
	n1, n2 = swap(&n1, &n2)

	fmt.Println(n1,n2)
}

//func getNum(num int) int {
//
//	if num == 1 || num == 2 {
//		return 1
//	}
//
//	return getNum(num - 1) + getNum(num - 2)
//}

func swap(n1 *int, n2 *int) (int, int) {
	*n1, *n2 = *n2, *n1
	return *n1, *n2
}