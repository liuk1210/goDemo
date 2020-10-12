package testgo

import "testing"

func Test_Add(t *testing.T){
	result := Add(1,2)
	if result!=3{
		t.Fatalf("测试失败，预期结果应该为3")
	}
	t.Logf("结果正确")
}