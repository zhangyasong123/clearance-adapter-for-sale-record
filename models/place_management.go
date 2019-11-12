package models

import (
	"clearance/clearance-adapter-for-sale-record/factory"
	"encoding/json"
	"errors"

	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

type Store struct {
	Id         int64       `json:"id"`
	TenantCode string      `json:"tenantCode"`
	Code       string      `json:"code"`
	Name       string      `json:"name"`
	Remark     string      `json:"remark"`
	ElandShops []ElandShop `json:"-" xorm:"-"`
}

type ElandShopInfo struct {
	ElandShopInfos []ElandShop `json:"elandShopInfos"`
}

type ElandShop struct {
	BrandCode string `json:"brandCode"`
	BrandId   int    `json:"brandId"`
	IsChief   bool   `json:"isChief"`
	ShopCode  string `json:"shopCode"`
}

func (Store) GetStore(storeId int64) (*Store, error) {
	var store Store
	queryBuilder := func() xorm.Interface {
		q := factory.GetPmEngine().ID(storeId)
		return q
	}
	has, err := queryBuilder().Get(&store)
	if err != nil {
		return nil, err
	}
	if !has {
		logrus.WithFields(logrus.Fields{
			"storeId": storeId,
		}).Error("Store not find!")
		return nil, errors.New("Store not find!")
	}

	elandShopInfo := ElandShopInfo{}
	json.Unmarshal([]byte(store.Remark), &elandShopInfo)
	store.ElandShops = elandShopInfo.ElandShopInfos
	return &store, nil
}
