package tasks

import (
	"context"
	"fmt"
	"github.com/NoahWTeng/todo-api-go/src/infra/helpers/pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (s *Database) FindAll(ctx context.Context) *pagination.Pages {
	collection := s.Methods.DB(ctx).Collection("tasks")

	var results []*Model

	withPagination, _ := ctx.Value("pagination").(*pagination.Pages)

	count, errCount := collection.CountDocuments(ctx, bson.D{})

	if errCount != nil {
		fmt.Println(errCount)
	}

	pagination.Update(withPagination, int(count))

	opts := options.Find().SetSort(bson.D{{"createdAt", 1}}).SetSkip(int64(withPagination.Page) - 1).SetLimit(int64(withPagination.Limit))
	cursor, errFind := collection.Find(ctx, bson.D{}, opts)

	if errFind != nil {
		withPagination.Items = []*Model{}
		return withPagination
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var elem Model
		errDecode := cursor.Decode(&elem)
		if errDecode != nil {
			withPagination.Items = []*Model{}
			return withPagination
		}
		results = append(results, &elem)
	}

	withPagination.Items = results
	return withPagination
}

func (s *Database) Create(ctx context.Context, task *Model) (*Model, error) {
	collection := s.Methods.DB(ctx).Collection(s.Collection)

	if err := task.Validation(); err != nil {
		return &Model{}, err
	}

	result, err := collection.InsertOne(ctx, task)

	if err != nil {
		return &Model{}, err
	}
	task.RawID = result.InsertedID.(primitive.ObjectID)
	return task, nil
}

func (s *Database) FindOne(ctx context.Context, id string) (*Model, error) {
	collection := s.Methods.DB(ctx).Collection(s.Collection)
	var result Model

	ObjectId, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(ctx, bson.D{{"_id", ObjectId}}).Decode(&result)

	if err != nil {
		return &Model{}, err
	}

	return &result, nil
}

func (s *Database) Update(ctx context.Context, task *Model, id string) (*Model, error) {
	collection := s.Methods.DB(ctx).Collection(s.Collection)

	if err := task.Validation(); err != nil {
		return &Model{}, err
	}

	var model Model
	ObjectId, _ := primitive.ObjectIDFromHex(id)

	_ = collection.FindOne(ctx, bson.D{{"_id", ObjectId}}).Decode(&model)

	model.UpdatedAt = time.Now()

	if task.Title != "" {
		model.Title = task.Title
	}

	if task.Status != "" {
		model.Status = task.Status
	}

	if task.Comment != "" {
		model.Comment = task.Comment
	}

	_, err := collection.UpdateOne(ctx, bson.M{"_id": ObjectId}, bson.D{{"$set", &model}})

	if err != nil {
		return &Model{}, err
	}

	return &model, nil
}

func (s *Database) Delete(ctx context.Context, id string) (int64, error) {
	collection := s.Methods.DB(ctx).Collection(s.Collection)
	ObjectId, _ := primitive.ObjectIDFromHex(id)
	result, err := collection.DeleteOne(ctx, bson.D{{"_id", ObjectId}})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
