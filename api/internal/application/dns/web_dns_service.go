package dns

import (
	dnsDTO "toolbox/internal/schema/dto/dns"
	dnsVO "toolbox/internal/schema/vo/dns"
	"toolbox/pkg/utils"
)

type WebDnsService struct {
	Repo Repository
}

func (r *WebDnsService) ServiceFindDnsList(orderBy string) ([]*dnsVO.WebDNSTableVO, error) {
	return r.Repo.GetDnsList(orderBy)
}

func (r *WebDnsService) ServiceDeleteDnsData(id int64) error {
	return r.Repo.DeleteDnsRow(id)
}

func (r *WebDnsService) ServiceSaveDnsData(data dnsDTO.WebSaveData) dnsVO.WebSaveDataVO {
	return r.Repo.SaveDnsData(data)
}

func (r *WebDnsService) ServiceNewDNSProxy() (*utils.DNSProxy, error) {
	return r.Repo.NewDNSProxy()
}
