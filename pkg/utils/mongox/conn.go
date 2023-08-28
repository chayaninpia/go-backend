package mongox

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type DBService struct {
	client   *mongo.Client
	DBName   string
	Database *mongo.Database
}

type Session struct {
	mongo.Session
}

type SessionContext mongo.SessionContext

var (
	// Private variable of database service, Key is database name.
	dbServices = make(map[string]*DBService)
)

func ClearCache() {
	// Close connection
	for _, dbService := range dbServices {
		dbService.Close()
	}

	// Clear cache
	dbServices = make(map[string]*DBService)
}

func (service *DBService) getCacheKey() string {
	return service.DBName
}

func (service *DBService) setCache() {
	dbServices[service.getCacheKey()] = service
}

func (service *DBService) GetCache() *DBService {
	return dbServices[service.getCacheKey()]
}

func (service *DBService) Client() *mongo.Client {
	return service.client
}

// New mongo service
func (service *DBService) NewService() error {
	if service.DBName == "" {
		return fmt.Errorf("database name is required")
	}

	// Already connected
	if service.client != nil && service.Database != nil {
		return nil
	}

	if err := service.newConnection(service.DBName); err != nil {
		// Error handler
		return err
	}

	// Cache
	service.setCache()

	return nil
}

func SetMockDB(dbName string, client *mongo.Client, database *mongo.Database) {
	mockService := &DBService{
		client:   client,
		DBName:   dbName,
		Database: database,
	}

	dbServices[mockService.getCacheKey()] = mockService
}

func (service *DBService) FindDatas(collection *mongo.Collection, result interface{}, filter interface{}, opts ...*options.FindOptions) error {
	return service.FindDatasWithcontext(context.Background(), collection, result, filter, opts...)
}

func (service *DBService) FindDatasWithcontext(ctx context.Context, collection *mongo.Collection, result interface{}, filter interface{}, opts ...*options.FindOptions) error {
	cursor, err := collection.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}

	if err := cursor.All(ctx, result); err != nil {
		return err
	}

	return nil
}

func (service *DBService) CountDocuments(collection *mongo.Collection, filter interface{}) (int64, error) {
	return service.CountDocumentsWithcontext(context.Background(), collection, filter)
}

