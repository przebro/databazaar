package store

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// ConnectionOptions - Options used to connect to database
type ConnectionOptions struct {
	Scheme  string
	Host    string
	Path    string
	Port    int
	Options map[string]string
}

// Common options for all drivers
const (
	UsernameOption   = "username"
	PasswordOption   = "password"
	RootCACertOption = "cacert"
	ClientCertOption = "clientcert"
	ClientKeyOption  = "clientkey"
	UntrustedOption  = "untrusted"
	HostnameOption   = "host"
)

var (
	errOptInvalidFormat error = errors.New("invalid connection string format")
	errOptMissingDriver error = errors.New("invalid connection string; missing driver name")
	errOptInvalidHost   error = errors.New("invalid connection string; invalid host:port format")
	errOptMissingHost   error = errors.New("invalid connection string format; missing host")
	errOptInvalidPort   error = errors.New("invalid connection string format; invalid port")
	errOptEmptyString   error = errors.New("empty connection string")
)

// BuildOptions - parses connection string and builds options
func BuildOptions(connectionString string) (ConnectionOptions, error) {

	//query string in format: driver;hostport;path?options

	var err error

	if connectionString == "" {
		return ConnectionOptions{}, errOptEmptyString
	}

	r := strings.Split(connectionString, ";")
	if len(r) != 3 {
		return ConnectionOptions{}, errOptInvalidFormat
	}

	if r[0] == "" {
		return ConnectionOptions{}, errOptMissingDriver
	}

	driver := r[0]

	host := ""
	port := 0
	strport := ""
	if r[1] != "" {

		host, strport, err = net.SplitHostPort(r[1])

		if err != nil {
			return ConnectionOptions{}, errOptInvalidHost
		}

		if host == "" {
			return ConnectionOptions{}, errOptMissingHost
		}

		port, err = strconv.Atoi(strport)
		if err != nil || port < 0 || port > 65535 {
			return ConnectionOptions{}, errOptInvalidPort
		}

	}

	var options = map[string]string{}
	path := ""

	if strings.Contains(r[2], "?") {
		pOptions := strings.Split(r[2], "?")
		path = pOptions[0]

		kvlist := strings.Split(pOptions[1], "&")
		for _, opt := range kvlist {
			kv := strings.SplitN(opt, "=", 2)
			if len(kv) != 2 {
				return ConnectionOptions{}, fmt.Errorf("invalid connection string;invalid option:%s", opt)
			}
			if kv[0] == "" || kv[1] == "" {
				return ConnectionOptions{}, fmt.Errorf("invalid connection string;invalid option:%s", opt)
			}
			options[kv[0]] = kv[1]
		}

	} else {
		path = r[2]
	}

	opt := ConnectionOptions{
		Scheme:  driver,
		Host:    host,
		Port:    port,
		Path:    path,
		Options: options,
	}

	return opt, nil
}
