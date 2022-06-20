package callbacks

import "github.com/zhangdapeng520/zdpgo_orm"

type BeforeCreateInterface interface {
	BeforeCreate(*zdpgo_orm.DB) error
}

type AfterCreateInterface interface {
	AfterCreate(*zdpgo_orm.DB) error
}

type BeforeUpdateInterface interface {
	BeforeUpdate(*zdpgo_orm.DB) error
}

type AfterUpdateInterface interface {
	AfterUpdate(*zdpgo_orm.DB) error
}

type BeforeSaveInterface interface {
	BeforeSave(*zdpgo_orm.DB) error
}

type AfterSaveInterface interface {
	AfterSave(*zdpgo_orm.DB) error
}

type BeforeDeleteInterface interface {
	BeforeDelete(*zdpgo_orm.DB) error
}

type AfterDeleteInterface interface {
	AfterDelete(*zdpgo_orm.DB) error
}

type AfterFindInterface interface {
	AfterFind(*zdpgo_orm.DB) error
}
