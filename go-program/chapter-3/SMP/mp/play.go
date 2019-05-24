package mp

import (
	"fmt"
)

type Player interface {
	Play(source string)
}

//Play 此处的 play 不是接口的 Play！！！
func Play(source, mtype string) {
	var p Player

	switch mtype {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("Unsupported music type", mtype)
		return
	}
	p.Play(source)
}
