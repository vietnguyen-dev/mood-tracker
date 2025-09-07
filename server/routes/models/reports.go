package models

import (
	"errors"
	"time"
)

type ReportRequest struct {
	Id        int        `json:"id"`
	UserId    string     `json:"user_id"`
	Report    string     `json:"report"`
	MoodData  string     `json:"mood_data"`
	StartDate time.Time  `json:"start_date"`
	EndDate   time.Time  `json:"end_date"`
	CreatedAt time.Time  `json:"created_date"`
	UpdatedAt *time.Time `json:"updated_date"`
	DeletedAt *time.Time `json:"deleted_date"`
}

func (e *ReportRequest) Validate() error {
	if e.Id == 0 {
		return errors.New("no id")
	}
	if e.UserId == "" {
		return errors.New("no user_id")
	}
	if e.Report == "" {
		return errors.New("no report")
	}
	if e.MoodData == "" {
		return errors.New("no mood_data")
	}
	if e.StartDate.IsZero() {
		return errors.New("no start_date")
	}
	if e.EndDate.IsZero() {
		return errors.New("no end_date")
	}
	if e.CreatedAt.IsZero() {
		return errors.New("no created_at")
	}
	if e.UpdatedAt != nil && e.CreatedAt.Before(*e.UpdatedAt) {
		return errors.New("no updated_at")
	}
	if e.DeletedAt != nil && e.CreatedAt.Before(*e.DeletedAt) {
		return errors.New("no deleted_at")
	}
	return nil
}
