package path

import (
	"os"
	"path/filepath"
)

// bookkeeper
var (
	BkHome                     = os.Getenv("BOOKKEEPER_HOME")
	BkConfigDir                = filepath.FromSlash(BkHome + "/conf")
	BkConfig                   = filepath.FromSlash(BkConfigDir + "/bk_server.conf")
	BkStandaloneConfig         = filepath.FromSlash(BkConfigDir + "/standalone.conf")
	BkOriginalConfig           = filepath.FromSlash(BkConfigDir + "/bk_server_original.conf")
	BkOriginalStandaloneConfig = filepath.FromSlash(BkConfigDir + "/standalone_original.conf")
)

// mate
var (
	BkMatePath              = filepath.FromSlash(BkHome + "/mate")
	BkScripts               = filepath.FromSlash(BkMatePath + "/scripts")
	BkInitScript            = filepath.FromSlash(BkScripts + "/init-bookkeeper.sh")
	BkStartScript           = filepath.FromSlash(BkScripts + "/start-bookkeeper.sh")
	BkStartStandaloneScript = filepath.FromSlash(BkScripts + "/start-bookkeeper-standalone.sh")
	BkCertDir               = filepath.FromSlash(BkMatePath + "/cert")
	ZkClientKeyCert         = filepath.FromSlash(BkCertDir + "/zk_client_key.jks")
	ZkClientTrustCert       = filepath.FromSlash(BkCertDir + "/zk_client_trust.jks")
	BkClientKeyCert         = filepath.FromSlash(BkCertDir + "/bk_client_key.jks")
	BkClientKeyPassword     = filepath.FromSlash(BkCertDir + "/bk_client_key.passwd")
	BkClientTrustCert       = filepath.FromSlash(BkCertDir + "/bk_client_trust.jks")
	BkClientTrustPassword   = filepath.FromSlash(BkCertDir + "/bk_client_trust.passwd")
	BkServerKeyCert         = filepath.FromSlash(BkCertDir + "/bk_server_key.jks")
	BkServerKeyPassword     = filepath.FromSlash(BkCertDir + "/bk_server_key.passwd")
	BkServerTrustCert       = filepath.FromSlash(BkCertDir + "/bk_server_trust.jks")
	BkServerTrustPassword   = filepath.FromSlash(BkCertDir + "/bk_server_trust.passwd")
)
