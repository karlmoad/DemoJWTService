package drawgames

import "errors"


type playerCard struct{
	Game  	string `json:"game"`
	Plays 	[]play `json:"plays"`
	Epoch 	int64  `json:"epoch"`
}

type play struct{
	Numbers 	map[string]int `json:"numbers"`
}

type DrawGamePrototype interface {
	Draw(plays int) (*playerCard, error)
}

func NewGame(game string) (DrawGamePrototype, error) {

	switch game {
	case "POWERBALL":
		pwrball := newPowerball()
		return pwrball, nil

	default:
		return nil, errors.New("invalid game")
	}
}



