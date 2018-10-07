package sdl

// Update information
type Update struct {
	Mouse        *Point
	IsMouseClick bool
}

// Updater interface for object in game screen who can interract with user
type Updater interface {
	Update(e Eventer)
}
