package adapter

import (
	"clearance/clearance-adapter-for-sale-record/models"
	"context"
	"fmt"
	"testing"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTransform(t *testing.T) {
	Convey("测试SrToClearanceETL的Transform方法", t, func() {
		saleData, _ := time.Parse("2006-01-02", "2019-07-18")
		source := models.AssortedSaleRecordAndDtls{
			AssortedSaleRecords: []models.AssortedSaleRecord{
				{
					OrderId:               1,
					StoreId:               1,
					TotalSalePrice:        200,
					TransactionCreateDate: saleData,
				},
			},
			AssortedSaleRecordDtls: []models.AssortedSaleRecordDtl{
				{
					Quantity:  2,
					SalePrice: 100,
					SkuId:     3,
				},
			},
		}
		saleTransactions, err := SrToClearanceETL{}.Transform(context.Background(), source)
		So(err, ShouldBeNil)
		saleTAndSaleTDtls := saleTransactions.(models.SaleTAndSaleTDtls)
		So(saleTAndSaleTDtls.SaleTransactions[0].OrderId, ShouldEqual, 1)
		So(saleTAndSaleTDtls.SaleTransactions[0].TotalSalePrice, ShouldEqual, 200.00)
		So(saleTAndSaleTDtls.SaleTransactions[0].StoreId, ShouldEqual, 1)
		So(saleTAndSaleTDtls.SaleTransactions[0].SaleDate, ShouldEqual, saleData)

		So(saleTAndSaleTDtls.SaleTransactionDtls[0].SkuId, ShouldEqual, 3)
		So(saleTAndSaleTDtls.SaleTransactionDtls[0].Quantity, ShouldEqual, 2)
		So(saleTAndSaleTDtls.SaleTransactionDtls[0].SalePrice, ShouldEqual, 100)
	})
}

func TestSrToClearanceForSaleRecordETL(t *testing.T) {
	Convey("测试SrToClearanceETL的Run方法", t, func() {
		Convey("可以把DATA 从sale-record导入到Clearance", func() {
			etl := buildSrToClearanceETL()
			ctx := context.Background()
			// data := map[string]string{"brandCode": "PB", "channelType": "POS", "startAt": "2019-08-27 05:42:08", "endAt": "2019-08-27 06:55:33"}
			data := map[string]string{"brandCode": "EE", "channelType": "POS", "startAt": "2019-09-21 16:47:00", "endAt": "2019-09-21 16:49:00"}
			etl.After(SrToClearanceETL{}.ReadyToLoad)
			err := etl.Run(context.WithValue(ctx, "data", data))
			fmt.Println(err)
			So(err, ShouldBeNil)
		})
	})
}
