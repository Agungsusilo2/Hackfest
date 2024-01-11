package helper

import "database/sql"

func ErrorTx(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollBack := tx.Rollback()
		if errRollBack != nil {
			PanicErr(errRollBack)
		}
		panic(err)
	} else {
		err := tx.Commit()
		PanicErr(err)
	}
}
