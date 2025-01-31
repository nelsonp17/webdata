package utils

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func NowPgx() pgtype.Timestamp {
	var ts pgtype.Timestamp
	err := ts.Scan(time.Now())
	if err != nil {
		fmt.Println("Error setting timestamp:", err)
	}
	return ts
}
func ParseTimeToPgx(t time.Time) pgtype.Timestamp {
	var ts pgtype.Timestamp
	err := ts.Scan(t)
	if err != nil {
		fmt.Println("Error setting timestamp:", err)
	}
	return ts
}

func ConvertTimestamp(pgTimestamp pgtype.Timestamp) (*time.Time, error) {
	if !pgTimestamp.Valid {
		return nil, fmt.Errorf("el timestamp es nulo")
	}
	return &pgTimestamp.Time, nil
}
