package mongorm

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect establishes a connection to the MongoDB server and returns the client.
//
// Parameters:
//   - uri: The connection string for the MongoDB server.
//
// Returns:
//   - *mongo.Client: The MongoDB client instance if successful.
//   - error: An error if the connection or ping fails.
func Connect(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to MongoDB: %w", err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("unable to ping MongoDB: %w", err)
	}

	fmt.Println("Connected to MongoDB successfully")
	return client, nil
}

// CreateOne inserts a single document into the specified collection.
//
// Parameters:
//   - ctx: The context for the operation.
//   - db: The MongoDB database instance.
//   - collectionName: The name of the collection to insert the document into.
//   - doc: The document to insert.
//
// Returns:
//   - error: An error if the insertion fails.
func (m *Model) CreateOne(ctx context.Context, db *mongo.Database, collectionName string, doc interface{}) error {
	collection := db.Collection(collectionName)

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return fmt.Errorf("failed to insert document: %w", err)
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return fmt.Errorf("unexpected type for InsertedID: %T", result.InsertedID)
	}

	m.ID = id
	return nil
}

// CreateAll inserts multiple documents into the specified collection.
//
// Parameters:
//   - ctx: The context for the operation.
//   - db: The MongoDB database instance.
//   - collectionName: The name of the collection to insert the documents into.
//   - docs: A slice of documents to insert.
//
// Returns:
//   - error: An error if the insertion of any document fails.
func (m *Model) CreateAll(ctx context.Context, db *mongo.Database, collectionName string, docs []interface{}) error {
	collection := db.Collection(collectionName)
	
	if len(docs) == 0 {
		return nil // No documents to insert
	}

	for i, doc := range docs {
		// Set CreatedAt and UpdatedAt only for the first document
		if i == 0 {
			m.CreatedAt = time.Now()
			m.UpdatedAt = time.Now()
		}

		_, err := collection.InsertOne(ctx, doc)
		if err != nil {
			return fmt.Errorf("failed to insert document %d: %w", i, err)
		}
	}
	return nil
}

// ReadOne retrieves a single document from the specified collection based on the filter.
//
// Parameters:
//   - ctx: The context for the operation.
//   - db: The MongoDB database instance.
//   - collectionName: The name of the collection to retrieve the document from.
//   - filter: The filter criteria to match the document.
//   - result: A pointer to the variable to store the retrieved document.
//
// Returns:
//   - error: An error if the document cannot be found or decoded.
func (m *Model) ReadOne(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, result interface{}) error {
	collection := db.Collection(collectionName)

	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return fmt.Errorf("failed to find document: %w", err)
	}

	return nil
}

// ReadAll retrieves multiple documents from the specified collection based on the filter.
//
// Parameters:
//   - ctx: The context for the operation.
//   - db: The MongoDB database instance.
//   - collectionName: The name of the collection to retrieve the documents from.
//   - filter: The filter criteria to match the documents.
//   - results: A pointer to the slice to store the retrieved documents.
//
// Returns:
//   - error: An error if the documents cannot be found or decoded.
func (m *Model) ReadAll(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, results interface{}) error {
	collection := db.Collection(collectionName)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to find documents: %w", err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, results); err != nil {
		return fmt.Errorf("failed to decode documents: %w", err)
	}

	return nil
}

// Update modifies an existing document in the specified collection based on the filter and update parameters.
//
// Parameters:
//   - ctx: The context for the operation.
//   - db: The MongoDB database instance.
//   - collectionName: The name of the collection to update the document in.
//   - filter: The filter criteria to match the document to update.
//   - update: The update parameters to apply to the document.
//
// Returns:
//   - error: An error if the update fails.
func (m *Model) Update(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, update interface{}) error {
	collection := db.Collection(collectionName)

	m.UpdatedAt = time.Now()

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update document: %w", err)
	}

	return nil
}

// DeleteOne removes a single document from the specified collection based on the filter.
//
// Parameters:
//   - ctx: The context for the operation.
//   - db: The MongoDB database instance.
//   - collectionName: The name of the collection to delete the document from.
//   - filter: The filter criteria to match the document to delete.
//
// Returns:
//   - error: An error if the deletion fails.
func (m *Model) DeleteOne(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}) error {
	collection := db.Collection(collectionName)

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete document: %w", err)
	}

	return nil
}

// DeleteAll removes multiple documents from the specified collection based on the filter.
//
// Parameters:
//   - ctx: The context for the operation.
//   - db: The MongoDB database instance.
//   - collectionName: The name of the collection to delete the documents from.
//   - filter: The filter criteria to match the documents to delete.
//
// Returns:
//   - error: An error if the deletion of any document fails.
func (m *Model) DeleteAll(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}) error {
	collection := db.Collection(collectionName)

	_, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete documents: %w", err)
	}

	return nil
}
