package main

import (
	"database/sql"
	"time"

	"github.com/darkluminance/higher-studies-application-tracker/go-server/internal/database"
	"github.com/google/uuid"
)

func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func ToNullInt(i int) sql.NullInt32 {
	return sql.NullInt32{Int32: int32(i), Valid: i != 0}
}

func ToNullBoolean(b bool) sql.NullBool {
	return sql.NullBool{Bool: b, Valid: true}
}

func ToNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: !t.IsZero()}
}

func ToNullReplyVibeEnum(s string) database.NullReplyVibeEnum {
	if s == "" {
		return database.NullReplyVibeEnum{Valid: false}
	}
	return database.NullReplyVibeEnum{ReplyVibeEnum: database.ReplyVibeEnum(s), Valid: true}
}

func convertStringArrayToUUIDArray(id_list []string) []uuid.UUID {
	var ids []uuid.UUID
	for _, id := range id_list {
		uuidID, _ := uuid.Parse(id)
		ids = append(ids, uuidID)
	}
	return ids
}
