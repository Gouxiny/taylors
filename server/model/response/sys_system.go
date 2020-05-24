package response

import "taylors/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
