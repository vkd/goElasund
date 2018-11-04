package sdl

// Stager of game
type Stager interface {
	Initer
	Updater
	Drawer
}

type Initer interface {
	Init()
}

// StageName of game
type StageName string

// Stages of game
const (
	StageNameMainMenu StageName = "main_menu"
	StageNameIncome   StageName = "income"
	StageNameQuit     StageName = "quit"
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
		s.active.Init()
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

// type stageConstructor struct {
// 	i Initer
// 	d Drawer
// 	u Updater
// }

// func (s *stageConstructor) Init() {
// 	if s != nil && s.i != nil {
// 		s.i.Init()
// 	}
// }

// func (s *stageConstructor) Update(e Eventer) {
// 	if s != nil && s.u != nil {
// 		s.u.Update(e)
// 	}
// }

// func (s *stageConstructor) Draw(draw *Draw) {
// 	if s != nil && s.d != nil {
// 		s.d.Draw(draw)
// 	}
// }

// var _ Stager = (*stageConstructor)(nil)

// func StageConstructor(ss ...interface{}) Stager {
// 	out := &stageConstructor{}
// 	for _, s := range ss {
// 		if si, ok := s.(Initer); ok {
// 			out.i = si
// 		}
// 		if su, ok := s.(Updater); ok {
// 			out.u = su
// 		}
// 		if sd, ok := s.(Drawer); ok {
// 			out.d = sd
// 		}
// 	}
// 	return out
// }
