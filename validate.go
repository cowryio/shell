// Contains validation method for stone
package stone

import "errors"
import "encoding/json"
import "fmt"
import "strings"
import "time"

// the unix time that indicates the time from when
// a meta.created_at time must start from
var START_TIME int64 = 1453975575


// Set the start time
func SetStartTime(t int64) {
	START_TIME = t
}


// Validate `meta` block
// * Must not contain unknown properties
// * A valid meta block must contain id, type and created_at properties
// * id must be string and 40 characters in length
// * type must be string
// * created_at must be an interger and a valid unix date in the past but not beyond a start/launch time
func ValidateMetaBlock(meta map[string]interface{}) error {

	var createdAt int64
	var err error

	// must reject unexpected properties
	accetableProps := []string{ "id", "type", "created_at" } 
	for prop, _ := range meta {
		if !InStringSlice(accetableProps, prop) {
			return errors.New(fmt.Sprintf("`%s` property is unexpected in `meta` block", prop))
		}
	}

	// must have expected properties
	props := []string{"id", "type", "created_at"}
	for _, prop := range props {
		if !HasKey(meta, prop) {
			return errors.New(fmt.Sprintf("`meta` block is missing `%s` property", prop))
		} 
	}

	// stone id must be a string
	if !IsStringValue(meta["id"]) {
		return errors.New("`meta.id` value type is invalid. Expects a string")
	}

	// stone id must be 40 characters in length
	if len(meta["id"].(string)) != 40 {
		return errors.New("`meta.id` must have 40 characters. Preferrable a UUIDv4 SHA1 hashed string")
	}
	
	// type must be string
	if !IsStringValue(meta["type"]) {
		return errors.New("`meta.type` value type is invalid. Expects a string")
	}

	// created_at must be a json number or a float or integer
	if !IsJSONNumber(meta["created_at"]) && !IsNumberValue(meta["created_at"]) {
		return errors.New("`meta.created_at` value type is invalid. Expects a number")
	}

	// created_at is json.Number, convert to int64
	if IsJSONNumber(meta["created_at"]) {
		createdAt, err = meta["created_at"].(json.Number).Int64()
		if err != nil {
			return errors.New("`meta.created_at` value type is invalid. Expects a number")
		}
	}

	// created_at is an integer
	if IsInt(meta["created_at"]) {
		createdAt = ToInt64(meta["created_at"])
	}
	
	// make time objects
	createdAtTime := UnixToTime(createdAt)
	startTime := UnixToTime(START_TIME)

	// date of creation cannot be before the start time
	if createdAtTime.Before(startTime) {
		return errors.New("`meta.created_at` value is too far in the past. Expects unix time on or after " + startTime.Format(time.RFC3339))
	}

	// date of creation cannot be a time in the future
	if createdAtTime.After(time.Now().UTC()) {
		return errors.New("`meta.created_at` value cannot be a unix time in the future")
	}

	return nil
}

// Validate `signature` block'
// * must contain only acceptable properties (meta, ownership, embeds)
// *`meta` signature must be present and must be a string type
// *`attributes` property must be string type if set
// *`ownership` property must be string type if set
// *`embeds` property must be string type if set
func ValidateSignaturesBlock(signatures map[string]interface{}) error {

	// must reject unexpected properties
	accetableProps := []string{ "meta", "ownership", "attributes", "embeds" } 
	for prop, _ := range signatures {
		if !InStringSlice(accetableProps, prop) {
			return errors.New(fmt.Sprintf("`%s` property is unexpected in `signatures` block", prop))
		}
	}

	// must have `meta` property
	if signatures["meta"] == nil {
		return errors.New("missing `signatures.meta` property")
	} else {
		// meta value type must be string
		if !IsStringValue(signatures["meta"]) {
			return errors.New("`signatures.meta` value type is invalid. Expects a string")
		}
	}

	// if signature has `ownership` property, it's value type must be string
	if signatures["ownership"] != nil {
		if !IsStringValue(signatures["ownership"]) {
			return errors.New("`signatures.ownership` value type is invalid. Expects a string")
		}
	}

	// if signature has `attributes` property, it's value type must be string
	if signatures["attributes"] != nil {
		if !IsStringValue(signatures["attributes"]) {
			return errors.New("`signatures.attributes` value type is invalid. Expects a string")
		}
	}

	// if signature has `embeds` property, it's value type must be string
	if signatures["embeds"] != nil {
		if !IsStringValue(signatures["embeds"]) {
			return errors.New("`signatures.embeds` value type is invalid. Expects a string")
		}
	}
	
	return nil
}

