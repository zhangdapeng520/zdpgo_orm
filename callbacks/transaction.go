package callbacks

import (
	"github.com/zhangdapeng520/zdpgo_orm"
)

func BeginTransaction(db *zdpgo_orm.DB) {
	if !db.Config.SkipDefaultTransaction && db.Error == nil {
		if tx := db.Begin(); tx.Error == nil {
			db.Statement.ConnPool = tx.Statement.ConnPool
			db.InstanceSet("gorm:started_transaction", true)
		} else if tx.Error == zdpgo_orm.ErrInvalidTransaction {
			tx.Error = nil
		} else {
			db.Error = tx.Error
		}
	}
}

func CommitOrRollbackTransaction(db *zdpgo_orm.DB) {
	if !db.Config.SkipDefaultTransaction {
		if _, ok := db.InstanceGet("gorm:started_transaction"); ok {
			if db.Error != nil {
				db.Rollback()
			} else {
				db.Commit()
			}

			db.Statement.ConnPool = db.ConnPool
		}
	}
}
