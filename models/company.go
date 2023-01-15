package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CompanyModel struct {
	Id          *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ComId       *string             `json:"comId" bson:"comId,omitempty"`
	ComName     *string             `json:"comName" bson:"comName,omitempty"`
	EmpAmount   *int                `json:"empAmount" bson:"empAmount,omitempty"`
	ComStatus   *int                `json:"comStatus" bson:"comStatus"`
	CreatedBy   *string             `json:"createdBy" bson:"createdBy,omitempty"`
	CreatedDate *primitive.DateTime `json:"createdDate" bson:"createdDate"`
	UpdatedBy   *string             `json:"updatedBy" bson:"updatedBy,omitempty"`
	UpdatedDate *primitive.DateTime `json:"updatedDate" bson:"updatedDate"`
}
