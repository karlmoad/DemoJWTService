package drawgames

import (
	"github.com/karlmoad/DemoJWTService/common"
	"fmt"
	"time"
)

const (
	powerballWhiteBalls = 69
	powerballRedBalls = 26
)



type powerball struct {
	whites	[]int
	reds 	[]int
}

func newPowerball() *powerball {
	p :=  &powerball{whites:make([]int,powerballWhiteBalls), reds:make([]int, powerballRedBalls)}
	// preload the entries into the lists
	for i := 0 ; i< powerballWhiteBalls; i++{
		p.whites[i]=i+1
	}

	for i :=0; i< powerballRedBalls; i++ {
		p.reds[i]=i+1
	}
	return p
}


func (p *powerball) Draw(plays int) (*playerCard, error) {

	card := &playerCard{Game:"POWERBALL",Epoch:time.Now().UnixNano()}

	for i := 0;  i< plays; i++ {
		playItem := play{Numbers:make(map[string]int)}

		wBalls := make([]int, powerballWhiteBalls)
		rBalls := make([]int, powerballRedBalls)
		copy(wBalls, p.whites)
		copy(rBalls, p.reds)

		// draw five white balls
		for wi:=1; wi<=5; wi++ {
			ball := common.Random(0,len(wBalls)-1)

			// assign the value to the play
			bn := fmt.Sprintf("%s_%d", "W",wi)
			playItem.Numbers[bn] = wBalls[ball]

			// remove the ball from the list so it cant be pulled again
			wBalls = append(wBalls[:ball], wBalls[ball+1:]...)
		}

		// draw the red ball
		red := common.Random(0,len(rBalls)-1)
		playItem.Numbers["R_1"] = rBalls[red]

		card.Plays = append(card.Plays, playItem)
		time.Sleep(time.Millisecond * 1)
	}

	return card, nil
}





