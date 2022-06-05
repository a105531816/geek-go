package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("王强", 0.32)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 1 {
			t.Fatalf("王强的排名预估是第1名，实际结果为：%v\n", randOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("王强的预期体脂率是0.32，实际是:%v\n", fatRateOfWQ)
		}
	}
	inputRecord("李静", 0.28)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 2 {
			t.Fatalf("王强的排名预估是第2名，实际结果为：%v\n", randOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("王强的预期体脂率是0.32，实际是:%v\n", fatRateOfWQ)
		}
	}
	{
		randOfLJ, fatRateOfLJ := getRand("李静")
		if randOfLJ != 1 {

			t.Fatalf("李静的排名预估是第1名，实际结果为：%v\n", randOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("李静的预期体脂率是0.28，实际是:%v\n", fatRateOfLJ)
		}
	}
}

func TestCase2(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("张伟", 0.38)
	inputRecord("李静", 0.28)
	{
		randOfLJ, fatRateOfLJ := getRand("李静")
		if randOfLJ != 1 {
			t.Fatalf("李静的排名预估是第1名，实际结果为：%v\n", randOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("李静的预期体脂率是0.28，实际是:%v\n", fatRateOfLJ)
		}
	}
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 2 {
			t.Fatalf("王强的排名预估是第2名，实际结果为：%v\n", randOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("王强的预期体脂率是0.38，实际是:%v\n", fatRateOfWQ)
		}
	}
	{
		randOfZW, fatRateOfLZW := getRand("张伟")
		if randOfZW != 2 {
			t.Fatalf("张伟的排名预估是第2名，实际结果为：%v\n", randOfZW)
		}
		if fatRateOfLZW != 0.38 {
			t.Fatalf("张伟的预期体脂率是0.38，实际是:%v\n", fatRateOfLZW)
		}
	}
}

func TestCase3(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("李静", 0.28)
	inputRecord("张伟")
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 2 {
			t.Fatalf("王强的排名预估是第2名，实际结果为：%v\n", randOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("王强的预期体脂率是0.38，实际是:%v\n", fatRateOfWQ)
		}
	}
	{
		randOfLJ, fatRateOfLJ := getRand("李静")
		if randOfLJ != 1 {
			t.Fatalf("李静的排名预估是第1名，实际结果为：%v\n", randOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("李静的预期体脂率是0.28，实际是:%v\n", fatRateOfLJ)
		}
	}
	{
		randOfZW, _ := getRand("张伟")
		if randOfZW != 3 {
			t.Fatalf("张伟的排名预估没有名次，实际结果为：%v\n", randOfZW)
		}
	}
}


