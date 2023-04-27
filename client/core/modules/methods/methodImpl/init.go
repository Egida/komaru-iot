package methodImpl

import "math/rand"

func Init() {

}

func randomChoice(yes []string) string {
	if len(yes) < 1 {
		return yes[0]
	}
	return yes[rand.Intn(len(yes))]
}
