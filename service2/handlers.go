package service2

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"os"
	"practice10/gen/proto"
)

type File struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
}

var logFile, _ = os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

// Установка вывода логов в файл и консоль
var logger = log.New(io.MultiWriter(os.Stdout, logFile), "", log.LstdFlags)

func GetFiles(client *mongo.Client, name string) ([]*proto.File, error) {
	// Получаем все файлы из MongoDB
	ctx := context.Background()
	collection := client.Database(name).Collection(name)
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	// Конвертируем в слайс структур File и кодируем в формат JSON
	var files []File
	for cur.Next(ctx) {
		var file File
		err := cur.Decode(&file)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	err = cur.Err()
	if err != nil {
		return nil, err
	}
	var convertedFiles []*proto.File
	for _, file := range files {
		convertedID := file.ID.Hex()
		file1 := proto.File{Id: convertedID, Name: file.Name, Description: file.Description}
		convertedFiles = append(convertedFiles, &file1)
	}

	return convertedFiles, nil
}

func GetFileInfo(client *mongo.Client, name string, id string) *proto.File {

	convertedId, _ := primitive.ObjectIDFromHex(id)

	// Get file from MongoDB
	ctx := context.Background()
	collection := client.Database(name).Collection(name)
	filter := bson.M{"_id": convertedId}
	var file File
	err := collection.FindOne(ctx, filter).Decode(&file)
	if err != nil {
		log.Println(err)

	}
	var file1 *proto.File
	file1 = &proto.File{Id: id, Name: file.Name, Description: file.Description}
	return file1
}

func UploadFile(client *mongo.Client, name string, filename string, description string) {

	// Create file document in MongoDB
	ctx := context.Background()
	collection := client.Database(name).Collection(name)
	_, err := collection.InsertOne(ctx, bson.M{
		"name":        filename,
		"description": description,
	})
	if err != nil {
		log.Println(err)
	}

	// Upload file to GridFS
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Println(err)

	}
	logger.Println("File was created")
}

func UpdateFile(client *mongo.Client, id string, filename string, description string, name string) {
	convertedId, _ := primitive.ObjectIDFromHex(id)
	// Update file document in MongoDB
	ctx := context.Background()
	collection := client.Database(name).Collection(name)
	filter := bson.M{"_id": convertedId}
	update := bson.M{
		"$set": bson.M{
			"name":        filename,
			"description": description,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(err)

	}
	logger.Println("File was updated")
}

func DeleteFile(client *mongo.Client, id string, name string) {
	convertedId, _ := primitive.ObjectIDFromHex(id)

	// Delete file document from MongoDB
	ctx := context.Background()
	collection := client.Database(name).Collection(name)
	filter := bson.M{"_id": convertedId}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)

	}
	logger.Println("File was deleted")
}
