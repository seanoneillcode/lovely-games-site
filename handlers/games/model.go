package games

type Game struct {
	Id           string
	Name         string
	Description  string
	Screenshot   string
	ReleaseState string // unreleased, restricted, released
}
