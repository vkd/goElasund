package sdl

// Stage of game
type Stage func()

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
	active Stage
	m      map[StageName]Stage
}

// Add stage
func (s *Stages) Add(name StageName, stage Stage) {
	if s.m == nil {
		s.m = make(map[StageName]Stage)
	}
	if s.active == nil {
		s.active = stage
	}
	s.m[name] = stage
}

// SetActive - change active stage
func (s *Stages) SetActive(name StageName) {
	next := s.m[name]
	if next != nil {
		s.active = next
	}
}

// Run active stage
func (s *Stages) Run() {
	s.active()
}
