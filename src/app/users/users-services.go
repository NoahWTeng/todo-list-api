package users

import (
	"context"
	"fmt"
	"github.com/NoahWTeng/todo-api-go/src/infra/helpers/pagination"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (s *Database) Create(ctx context.Context, user *Model) (*Model, error) {
	collection := s.Methods.DB(ctx).Collection(s.Collection)

	if err := user.CreateValidation(); err != nil {
		return &Model{}, err
	}

	passwordHashing(&user.Password)

	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		return &Model{}, err
	}
	user.RawID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *Database) FindOne(ctx context.Context, id string, email string) (*Model, error) {
	collection := s.Methods.DB(ctx).Collection(s.Collection)
	var result Model

	if id != "" {
		ObjectId, _ := primitive.ObjectIDFromHex(id)
		err := collection.FindOne(ctx, bson.D{{"_id", ObjectId}}).Decode(&result)

		if err != nil {
			return &Model{}, err
		}
	}

	if email != "" {
		err := collection.FindOne(ctx, bson.D{{"email", email}}).Decode(&result)
		if err != nil {
			return &Model{}, err
		}
	}

	return &result, nil
}

func (s *Database) FindAll(ctx context.Context) *pagination.Pages {
	collection := s.Methods.DB(ctx).Collection(s.Collection)

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

func (s *Database) Update(ctx context.Context, user *Model, id string) (*Model, error) {
	collection := s.Methods.DB(ctx).Collection(s.Collection)

	if err := user.UpdateValidation(); err != nil {
		return &Model{}, err
	}

	var model Model
	ObjectId, _ := primitive.ObjectIDFromHex(id)

	_ = collection.FindOne(ctx, bson.D{{"_id", ObjectId}}).Decode(&model)

	model.UpdatedAt = time.Now()

	if user.Name != "" {
		model.Name = user.Name
	}

	if user.Email != "" {
		model.Email = user.Email
	}

	if user.Password != "" {
		passwordHashing(&user.Password)
		model.Password = user.Password
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

func (s *Database) SignIn(ctx context.Context, login *Login) (string, error) {
	collection := s.Methods.DB(ctx).Collection(s.Collection)

	var user Model

	if err := login.LoginValidation(); err != nil {
		return "", err
	}

	err := collection.FindOne(ctx, bson.D{{"email", login.Email}}).Decode(&user)

	if err != nil {
		return "", errors.New("Email or password is not exists!")
	}

	isAuth := comparePasswords(user.Password, login.Password)
	if !isAuth {
		return "", errors.New("Email or password is not exists!")
	}
	// Declare the expiration time of the token
	// here, we have kept it as 60 minutes
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, errToken := token.SignedString(JwtKey)
	if errToken != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", errors.New("Error generate JWT key!")
	}

	return tokenString, nil
}
