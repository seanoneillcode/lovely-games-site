package games

type Game struct {
	Id               string
	Name             string
	Description      string
	ShortDescription string
	DescriptionFile  string
	Screenshot       string
	HeaderImage      string
	ReleaseState     string // unreleased, restricted, released
	GameFile         string
	FrameWidth       int
	FrameHeight      int
}
