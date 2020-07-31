package play

import (
	"fmt"
)

func Play(source, mtype string) {
	var p Player
	switch mtype {
	case "mp3":
		p = MP3Player{}
	case "wav":
		p = WAVPlayer{}
	default:
		fmt.Println("unsupport music type", mtype)
		return
	}
	p.Play(source)
}

type Player interface {
	Play(source string)
}
