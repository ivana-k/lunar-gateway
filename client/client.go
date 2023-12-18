package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fullstorydev/grpcurl"
	"github.com/goccy/go-json"
	"github.com/gorilla/mux"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Client struct {
	refClient       *grpcreflect.Client
	cc              *grpc.ClientConn
	descSource      grpcurl.DescriptorSource
	grpcServiceName string
}

func (client *Client) InvokeGrpcMethod(writer http.ResponseWriter, req *http.Request) {
	descSource := client.descSource
	mtdName := req.Context().Value("mtdName").(string)
	reader := prepareReader(req)
	headers := prepareHeaders(req.Header)

	var resultBuffer bytes.Buffer
	rf, formatter, _ := grpcurl.RequestParserAndFormatter(grpcurl.Format("json"), descSource, reader, grpcurl.FormatOptions{})
	h := &grpcurl.DefaultEventHandler{
		Out:            &resultBuffer,
		Formatter:      formatter,
		VerbosityLevel: 0,
	}

	err := grpcurl.InvokeRPC(context.Background(), descSource, client.cc, client.grpcServiceName+"/"+mtdName, headers, h, rf.Next)
	if err != nil {
		log.Printf("[Error] %+v", err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	httpCode := statusCodes[h.Status.Code()]
	if h.Status.Message() != "" {
		http.Error(writer, h.Status.Message(), httpCode)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(httpCode)
	writer.Write(resultBuffer.Bytes())
}

func prepareReader(req *http.Request) io.Reader {
	var reader io.Reader
	params := mux.Vars(req)
	var reqBodyMap map[string]interface{}
	json.NewDecoder(req.Body).Decode(&reqBodyMap)
	if reqBodyMap != nil {
		for k, v := range params {
			reqBodyMap[k] = v
		}
		reqBodyBytes, _ := json.Marshal(reqBodyMap)
		reader = bytes.NewReader(reqBodyBytes)
	} else {
		paramsBytes, _ := json.Marshal(params)
		reader = bytes.NewReader(paramsBytes)
	}

	return reader
}

func prepareHeaders(headers http.Header) []string {
	preparedHeaders := make([]string, 0, len(headers))
	for k, v := range headers {
		if k == "Connection" {
			continue
		}
		preparedHeaders = append(preparedHeaders, fmt.Sprintf("%s: %s", k, v[0]))
	}
	return preparedHeaders
}

var statusCodes = map[codes.Code]int{
	codes.Internal:         http.StatusBadRequest,
	codes.Unauthenticated:  http.StatusUnauthorized,
	codes.PermissionDenied: http.StatusForbidden,
	codes.Unimplemented:    http.StatusNotFound,
	codes.Unavailable:      http.StatusServiceUnavailable,
	codes.OK:               http.StatusOK,
}
