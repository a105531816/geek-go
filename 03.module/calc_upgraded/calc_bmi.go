package calc

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	if tall <= 0 {
		panic("身高不能为0或者复数！")
	}
	return weight / (tall * tall), nil
}
