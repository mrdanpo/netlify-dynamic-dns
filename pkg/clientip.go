package pkg

import (
	"io/ioutil"

	"github.com/oscartbeaumont/netlify-dynamic-dns/internal"
)

// These Are The Web Services Used To Determain Your IP (I couldn't get local detection working well)
const (
	ipv4ApiEndpoint = "https://v4.ident.me/"
	ipv6ApiEndpoint = "https://v6.ident.me/"
)

// GetPublicIPv4 returns your public IPv4 addresss as a string
func GetPublicIPv4() (ip string, err error) {
	res, err := internal.Client.Get(ipv4ApiEndpoint)
	if err != nil {
		return "", err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), err
}

// GetPublicIPv6 returns your public IPv6 addresss as a string
func GetPublicIPv6() (ip string, err error) {
	res, err := internal.Client.Get(ipv6ApiEndpoint)
	if err != nil {
		return "", err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), err
}
