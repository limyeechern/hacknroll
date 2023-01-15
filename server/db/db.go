package db

import (
	"context"
	"os"
	"server/pkg/models"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	log "github.com/sirupsen/logrus"
)

var database *mongo.Database

func InitDB(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}
	log.Info("environment variables loaded")
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("URI")))
	if err != nil {
		log.Fatal("Fatal error when instantiating mongo client", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 1000 * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Fatal error when client connects to context after instantiating mongo client", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil{
		log.Fatal("Fatal error when mongo client pings", err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil{
		log.Fatal("Fatal error when enumerating databases", err)
	} 
	database = client.Database("mycommunity")
	log.Info("databases are",databases)
	log.Info("no errors in instantiating database clients")
	// defer client.Disconnect(ctx)
}


func NewPost(p models.NewPost) error{

	res, err := database.Collection("posts").InsertOne(context.TODO(), p)

	if err != nil {
		log.Error("Insert Error : ", err)
		return err
	}

	log.Info("res is ", res)

	return nil
}

func GetFeed(page int) ([]models.Post, models.Pagination){
	var result []models.Post

	// Set the number of documents to display per page
	itemsPerPage := 3

	// Calculate the number of documents to skip
	skip := itemsPerPage * (page - 1)

	// Set up the options for the paginated query
	findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(itemsPerPage)).SetSort((bson.M{"_id" : -1}))

	cursor, err := database.Collection("posts").Find(context.TODO(), bson.M{}, findOptions)
	count, err := database.Collection("posts").CountDocuments(context.TODO(), bson.M{})


	if err != nil {
		log.Error("Error while getting profile by id: ", err)
	}

	for cursor.Next(context.TODO()) {
		var post models.Post
		err := cursor.Decode(&post)
		if err != nil {
			log.Error("Error from cursor.Decode while getting profile by id: ", err)
		}
		result = append(result, post)
	}

	pagination := models.Pagination{
		Page: page,
		PerPage: itemsPerPage,
		LastPage: int(count) / itemsPerPage,
	}

	return result, pagination
}


// func UpdateProfile(p models.Profile, id string) error{

// 	objectId , err := primitive.ObjectIDFromHex(id)

// 	if err != nil {
// 		log.Error("Object id from hex error : ", err)
// 		return err
// 	}

// 	filter := bson.M{"_id": objectId}
// 	update := bson.M{"$set": bson.M{
// 		"studies" : p.Studies,
// 		"header" : p.Header,
// 		"introduction" : p.Introduction,
// 		"profiletags.interests" : p.ProfileTags.Interests,
// 		"profiletags.modules" : p.ProfileTags.Modules,
// 		"profilepretexts" : p.ProfilePreTexts,
// 	}}

// 	_, err = database.Collection("profile").UpdateOne(context.TODO(), filter, update)

// 	if err != nil {
// 		log.Error("Insert Error : ", err)
// 		return err
// 	}

// 	log.Info("Profile data updated!!!")
// 	return nil
// }

// func NewProfile(p models.NewProfile) (primitive.ObjectID, error){
// 	res, err := database.Collection("profile").InsertOne(context.TODO(), p)
// 	if err != nil {
// 		log.Error("Insert Error : ", err)
// 	}

// 	log.Info("New profile added into mongodb", p.LinkedInId)
// 	log.Info("res is", res)
// 	// log.Info("profile's id is "fe res.InsertedID.(primitive.ObjectID).String())

// 	r := (res.InsertedID.(primitive.ObjectID))

// 	database.Collection("history").InsertOne(context.TODO(), 
// 												bson.M{"_id" : res.InsertedID.(primitive.ObjectID),
// 														"history" : []string{} })

// 	return r, err
// }

// func GetProfileById(id string) models.Profile{
// 	var result models.Profile
// 	objectId , err := primitive.ObjectIDFromHex(id)
// 	if err != nil{
// 		log.Error("Id is ", id)
// 		log.Error("Invalid id while getting profile by id", err)
// 	}
// 	err = database.Collection("profile").FindOne(context.TODO(), bson.M{"_id" : objectId}).Decode(&result)
// 	if err != nil {
// 		log.Error("Error while getting profile by id: ", err)
// 	}
// 	return result
// }

// func CheckProfileByLinkedInId(id string) (models.Profile, error){
// 	var p models.Profile
// 	err := database.Collection("profile").FindOne(context.TODO(), bson.M{"profilelinkedin.linkedinid": id}).Decode(&p)
// 	if err != nil{
// 		return p, err
// 	}
// 	return p, nil
// }

// func AddNewTopic(topic models.NewTopic) (primitive.ObjectID, error) {
// 	res, err := database.Collection("topic").InsertOne(context.TODO(), topic)
// 	if err != nil {
// 		log.Error("Insert Error : ", err)
// 	}

// 	r := (res.InsertedID.(primitive.ObjectID))

// 	return r, err
// }

// func GetTopicsInConnect(page int, mongoId string) ([] models.Topic, models.Pagination){
// 	topicColl := database.Collection("topic")
// 	var topics [] models.Topic

// 	// Set the number of documents to display per page
// 	itemsPerPage := 3

// 	// Calculate the number of documents to skip
// 	skip := itemsPerPage * (page - 1)

// 	// Set up the options for the paginated query
// 	findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(itemsPerPage)).SetSort((bson.M{"_id" : -1}))

// 	objectId, err := primitive.ObjectIDFromHex(mongoId)
// 	if err != nil {
// 		log.Error("cannot find objectId")
// 	}

// 	var history models.History

// 	database.Collection("history").FindOne(context.TODO(), bson.M{"_id" : objectId}).Decode(&history)

// 	history.History = append(history.History, objectId)
	
// 	cursor, err := topicColl.Find(context.TODO(), bson.M{"participants" : bson.M{ "$nin" : []primitive.ObjectID{objectId}}, "_id" : bson.M{"$nin": history.History}}, findOptions)
	
// 	if err != nil {
// 		log.Error("Error while finding data from profiles collection while getting profile in connect page: ", err)
// 	}

// 	count, err := topicColl.CountDocuments(context.TODO(), bson.M{})

// 	if err != nil {
// 		log.Error("Error while finding data from profiles collection while counting profiles in connect page: ", err)
// 	}

// 	for cursor.Next(context.TODO()) {
// 		var topic models.Topic
// 		err := cursor.Decode(&topic)
// 		if err != nil {
// 			log.Error("Error from cursor.Decode while getting profile by id: ", err)
// 		}

// 		for _, profileId := range topic.Participants {
// 			var profile models.Profile
// 			_ = database.Collection("profile").FindOne(context.TODO(), bson.M{"_id" : profileId}).Decode(&profile)
// 			topic.ParticipantsData = append(topic.ParticipantsData, profile.ProfileLinkedIn)
// 		}

// 		topics = append(topics, topic)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		log.Error("Error from cursor while getting profile by id: ", err)
// 	}
	
// 	cursor.Close(context.TODO())

// 	pagination := models.Pagination{
// 		Page: page,
// 		PerPage: itemsPerPage,
// 		LastPage: int(count) / itemsPerPage,
// 	}

// 	return topics, pagination
// }

// func AddInteraction(interaction models.Interaction) error {
	
// 	test := database.Collection("interaction").FindOne(context.TODO(), bson.M{"self" : interaction.OtherUserId, "other_user_id" : interaction.Self})
// 	//Check if a vice versa already exists in database

// 	var existingInteraction models.Interaction
// 	test.Decode(&existingInteraction)

// 	if existingInteraction.TopicId.IsZero() { //if its zero, ie, a vice versa does not exist in database

// 		_, err := database.Collection("interaction").InsertOne(context.TODO(), interaction)
// 		database.Collection("history").UpdateOne(context.TODO(), bson.M{"_id" : interaction.Self}, bson.M{"$addToSet" : bson.M{"history" : interaction.TopicId}})
// 		//add in the current interaction
// 		if err != nil{
// 			return err
// 		}
// 	} else {//else if not zero, ie, a vice versa exists in database
// 		var inter models.Interaction
// 		database.Collection("interaction").FindOne(context.TODO(), bson.M{"self" : interaction.OtherUserId, "other_user_id" : interaction.Self}).Decode(&inter)

// 		var otherTopic models.Topic
// 		var thisTopic models.Topic
// 		database.Collection("topic").FindOne(context.TODO(), bson.M{"_id" : inter.TopicId}).Decode(&otherTopic)
// 		database.Collection("topic").FindOne(context.TODO(), bson.M{"_id" : interaction.TopicId}).Decode(&thisTopic)

// 		database.Collection("chat").InsertOne(context.TODO(), models.NewChat{
// 			IsOpened: false,
// 			Participants: [] primitive.ObjectID{interaction.Self, interaction.OtherUserId},
// 			UnopenedSettings: []models.UnopenedSettings{otherTopic.UnopenedSettings, thisTopic.UnopenedSettings},
// 			DisplayTitle: map[primitive.ObjectID]string{interaction.Self : thisTopic.UnopenedSettings.Title, interaction.OtherUserId : otherTopic.UnopenedSettings.Title },
// 			Expiry: time.Now().AddDate(0,0,14),
// 		})

// 		_, err := database.Collection("interaction").DeleteMany(context.TODO(), bson.M{"self" : interaction.OtherUserId, "other_user_id" : interaction.Self})
// 		// delete all (may be multiple) vice versa interactions as a match has already occurred
// 		if err != nil {
// 			return err
// 		}
// 		database.Collection("history").UpdateOne(context.TODO(), bson.M{"_id" : interaction.Self}, bson.M{"$addToSet" : bson.M{"history" : interaction.TopicId}})
// 		// prevent user from seeing the same topics user has liked
// 	}
// 	return nil
// }

// func GetAllChats(id string) []models.ReducedChat {
// 	objectId , err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		log.Error("error while getting objectId is ", err)
// 	}

// 	cursor, err := database.Collection("chat").Find(context.TODO(), bson.M{"participants" : bson.M{ "$in" : []primitive.ObjectID{objectId}}})
	
// 	if err != nil {
// 		log.Error("Error while finding chats from chat collection ", err)
// 	}

// 	var allChats []models.ReducedChat

// 	for cursor.Next(context.TODO()) {
// 		var chat models.ReducedChat
// 		err := cursor.Decode(&chat)
// 		if err != nil {
// 			log.Error("Error from cursor.Decode while getting profile by id: ", err)
// 		}

// 		allChats = append(allChats, chat)
// 	}

// 	log.Info("all chats are ", allChats)
// 	return allChats
// }

// // 
// // 

// func GetUser(userId string) bool {
// 	userIdHex, err := primitive.ObjectIDFromHex(userId)
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}

