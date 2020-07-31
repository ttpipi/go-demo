package helper

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"google.golang.org/grpc/credentials"
)

func GetServerCreds() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("../cert/server.pem", "../cert/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("../cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequestClientCert,
		ClientCAs:    certPool,
	})
	return creds
}

func GetClientCreds() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("../cert/client.pem", "../cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("../cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
	return creds
}
