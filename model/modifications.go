package model

type ModificationType string

const (
	ModificationTypeMain     ModificationType = "MAIN"
	ModificationTypePosition ModificationType = "POSITION"
	ModificationTypeSequence ModificationType = "SEQUENCE"
)