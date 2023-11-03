/**
 *
 * @package       split
 * @author        YuanZhiGang <zackyuan@yeah.net>
 * @version       1.0.0
 * @copyright (c) 2013-2023, YuanZhiGang
 */

package split

import (
	"reflect"
	"testing"
)

//func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
//	got := Split("ssssfdagsgdfgdddfffgggtthdddddd", "d") // 程序结果
//	want := []string{"ssssf", "agsg", "fg", "fffgggtth"} // 期望结果
//
//	if !reflect.DeepEqual(got, want) { // 因为slice不能直接比较，借助反射包中的方法比较
//		t.Errorf("want:%v,got:%v", want, got)
//	}
//}

func TestMoreSplit(t *testing.T) {
	got := Split("to divide, or to make a group of people divide, into smaller groups that have very different opinions", "vi")   // 程序结果
	want := []string{"to di", "de, or to make a group of people di", "de, into smaller groups that have very different opinions"} // 期望结果

	if !reflect.DeepEqual(got, want) { // 因为slice不能直接比较，借助反射包中的方法比较
		t.Errorf("want:%v,got:%v", want, got)
	}
}
