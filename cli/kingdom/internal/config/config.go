package config

import (
	"go.zenithar.org/pkg/platform/diagnostic"
	"go.zenithar.org/pkg/platform/jaeger"
	"go.zenithar.org/pkg/platform/prometheus"
)

// Configuration contains spotigraph settings
type Configuration struct {
	Debug struct {
		Enable bool `toml:"enable" default:"false" comment:"allow debug mode"`
	} `toml:"Debug" comment:"###############################\n Debug \n##############################"`

	Instrumentation struct {
		Network    string `toml:"network" default:"tcp" comment:"Network class used for listen (tcp, tcp4, tcp6, unixsocket)"`
		Listen     string `toml:"listen" default:":5556" comment:"Listen address for instrumentation server"`
		Diagnostic struct {
			Enabled bool              `toml:"enabled" default:"true" comment:"Enable diagnostic handlers"`
			Config  diagnostic.Config `toml:"Config" comment:"Diagnostic settings"`
		} `toml:"Diagnostic" comment:"###############################\n Diagnotic Settings \n##############################"`
		Logs struct {
			Level     string `toml:"level" default:"warn" comment:"Log level: debug, info, warn, error, dpanic, panic, and fatal"`
			SentryDSN string `toml:"sentryDSN" comment:"Sentry DSN"`
		} `toml:"Logs" comment:"###############################\n Logs Settings \n##############################"`
		Prometheus struct {
			Enabled bool              `toml:"enabled" default:"true" comment:"Enable metric instrumentation"`
			Config  prometheus.Config `toml:"Config" comment:"Prometheus settings"`
		} `toml:"Prometheus" comment:"###############################\n Prometheus exporter \n##############################"`
		Jaeger struct {
			Enabled bool          `toml:"enabled" default:"true" comment:"Enable trace instrumentation"`
			Config  jaeger.Config `toml:"Config" comment:"Jaeger settings"`
		} `toml:"Jaeger" comment:"###############################\n Jaeger exporter \n##############################"`
	} `toml:"Instrumentation" comment:"###############################\n Instrumentation \n##############################"`

	Core struct {
		AutoMigrate bool   `toml:"-" default:"false"`
		Type        string `toml:"type" default:"postgresql" comment:"Database connector to use: rethinkdb."`
		Hosts       string `toml:"hosts" default:"127.0.0.1:5432" comment:"Database hosts (comma separated)"`
		Database    string `toml:"database" default:"kingdom" comment:"Database namespace"`
		Username    string `toml:"username" default:"" comment:"Database connection username"`
		Password    string `toml:"password" default:"" comment:"Database connection password"`
	} `toml:"Core" comment:"###############################\n Core \n##############################"`

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
	}
}
