package response

import "taylors/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
