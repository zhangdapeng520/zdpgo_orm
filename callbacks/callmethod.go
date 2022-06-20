package callbacks

import (
	"reflect"

	"github.com/zhangdapeng520/zdpgo_orm"
)

func callMethod(db *zdpgo_orm.DB, fc func(value interface{}, tx *zdpgo_orm.DB) bool) {
	tx := db.Session(&zdpgo_orm.Session{NewDB: true})
	if called := fc(db.Statement.ReflectValue.Interface(), tx); !called {
		switch db.Statement.ReflectValue.Kind() {
		case reflect.Slice, reflect.Array:
			db.Statement.CurDestIndex = 0
			for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
				fc(reflect.Indirect(db.Statement.ReflectValue.Index(i)).Addr().Interface(), tx)
				db.Statement.CurDestIndex++
			}
		case reflect.Struct:
			fc(db.Statement.ReflectValue.Addr().Interface(), tx)
		}
	}
}
