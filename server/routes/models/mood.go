package models

import (
	"database/sql"
	"errors"
)

type Mood struct {
	ID int `json:"id"`
	Mood int `json:"mood"`
	Note string `json:"note"`
	UserId string `json:"user_id"`
	CreatedAt sql.NullString `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
	DeletedAt sql.NullString `json:"deleted_at"`
}

type MoodRequest struct {
	Mood int `json:"mood"`
	Note string `json:"notes"`
	UserId int `json:"user_id"`
}

func (e* MoodRequest) Validate() error {
	if e.Mood < 1 || e.Mood > 10 {
		return errors.New("mood must be between 1 and 10")
	}
	if e.Note == "" {
		return errors.New("note is required")
	}
	if e.UserId == 0 {
		return errors.New("note is required")	
	}
	return nil
}

type EditMoodRequest struct {
	ID int `json:"id"`
	Mood int `json:"mood"`
	Notes string `json:"notes"`
}

func (e* EditMoodRequest) Validate() error {
	if e.ID == 0 {
		return errors.New("id is required")
	}
	if e.Mood < 1 || e.Mood > 10 {
		return errors.New("mood must be between 1 and 10")
	}
	if e.Notes == "" {
		return errors.New("notes is required")
	}
	return nil
}