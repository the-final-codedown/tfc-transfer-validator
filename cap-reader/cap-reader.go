package cap_reader

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"

	//cap "github.com/the-final-codedown/tfc-cap-updater/proto/tfc/cap/updater"

	"log"
	"net/http"
	"time"
)

type CapReader struct {
	collection *mongo.Collection
}

type CapReaderInterface interface {
	GetCap(id int64) (int64, error)
}

var client *mongo.Client

func InitializeReader(mongoAddress string) *CapReader {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoAddress))
	if err != nil {
		log.Panic(err)
	}
	capCollection := client.Database("tfc").Collection("cap")
	return &CapReader{collection: capCollection}
}

func DisconnectReader() {
	defer client.Disconnect(context.TODO())

}

func (repository *CapReader) GetCap(id string) (int32, error) {
	accountFilter := bson.D{{"_id", id}}
	result := struct {
		AccountID int64
		Value     int32
	}{}
	err := repository.collection.FindOne(context.Background(), accountFilter).Decode(&result)
	if err != nil {

		_ = CreateCap(id)
	}
	return result.Value, err
}

func CreateCap(id string) error {

	resp, err := http.Get("http://app:8080/accounts/" + id + "/cap")
	if err != nil {
		log.Println(err)
		return err
	}

	result := struct {
		Money int32
		AmountSlidingWindow     int32
	}{}
	s1, _:= ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(s1, &result); err != nil {
		log.Println(err)
	}
	//println(resp.Body.Read(result))
	println("Response body : " + fmt.Sprint(result.Money) + " " + fmt.Sprint(result.AmountSlidingWindow))

	/*_ = cap.CapDownscale{
		AccountID: id,
		Value:     0,
	}*/
	return err
}
