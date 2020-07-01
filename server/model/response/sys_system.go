package response

import "taylors/model/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
