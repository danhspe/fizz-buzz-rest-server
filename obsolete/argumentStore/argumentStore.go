package argumentStore

import "github.com/danhspe/fizz-buzz-rest-server/models/arguments"

type ArgumentStore = map[arguments.Arguments]int

func New() ArgumentStore {
	return make(ArgumentStore)
}
