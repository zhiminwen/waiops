package waiops

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/scram"
)

func tlsConfigWithRootCA(pemFile string) (*tls.Config, error) {
	rootCA, err := x509.SystemCertPool()
	if err != nil {
		log.Printf("Failed to load system cert:%v", err)
		return nil, err
	}

	if rootCA == nil {
		log.Printf("root ca is nil. Creating new cert pool")
		rootCA = x509.NewCertPool()
	}

	cert, err := os.ReadFile(pemFile)
	if err != nil {
		log.Printf("Failed to read the cert: %v", err)
		return nil, err
	}
	ok := rootCA.AppendCertsFromPEM(cert)
	if !ok {
		log.Printf("The cert of %s is not added.", pemFile)
		return nil, fmt.Errorf("Check the certificate %s", pemFile)
	}

	return &tls.Config{
		RootCAs: rootCA,
	}, nil
}

// TLS enabled
func NewSASL512Client(caPemFile, sasl_user, sals_pass string, opts ...kgo.Opt) (*kgo.Client, error) {
	// Create a new SASL/SCRAM-SHA-512 client.
	/*
		client := try.E1(NewSASL512Client(
			config.String("kafka.caCert"),
			config.String("kafka.user"),
			config.String("kafka.password"),
			kgo.SeedBrokers(config.Strings("kafka.brokers")...),
			kgo.DefaultProduceTopic("cp4waiops-cartridge.irdatalayer.replay.alerts"),
		))
	*/

	tlsConfig, err := tlsConfigWithRootCA(caPemFile)
	if err != nil {
		return nil, err
	}

	myOpts := []kgo.Opt{
		kgo.DialTLSConfig(tlsConfig),
		kgo.SASL(scram.Auth{
			User: sasl_user,
			Pass: sals_pass,
		}.AsSha512Mechanism()),
	}
	myOpts = append(myOpts, opts...)
	client, err := kgo.NewClient(myOpts...)
	if err != nil {
		log.Printf("Failed to create kafka client:%v", err)
		return nil, err
	}
	return client, nil
}
