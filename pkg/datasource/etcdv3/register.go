package etcdv3

import (
	"github.com/douyu/jupiter/pkg/client/etcdv3"
	"github.com/douyu/jupiter/pkg/conf"
	"github.com/douyu/jupiter/pkg/datasource/manager"
	"github.com/douyu/jupiter/pkg/flag"
	"github.com/douyu/jupiter/pkg/util/xnet"
	"github.com/douyu/jupiter/pkg/xlog"
)

// DataSourceEtcdv3 defines etcdv3 scheme
const DataSourceEtcdv3 = "etcdv3"

func init() {
	manager.Register(DataSourceEtcdv3, func() conf.DataSource {
		var (
			configAddr = flag.String("config")
		)
		if configAddr == "" {
			xlog.Panic("new apollo dataSource, configAddr is empty")
			return nil
		}
		// configAddr is a string in this format:
		// etcdv3://ip:port?basicAuth=true&username=XXX&password=XXX&key=XXX&certFile=XXX&keyFile=XXX&caCert=XXX&secure=XXX

		urlObj, err := xnet.ParseURL(configAddr)
		if err != nil {
			xlog.Panic("parse configAddr error", xlog.FieldErr(err))
			return nil
		}
		etcdConf := etcdv3.DefaultConfig()
		etcdConf.Endpoints = []string{urlObj.Host}

		etcdConf.Secure = urlObj.QueryBool("secure", false)
		etcdConf.BasicAuth = urlObj.QueryBool("basicAuth", false)

		etcdConf.CertFile = urlObj.Query().Get("certFile")
		etcdConf.KeyFile = urlObj.Query().Get("keyFile")
		etcdConf.CaCert = urlObj.Query().Get("caCert")
		etcdConf.UserName = urlObj.Query().Get("username")
		etcdConf.Password = urlObj.Query().Get("password")
		return NewDataSource(etcdConf.MustBuild(), urlObj.Query().Get("key"))
	})
}
