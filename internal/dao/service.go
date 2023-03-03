package dao

import "github.com/lllllan-fv/gateway-proxy/internal/models"

func ListService(loadType int) (list []*models.GatewayServiceInfo) {
	db := models.GetDB().Model(&models.GatewayServiceInfo{})

	if loadType != 0 {
		db = db.Where("load_type = ?", loadType)
	}

	db.Find(&list)
	return
}
