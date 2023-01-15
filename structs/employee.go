package structs

type EmployeeReq struct {
	Id          *string `json:"id"`
	EmpId       *string `json:"empId"`
	Keyword     *string `json:"keyword"`
	EmpGender   *string `json:"empGender"`
	EmpStatus   *int    `json:"empStatus"`
	EmpName     *string `json:"empName"`
	EmpPosition *string `json:"empPosition"`
	EmpEmail    *string `json:"empEmail"`
	EmpCompany  *string `json:"empCompany"`
	CreatedBy   *string `json:"createdBy"`
	UpdatedBy   *string `json:"updatedBy"`
	PerPage     *int64  `json:"perPage"`
	Page        *int64  `json:"page"`
}
