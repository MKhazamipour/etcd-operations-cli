package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Certs struct {
	etcdCA   string
	etcdCert string
	etcdkey  string
}

func certsFile() *Certs {
	caCert := viper.GetString("caPath.cacert")
	cert := viper.GetString("caPath.cert")
	key := viper.GetString("caPath.Key")
	certs := Certs{
		etcdCA:   caCert,
		etcdCert: cert,
		etcdkey:  key,
	}

	return &certs
}

func tlsConfig() *tls.Config {
	c := certsFile()
	pool := x509.NewCertPool()
	capem, err := os.ReadFile(c.etcdCA)
	if err != nil {
		log.Fatal(err)
	}
	if !pool.AppendCertsFromPEM(capem) {
		log.Fatal("error: failed to add ca to cert pool")
	}
	cert, err := tls.LoadX509KeyPair(c.etcdCert, c.etcdkey)
	if err != nil {
		log.Fatal(err)
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}
	return tlsConfig
}

// returns etcd client to interact with etcd
func newClient() *clientv3.Client {
	etcdep := viper.GetStringSlice("etcd.endpoints")
	tlscfg := tlsConfig()
	config := clientv3.Config{
		Endpoints:   etcdep,
		DialTimeout: 5 * time.Second,
		TLS:         tlscfg,
	}
	client, err := clientv3.New(config)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
