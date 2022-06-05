package calc

import "testing"

func TestCalcBMI(t *testing.T) {
	inputHeight, inputWeight := 1.0, 1.0
	expectedOutput := 1.1
	t.Logf("开始计算，输入身高：%v，体重：%v，结果：%v", inputHeight, inputWeight, expectedOutput)
	actualOutput, err := CalcBMI(inputHeight, inputWeight)
	t.Logf("实际得到：%v", actualOutput)
	if err != nil {
		t.Errorf("出错：%v", err)
	}
	if actualOutput != expectedOutput {
		t.Fatalf("实际结果与预测结果不一致！实际结果：%v", actualOutput)
	}
}
