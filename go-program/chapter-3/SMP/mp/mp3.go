package mp

import (
	"fmt"
	"time"
)

//MP3Player doc
type MP3Player struct {
	stat     int
	progress int
}

//Play test
func (p *MP3Player) Play(source string) {
	fmt.Println("Playing MP3 Music", source)

	p.progress = 0

	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond) // 模拟播放
		fmt.Println(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing", source)
}
