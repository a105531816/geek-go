package main

type Assets struct {
	assets []Asset
}

func (a *Assets) DoStartWrok() {
	for _, i2 := range a.assets {
		if d, ok := i2.(Door); ok {
			d.Unlock()
			d.Open()
		}
	}
}

func (a *Assets) DoStopWrok() {
	for _, i2 := range a.assets {
		if d, ok := i2.(Door); ok {
			d.Close()
			d.Lock()
		}
	}
}
