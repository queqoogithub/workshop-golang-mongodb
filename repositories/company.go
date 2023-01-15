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

func InsertCompanyRepo(req models.CompanyModel) (*primitive.ObjectID, error) {
	coll, client, ctx, cancel, collErr := baseDB.OpenCollection("company")
	defer cancel()
	defer client.Disconnect(ctx)

	if collErr != nil {
		return nil, collErr
	}

	mod := mongo.IndexModel{ // ทำ index เพื่อให้ access data ได้ไวขึ้น
		Keys: bson.M{
			"comId": 1,
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

func UpdateCompanyRepo(req models.CompanyModel) (*int64, error) {
	coll, client, ctx, cancel, collErr := baseDB.OpenCollection("company")
	defer cancel()
	defer client.Disconnect(ctx)

	if collErr != nil {
		return nil, collErr
	}

	var updateCount int64
	param := bson.M{}
	param["comName"] = req.ComName
	param["comStatus"] = req.ComStatus
	param["empAmount"] = req.EmpAmount
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

func GetCompanyRepo(req st.CompanyReq) (*models.CompanyModel, error) {
	coll, client, ctx, cancel, collErr := baseDB.OpenCollection("company")
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
	if req.ComId != nil {
		param["comId"] = *req.ComId
	}
	if req.ComStatus != nil {
		param["EmpStatus"] = *req.ComStatus
	}

	var obj models.CompanyModel
	if len(param) > 0 {
		resultErr := coll.FindOne(ctx, param).Decode(&obj)
		if resultErr != nil {
			return nil, resultErr
		}
	}

	return &obj, nil
}

func ListCompanyRepo(param st.CompanyReq) ([]models.CompanyModel, *int64, *int64, error) {
	coll, client, ctx, cancel, collErr := baseDB.OpenCollection("company")
	defer cancel()
	defer client.Disconnect(ctx)
	if collErr != nil {
		return nil, nil, nil, collErr
	}

	filter := bson.M{}
	if param.Keyword != nil && *param.Keyword != "" {
		var opts []bson.M
		opts = append(opts, bson.M{"comName": bson.M{"$regex": primitive.Regex{Pattern: *param.Keyword, Options: "i"}}})
		opts = append(opts, bson.M{"comStatus": bson.M{"$regex": primitive.Regex{Pattern: *param.Keyword, Options: "i"}}})
		opts = append(opts, bson.M{"empAmount": bson.M{"$regex": primitive.Regex{Pattern: *param.Keyword, Options: "i"}}})
		filter["$or"] = opts
	}

	if param.ComId != nil {
		filter["comId"] = *param.ComId
	}

	if param.ComStatus != nil {
		filter["empStatus"] = *param.ComStatus
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

	listData := []models.CompanyModel{}
	if resultErr := result.All(ctx, &listData); resultErr != nil {
		return nil, nil, nil, resultErr
	}

	totalPages := int64(math.Ceil(float64(totalResults) / float64(limit)))

	return listData, &totalResults, &totalPages, nil
}

func DeleteCompanyRepo(id *string) (*int64, error) {
	coll, client, ctx, cancel, collErr := baseDB.OpenCollection("company")
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
