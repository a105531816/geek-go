package main

import "geek/pkg-bak/apis"

type InterFace interface {
	ReadPersopnalINformation() (*apis.Person, error)
}