// Validate ownership block.
// * Must not contain unknown properties
// * A valid ownership block can only contain ref_id, type, sole and status properties.
// * ownership.ref_id must be set and value type must be string
// * ref_id must be equal to the meta id
// * A valid ownership block can only contain type, sole and status properties.
// * ownership.type must be set, value type must be a string and value must be known
// * - if ownership.type is 'sole':
// * - ownership.sole must be set to an object
// * - ownership.sole.address_id must be set and it must be a string
// * ownership.status is optional, but if set
// *  - ownership.status must be a string value. The value must also be known
func ValidateOwnershipBlock(ownership map[string]interface{}, metaID string) error {

	// must reject unexpected properties
	accetableProps := []string{ "ref_id", "type", "sole", "status" } 
	for prop, _ := range ownership {
		if !InStringSlice(accetableProps, prop) {
			return errors.New(fmt.Sprintf("`%s` property is unexpected in `ownership` block", prop))
		}
	}

	// `ref_id` property must be set
	if ownership["ref_id"] == nil {
		return errors.New("`ownership` block is missing `ref_id` property")
	}

	// `ref_id` property must be a string value
	if !IsStringValue(ownership["ref_id"]) {
		return errors.New("`ownership.ref_id` value type is invalid. Expects string value")
	}

	// `ref_id` property must be equal to meta id 
	if ownership["ref_id"].(string) != metaID {
		return errors.New("`ownership.ref_id` not equal to `meta.id`");
	}
 
	// `type` property must be set
	if ownership["type"] == nil {
		return errors.New("`ownership` block is missing `type` property")
	}

	// type property must have string value
	if !IsStringValue(ownership["type"]) {
		return errors.New("`ownership.type` value type is invalid. Expects a string")
	}
	
	// type property value must be known
	acceptableValues := []string{"sole"}
	if !InStringSlice(acceptableValues, ownership["type"].(string)) {
		return errors.New("`ownership.type` property has unexpected value")
	}

	// if ownership.type is `sole`, `sole` property is required
	if ownership["type"].(string) == "sole" && ownership["sole"] == nil {
		return errors.New("`ownership` block is missing `sole` property")
	} else {

		// `sole` property must be a map
		if !IsMapOfAny(ownership["sole"]) {
			return errors.New("`ownership.sole` value type is invalid. Expects a JSON object")
		}

		// `sole` property must have `address_id` property
		soleProperty := ownership["sole"].(map[string]interface{})
		if soleProperty["address_id"] == nil {
			return errors.New("`ownership.sole` property is missing `address_id` property")
		}

		// `sole.address_id` value type must be string
		if !IsStringValue(soleProperty["address_id"]) {
			return errors.New("`ownership.sole.address_id` value type is invalid. Expects a string")
		}
	}

	// `status` property is optional, but if set, it's type must be string 
	// and must have acceptable values
	if ownership["status"] != nil {
		if !IsStringValue(ownership["status"]) {
			return errors.New("`ownership.status` value type is invalid. Expects a string")
		}
		if !InStringSlice([]string{ "transferred" }, ownership["status"].(string)) {
			return errors.New("`ownership.status` property has unexpected value")
		}
	}
	
	return nil
}


// Validate attributes block.
// * Accept only `ref_id` and `data` properties 
// * `ref_id` property must be provided
// * `ref_id` property must be a string
// * `ref_id` property must equal meta id (meta.id property)
// * `data` propert must be set
func ValidateAttributesBlock(attributes map[string]interface{}, metaID string) error {

	// must reject unexpected properties
	for prop, _ := range attributes {
		if !InStringSlice([]string{ "ref_id", "data" }, prop) {
			return errors.New(fmt.Sprintf("`%s` property is unexpected in `attributes` block", prop))
		}
	}

	// `ref_id` property must be set
	if attributes["ref_id"] == nil {
		return errors.New("`attributes` block is missing `ref_id` property")
	}

	// `ref_id` property must be a string value
	if !IsStringValue(attributes["ref_id"]) {
		return errors.New("`attributes.ref_id` value type is invalid. Expects string value")
	}

	// `ref_id` property must be equal to meta id 
	if attributes["ref_id"].(string) != metaID {
		return errors.New("`attributes.ref_id` not equal to `meta.id`");
	}

	// `data` property must be provided
	if (attributes["data"] == nil) {
		return errors.New("`attributes` block is missing `data` property");
	}

	return nil
}


