package limit

package sentinel

import (
"context"
sentinel_api "github.com/alibaba/sentinel-golang/api"
"github.com/alibaba/sentinel-golang/core/base"
"github.com/alibaba/sentinel-golang/core/flow"
"github.com/micro/go-micro/server"
"log"
)

//
//func InitSentinel() {
//	err := sentinel_api.InitDefault()
//	if err != nil {
//		log.Fatalf("Unexpected error: %+v", err)
//	}
//
//	_, err = flow.LoadRules([]*flow.FlowRule{
//		{
//			Resource:        "list-limit",
//			MetricType:      flow.QPS,
//			Count:           10,
//			ControlBehavior: flow.Reject,
//		},
//	})
//	if err != nil {
//		log.Fatalf("Unexpected error: %+v", err)
//		return
//	}
//}

