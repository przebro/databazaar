package store

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

//ConnectionOptions - Options used to connect to database
type ConnectionOptions struct {
	Scheme  string
	Host    string
	Path    string
	Port    int
	Options map[string]string
}

//Common options for all drivers
const (
	UsernameOption   = "username"
	PasswordOption   = "password"
	RootCACertOption = "cacert"
	ClientCertOption = "clientcert"
	ClientKeyOption  = "clientkey"
	UntrustedOption  = "untrusted"
)

var (
	errOptInvalidFormat error = errors.New("invalid connection string format")
	errOptMissingDriver error = errors.New("invalid connection string; missing driver name")
	errOptInvalidHost   error = errors.New("invalid connection string; invalid host:port format")
	errOptMissingHost   error = errors.New("invalid connection string format; missing host")
	errOptInvalidPort   error = errors.New("invalid connection string format; invalid port")
	errOptEmptyString   error = errors.New("empty connection string")
)

//BuildOptions - parses connection string and builds options
func BuildOptions(connectionString string) (ConnectionOptions, error) {

	if connectionString == "" {
		return ConnectionOptions{}, errOptEmptyString
	}

	sopts := strings.SplitN(connectionString, ";", 2)

	if len(sopts) != 2 {
		return ConnectionOptions{}, errOptInvalidFormat
	}
	if sopts[0] == "" {
		return ConnectionOptions{}, errOptMissingDriver
	}
	driver := sopts[0]
	sopts = strings.SplitN(sopts[1], "/", 2)

	var host, strport string
	var port int
	var path string
	var err error

	if sopts[0] != "" {

		host, strport, err = net.SplitHostPort(sopts[0])

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

	if len(sopts) == 2 {

		sopts = strings.SplitN(sopts[1], "?", 2)
		path = sopts[0]

	}

	if len(sopts) == 2 {

		kvlist := strings.Split(sopts[1], "&")
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