// Validate embeds block.
// Rules:
// * expects a json object as parameter
// * must not contain expected properties other than `ref_id` and `data`
// * ref_id must be set and should have a string value
// * ref_id must be equal to meta id
// * data property must be set and value type must be an array of json objects
func ValidateEmbedsBlock(embeds map[string]interface{}, metaID string) error {

	// must reject unexpected properties
	for prop, _ := range embeds {
		if !InStringSlice([]string{ "ref_id", "data" }, prop) {
			return errors.New(fmt.Sprintf("`%s` property is unexpected in `embeds` block", prop))
		}
	}

	// `ref_id` property must be set
	if embeds["ref_id"] == nil {
		return errors.New("`embeds` block is missing `ref_id` property")
	}

	// `ref_id` property must be a string value
	if !IsStringValue(embeds["ref_id"]) {
		return errors.New("`embeds.ref_id` value type is invalid. Expects string value")
	}

	// `ref_id` property must be equal to meta id 
	if embeds["ref_id"].(string) != metaID {
		return errors.New("`embeds.ref_id` not equal to `meta.id`");
	}

	// `data` property must be provided
	if (embeds["data"] == nil) {
		return errors.New("`embeds` block is missing `data` property");
	}

	// `data` property must be a map
	if !IsSlice(embeds["data"]) || !ContainsOnlyMapType(embeds["data"].([]interface{})) {
		return errors.New("`embeds.data` value type is invalid. Expects a slice of JSON objects")
	}

	allEmbeds := embeds["data"].([]interface{})

	// validate each embeds. To prevent continues validaton of child embeds,  
	// we will remove the `embeds` block before calling Validate() for every objects, 
	// Reassigning the values of the `embeds` block after validation.
	for i, embed := range allEmbeds {
		
		var embedsClone map[string]interface{}
		item := embed.(map[string]interface{})
		
		// Ensure the item has a embeds block set.
		// If so, temporary remove embeds property of the object
		if item["embeds"] != nil {
			embedsClone = item["embeds"].(map[string]interface{})
			item["embeds"] = map[string]interface{}{}
		}

		if err := Validate(item); err != nil {
			return errors.New(fmt.Sprintf("unable to validate embed at index %d. Reason: %s", i, err.Error()))
		}

		// reassign stone's embeds
		item["embeds"] = embedsClone
	}

	return nil
}

// Validate a stone. This function ensures 
// the existence of mandatory stone properties and attributes.
func Validate(stoneData interface{}) error {

	var metaID string

	// parse stone data to map[string]interface{} stoneData is string
	var data map[string]interface{}
	switch d := stoneData.(type) {
	case string:
		decoder := json.NewDecoder(strings.NewReader(d))
		decoder.UseNumber();
		if err := decoder.Decode(&data); err != nil {
	        return errors.New("unable to parse json string");
	    }
	    break;
	case map[string]interface{}:
		data = d
		break
	default:
		return errors.New("unsupported parameter type");
	}

    // must have `meta` block
    if data["meta"] == nil {
    	return errors.New("missing `meta` block")
    } else {
		if !IsMapOfAny(data["meta"]) {
			return errors.New("`meta` block value type is invalid. Expects a JSON object")
		}
		if  err := ValidateMetaBlock(data["meta"].(map[string]interface{})); err != nil {
			return err
		}
		metaBlock := data["meta"].(map[string]interface{})
		metaID = metaBlock["id"].(string)
	}

    // if `ownership` block exists, it must be a map
    if data["ownership"] != nil {
    	if !IsMapOfAny(data["ownership"]) {
    		return errors.New("`ownership` block value type is invalid. Expects a JSON object")
    	} 
		if !IsMapEmpty(data["ownership"].(map[string]interface{})) {
    		if err := ValidateOwnershipBlock(data["ownership"].(map[string]interface{}), metaID); err != nil {
				return err
			}
    	}
    }

    // if `attributes` block exists, it must be a map
    if data["attributes"] != nil {
    	if !IsMapOfAny(data["attributes"]) {
    		return errors.New("`attributes` block value type is invalid. Expects a JSON object")
    	}
		if !IsMapEmpty(data["attributes"].(map[string]interface{})) {
    		if err := ValidateAttributesBlock(data["attributes"].(map[string]interface{}), metaID); err != nil {
				return err
			}
    	}
    }

    // if `embeds` block exists, it must be a map
    if data["embeds"] != nil {
    	if !IsMapOfAny(data["embeds"]) {
    		return errors.New("`embeds` block value type is invalid. Expects a JSON object")
    	}
    	if !IsMapEmpty(data["embeds"].(map[string]interface{})) {
    		if err := ValidateEmbedsBlock(data["embeds"].(map[string]interface{}), metaID); err != nil {
				return err
			}
    	}
    }

    return nil
}