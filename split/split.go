/**
 *
 * @package       split
 * @author        YuanZhiGang <zackyuan@yeah.net>
 * @version       1.0.0
 * @copyright (c) 2013-2023, YuanZhiGang
 */

package split

import (
	"strings"
)

//// 定义一个按照指定字符串拆分字符串的方法(只能处理单字符，需要换一种方法)
//func Split(s, sep string) []string { // s字符串按照sep分割
//	mySlice := make([]string, 0, len(s)) // 最终结果切片
//	temp := make([]string, 0, len(s))    // 中间存储切片
//
//	for _, v := range s {
//
//		if reflect.DeepEqual(string(v), sep) { // 比较类型与值是否相同
//			// 使用strings.Join函数将字符切片转换为字符串
//			// str := strings.Join(chars, "")
//
//			str := strings.Join(temp, "")
//			if len(str) != 0 { // 过滤连续出现拆分关键词时,临时切片保存的数据
//
//				mySlice = append(mySlice, str)
//			}
//
//			temp = make([]string, 0, len(s)) // 清空中间存储切片
//		} else {
//
//			temp = append(temp, string(v)) //字符转字符串
//		}
//	}
//
//	// 判断切片是否为空
//	if len(temp) != 0 {
//		str := strings.Join(temp, "")
//		mySlice = append(mySlice, str)
//	}
//
//	fmt.Println(mySlice)
//
//	return mySlice
//}

func Split(s, sep string) (result []string) { // s字符串按照sep分割
	// strings.Index函数用于在一个字符串中查找另一个字符串的首次出现位置。
	//它返回被查找字符串在源字符串中首次出现的索引，如果未找到则返回-1。
	i := strings.Index(s, sep)

	for i > -1 {

		if strings.TrimSpace(s[:i]) != "" {
			result = append(result, s[:i]) // s[:i] 范围是0 ~ i-1,含首不含尾
			//fmt.Println("sss:", s[:i], reflect.ValueOf(s[:i]).Type(), reflect.ValueOf(s[:i]).String())
		}

		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}

	if strings.TrimSpace(s) != "" {
		result = append(result, s)
		//fmt.Println("sss:", s, reflect.ValueOf(s).Type(), reflect.ValueOf(s).String())
	}

	return result
}
