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

func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("ssssfdagsgdfgdddfffgggtthdddddd", "d") // 程序结果
	want := []string{"ssssf", "agsg", "fg", "fffgggtth"} // 期望结果

	if !reflect.DeepEqual(got, want) { // 因为slice不能直接比较，借助反射包中的方法比较
		t.Errorf("want:%v,got:%v", want, got)
	}
}

func TestMoreSplit(t *testing.T) {
	got := Split("to divide, or to make a group of people divide, into smaller groups that have very different opinions", "vi")   // 程序结果
	want := []string{"to di", "de, or to make a group of people di", "de, into smaller groups that have very different opinions"} // 期望结果

	if !reflect.DeepEqual(got, want) { // 因为slice不能直接比较，借助反射包中的方法比较
		t.Errorf("want:%v,got:%v", want, got)
	}
}

func TestYuanSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}

	tests := map[string]test{
		"tets1": {input: "hdlghsrtlksfdgtjhs;othr", sep: "h", want: []string{"dlg", "srtlksfdgtj", "s;ot", "r"}},
		"tets2": {input: "dbhfdjhgskjddfks", sep: "dd", want: []string{"dbhfdjhgskj", "fks"}},
		"tets3": {input: "bdsgdufgsdfhfofghoisfgjpgfdhjosdfihgilty", sep: "fg", want: []string{"bdsgdu", "sdfhfo", "hois", "jpgfdhjosdfihgilty"}},
		"tets4": {input: "枯藤老树昏鸦，小桥流水人家，古道西风瘦马。夕阳西下，断肠人在天涯。", sep: "人", want: []string{"枯藤老树昏鸦，小桥流水", "家，古道西风瘦马。夕阳西下，断肠", "在天涯。"}},
	}

	for k, tc := range tests {
		t.Run(k, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)        // 程序结果
			if !reflect.DeepEqual(got, tc.want) { // 因为slice不能直接比较，借助反射包中的方法比较
				//t.Errorf("name:%s want:%#v,got:%#v", k, tc.want, got)
				t.Errorf("want:%#v,got:%#v", tc.want, got)
			}
		})
		//got := Split(tc.input, tc.sep)        // 程序结果
		//if !reflect.DeepEqual(got, tc.want) { // 因为slice不能直接比较，借助反射包中的方法比较
		//	t.Errorf("name:%s want:%#v,got:%#v", k, tc.want, got)
		//}
	}
}
