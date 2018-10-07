package sdl

// Stager of game
type Stager interface {
	Updater
	Drawer
}

// StageName of game
type StageName string

const (
	// StageNameMainMenu of game
	StageNameMainMenu StageName = "main_menu"
	// StageNameIncome - 1st stage
	StageNameIncome StageName = "income"
)

// Stages - allow to change game stages
type Stages struct {
	active Stager
	m      map[StageName]Stager
}

// StageManager - control active stage
type StageManager interface {
	Next(name StageName)
}

// Add stage
func (s *Stages) Add(name StageName, stage Stager) {
	if s.m == nil {
		s.m = make(map[StageName]Stager)
	}
	if s.active == nil {
		s.active = stage
	}
	s.m[name] = stage
}

// Next - change active stage
func (s *Stages) Next(name StageName) {
	next := s.m[name]
	if next != nil {
		s.active = next
	}
}

// Update active stage
func (s *Stages) Update(e Eventer) {
	s.active.Update(e)
}

// Draw active stage
func (s *Stages) Draw(draw *Draw) {
	s.active.Draw(draw)
}
