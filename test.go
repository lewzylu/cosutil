package main

import (
	//	"bufio"
	"fmt"
	//"os"
	//"io/ioutil"
	"strings"
	"github.com/aws/aws-sdk-go/aws"
	"	"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	//"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	//set endpoint
	myCustomResolver := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		if service == endpoints.S3ServiceID {
			return endpoints.ResolvedEndpoint{
				URL:           "http://cos.cn-south.myqcloud.com",
				SigningRegion: "cos.cn-south",
			}, nil
		}
		return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
	}
	creds := credentials.NewStaticCredentials("AKID15IsskiBQKTZbAo6WhgcBqVls9SmuG00", "ciivKvnnrMvSvQpMAWuIz12pThGGlWRW", "")

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials:      creds,
		Region:           aws.String("cos.cn-south"),
		EndpointResolver: endpoints.ResolverFunc(myCustomResolver),
	}))
	bucket := aws.String("lewzylu099-1252448703")
	svc := s3.New(sess)
	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(strings.NewReader("helloworld.go.exe")),
		Bucket: bucket,
		Key: aws.String("111"),
		
	}
	result, err := svc.PutObject(input)
	fmt.Println(result)
	fmt.Println(err)
}