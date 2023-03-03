package models

import "time"

const (
	HttpLoadType = iota + 1
	TcpLoadType
	GrpcLoadType
)

type GatewayServiceInfo struct {
	ID                     uint `gorm:"primarykey"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
	LoadType               int    `gorm:"not null;default:0;comment:'负载类型 1=http 2=tcp 3=grpc'"`
	ServiceName            string `gorm:"not null;default:'';size:255;comment:'服务名称 6-128 数字字母下划线'"`
	ServiceDesc            string `gorm:"not null;default:'';size:255;comment:'服务描述'"`
	OpenAuth               int    `gorm:"not null;default:0;comment:'(access_control)是否开启权限 1=开启'"`
	BlackList              string `gorm:"not null;default:'';type:text CHARACTER SET utf8 COLLATE utf8_general_ci;comment:'(access_control)黑名单ip'"`
	WhiteList              string `gorm:"not null;default:'';type:text CHARACTER SET utf8 COLLATE utf8_general_ci;comment:'(access_control)白名单ip'"`
	WhiteHostName          string `gorm:"not null;default:'';type:text CHARACTER SET utf8 COLLATE utf8_general_ci;comment:'(access_control)白名单主机'"`
	ClientIPFlowLimit      int    `gorm:"not null;default:0;comment:'(access_control)客户端ip限流'"`
	ServiceFlowLimit       int    `gorm:"not null;default:0;comment:'(access_control)服务端限流'"`
	HeaderTransfor         string `gorm:"not null;default:'';type:text CHARACTER SET utf8 COLLATE utf8_general_ci;comment:'(http)header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔'"`
	RuleType               int    `gorm:"not null;default:0;comment:'(http)匹配类型 0=url前缀url_prefix 1=域名domain '"`
	Rule                   string `gorm:"not null;default:'';size:255;comment:'(http)type=domain表示域名 type=url_prefix时表示url前缀'"`
	NeedHTTPS              int    `gorm:"not null;default:0;comment:'(http)支持https 1=支持'"`
	NeedStripURI           int    `gorm:"not null;default:0;comment:'(http)启用strip_uri 1=启用'"`
	NeedWebSocket          int    `gorm:"not null;default:0;comment:'(http)是否支持websocket 1=支持'"`
	URLRewrite             string `gorm:"not null;default:'';type:text CHARACTER SET utf8 COLLATE utf8_general_ci;comment:'(http)url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔'"`
	Port                   int    `gorm:"not null;default:0;comment:'(grpc/tcp)端口'"`
	CheckMethod            int    `gorm:"not null;default:0;comment:'(load_balance)检查方法 0=tcpchk,检测端口是否握手成功'"`
	CheckTimeout           int    `gorm:"not null;default:0;comment:'(load_balance)check超时时间,单位s'"`
	CheckInterval          int    `gorm:"not null;default:0;comment:'(load_balance)检查间隔, 单位s'"`
	RoundType              int    `gorm:"not null;default:2;comment:'(load_balance)轮询方式 1=random 2=round-robin 3=weight_round-robin 4=ip_hash'"`
	IPList                 string `gorm:"not null;default:'';type:text CHARACTER SET utf8 COLLATE utf8_general_ci;comment:'(load_balance)ip列表'"`
	WeightList             string `gorm:"not null;default:'';type:text CHARACTER SET utf8 COLLATE utf8_general_ci;comment:'(load_balance)权重列表'"`
	ForbidList             string `gorm:"not null;default:'';type:text CHARACTER SET utf8 COLLATE utf8_general_ci;comment:'(load_balance)禁用ip列表'"`
	UpstreamConnectTimeout int    `gorm:"not null;default:0;comment:'(load_balance)建立连接超时, 单位s'"`
	UpstreamHeaderTimeout  int    `gorm:"not null;default:0;comment:'(load_balance)获取header超时, 单位s'"`
	UpstreamIdleTimeout    int    `gorm:"not null;default:0;comment:'(load_balance)链接最大空闲时间, 单位s'"`
	UpstreamMaxIdle        int    `gorm:"not null;default:0;comment:'(load_balance)最大空闲链接数'"`
}

func (GatewayServiceInfo) TableName() string {
	return "gateway_service_info"
}
