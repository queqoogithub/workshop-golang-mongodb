package repositories

import (
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gosoft.co.th/workshop-api/baseDB"
	"gosoft.co.th/workshop-api/models"
	st "gosoft.co.th/workshop-api/structs"
)

func InsertEmployeeRepo(req models.EmployeeModel) (*primitive.ObjectID, error) {
	coll, client, ctx, cancel, collErr := baseDB.OpenCollection("employee")
	defer cancel()
	defer client.Disconnect(ctx)

	if collErr != nil {
		return nil, collErr
	}

	mod := mongo.IndexModel{
		Keys: bson.M{
			"empId": 1,
		},
		Options: options.Index().SetUnique(true),
	}

	_, indxErr := coll.Indexes().CreateOne(ctx, mod)
	if indxErr != nil {
		return nil, indxErr
	}

	var id primitive.ObjectID
	curDate := primitive.NewDateTimeFromTime(time.Now())
	req.CreatedDate = &curDate
	req.UpdatedDate = &curDate
	result, insertErr := coll.InsertOne(ctx, req)
	if insertErr != nil {
		return nil, insertErr
	}

	if result != nil {
		id = result.InsertedID.(primitive.ObjectID)
	}

	return &id, nil
}

func UpdateEmployeeRepo(req models.EmployeeModel) (*int64, error) {
	coll, client, ctx, cancel, collErr := baseDB.OpenCollection("employee")
	defer cancel()
	defer client.Disconnect(ctx)

	if collErr != nil {
		return nil, collErr
	}

	var updateCount int64
	param := bson.M{}
	param["empName"] = req.EmpName
	param["empPosition"] = req.EmpPosition
	param["empGender"] = req.EmpGender
	param["empEmail"] = req.EmpEmail
	param["empCompany"] = req.EmpCompany
	param["empStatus"] = req.EmpStatus
	param["updatedBy"] = req.UpdatedBy
	param["updatedDate"] = primitive.NewDateTimeFromTime(time.Now())
	updateResult, updateErr := coll.UpdateOne(ctx, bson.M{"_id": req.Id}, bson.M{"$set": param})
	if updateErr != nil {
		return nil, updateErr
	}

	if updateResult != nil {
		updateCount = updateResult.ModifiedCount
	}
	return &updateCount, nil
}

func GetEmployeeRepo(req st.EmployeeReq) (*models.EmployeeModel, error) {
	coll, client, ctx, cancel, collErr := baseDB.OpenCollection("employee")
	defer cancel()
	defer client.Disconnect(ctx)
	if collErr != nil {
		return nil, collErr
	}

	param := bson.M{}
	if req.Id != nil {
		objId, _ := primitive.ObjectIDFromHex(*req.Id)
		param["_id"] = objId
	}
	if req.EmpId != nil {
		param["empId"] = *req.EmpId
	}
	if req.EmpStatus != nil {
		param["EmpStatus"] = *req.EmpStatus
	}

	var obj models.EmployeeModel
	if len(param) > 0 {
		resultErr := coll.FindOne(ctx, param).Decode(&obj)
		if resultErr != nil {
			return nil, resultErr
		}
	}

	return &obj, nil
}

func ListEmployeeRepo(param st.EmployeeReq) ([]models.EmployeeModel, *int64, *int64, error) {
	coll, client, ctx, cancel, collErr := baseDB.OpenCollection("employee")
	defer cancel()
	defer client.Disconnect(ctx)
	if collErr != nil {
		return nil, nil, nil, collErr
	}

	filter := bson.M{}
	if param.Keyword != nil && *param.Keyword != "" {
		var opts []bson.M
		opts = append(opts, bson.M{"empName": bson.M{"$regex": primitive.Regex{Pattern: *param.Keyword, Options: "i"}}})
		opts = append(opts, bson.M{"empPosition": bson.M{"$regex": primitive.Regex{Pattern: *param.Keyword, Options: "i"}}})
		opts = append(opts, bson.M{"empEmail": bson.M{"$regex": primitive.Regex{Pattern: *param.Keyword, Options: "i"}}})
		opts = append(opts, bson.M{"empCompany": bson.M{"$regex": primitive.Regex{Pattern: *param.Keyword, Options: "i"}}})
		filter["$or"] = opts
	}

	if param.EmpId != nil {
		filter["empId"] = *param.EmpId
	}

	if param.EmpStatus != nil {
		filter["empStatus"] = *param.EmpStatus
	}

	var limit int64 = 30
	var skip int64 = 0

	if param.PerPage != nil {
		limit = *param.PerPage
	}

	if param.Page != nil {
		skip = (*param.Page - 1) * limit
	}

	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)
	findOptions.SetSort(bson.M{"updatedDate": 1})

	totalResults, _ := coll.CountDocuments(ctx, filter)
	result, findErr := coll.Find(ctx, filter, findOptions)
	if findErr != nil {
		return nil, nil, nil, findErr
	}

	listData := []models.EmployeeModel{}
	if resultErr := result.All(ctx, &listData); resultErr != nil {
		return nil, nil, nil, resultErr
	}

	totalPages := int64(math.Ceil(float64(totalResults) / float64(limit)))

	return listData, &totalResults, &totalPages, nil
}

func DeleteEmployeeRepo(id *string) (*int64, error) {
	coll, client, ctx, cancel, collErr := baseDB.OpenCollection("employee")
	defer cancel()
	defer client.Disconnect(ctx)
	if collErr != nil {
		return nil, collErr
	}

	objectId, objectErr := primitive.ObjectIDFromHex(*id)
	if objectErr != nil {
		return nil, objectErr
	}

	var deleteCount int64
	deleteResult, deleteErr := coll.DeleteOne(ctx, bson.M{"_id": objectId})
	if deleteErr != nil {
		return nil, deleteErr
	}

	if deleteResult != nil {
		deleteCount = deleteResult.DeletedCount
	}
	return &deleteCount, nil
}
