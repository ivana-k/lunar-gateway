package server

import (
	"encoding/json"
	bPb "github.com/c12s/blackhole/pb"
	cPb "github.com/c12s/celestial/pb"
	"github.com/c12s/lunar-gateway/model"
	"io"
	"log"
	"net/http"
)

const (
	BLACKHOLE = "blackhole"
	CELESTIAL = "celestial"
)

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to encode a JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Failed to write the response body: %v", err)
		return
	}
}

func sendErrorMessage(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(status)
	io.WriteString(w, msg)
}

func mutateToProto(data *model.MutateRequest) *bPb.PutReq {
	req := &bPb.PutReq{}
	return req
}

func mutateNSToProto(data *model.NMutateRequest) *bPb.PutReq {
	req := &bPb.PutReq{}
	return req
}

func listToProto(data map[string]string) *cPb.ListReq {
	labels := []*cPb.KV{}
	for k, v := range data {
		l := &cPb.KV{
			Key:   k,
			Value: v,
		}
		labels = append(labels, l)
	}

	return &cPb.ListReq{
		Labels: labels,
		Kind:   cPb.ReqKind_CONFIGS,
	}
}

func RequestToProto(req interface{}, data ...interface{}) {
	switch castReq := req.(type) {
	case model.MutateRequest:
		data[0] = mutateToProto(&castReq)
	case model.NMutateRequest:
	case map[string]string:
		data[0] = listToProto(castReq)
	default:
		data[0] = nil
	}
}
