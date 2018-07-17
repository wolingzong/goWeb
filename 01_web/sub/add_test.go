// add_test.go  ==> 文件名 _test.go 结尾的默认为测试代码文件
package models

import (
	"testing"
)

//Test开头的默认为测试方法
func TestAdd(t *testing.T) {
	//arrange
	var x, y, res int
	x = 2
	y = 3

	//act
	res = Add(x, y)

	//assert
	if res != 5 {
		t.Fatal("Add的结果不正确")
	}
}
