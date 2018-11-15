package warning

import (
	"encoding/json"
	"log"

	"github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
	"github.com/TerrexTech/uuuid"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

// Insert handles "insert" events.
func Insert(collection *mongo.Collection, event *model.Event) *model.Document {
	warn := &Warning{}
	err := json.Unmarshal(event.Data, warn)
	if err != nil {
		err = errors.Wrap(err, "Insert: Error while unmarshalling Event-data")
		log.Println(err)
		return &model.Document{
			AggregateID:   event.AggregateID,
			CorrelationID: event.CorrelationID,
			Error:         err.Error(),
			ErrorCode:     InternalError,
			EventAction:   event.EventAction,
			ServiceAction: event.ServiceAction,
			UUID:          event.UUID,
		}
	}

	if warn.ItemID == (uuuid.UUID{}) {
		err = errors.New("missing ItemID")
		err = errors.Wrap(err, "Insert")
		log.Println(err)
		return &model.Document{
			AggregateID:   event.AggregateID,
			CorrelationID: event.CorrelationID,
			Error:         err.Error(),
			ErrorCode:     InternalError,
			EventAction:   event.EventAction,
			ServiceAction: event.ServiceAction,
			UUID:          event.UUID,
		}
	}

	insertResult, err := collection.InsertOne(warn)
	if err != nil {
		err = errors.Wrap(err, "Insert: Error Inserting Warning into Mongo")
		log.Println(err)
		return &model.Document{
			AggregateID:   event.AggregateID,
			CorrelationID: event.CorrelationID,
			Error:         err.Error(),
			ErrorCode:     DatabaseError,
			EventAction:   event.EventAction,
			ServiceAction: event.ServiceAction,
			UUID:          event.UUID,
		}
	}
	insertedID, assertOK := insertResult.InsertedID.(objectid.ObjectID)
	if !assertOK {
		err = errors.New("error asserting InsertedID from InsertResult to ObjectID")
		err = errors.Wrap(err, "Insert")
		log.Println(err)
		return &model.Document{
			AggregateID:   event.AggregateID,
			CorrelationID: event.CorrelationID,
			Error:         err.Error(),
			ErrorCode:     InternalError,
			EventAction:   event.EventAction,
			ServiceAction: event.ServiceAction,
			UUID:          event.UUID,
		}
	}

	warn.ID = insertedID
	result, err := json.Marshal(warn)
	if err != nil {
		err = errors.Wrap(err, "Insert: Error marshalling Warning Insert-result")
		log.Println(err)
		return &model.Document{
			AggregateID:   event.AggregateID,
			CorrelationID: event.CorrelationID,
			Error:         err.Error(),
			ErrorCode:     InternalError,
			EventAction:   event.EventAction,
			ServiceAction: event.ServiceAction,
			UUID:          event.UUID,
		}
	}

	return &model.Document{
		AggregateID:   event.AggregateID,
		CorrelationID: event.CorrelationID,
		Result:        result,
		EventAction:   event.EventAction,
		ServiceAction: event.ServiceAction,
		UUID:          event.UUID,
	}
}
