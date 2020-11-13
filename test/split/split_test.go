package split

// 单元测试的测试文件以**_test.go
import (
	"reflect"
	"testing"
)

// //TestSplit 单个测试用例
// func TestSplit(t *testing.T) {
// 	got := Split("test", "s")
// 	want := []string{"te", "t"}
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("want: %v, got: %v", want, got)
// 	}
// }

func TestSplit(t *testing.T) {
	// 测试函数以Test加函数名开头
	type test struct {
		input string
		sep   string
		want  []string
	}
	// 测试组
	tests := map[string]test{
		"chinese": test{input: "中文测试例子", sep: "测试", want: []string{"中文", "例子"}},
		"english": test{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
	}
	for key, testCase := range tests {
		t.Run(key, func(t *testing.T) {
			// 也可以不使用Run这个函数，调用run可以显示更多信息，命令行测试时需要使用go test -v
			// 还可以单独跑其中一个测试用例，例如go test -run=Split/chinese
			got := Split(testCase.input, testCase.sep)
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("TestCase: %v failed, want: %v, got: %v", key, testCase.want, got)
			}
		})
	}

}

// 性能测试
func BenchmarkSplit(b *testing.B) {
	// go test -bench=Split -benchmem 测试可以输出占用的时间和地址申请次数
	// 减少地址申请次数
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}
