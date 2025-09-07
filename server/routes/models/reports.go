package models

import (
	"database/sql"
	"errors"
)

type ReportRequest struct {
	Id          int            `json:"id"`
	UserId      string         `json:"user_id"`
	Report      string         `json:"report"`
	MoodData    string         `json:"mood_data"`
	StartDate   sql.NullString `json:"start_date"`
	EndDate     sql.NullString `json:"end_date"`
	CreatedDate sql.NullString `json:"created_date"`
	UpdatedDate sql.NullString `json:"updated_date"`
	DeletedDate sql.NullString `json:"deleted_date"`
}

func (e *ReportRequest) Validate() error {
	if e.Id == 0 {
		return errors.New("no id")
	}
	if e.UserId == "" {
		return errors.New("no id")
	}
	return nil
}
