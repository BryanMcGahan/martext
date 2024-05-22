package terminal

import (
	"golang.org/x/sys/unix"
)

type Terminal struct {
	originalState *unix.Termios
	currentState  *unix.Termios
	fd            int
}

func Init(fd int) (*Terminal, error) {
	var term *Terminal = &Terminal{}

	termios, err := unix.IoctlGetTermios(fd, unix.TIOCGETA)
	if err != nil {
		return nil, err
	}

	term.originalState, term.currentState = termios, termios

	return term, nil
}

func (term *Terminal) MakeRaw() (*Terminal, error) {
	term.currentState.Iflag &^= unix.BRKINT | unix.ICRNL | unix.INPCK | unix.ISTRIP | unix.IXON
	term.currentState.Oflag &^= unix.OPOST
	term.currentState.Cflag |= unix.CS8
	term.currentState.Lflag &^= unix.ECHO | unix.ICANON | unix.ISIG | unix.IEXTEN
	term.currentState.Cflag &^= unix.CSIZE | unix.PARENB
	term.currentState.Cc[unix.VMIN] = 0
	term.currentState.Cc[unix.VTIME] = 1

	if err := unix.IoctlSetTermios(term.fd, unix.TIOCSETA, term.currentState); err != nil {
		return nil, err
	}
	return term, nil
}

func (term *Terminal) Restore() error {
	return unix.IoctlSetTermios(term.fd, unix.TIOCSETA, term.originalState)
}
