package server

import (
	"encoding/json"
	pb "github.com/c12s/celestial/pb"
	// "github.com/c12s/lunar-gateway/model"
	"io"
	"log"
	"net/http"
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

// func mutateToProto(data *model.MutateRequest) *pb.MutateReq {
// 	content := []*pb.Content{}
// 	payload := []*pb.Payload{}

// 	for k, v := range data.Data {
// 		p := &pb.Payload{
// 			Name:   k,
// 			Params: v,
// 		}
// 		payload = append(payload, p)
// 	}

// 	for k, v := range data.Region {
// 		labels := []*pb.KV{}
// 		for lk, lv := range v.Selector {
// 			label := &pb.KV{
// 				Key:   lk,
// 				Value: lv,
// 			}
// 			labels = append(labels, label)
// 		}

// 		c := &pb.Content{
// 			Region:   k,
// 			Clusters: v.Cluster,
// 			Labels:   labels,
// 			Data:     payload,
// 		}
// 		content = append(content, c)
// 	}

// 	req := &pb.MutateRequest{
// 		Content: content,
// 		Kind:    pb.ReqKind_CONFIGS,
// 	}

// 	return req
// }

func listToProto(data map[string]string) *pb.ListReq {
	labels := []*pb.KV{}
	for k, v := range data {
		l := &pb.KV{
			Key:   k,
			Value: v,
		}
		labels = append(labels, l)
	}

	return &pb.ListReq{
		Labels: labels,
		Kind:   pb.ReqKind_CONFIGS,
	}
}

// func RequestToProto(req interface{}, data ...interface{}) {
// 	switch castReq := req.(type) {
// 	case model.MutateReq:
// 		data[0] = mutateToProto(&castReq)
// 	case map[string]string:
// 		data[0] = listToProto(castReq)
// 	default:
// 		data[0] = nil
// 	}
// }
