package model

//Status enum type
type Status int

//Status constants
const (
	todo Status = iota
	inProgress
	done
)

func (s Status) String() string {
	return [...]string{"TODO", "IN_PROGRESS", "DONE"}[s]
}