// 	filter := bson.D{{Key: "_id", Value: userIdHex}}
// 	err2 := database.Collection("profile").FindOne(context.Background(), filter)
// 	if err2 != nil {
// 		if err2.Err() == mongo.ErrNoDocuments {
// 			fmt.Println("No user")
// 			return false
// 		}
// 	}
// 	fmt.Println(userId)
// 	return true
// }

// func SendMessage(req models.SendMessageRes) {

// 	data := models.SendMessageDB{
// 		UserId:      req.UserId,
// 		RoomId:      req.RoomId,
// 		Content:     req.Content,
// 		PrevMessage: "",
// 		ContentType: req.ContentType,
// 		TimeStamp:   req.Timestamp,
// 		MessageId:   req.MessageId,
// 	}

// 	_, err := database.Collection("message").InsertOne(context.Background(), data)
// 	if err != nil {
// 		fmt.Println("Insert Error : ", err)
// 	}

// 	roomId, _ := primitive.ObjectIDFromHex(req.RoomId.Hex())
// 	filter := bson.D{{Key: "_id", Value: roomId}}
// 	update := bson.D{{Key: "$set", Value: bson.D{{Key: "latestmessageid", Value: req.MessageId}, {Key: "latestmessage", Value: req.Content}}}}

// 	val, err := database.Collection("chat").UpdateOne(context.Background(), filter, update)
// 	fmt.Println(val, err)
// 	fmt.Println("Message Sent")

