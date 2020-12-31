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
	UsernameOption = "username"
	PasswordOption = "password"
)

var (
	errOptInvalidFormat  error = errors.New("invalid connection string format")
	errOptMissingDriver  error = errors.New("invalid connection string; missing driver name")
	errOptInvalidHost    error = errors.New("invalid connection string; invalid host address")
	errOptMissingHost    error = errors.New("invalid connection string format; missing host")
	errOptInvalidPort    error = errors.New("invalid connection string format; invalid port")
	erroOptInvalidOption error
)

//BuildOptions - parses connection string and builds options
func BuildOptions(connectionString string) (ConnectionOptions, error) {

	sopts := strings.SplitN(connectionString, ";", 2)

	if len(sopts) != 2 {
		return ConnectionOptions{}, errOptInvalidFormat
	}
	if sopts[0] == "" {
		return ConnectionOptions{}, errOptMissingDriver
	}
	driver := sopts[0]
	sopts = strings.SplitN(sopts[1], "/", 2)

	host, strport, err := net.SplitHostPort(sopts[0])

	if err != nil {
		return ConnectionOptions{}, errOptInvalidHost
	}

	if host == "" {
		return ConnectionOptions{}, errOptMissingHost
	}

	port, err := strconv.Atoi(strport)
	if err != nil || port < 0 || port > 65535 {
		return ConnectionOptions{}, errOptInvalidPort
	}

	sopts = strings.SplitN(sopts[1], "?", 2)

	path := sopts[0]
	options := map[string]string{}

	if len(sopts) == 2 {

		kvlist := strings.Split(sopts[1], "&")
		for _, opt := range kvlist {
			kv := strings.SplitN(opt, "=", 2)
			if len(kv) != 2 {
				erroOptInvalidOption = fmt.Errorf("invalid connection string;invalid option:%s", opt)
				return ConnectionOptions{}, erroOptInvalidOption

			}
			if kv[0] == "" || kv[1] == "" {
				erroOptInvalidOption = fmt.Errorf("invalid connection string;invalid option:%s", opt)
				return ConnectionOptions{}, erroOptInvalidOption
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
