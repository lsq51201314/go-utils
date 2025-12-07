package gformat

import (
	"database/sql"
	"time"

	"github.com/lsq51201314/go-utils/gtime"
)

func Time(t time.Time) string {
	return t.In(gtime.Location).Format("2006-01-02 15:04:05")
}

func Date(t time.Time) string {
	return t.In(gtime.Location).Format("2006-01-02")
}

func GetTimeNull(t string) sql.NullTime {
	r := sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}
	if n, err := time.ParseInLocation("2006-01-02 15:04:05", t, gtime.Location); err == nil {
		r.Time = n
		r.Valid = true
	}
	return r
}

func GetNullTime(t sql.NullTime) string {
	if !t.Valid {
		return ""
	}
	return Time(t.Time)
}

func GetDateNull(t string) sql.NullTime {
	r := sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}
	if n, err := time.ParseInLocation("2006-01-02", t, gtime.Location); err == nil {
		r.Time = n
		r.Valid = true
	}
	return r
}

func GetNullDate(t sql.NullTime) string {
	if !t.Valid {
		return ""
	}
	return Date(t.Time)
}
