package main

func main() {
	dd := &downloadFromNetDisk{
		secret:   &dyPeople{mobleNumber: "fdsafasdfasdf"},
		filePath: "接口编程.ppt",
	}
	dd.DownloadFile()
}

type dymeNume interface {
	getNumber() string
}

type dyPeople struct {
	mobleNumber string
} 

func (dy *dyPeople) getNumber() string {
	return "dafasdf"
}
