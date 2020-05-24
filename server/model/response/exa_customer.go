package response

import "taylors/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
