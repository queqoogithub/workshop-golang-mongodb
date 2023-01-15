package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gosoft.co.th/workshop-api/models"
	repo "gosoft.co.th/workshop-api/repositories"
	st "gosoft.co.th/workshop-api/structs"
)

func InsertEmployeeSv(data models.EmployeeModel) (*string, error) {
	var id *string = nil
	objId, err := repo.InsertEmployeeRepo(data)
	if err != nil {
		return nil, err
	}
	if objId != nil {
		tmpId := objId.Hex()
		id = &tmpId
	}
	return id, nil
}
func UpdateEmployeeSv(param st.EmployeeReq) (*int64, error) {
	var id primitive.ObjectID
	if param.Id != nil {
		objId, _ := primitive.ObjectIDFromHex(*param.Id)
		id = objId
	}
	data := models.EmployeeModel{
		Id:          &id,
		EmpName:     param.EmpName,
		EmpPosition: param.EmpPosition,
		EmpGender:   param.EmpGender,
		EmpEmail:    param.EmpEmail,
		EmpCompany:  param.EmpCompany,
		EmpStatus:   param.EmpStatus,
		UpdatedBy:   param.UpdatedBy,
	}
	return repo.UpdateEmployeeRepo(data)
}
func GetEmployeeSv(param st.EmployeeReq) (*models.EmployeeModel, error) {
	return repo.GetEmployeeRepo(param)
}
func GetEmployeesSv(param st.EmployeeReq) ([]models.EmployeeModel, *int64, *int64, error) {
	return repo.ListEmployeeRepo(param)
}
func DeleteEmployeeSv(param st.EmployeeReq) (*int64, error) {
	return repo.DeleteEmployeeRepo(param.Id)
}
