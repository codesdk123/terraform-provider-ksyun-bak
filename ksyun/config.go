package ksyun

import (
	"github.com/ksc/ksc-sdk-go/ksc"
	"github.com/ksc/ksc-sdk-go/ksc/utils"
	"github.com/ksc/ksc-sdk-go/service/eip"
	"github.com/ksc/ksc-sdk-go/service/epc"
	"github.com/ksc/ksc-sdk-go/service/kcm"
	"github.com/ksc/ksc-sdk-go/service/kcsv1"
	"github.com/ksc/ksc-sdk-go/service/kcsv2"
	"github.com/ksc/ksc-sdk-go/service/kec"
	"github.com/ksc/ksc-sdk-go/service/krds"
	"github.com/ksc/ksc-sdk-go/service/sks"
	"github.com/ksc/ksc-sdk-go/service/slb"
	"github.com/ksc/ksc-sdk-go/service/sqlserver"
	"github.com/ksc/ksc-sdk-go/service/vpc"
)

// Config is the configuration of ksyun meta data
type Config struct {
	AccessKey string
	SecretKey string
	Region    string
	Insecure  bool
}

// Client will returns a client with connections for all product
func (c *Config) Client() (*KsyunClient, error) {
	var client KsyunClient
	//init ksc client info
	client.region = c.Region
	cli := ksc.NewClient(c.AccessKey, c.SecretKey)
	cfg := &ksc.Config{
		Region: &c.Region,
	}
	url := &utils.UrlInfo{
		UseSSL: false,
		Locate: false,
	}
	client.vpcconn = vpc.SdkNew(cli, cfg, url)
	client.eipconn = eip.SdkNew(cli, cfg, url)
	client.slbconn = slb.SdkNew(cli, cfg, url)
	client.kecconn = kec.SdkNew(cli, cfg, url)
	client.sqlserverconn = sqlserver.SdkNew(cli, cfg, url)
	client.krdsconn = krds.SdkNew(cli, cfg, url)
	client.kcmconn = kcm.SdkNew(cli, cfg, url)
	client.sksconn = sks.SdkNew(cli, cfg, url)
	client.kcsv1conn = kcsv1.SdkNew(cli, cfg, url)
	client.kcsv2conn = kcsv2.SdkNew(cli, cfg, url)
	client.epcconn = epc.SdkNew(cli, cfg, url)
	return &client, nil
}
