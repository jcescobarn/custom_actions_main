package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

func main() {
	fmt.Println("setting Credentials...")
	accessKey := os.Getenv("access_key_id")
	secretKey := os.Getenv("secret_access_key")
	distributionID := os.Getenv("distribution_id")
	invalidationPaths := os.Getenv("invalidations_path")
	array_paths := strings.Split(invalidationPaths, ",")
	quantity := int32(1)
	Paths := types.Paths{
		Quantity: &quantity,
		Items:    array_paths,
	}
	invalidationBatchVar := types.InvalidationBatch{
		CallerReference: &distributionID,
		Paths:           &Paths,
	}

	// Validar si hay valores para las credenciales
	if accessKey == "" || secretKey == "" || distributionID == "" || invalidationPaths == "" {
		fmt.Println("Error: AWS_ACCESS_KEY_ID y AWS_SECRET_ACCESS_KEY deben estar configurados")
		return
	}
	fmt.Println("Credentials Loaded")

	fmt.Println("Establishing connection with aws...")
	config, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     accessKey,
				SecretAccessKey: secretKey,
			},
		}),
		config.WithRegion("us-east-1"),
	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sucessfull connection")
	client := cloudfront.NewFromConfig(config)

	params := &cloudfront.CreateInvalidationInput{
		DistributionId:    &distributionID,
		InvalidationBatch: &invalidationBatchVar,
	}
	fmt.Println("Creating cloudfront invalidation")
	result, err := client.CreateInvalidation(context.TODO(), params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
