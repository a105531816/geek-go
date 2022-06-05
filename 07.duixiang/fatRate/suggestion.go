package main

func GetFatRateSuggestion() *fatRageSuggestion {
	return &fatRageSuggestion{
		suggArr: [][][]int{
			{ //男
				{ //18-39
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
				},
				{ //40-59

				},
				{ //60+

				},
			},
			{ //女
				{ //18-39

				},
				{ //40-59

				},
				{ //60+

				},
			},
		},
	}
}

type fatRageSuggestion struct {
	suggArr [][][]int
}

func (f *fatRageSuggestion) GetSuggestion(person *Person) string {
	sexIdx := f.getIndexOfSex(person.sex)
	ageIdx := f.getIndexOfAge(person.age)
	maxFRSupported := len(f.suggArr[sexIdx][ageIdx]) - 1
	frIdx := int(person.fatRate * 100)
	if frIdx > maxFRSupported {
		frIdx = maxFRSupported
	}
	suggIdx := f.suggArr[sexIdx][ageIdx][frIdx]
	return f.translateResult(suggIdx)
	return ""
}

func (f *fatRageSuggestion) getIndexOfSex(sex string) int {
	if sex == "男" {
		return 0
	}
	return 0
}

func (f *fatRageSuggestion) getIndexOfAge(age int) int {
	switch {
	case age >= 18 && age <= 39:
		return 0
	case age >= 40 && age <= 59:
		return 0
	case age >= 60:
		return 0
	default:
		return 0
	}
	return 0
}

func (f *fatRageSuggestion) translateResult(idx int) string {
	switch idx {
	case 0:
		return "偏瘦"
	case 1:
		return "标准"
	case 2:
		return "偏胖"
	case 3:
		return "肥胖"
	case 5:
		return "严重肥胖"
	default:
		return "未知"
	}
}
