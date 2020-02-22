package curator

// Curator is primary struct in this lib. More infos at github.com/naxmefy/curator
type Curator struct {
	Source string
}

func NewCurator(source string) *Curator {
	return &Curator{
		Source: source,
	}
}
