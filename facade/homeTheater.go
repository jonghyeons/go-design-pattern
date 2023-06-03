package facade

import "fmt"

type HomeTheaterFacade struct {
	Amp           Amplifier
	Tuner         Tuner
	Player        StreamingPlayer
	Projector     Projector
	TheaterLights TheaterLights
	Screen        Screen
	PopcornPopper PopcornPopper
}

type Amplifier struct{}

func (a *Amplifier) On()               {}
func (a *Amplifier) Off()              {}
func (a *Amplifier) SetSurroundSound() {}

type Tuner struct{}
type StreamingPlayer struct{}

func (sp *StreamingPlayer) On()   {}
func (sp *StreamingPlayer) Play() {}
func (sp *StreamingPlayer) Stop() {}
func (sp *StreamingPlayer) Off()  {}

type Projector struct{}

func (p *Projector) On()             {}
func (p *Projector) Off()            {}
func (p *Projector) WideScreenMode() {}

type TheaterLights struct{}

func (tl *TheaterLights) On()  {}
func (tl *TheaterLights) Off() {}
func (tl *TheaterLights) Dim() {}

type Screen struct{}

func (s *Screen) Up()   {}
func (s *Screen) Down() {}

type PopcornPopper struct{}

func (pp *PopcornPopper) On()  {}
func (pp *PopcornPopper) Pop() {}
func (pp *PopcornPopper) Off() {}

func (htf *HomeTheaterFacade) Constructor(amp Amplifier, tuner Tuner, player StreamingPlayer, projector Projector, theaterLights TheaterLights, screen Screen, popcornPopper PopcornPopper) {
	htf.Amp = amp
	htf.Tuner = tuner
	htf.Player = player
	htf.Projector = projector
	htf.TheaterLights = theaterLights
	htf.Screen = screen
	htf.PopcornPopper = popcornPopper
}

func (htf *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println(movie, "영화 볼 준비 중")
	htf.PopcornPopper.On()
	htf.PopcornPopper.Pop()
	htf.TheaterLights.Dim()
	htf.Screen.Down()
	htf.Projector.On()
	htf.Projector.WideScreenMode()
	htf.Amp.On()
	htf.Player.On()
	htf.Player.Play()
}
func (htf *HomeTheaterFacade) EndMovie() {
	fmt.Println("홈시어터를 끄는 중")
	htf.PopcornPopper.Off()
	htf.TheaterLights.On()
	htf.Screen.Up()
	htf.Projector.Off()
	htf.Amp.Off()
	htf.Player.Stop()
	htf.Player.Off()
}
