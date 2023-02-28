package bk

import (
	"bookkeeper_mate_go/pkg/config"
	"bookkeeper_mate_go/pkg/path"
	"github.com/paas-mate/gutil"
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
	// data
	configProp.Set("ledgerDirectories", path.BkDataDir)
	configProp.SetInt64("logSizeLimit", config.LogSizeLimit)
	// journal
	configProp.Set("journalDirectories", path.BkJournalDir)
	configProp.SetInt("journalMaxSizeMB", config.JournalMaxSizeMB)
	configProp.SetInt("journalMaxBackups", config.JournalMaxBackups)
	configProp.SetBool("journalSyncData", config.JournalSyncData)
	configProp.SetBool("journalWriteData", config.JournalWriteData)
	configProp.SetInt("journalPreAllocSizeMB", config.JournalPreAllocSizeMB)
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
		// http
		prop.Set("httpServerEnabled", "true")
		prop.Set("httpServerPort", "8080")
		prop.Set("statsProviderClass", "org.apache.bookkeeper.stats.prometheus.PrometheusMetricsProvider")
		prop.Set("httpServerClass", "org.apache.bookkeeper.http.vertx.VertxHttpServer")
	}
}
