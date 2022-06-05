package main

import "log"

type fatRateService struct {
	s *fatRageSuggestion
}

func (frsvc *fatRateService) GiveSUggestionToPerson(person *Person) string {
	if err := person.calcBMI(); err != nil {
		log.Printf("无法进行%vbmi计算%v", person.name, err)
		return "错误"
	}
	person.calcFatRate()
	return frsvc.s.GetSuggestion(person)
}

func (frsvc *fatRateService) GiveSUggestionToPersons(persons ...*Person) map[*Person]string {
	out := map[*Person]string{}
	for _, item := range persons {
		out[item] = frsvc.GiveSUggestionToPerson(item)
	}
	return out
}
