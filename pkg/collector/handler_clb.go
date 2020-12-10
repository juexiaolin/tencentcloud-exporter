package collector

import (
	"github.com/go-kit/kit/log"
	"github.com/tencentyun/tencentcloud-exporter/pkg/metric"
	"github.com/tencentyun/tencentcloud-exporter/pkg/util"
)

const (
	ClbNamespace     = "QCE/LB_PUBLIC"
	ClbInstanceidKey = "vip"
)

func init() {
	registerHandler(ClbNamespace, defaultHandlerEnabled, NewClbHandler)
}

type clbHandler struct {
	baseProductHandler
}

func (h *clbHandler) IsMetricMetaVaild(meta *metric.TcmMeta) bool {
	if !util.IsStrInList(meta.SupportDimensions, ClbInstanceidKey) {
		meta.SupportDimensions = append(meta.SupportDimensions, ClbInstanceidKey)
	}

	return true
}

func (h *clbHandler) GetNamespace() string {
	return ClbNamespace
}

func (h *clbHandler) IsMetricVaild(m *metric.TcmMetric) bool {
	return true
}

func NewClbHandler(c *TcProductCollector, logger log.Logger) (handler ProductHandler, err error) {
	handler = &clbHandler{
		baseProductHandler{
			monitorQueryKey: ClbInstanceidKey,
			collector:       c,
			logger:          logger,
		},
	}
	return

}
