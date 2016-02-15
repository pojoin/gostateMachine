package gostateMachine

import (
	"testing"
)

//未提交->未审核->审核通过
//			|->审核未通过
func createQjStateMachine() StateMachine {
	qj := NewStateMachine("qj", "1")
	wtj := NewState("未提交")
	wsh := NewState("未审核")
	shtg := NewState("审核通过")
	shwtg := NewState("审核未通过")
	qj.AddTransition("tjsh", wtj, wsh)
	qj.AddTransition("tjtg", wsh, shtg)
	qj.AddTransition("tjwtg", wsh, shwtg)
	return qj
}

//测试请假，从未提交到为审核
func TestStateMachine(t *testing.T) {
	qj := createQjStateMachine()
	tjsh, err := qj.GetTransitionByName("tjsh")
	if err != nil {
		t.Error(err.Error())
	}
	result := tjsh.execute(make(map[string]interface{}))

	if result.GetStateValue() != "未审核" {
		t.Fatal("result", result.GetStateValue())
	}
	t.Log("result", result.GetStateValue())

}
