package ksyun

import (
	"github.com/ksc/ksc-sdk-go/service/eip"
	"github.com/ksc/ksc-sdk-go/service/kcm"
	"github.com/ksc/ksc-sdk-go/service/kec"
	"github.com/ksc/ksc-sdk-go/service/sks"
	"github.com/ksc/ksc-sdk-go/service/slb"
	"github.com/ksc/ksc-sdk-go/service/sqlserver"
	"github.com/ksc/ksc-sdk-go/service/vpc"
	"github.com/ksc/ksc-sdk-go/service/kcsv1"
	"github.com/ksc/ksc-sdk-go/service/kcsv2"
)

type KsyunClient struct {
	region  string
	eipconn *eip.Eip
	slbconn *slb.Slb
	vpcconn *vpc.Vpc
	kecconn *kec.Kec
	sqlserverconn *sqlserver.Sqlserver
	kcmconn *kcm.Kcm
	sksconn *sks.Sks
	kcsv1conn 	*kcsv1.Kcsv1
	kcsv2conn 	*kcsv2.Kcsv2
}
