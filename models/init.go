package models

import "github.com/go-xorm/xorm"

func InitDb(db *xorm.Engine) error {
	return db.Sync(new(SaleTransaction), new(SaleTransactionDtl), new(SaleRecordIdSuccessMapping), new(SaleRecordIdFailMapping))
}
