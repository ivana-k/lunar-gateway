package server

import (
	"encoding/json"
	bPb "github.com/c12s/blackhole/pb"
	cPb "github.com/c12s/celestial/pb"
	"github.com/c12s/lunar-gateway/model"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	BLACKHOLE = "blackhole"
	CELESTIAL = "celestial"

	all = "all"
	any = "any"

	file   = "file"
	env    = "env"
	action = "action"

	at_once        = "AtOnce"
	rolling_update = "RollingUpdate"

	compare = "compare"
	labels  = "labels"
	sep     = ":"
	kind    = "kind"
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

func cKind(kind string) bPb.CompareKind {
	switch kind {
	case all:
		return bPb.CompareKind_ALL
	case any:
		return bPb.CompareKind_ANY
	default:
		return -1
	}
}

func pKind(kind string) bPb.PayloadKind {
	switch kind {
	case file:
		return bPb.PayloadKind_FILE
	case env:
		return bPb.PayloadKind_ENV
	case action:
		return bPb.PayloadKind_ACTION
	default:
		return -1
	}
}

func sKind(kind string) bPb.StrategyKind {
	switch kind {
	case at_once:
		return bPb.StrategyKind_AT_ONCE
	case rolling_update:
		return bPb.StrategyKind_ROLLING_UPDATE
	default:
		return -1
	}
}

func mutateToProto(data *model.MutateRequest) *bPb.PutReq {
	tasks := []*bPb.PutTask{}
	for _, region := range data.Regions {
		for _, cluster := range region.Clusters {
			labels := []*bPb.KV{}
			for k, v := range cluster.Selector.Labels {
				labels = append(labels, &bPb.KV{Key: k, Value: v})
			}

			payload := []*bPb.Payload{}
			for _, entry := range cluster.Payload {
				values := []*bPb.KV{}
				for k, v := range entry.Content {
					values = append(values, &bPb.KV{Key: k, Value: v})
				}
				pld := &bPb.Payload{
					Kind:  pKind(entry.Kind),
					Value: values,
				}
				payload = append(payload, pld)
			}

			task := &bPb.PutTask{
				RegionId:  region.ID,
				ClusterId: cluster.ID,
				Strategy: &bPb.Strategy{
					Type: cluster.Strategy.Type,
					Kind: sKind(cluster.Strategy.Kind),
				},
				Selector: &bPb.Selector{
					Kind:   cKind(cluster.Selector.Compare[kind]),
					Labels: labels,
				},
				Payload: payload,
			}
			tasks = append(tasks, task)
		}
	}

	return &bPb.PutReq{
		Version: data.Version,
		UserId:  data.Request,
		Mtdata: &bPb.Metadata{
			TaskName:            data.MTData.TaskName,
			Timestamp:           data.MTData.Timestamp,
			Namespace:           data.MTData.Namespace,
			ForceNamespaceQueue: data.MTData.ForceNSQueue,
			Queue:               data.MTData.Queue,
		},
		Tasks: tasks,
	}
}

func mutateNSToProto(data *model.NMutateRequest) *bPb.PutReq {
	labels := []*bPb.KV{}
	for k, v := range data.Labels {
		labels = append(labels, &bPb.KV{Key: k, Value: v})
	}

	return &bPb.PutReq{
		Version: data.Version,
		UserId:  data.Request,
		Mtdata: &bPb.Metadata{
			TaskName:            data.MTData.TaskName,
			Timestamp:           data.MTData.Timestamp,
			Namespace:           data.MTData.Namespace,
			ForceNamespaceQueue: data.MTData.ForceNSQueue,
			Queue:               data.MTData.Queue,
		},
		Extras: labels,
	}
}

func listToProto(data map[string][]string) *cPb.ListReq {
	labels := []*cPb.KV{}
	for _, slabels := range data[labels] {
		pair := strings.Split(slabels, sep)
		l := &cPb.KV{
			Key:   pair[0],
			Value: pair[1],
		}
		labels = append(labels, l)
	}

	return &cPb.ListReq{
		Labels:  labels,
		Compare: data[compare][0],
	}
}

func RequestToProto(req interface{}, data interface{}) {
	switch castReq := req.(type) {
	case model.MutateRequest:
		data = mutateToProto(&castReq)
	case model.NMutateRequest:
	case map[string][]string:
		data = listToProto(castReq)
	default:
		data = nil
	}
}
