package config

import "go.zenithar.org/pkg/platform"

// Configuration contains spotigraph settings
type Configuration struct {
	Debug struct {
		Enable bool `toml:"enable" default:"false" comment:"allow debug mode"`
	} `toml:"Debug" comment:"###############################\n Debug \n##############################"`

	Instrumentation platform.InstrumentationConfig `toml:"Instrumentation" comment:"###############################\n Instrumentation \n##############################"`

	Core struct {
		AutoMigrate bool   `toml:"-" default:"false"`
		Type        string `toml:"type" default:"postgresql" comment:"Database connector to use: rethinkdb."`
		Hosts       string `toml:"hosts" default:"postgresql://kingdom:changeme@localhost:5432/kingdom?driver=pgx" comment:"Database hosts (comma separated)"`
		Database    string `toml:"database" default:"kingdom" comment:"Database namespace"`
		Username    string `toml:"username" default:"kingdom" comment:"Database connection username"`
		Password    string `toml:"password" default:"changeme" comment:"Database connection password"`
	} `toml:"Core" comment:"###############################\n Core \n##############################"`

	Security struct {
		KeeperURL   string `toml:"keeperURL" default:"base64key://smGbjm71Nxd1Ig5FS0wj9SlbzAIrnolCz9bQQ6uAhl4=" comment:"URL of secret keeper"`
		PasswordKey string `toml:"passwordKey" default:"7vcfai7g8DvLjFeaFKtUBaTvmuql+RNG1X+zqZ390duK1BkQjC7FIuVHZD2LoSdFoOEukVqwAsoQjYjU2KL4e/ktAJdDucNv" comment:"Password at-rest encryption key"`
	} `toml:"Security" comment:"###############################\n Security \n##############################"`

	Server struct {
		Network string `toml:"network" default:"tcp" comment:"Network class used for listen (tcp, tcp4, tcp6, unixsocket)"`
		Listen  string `toml:"listen" default:":5555" comment:"Listen address for gRPC server"`
		UseTLS  bool   `toml:"useTLS" default:"false" comment:"Enable TLS listener"`
		TLS     struct {
			CertificatePath              string `toml:"certificatePath" default:"" comment:"Certificate path"`
			PrivateKeyPath               string `toml:"privateKeyPath" default:"" comment:"Private Key path"`
			CACertificatePath            string `toml:"caCertificatePath" default:"" comment:"CA Certificate Path"`
			ClientAuthenticationRequired bool   `toml:"clientAuthenticationRequired" default:"false" comment:"Force client authentication"`
		} `toml:"TLS" comment:"TLS Socket settings"`
	} `toml:"Server" comment:"###############################\n Server \n##############################"`
}
