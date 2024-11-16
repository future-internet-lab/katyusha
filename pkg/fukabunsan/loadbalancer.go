package fukabunsan // 負荷分散 - ふかぶんさん - Load Balancing

import (
	"reflect"

	"github.com/bonavadeur/katyusha/pkg/bonalib"
	"github.com/bonavadeur/katyusha/pkg/hashi"
)

type LoadBalancer struct {
	lbBridge *hashi.Hashi
}

func NewLoadBalancer() *LoadBalancer {
	newLoadBalancerServer := &LoadBalancer{}

	newLoadBalancerServer.lbBridge = hashi.NewHashi(
		"lbBridge",
		hashi.HASHI_TYPE_SERVER,
		BASE_PATH+"/lb-bridge",
		bonalib.Cm2Int("katyusha-threads"),
		reflect.TypeOf(LBRequest{}),
		reflect.TypeOf(LBResponse{}),
		newLoadBalancerServer.LBResponseAdapter,
	)

	return newLoadBalancerServer
}

func (lb *LoadBalancer) LBResponseAdapter(params ...interface{}) (interface{}, error) {
	lbRequest := params[0].(*LBRequest)

	target, err := lb.loadBalance(lbRequest)
	if err != nil {
		panic(err)
	}

	return &LBResponse{Target: target}, nil
}

func (lb *LoadBalancer) loadBalance(lbRequest *LBRequest) (string, error) {
	target := lbRequest.Targets[0]
	return target, nil
}