// }

// func ReplyMessage(req models.ReplyMessageRes) {

// 	data := models.SendMessageDB{
// 		UserId:      req.UserId,
// 		RoomId:      req.RoomId,
// 		Content:     req.Content,
// 		ContentType: req.ContentType,
// 		TimeStamp:   req.Timestamp,
// 		PrevMessage: req.PrevMessage,
// 		MessageId:   req.MessageId,
// 	}

// 	roomId, _ := primitive.ObjectIDFromHex(req.RoomId.Hex())
// 	filter := bson.D{{Key: "_id", Value: roomId}}
// 	update := bson.D{{Key: "$set", Value: bson.D{{Key: "latestmessageid", Value: req.MessageId}, {Key: "latestmessage", Value: req.Content}}}}

// 	_, err := database.Collection("message").InsertOne(context.Background(), data)
// 	if err != nil {
// 		fmt.Println("Insert Error : ", err)
// 	}

// 	val, err := database.Collection("chat").UpdateOne(context.Background(), filter, update)
// 	fmt.Println(val, err)
// 	fmt.Println("Message Sent")

// }

// func DeleteMessage(req models.DeleteMessageReq) {

// 	fmt.Println(req.MessageId)
// 	filter := bson.D{{Key: "messageid", Value: req.MessageId}}
// 	val, err := database.Collection("messages").DeleteOne(context.Background(), filter)
// 	if err != nil {
// 		fmt.Println("Delete Message Error", err)
// 	}
// 	fmt.Println(val)

// 	filter2 := bson.D{{Key: "_id", Value: req.RoomId}, {Key: "latestmessageid", Value: req.MessageId}}
// 	update := bson.D{{Key: "$set", Value: bson.D{{Key: "latestmessageid", Value: ""}, {Key: "latestmessage", Value: ""}}}}
// 	val2, err := database.Collection("chats").UpdateOne(context.Background(), filter2, update)
// 	if err != nil {
// 		fmt.Println("Delete Message Error", err)
// 	}

// 	fmt.Println(val2, err)

// }

// func UpdateMessage(req models.UpdateMessageReq) {

// 	filter := bson.D{{Key: "messageid", Value: req.MessageId}}
// 	update := bson.D{{Key: "$set", Value: bson.D{{Key: "content", Value: req.Content}, {Key: "contenttype", Value: req.ContentType}}}}
// 	_, err := database.Collection("messages").UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		fmt.Println("Delete Message Error", err)
// 	}

// 	filter2 := bson.D{{Key: "_id", Value: req.RoomId}, {Key: "latestmessageid", Value: req.MessageId}}
// 	update2 := bson.D{{Key: "$set", Value: bson.D{{Key: "latestmessageid", Value: req.MessageId}, {Key: "latestMessage", Value: req.Content}}}}
// 	val, err := database.Collection("chats").UpdateOne(context.Background(), filter2, update2)
// 	if err != nil {
// 		fmt.Println("Update Message Error", err)
// 	}

// 	fmt.Println(val, err)

// }