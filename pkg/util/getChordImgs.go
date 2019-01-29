package util

import (
	"io/ioutil"
)

func GetChordImages(name string) []string{
	local := "static/chords/"+name
	files , _ := ioutil.ReadDir(local)
	imgArr := make([]string,len(files),6)
	for i,f := range files{
		imgArr[i] = local+"/"+f.Name()
	}
	return imgArr
}
