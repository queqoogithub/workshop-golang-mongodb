package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type EmployeeModel struct {
	Id          *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	EmpId       *string             `json:"empId" bson:"empId,omitempty"`
	EmpName     *string             `json:"empName" bson:"empName,omitempty"`
	EmpPosition *string             `json:"empPosition" bson:"empPosition,omitempty"`
	EmpGender   *string             `json:"empGender" bson:"empGender,omitempty"`
	EmpEmail    *string             `json:"empEmail" bson:"empEmail,omitempty"`
	EmpCompany  *string             `json:"empCompany" bson:"empCompany,omitempty"`
	EmpStatus   *int                `json:"empStatus" bson:"empStatus"`
	CreatedBy   *string             `json:"createdBy" bson:"createdBy,omitempty"`
	CreatedDate *primitive.DateTime `json:"createdDate" bson:"createdDate"`
	UpdatedBy   *string             `json:"updatedBy" bson:"updatedBy,omitempty"`
	UpdatedDate *primitive.DateTime `json:"updatedDate" bson:"updatedDate"`
}
