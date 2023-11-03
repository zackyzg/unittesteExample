/**
 *
 * @package       split
 * @author        YuanZhiGang <zackyuan@yeah.net>
 * @version       1.0.0
 * @copyright (c) 2013-2023, YuanZhiGang
 */

package split

import (
	"fmt"
	"reflect"
	"strings"
)

// 定义一个按照指定字符串拆分字符串的方法
func Split(s, sep string) []string { // s字符串按照sep分割
	mySlice := make([]string, 0, len(s)) // 最终结果切片
	temp := make([]string, 0, len(s))    // 中间存储切片

	num := len(sep)

	for _, v := range s {

		if reflect.DeepEqual(string(v), sep) { // 比较类型与值是否相同
			// 使用strings.Join函数将字符切片转换为字符串
			// str := strings.Join(chars, "")

			str := strings.Join(temp, "")
			if len(str) != 0 { // 过滤连续出现拆分关键词时,临时切片保存的数据

				mySlice = append(mySlice, str)
			}

			temp = make([]string, 0, len(s)) // 清空中间存储切片
		} else {

			temp = append(temp, string(v)) //字符转字符串
		}
	}

	// 判断切片是否为空
	if len(temp) != 0 {
		str := strings.Join(temp, "")
		mySlice = append(mySlice, str)
	}

	fmt.Println(mySlice)

	return mySlice
}
