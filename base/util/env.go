package util

import "strings"

// Env defines for configuration
type (
	// Environment defines public interfaces
	Environment interface {
		// IsProd returns determined product environment mode
		IsProd() bool
		// IsSentry returns
		IsSentry() bool
		// IsDebug returns
		IsDebug() bool
		// IsProfiler returns
		IsProfiler() bool
		// IsUserFilter returns
		IsUserFilter() bool
		// IsAuth returns
		IsAuth(ipaddr string) bool
		// EnvString gets property variable
		EnvString(prop string) string // must be implement in children
		// EnvInt gets property variable
		EnvInt(prop string) int // must be implement in children
	}

	Env struct {
		// GWURI listens a port as grpc gateway server
		GWURI string `envconfig:"APICUBE_URI" default:"http://0.0.0.0:50000"`
		// URI listens a port as grpc server
		URI string `envconfig:"APICUBE_URI" default:"grpc://0.0.0.0:50001"`

		// RDBURI is set server host and port with db number, that's like DSN
		RDBURI string `envconfig:"APICUBE_RDBURI" default:"redis://127.0.0.1:6379/10"`

		// DLMURI is set distributed lock server host and port with db number, that's like DSN
		DLMURI string `envconfig:"APICUBE_DLMURI" default:"redis://127.0.0.1:6379/9"`

		// URI is set server host and port, which means the same as FQDN
		ESURI string `envconfig:"APICUBE_ESURL" default:"http://127.0.0.1:9200"`

		// FURI is storage uri e.g. s3://data_bucket/path/data.flac or file://
		FURI string `envconfig:"APICUBE_FURI" default:"file://./storage/data.flac"`

		// BQGSURI is bigquery loader storage
		BQGSURI string `envconfig:"APICUBE_BQGSURI" default:"gs://inolab-bigquery-development/apicube/table.json"`

		// DSN is mysql data source name
		DSN string `envconfig:"APICUBE_DSN" default:"root:@tcp(127.0.0.1:3306)/apicube?parseTime=true&loc=Asia%2FTokyo"`

		// Mode defines project environment mode
		Mode string `envconfig:"APICUBE_MODE" default:"development"`

		// Debug controls
		Debug string `envconfig:"APICUBE_DEBUG" default:""` // debug|pprof|something

		// IPFilter blocks not allowed ipaddresses
		IPFilter string `envconfig:"APICUBE_IPFILTER" default:"0.0.0.0,1.1.1.1"`

		// GCProject: GOOGLE_PROJECT, GCLOUD_PROJECT
		GCProject string `envconfig:"GOOGLE_PROJECT" default:""`

		// SentryDSN: SENTRY_DSN
		SentryDSN string `envconfig:"SENTRY_DSN" default:""`
	}
)

// IsProd returns determined product environment mode
func (e *Env) IsProd() bool {
	return e.Mode == "production"
}

// IsSentry returns
func (e *Env) IsSentry() bool {
	return e.SentryDSN != ""
}

// IsDebug returns
func (e *Env) IsDebug() bool {
	return !e.IsProd() || e.Debug == "debug"
}

// IsProfiler returns
func (e *Env) IsProfiler() bool {
	return e.Debug == "pprof"
}

// IsUserFilter returns
func (e *Env) IsUserFilter() bool {
	return e.Debug != "debug"
}

// IsAuth returns
func (e *Env) IsAuth(ipaddr string) bool {
	for _, ip := range strings.Split(e.IPFilter, ",") {
		if ipaddr == ip {
			return true
		}
	}

	return false
}
