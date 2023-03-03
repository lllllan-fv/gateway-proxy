package service

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lllllan-fv/gateway-proxy/internal/dao"
	"github.com/lllllan-fv/gateway-proxy/internal/models"
)

func HTTPAccessMode(c *gin.Context) (*models.GatewayServiceInfo, error) {
	host := c.Request.Host
	host = host[0:strings.Index(host, ":")]
	path := c.Request.URL.Path

	for _, service := range dao.ListService(models.HttpLoadType) {
		if service.RuleType == models.PrefixURLHTTPRuleType {
			if service.Rule == host {
				return service, nil
			}
		} else {
			if strings.HasPrefix(path, service.Rule) {
				return service, nil
			}
		}
	}

	return nil, errors.New("not matched service")
}
