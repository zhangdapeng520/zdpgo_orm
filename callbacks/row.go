package callbacks

import (
	"github.com/zhangdapeng520/zdpgo_orm"
)

func RowQuery(db *zdpgo_orm.DB) {
	if db.Error == nil {
		BuildQuerySQL(db)
		if db.DryRun {
			return
		}

		if isRows, ok := db.Get("rows"); ok && isRows.(bool) {
			db.Statement.Settings.Delete("rows")
			db.Statement.Dest, db.Error = db.Statement.ConnPool.QueryContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
		} else {
			db.Statement.Dest = db.Statement.ConnPool.QueryRowContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
		}

		db.RowsAffected = -1
	}
}
