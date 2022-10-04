package bk

import (
	"bookkeeper_mate_go/pkg/config"
	"bookkeeper_mate_go/pkg/path"
	"bookkeeper_mate_go/pkg/util"
	"github.com/paas-mate/gutil"
	"go.uber.org/zap"
	"os"
)

func Config() error {
	if !config.ClusterEnable {
		return ConfigStandalone()
	} else {
		return ConfigCluster()
	}
}

func ConfigCluster() error {
	configProp, err := gutil.ConfigPropFromFile(path.BkOriginalConfig)
	if err != nil {
		return err
	}
	configCommon(configProp)
	if config.KubernetesCluster {
		configProp.Set("advertisedAddress", os.Getenv("HOSTNAME")+".bookkeeper")
		configProp.Set("useHostNameAsBookieID", "true")
	}
	configProp.Set("listeningInterface", "eth0")
	// http
	configProp.Set("httpServerEnabled", "true")
	configProp.Set("httpServerPort", "8080")
	configProp.Set("statsProviderClass", "org.apache.bookkeeper.stats.prometheus.PrometheusMetricsProvider")
	configProp.Set("httpServerClass", "org.apache.bookkeeper.http.vertx.VertxHttpServer")
	// data dir
	configProp.Set("ledgerDirectories", path.BkDataDir)
	configProp.Set("journalDirectories", path.BkJournalDir)
	// zookeeper
	configProp.Set("zkServers", config.ZkAddress)
	configProp.Set("metadataServiceUri", "zk+hierarchical://"+config.MetaDataServiceUri+"/ledgers")
	return configProp.Write(path.BkConfig)
}

func ConfigStandalone() error {
	configProp, err := gutil.ConfigPropFromFile(path.BkOriginalStandaloneConfig)
	if err != nil {
		return err
	}
	configCommon(configProp)
	configProp.Set("advertisedAddress", config.AdvertisedAddress)
	return configProp.Write(path.BkStandaloneConfig)
}

func configCommon(prop *gutil.ConfigProperties) {
	if config.BkTlsEnable {
		// client
		prop.SetBool("tlsClientAuthentication", true)
		prop.Set("clientKeyStore", path.BkClientKeyCert)
		prop.Set("clientKeyStorePasswordPath", path.BkClientKeyPassword)
		prop.Set("clientTrustStore", path.BkClientTrustCert)
		prop.Set("clientTrustStorePasswordPath", path.BkClientTrustPassword)
		// server
		prop.Set("tlsProvider", "JDK")
		prop.Set("tlsProviderFactoryClass", "org.apache.bookkeeper.tls.TLSContextFactory")
		prop.Set("tlsKeyStoreType", "JKS")
		prop.Set("tlsKeyStore", path.BkServerKeyCert)
		prop.Set("tlsKeyStorePasswordPath", path.BkServerKeyPassword)
		prop.Set("tlsTrustStoreType", "JKS")
		prop.Set("tlsTrustStore", path.BkServerTrustCert)
		prop.Set("tlsTrustStorePasswordPath", path.BkServerTrustPassword)
	}
}

func Start() error {
	if config.ClusterEnable {
		stdout, stderr, err := gutil.CallScript(path.BkStartScript)
		util.Logger().Error("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
		return err
	} else {
		stdout, stderr, err := gutil.CallScript(path.BkStartStandaloneScript)
		util.Logger().Error("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
		return err
	}
}
