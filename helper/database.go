// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

import (
	"database/sql"
)

// CloseDatabaseWithError closes the database after an error.
func CloseDatabaseWithError(db *sql.DB, err error) error { _ = "STUB: not implemented"; return nil }

// CloseDatabaseRows closes the database rows.
func CloseDatabaseRows(rows *sql.Rows) { _ = "STUB: not implemented"; return }
