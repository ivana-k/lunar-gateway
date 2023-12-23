package client

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	iam "iam-service/proto1"
)

func InterceptRequest(vault_token string) (string, error) {
	conn, err := grpc.Dial("iam-service:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := iam.NewAuthServiceClient(conn)

	getResp, err := client.VerifyToken(context.Background(), &iam.Token{
			Token: vault_token,
	}) 
	if err != nil {
		fmt.Println(err)
		return "", err
	} else {
		return getResp.Token.Jwt, nil
	}
}

