package status

//Status enum type
type Status int

//Status constants
type list struct{
	todo Status = iota
	inProgress
	done
}

func (s Status) String() string {
	return [...]string{"TODO", "IN_PROGRESS", "DONE"}[s]
}

// Enum for public use
var Enum = &list{ 
    todo: 0,
    inProgress: 1,
    done: 2,
}
