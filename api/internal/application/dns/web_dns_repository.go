package dns

import (
	dnsDTO "toolbox/internal/schema/dto/dns"
	dnsVO "toolbox/internal/schema/vo/dns"
	"toolbox/pkg/utils"
)

type Repository interface {
	GetDnsList(orderBy string) ([]*dnsVO.WebDNSTableVO, error)
	DeleteDnsRow(id int64) error
	SaveDnsData(data dnsDTO.WebSaveData) dnsVO.WebSaveDataVO
	NewDNSProxy() (*utils.DNSProxy, error)
}
