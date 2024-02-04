package ddltable

type Migration struct {
	Up   string
	Down string
}

func (m Migration) String() string {
	return m.Up + "\n" + m.Down
}
