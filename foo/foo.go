package foo

type A struct {
	name string
}

type As struct {
	*A

	enabled bool
}

type State struct {
	enabled bool
}

func handleWithStructState(as []*A) []*As {
	ases := make([]*As, len(as))
	for i, a := range as {
		ases[i] = &As{A: a}
		if i%50 == 0 {
			ases[i].enabled = true
		}
	}

	return ases
}
func handleWithMapState(as []*A) map[*A]*State {
	ases := make(map[*A]*State)
	for i, a := range as {
		ases[a] = &State{}
		if i%50 == 0 {
			ases[a].enabled = true
		}
	}

	return ases
}