func (service *DBService) CountDocumentsWithcontext(ctx context.Context, collection *mongo.Collection, filter interface{}) (int64, error) {
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (service *DBService) FindFirstData(collection *mongo.Collection, result interface{}, filter interface{}, opts ...*options.FindOneOptions) error {
	return service.FindFirstDataWithContext(context.Background(), collection, result, filter, opts...)
}

func (service *DBService) FindFirstDataWithContext(ctx context.Context, collection *mongo.Collection, result interface{}, filter interface{}, opts ...*options.FindOneOptions) error {
	singleResult := collection.FindOne(ctx, filter, opts...)
	if err := singleResult.Err(); err != nil {
		return err
	}

	return singleResult.Decode(result)
}

func (service *DBService) FindLastData(collection *mongo.Collection, result interface{}) error {
	return service.FindLastDataWithContext(context.Background(), collection, result)
}

func (service *DBService) FindLastDataWithContext(ctx context.Context, collection *mongo.Collection, result interface{}) error {
	findOptions := options.FindOne()
	findOptions.SetSort(
		bson.D{
			primitive.E{
				Key:   "_id",
				Value: -1,
			},
		},
	)

	singleResult := collection.FindOne(ctx, bson.D{}, findOptions)
	if err := singleResult.Err(); err != nil {
		return err
	}

	return singleResult.Decode(result)
}

func (service *DBService) InsertOne(collection *mongo.Collection, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return service.InsertOneWithContext(context.Background(), collection, document, opts...)
}

func (service *DBService) InsertOneWithContext(ctx context.Context, collection *mongo.Collection, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	result, err := collection.InsertOne(ctx, document, opts...)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (service *DBService) InsertMany(collection *mongo.Collection, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return service.InsertManyWithContext(context.Background(), collection, documents, opts...)
}

func (service *DBService) InsertManyWithContext(ctx context.Context, collection *mongo.Collection, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	result, err := collection.InsertMany(ctx, documents, opts...)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (service *DBService) UpdateByID(collection *mongo.Collection, id interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return service.UpdateByIDWithContext(context.Background(), collection, id, update, opts...)
}

func (service *DBService) UpdateByIDWithContext(ctx context.Context, collection *mongo.Collection, id interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	objectToUpdate := bson.M{"$set": update}

	result, err := collection.UpdateByID(ctx, id, objectToUpdate, opts...)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (service *DBService) UpdateOneByStruct(collection *mongo.Collection, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return service.UpdateOneByStructWithContext(context.Background(), collection, filter, update, opts...)
}

func (service *DBService) UpdateOneByStructWithContext(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	objectToUpdate := bson.M{"$set": update}

	result, err := collection.UpdateOne(ctx, filter, objectToUpdate, opts...)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (service *DBService) UpdateOneByScript(collection *mongo.Collection, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return service.UpdateOneByScriptWithContext(context.Background(), collection, filter, update, opts...)
}

func (service *DBService) UpdateOneByScriptWithContext(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	result, err := collection.UpdateOne(ctx, filter, update, opts...)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (service *DBService) UpdateMany(collection *mongo.Collection, filter interface{}, document interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return service.UpdateManyWithContext(context.Background(), collection, filter, document, opts...)
}

func (service *DBService) UpdateManyWithContext(ctx context.Context, collection *mongo.Collection, filter interface{}, document interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	result, err := collection.UpdateMany(ctx, filter, document, opts...)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (service *DBService) UpdateManyByScriptWithContext(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	result, err := collection.UpdateMany(ctx, filter, update, opts...)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (service *DBService) DeleteOne(collection *mongo.Collection, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return service.DeleteOneWithContext(context.Background(), collection, filter, opts...)
}

func (service *DBService) DeleteOneWithContext(ctx context.Context, collection *mongo.Collection, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	result, err := collection.DeleteOne(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (service *DBService) DeleteMany(collection *mongo.Collection, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return service.DeleteManyWithContext(context.Background(), collection, filter, opts...)
}

func (service *DBService) DeleteManyWithContext(ctx context.Context, collection *mongo.Collection, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	result, err := collection.DeleteMany(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (service *DBService) Aggregate(collection *mongo.Collection, document interface{}, pipeline interface{}, opts ...*options.AggregateOptions) error {
	return service.AggregateWithContext(context.Background(), collection, document, pipeline, opts...)
}

func (service *DBService) AggregateWithContext(ctx context.Context, collection *mongo.Collection, document interface{}, pipeline interface{}, opts ...*options.AggregateOptions) error {
	cursor, err := collection.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return err
	}

	if err := cursor.All(ctx, document); err != nil {
		return err
	}

	return nil
}

func (service *DBService) Close() {
	if service != nil && service.client != nil {
		service.disconnect()
	}
}

func (service *DBService) newConnection(dbName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoUri := viper.GetString("mongodb.uri")

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongoUri),
	)
	if err != nil {
		log.Println("Error connection mongoDB :", err)

		return err
	}

	// Cache to struct
	service.client = client

	// Ping connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Println("Ping mongoDB server error :", err)

		return err
	}

	// Choose database
	service.Database = service.client.Database(dbName)

	log.Println("Mongo connected.")

	return nil
}

func (service *DBService) disconnect() {
	if err := service.client.Disconnect(context.TODO()); err != nil {
		log.Println("Error disconnect mongoDB :", err)
	}

	log.Println("Mongo disconnected.")
}

func IsErrNoDocuments(err error) bool {
	return err == mongo.ErrNoDocuments
}

func (service *DBService) StartSession() (Session, error) {
	option := &options.SessionOptions{
		DefaultReadConcern:    readconcern.Snapshot(),
		DefaultWriteConcern:   writeconcern.New(writeconcern.WMajority()),
		DefaultReadPreference: readpref.Primary(),
	}

	session, err := service.Client().StartSession(option)
	if err != nil {
		return Session{}, err
	}

	return Session{session}, nil
}

func (session Session) DoTransaction(ctx context.Context, fn func(SessionContext) (interface{}, error)) (interface{}, error) {
	callback := func(sctx mongo.SessionContext) (interface{}, error) {
		return fn(sctx)
	}
	return session.WithTransaction(ctx, callback)
}

func (session Session) Close(ctx context.Context) {
	session.EndSession(ctx)
}

func GetPushOrUpdateToArrayUpdateInfo(arrayField, keyCompareField string, keyCompareValue, updateData any, createData bson.M, anotherFieldValues ...bson.M) []bson.M {
	updateCmd := []bson.M{
		{
			"$set": bson.M{
				arrayField: bson.M{
					"$cond": []any{
						bson.M{
							"$not": []string{
								fmt.Sprintf("$%s", arrayField),
							},
						}, // If field not exists
						[]any{
							createData,
						}, // Then create data in array
						bson.M{
							"$cond": []bson.M{
								{
									"$in": []any{
										keyCompareValue,
										fmt.Sprintf("$%s.%s", arrayField, keyCompareField),
									},
								},
								{
									"$map": bson.M{
										"input": fmt.Sprintf("$%s", arrayField),
										"in": bson.M{
											"$cond": []any{
												bson.M{
													"$eq": []any{
														fmt.Sprintf("$$this.%s", keyCompareField),
														keyCompareValue,
													},
												},
												updateData,
												"$$this",
											},
										},
									},
								},
								{
									"$concatArrays": []any{
										fmt.Sprintf("$%s", arrayField),
										[]bson.M{
											createData,
										},
									},
								},
							},
						}, // Else push or update to array
					},
				},
			},
		},
	}

	for _, anotherFieldValue := range anotherFieldValues {
		for field, value := range anotherFieldValue {
			updateCmd[0]["$set"].(bson.M)[field] = value
		}
	}

	return updateCmd
}
