package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"flag"
	"log"
	"math/big"
	"os"
	"time"
)

var (
	// TODO: forceF = flag.Bool("f", false, "Forces overwriting of existing file(s)")
	crtF = flag.String("crt", "localhost.crt", "Output certificate file path")
	keyF = flag.String("key", "localhost.key", "Output private key file path")
	cnF  = flag.String("cn", "localhost", "Common name (CN) field to use during certificate generation.")
)

func main() {

	flag.Parse()

	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		log.Fatalf("Error during private key generation: %v", err)
	}

	inCertificate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"localhost"},
		},
		IsCA:                  true,
		DNSNames:              []string{*cnF},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24 * 365 * 2),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &inCertificate, &inCertificate, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Fatalf("Error during CreateCertificate: %v", err)
	}

	f, err := os.OpenFile(*crtF, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("Error opening certificate file: %v", err)
	}
	defer f.Close()
	err = pem.Encode(f, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if err != nil {
		log.Printf("Error calling pem.Encode for certificate: %v", err)
	}
	f.Close()

	privateKeyB, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		log.Printf("Error calling x509.MarshalECPrivateKey: %v", err)
	}

	f, err = os.OpenFile(*keyF, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("Error opening key file: %v", err)
	}
	defer f.Close()
	err = pem.Encode(f, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privateKeyB})
	if err != nil {
		log.Printf("Error calling pem.Encode for key: %v", err)
	}
	f.Close()

}
