package ssl

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/FoxComm/libs/Godeps/_workspace/src/github.com/FoxComm/goauth2/oauth"
)

func init() {
	InjectCertificates()
}

func InjectCertificates() {
	transport, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		log.Printf("Unexpected underlying type of http.DefaultTransport, aborting certificate injection")
		return
	}

	if transport.TLSClientConfig == nil {
		transport.TLSClientConfig = &tls.Config{}
	}

	if transport.TLSClientConfig.RootCAs != nil {
		log.Printf("RootCAs is not nil, cannot inject certificates")
		return
	}
	transport.TLSClientConfig.RootCAs = oauth.Pool
}
