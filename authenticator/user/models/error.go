package models

type Error struct {
	Validation, Internal []error
}
