package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func main() {
	var err error

	router := gin.Default()

	// Serve frontend static files and exclude /api routes
	router.Use(func(context *gin.Context) {
		if strings.HasPrefix(context.Request.URL.Path, "/api") {
			context.Next()
		} else {
			http.ServeFile(context.Writer, context.Request, "./website/build"+context.Request.URL.Path)
			context.Abort()
		}
	})

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	tlsConfig := loadTls()

	server := &http.Server{
		Addr:      ":8080",
		Handler:   router,
		TLSConfig: tlsConfig,
	}

	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatalln(err)
	}
}

func loadTls() *tls.Config {
	rawContent := getAwsSslCertificate()

	log.Println(rawContent)

	parts := strings.SplitAfter(rawContent, "-----END PRIVATE KEY-----")
	if len(parts) != 2 {
		log.Fatalf("failed to split the private key and certificate")
	}
	privateKey := parts[0]
	certificate := parts[1]

	cert, err := tls.X509KeyPair([]byte(certificate), []byte(privateKey))
	if err != nil {
		log.Fatalf("failed to parse key pair: %v", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM([]byte(certificate)) {
		log.Fatalf("failed to append certificate to pool")
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	}
}

func getAwsSslCertificate() string {
	secretName := "Webserver-SSL"
	region := "eu-north-1"

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	svc := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatal(err.Error())
	}

	var secretString string = *result.SecretString
	return secretString
}
