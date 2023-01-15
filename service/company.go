package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gosoft.co.th/workshop-api/models"
	repo "gosoft.co.th/workshop-api/repositories"
	st "gosoft.co.th/workshop-api/structs"
)

func InsertCompanySv(data models.CompanyModel) (*string, error) {
	var id *string = nil
	objId, err := repo.InsertCompanyRepo(data)
	if err != nil {
		return nil, err
	}
	if objId != nil {
		tmpId := objId.Hex()
		id = &tmpId
	}
	return id, nil
}
func UpdateCompanySv(param st.CompanyReq) (*int64, error) {
	var id primitive.ObjectID
	if param.Id != nil {
		objId, _ := primitive.ObjectIDFromHex(*param.Id)
		id = objId
	}
	data := models.CompanyModel{
		Id:        &id,
		ComName:   param.ComName,
		ComStatus: param.ComStatus,
		EmpAmount: param.EmpAmount,
		UpdatedBy: param.UpdatedBy,
	}
	return repo.UpdateCompanyRepo(data)
}
func GetCompanySv(param st.CompanyReq) (*models.CompanyModel, error) {
	return repo.GetCompanyRepo(param)
}
func GetCompaniesSv(param st.CompanyReq) ([]models.CompanyModel, *int64, *int64, error) {
	return repo.ListCompanyRepo(param)
}
func DeleteCompanySv(param st.CompanyReq) (*int64, error) {
	return repo.DeleteCompanyRepo(param.Id)
}
