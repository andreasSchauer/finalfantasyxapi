package helpers

import "database/sql"



func NullInt32IsZero(n sql.NullInt32) bool {
	return n.Valid == false && n.Int32 == 0
}

func NullBoolIsZero(n sql.NullBool) bool {
	return n.Valid == false && n.Bool == false
}
