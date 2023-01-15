package structs

type CompanyReq struct {
	Id        *string `json:"id"`
	ComId     *string `json:"comId"`
	Keyword   *string `json:"keyword"`
	ComName   *string `json:"comName"`
	ComStatus *int    `json:"comStatus"`
	EmpAmount *int    `json:"empAmount"`
	CreatedBy *string `json:"createdBy"`
	UpdatedBy *string `json:"updatedBy"`
	PerPage   *int64  `json:"perPage"`
	Page      *int64  `json:"page"`
}
