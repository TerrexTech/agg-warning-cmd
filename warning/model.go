package warning

import (
	"encoding/json"

	util "github.com/TerrexTech/go-commonutils/commonutil"

	"github.com/TerrexTech/uuuid"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

// AggregateID is the global AggregateID for Inventory Aggregate.
const AggregateID int8 = 18

type Warning struct {
	ID            objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	WarningID     uuuid.UUID        `bson:"warningID,omitempty" json:"warningID,omitempty"`
	ItemID        uuuid.UUID        `bson:"itemID,omitempty" json:"itemID,omitempty"`
	SKU           string            `bson:"sku,omitempty" json:"sku,omitempty"`
	Name          string            `bson:"name,omitempty" json:"name,omitempty"`
	SoldWeight    float64           `bson:"soldWeight,omitempty" json:"soldWeight,omitempty"`
	TotalWeight   float64           `bson:"totalWeight,omitempty" json:"totalWeight,omitempty"`
	UnsoldWeight  float64           `bson:"unsoldWeight,omitempty" json:"unsoldWeight,omitempty"`
	Lot           string            `bson:"lot,omitempty" json:"lot,omitempty"`
	WarningActive bool              `bson:"warningActive,omitempty" json:"warningActive,omitempty"`
	Timestamp     int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Ethylene      float64           `bson:"ethylene,omitempty" json:"ethylene,omitempty"`
	CarbonDioxide float64           `bson:"carbonDioxide,omitempty" json:"carbonDioxide,omitempty"`
	ProjectedDate int64             `bson:"projectedDate,omitempty" json:"projectedDate,omitempty"`
	Status        string            `bson:"status,omitempty" json:"status,omitempty"`
}

func (i Warning) MarshalBSON() ([]byte, error) {
	in := map[string]interface{}{
		"warningID":     i.WarningID.String(),
		"itemID":        i.ItemID.String(),
		"sku":           i.SKU,
		"name":          i.Name,
		"soldWeight":    i.SoldWeight,
		"totalWeight":   i.TotalWeight,
		"timestamp":     i.Timestamp,
		"unsoldWeight":  i.UnsoldWeight,
		"warningActive": i.WarningActive,
		"lot":           i.Lot,
		"ethylene":      i.Ethylene,
		"carbonDioxide": i.CarbonDioxide,
		"projectedDate": i.ProjectedDate,
		"status": i.Status,
	}

	if i.ID != objectid.NilObjectID {
		in["_id"] = i.ID
	}
	return bson.Marshal(in)
}

// MarshalJSON returns bytes of JSON-type.
func (i Warning) MarshalJSON() ([]byte, error) {
	in := map[string]interface{}{
		"warningID":     i.WarningID.String(),
		"itemID":        i.ItemID.String(),
		"sku":           i.SKU,
		"name":          i.Name,
		"soldWeight":    i.SoldWeight,
		"totalWeight":   i.TotalWeight,
		"timestamp":     i.Timestamp,
		"unsoldWeight":  i.UnsoldWeight,
		"warningActive": i.WarningActive,
		"lot":           i.Lot,
		"ethylene":      i.Ethylene,
		"carbonDioxide": i.CarbonDioxide,
		"projectedDate": i.ProjectedDate,
		"status": i.Status,
	}

	if i.ID != objectid.NilObjectID {
		in["_id"] = i.ID.Hex()
	}
	return json.Marshal(in)
}

func (i *Warning) UnmarshalBSON(in []byte) error {
	m := make(map[string]interface{})
	err := bson.Unmarshal(in, m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	err = i.unmarshalFromMap(m)
	return err
}

func (i *Warning) UnmarshalJSON(in []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(in, &m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	err = i.unmarshalFromMap(m)
	return err
}

// unmarshalFromMap unmarshals Map into Inventory.
func (i *Warning) unmarshalFromMap(m map[string]interface{}) error {
	var err error
	var assertOK bool

	// Hoping to discover a better way to do this someday
	if m["_id"] != nil {
		i.ID, assertOK = m["_id"].(objectid.ObjectID)
		if !assertOK {
			i.ID, err = objectid.FromHex(m["_id"].(string))
			if err != nil {
				err = errors.Wrap(err, "Error while asserting ObjectID")
				return err
			}
		}
	}

	if m["warningID"] != nil {
		i.WarningID, err = uuuid.FromString(m["warningID"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error while asserting warningID")
			return err
		}
	}

	if m["itemID"] != nil {
		i.ItemID, err = uuuid.FromString(m["itemID"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error while asserting ItemID")
			return err
		}
	}

	if m["lot"] != nil {
		i.Lot, assertOK = m["lot"].(string)
		if !assertOK {
			return errors.New("Error while asserting Lot")
		}
	}
	if m["name"] != nil {
		i.Name, assertOK = m["name"].(string)
		if !assertOK {
			return errors.New("Error while asserting Name")
		}
	}

	if m["warningActive"] != nil {
		i.WarningActive, assertOK = m["warningActive"].(bool)
		if !assertOK {
			return errors.New("Error while asserting warningActive")
		}
	}

	if m["sku"] != nil {
		i.SKU, assertOK = m["sku"].(string)
		if !assertOK {
			return errors.New("Error while asserting Sku")
		}
	}
	if m["soldWeight"] != nil {
		i.SoldWeight, err = util.AssertFloat64(m["soldWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting SoldWeight")
			return err
		}
	}
	if m["timestamp"] != nil {
		i.Timestamp, err = util.AssertInt64(m["timestamp"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting Timestamp")
			return err
		}
	}
	if m["totalWeight"] != nil {
		i.TotalWeight, err = util.AssertFloat64(m["totalWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting TotalWeight")
			return err
		}
	}
	if m["unsoldWeight"] != nil {
		i.UnsoldWeight, err = util.AssertFloat64(m["unsoldWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting unsoldWeight")
			return err
		}
	}
	if m["ethylene"] != nil {
		i.Ethylene, err = util.AssertFloat64(m["ethylene"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting Ethylene")
			return err
		}
	}
	if m["carbonDioxide"] != nil {
		i.CarbonDioxide, err = util.AssertFloat64(m["carbonDioxide"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting carbonDioxide")
			return err
		}
	}
	if m["projectedDate"] != nil {
		i.ProjectedDate, err = util.AssertInt64(m["projectedDate"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting Projected Date")
			return err
		}
	}

	if m["status"] != nil {
		i.Status, assertOK = m["status"].(string)
		if !assertOK {
			return errors.New("Error while asserting Status")
		}
	}

	return nil
}
