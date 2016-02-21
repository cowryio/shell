// Contains validation method for seed
package seed

import "errors"
import "encoding/json"
import "fmt"
import "time"

// Validate `meta` block
// * A valid meta block must contain seed_id, seed_type and created_at properties
// * seed_id must be string and 40 characters in length
// * seed_type must be string
// * created_at must be a valid unix date in the past but not beyond a start/launch time
func ValidateMetaBlock(meta map[string]interface{}) error {

	// must reject unexpected properties
	accetableProps := []string{ "seed_id", "seed_type", "created_at" } 
	for prop, _ := range meta {
		if !InStringSlice(accetableProps, prop) {
			return errors.New(fmt.Sprintf("`%s` property is unexpected in `meta` block", prop))
		}
	}

	// must have expected properties
	props := []string{"seed_id", "seed_type", "created_at"}
	for _, prop := range props {
		if !HasKey(meta, prop) {
			return errors.New(fmt.Sprintf("`meta` block is missing `%s` property", prop))
		} 
	}

	// seed id must be a string and should be 40 characters in length
	if !IsStringValue(meta["seed_id"]) {
		return errors.New("`meta.seed_id` value type is invalid. Expects string value")
	} else {
		if len(meta["seed_id"].(string)) != 40 {
			return errors.New("`meta.seed_id` must have 40 characters. Preferrable a UUIDv4 SHA1 hashed string")
		}
	}

	// seed_type must be string
	if !IsStringValue(meta["seed_type"]) {
		return errors.New("`meta.seed_type` value type is invalid. Expects string value")
	}

	// created_at should be a number
	if !IsNumberValue(meta["created_at"]) {
		return errors.New("`created_at` value type is invalid. Expects a number")
	} else {

		// the unix time that indicates the time from when
		// a meta.created_at time must start from
		START_TIME := ToInt64(Env("LAUNCH_TIME", "1453975575")) 

		// created_date should not be too old or in the future
		createdAt := UnixToTime(ToInt64(meta["created_at"]))
		startTime := UnixToTime(START_TIME)

		if createdAt.Before(startTime) {
			return errors.New("`created_at` value is too far in the past. Expects unix time on or after " + startTime.Format(time.RFC3339))
		} else if createdAt.After(time.Now().UTC()) {
			return errors.New("`created_at` value cannot be a unix time in the future")
		}
	}

	return nil
}

// Validate `signature` block'
// * `meta` signature must be present and must be a string type
// * `attributes` property must be string type if set
// * `ownership` property must be string type if set
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
			return errors.New("`signatures.meta` value type is invalid. Expects string value")
		}
	}

	// if signature has `ownership` property, it's value type must be string
	if signatures["ownership"] != nil {
		if !IsStringValue(signatures["ownership"]) {
			return errors.New("`signatures.ownership` value type is invalid. Expects string value")
		}
	}

	// if signature has `attributes` property, it's value type must be string
	if signatures["attributes"] != nil {
		if !IsStringValue(signatures["attributes"]) {
			return errors.New("`signatures.attributes` value type is invalid. Expects string value")
		}
	}

	// if signature has `embeds` property, it's value type must be string
	if signatures["embeds"] != nil {
		if !IsStringValue(signatures["embeds"]) {
			return errors.New("`signatures.embeds` value type is invalid. Expects string value")
		}
	}
	
	return nil
}

// Validate ownership block
// * `type` property must exist and must contain acceptable value
func ValidateOwnershipBlock(ownership map[string]interface{}) error {

	// must reject unexpected properties
	accetableProps := []string{ "type", "sole", "status" } 
	for prop, _ := range ownership {
		if !InStringSlice(accetableProps, prop) {
			return errors.New(fmt.Sprintf("`%s` property is unexpected in `ownership` block", prop))
		}
	}

	// `type` property must be set
	if ownership["type"] != nil {
		acceptableValues := []string{"sole"}
		if !InStringSlice(acceptableValues, ownership["type"].(string)) {
			return errors.New("`ownership.type` property has unexpected value")
		}
	} else {
		return errors.New("`ownership` block is missing `type` property")
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
			return errors.New("`ownership.sole.address_id` value type is invalid. Expects string value")
		}
	}

	// `status` property is optional, but if set, it's type must be string 
	// and must have acceptable values
	if ownership["status"] != nil {
		if !IsStringValue(ownership["status"]) {
			return errors.New("`ownership.status` value type is invalid. Expects string value")
		} else {
			acceptableValues := []string{ "transferred" }
			if !InStringSlice(acceptableValues, ownership["status"].(string)) {
				return errors.New("`ownership.status` property has unexpected value")
			}
		}
	}
	
	return nil
}

// Validate a seed. This function ensures 
// the existence of mandatory seed properties and attributes.
// TODO: if ownership block exists, signatures block must have ownership property
func Validate(seedData interface{}) error {

	// parse seed data to map[string]interface{} seedData is string
	var data map[string]interface{}
	switch d := seedData.(type) {
	case string:
		if err := json.Unmarshal([]byte(d), &data); err != nil {
	        return errors.New("unable to parse json string");
	    }
	    break;
	case map[string]interface{}:
		data = d
		break
	default:
		errors.New("unsupported seed data type. Requires seed data in JSON string or golang map[string]interface{}");
	}

    // must have `meta` block
    if data["meta"] == nil {
    	return errors.New("missing `meta` block")
    } else {
    	switch meta := data["meta"].(type) {
    	case map[string]interface{}:
    		if  err := ValidateMetaBlock(meta); err != nil {
    			return err
    		}
    		break
    	default:
    		return errors.New("`meta` block value type is invalid. Expects a JSON object")
    	}
    }
    
    // must have `signatures` block
    if data["signatures"] == nil {
    	return errors.New("missing `signatures` block")
    } else {
    	if IsMapOfAny(data["signatures"]) {
    		var signatures = data["signatures"].(map[string]interface{})
    		if err := ValidateSignaturesBlock(signatures); err != nil {
    			return err
    		}
		} else {
			return errors.New("`signature` block value type is invalid. Expects a JSON object")	
		}
    }

    // if `ownership` block exists, it must be a map
    if data["ownership"] != nil {
    	if !IsMapOfAny(data["ownership"]) {
    		return errors.New("`ownership` block value type is invalid. Expects a JSON object")
    	} else {
    		if !IsMapEmpty(data["ownership"].(map[string]interface{})) {
	    		var signatures = data["signatures"].(map[string]interface{})
	    		if signatures["ownership"] == nil {
	    			return errors.New("missing `ownership` property in `signatures` block")
	    		} else {
	    			if err := ValidateOwnershipBlock(data["ownership"].(map[string]interface{})); err != nil {
	    				return err
	    			}
	    		}
	    	}
    	}
    }

    // if `attributes` block exists, it must be a map
    if data["attributes"] != nil {
    	if !IsMapOfAny(data["attributes"]) {
    		return errors.New("`attributes` block value type is invalid. Expects a JSON object")
    	} else {
    		if !IsMapEmpty(data["attributes"].(map[string]interface{})) {
	    		// `signatures` block must have `attributes` property
	    		var signatures = data["signatures"].(map[string]interface{})
	    		if signatures["attributes"] == nil {
	    			return errors.New("missing `attributes` property in `signatures` block")
	    		}
	    	}
    	}
    }

    // if `embeds` block exists, it must be a slice of maps
    if data["embeds"] != nil {

    	if !IsSlice(data["embeds"]) || !ContainsOnlyMapType(data["embeds"].([]interface{}))  {
    		return errors.New("`embeds` block value type is invalid. Expects a list of only JSON objects")
    	}

    	embeds := data["embeds"].([]interface{})
    	if len(embeds) == 0 {
    		return nil
    	}

    	// ensure `embeds` signature exists
    	var signatures = data["signatures"].(map[string]interface{})
    	if signatures["embeds"] == nil {
			return errors.New("missing `embeds` property in `signatures` block")
		}

		// validate each seeds in the embeds block. Prevent validaton of the individual seeds' embeds
		// by empting their `embeds` block before calling Validate() on them. Reassign their `embeds` block
		// back after validation.
		for i, embed := range embeds {
			
			var embedsClone []interface{}
			seed := embed.(map[string]interface{})
			
			// Ensure the seed has an `embed` block and the value type is a slice.
			// Then temporary remove embeds property in the seed as we aren't interested in validating deeper levels
			if seed["embeds"] != nil && IsSlice(seed["embeds"].([]interface{}))  {
				embedsClone = CloneSliceOfInterface(seed["embeds"].([]interface{}))
				seed["embeds"] = []interface{}{}
			}

			if err := Validate(seed); err != nil {
				return errors.New(fmt.Sprintf("unable to validate embed at index %d. Reason: %s", i, err.Error()))
			}

			// reassign seed's embeds
			seed["embeds"] = embedsClone
		}
    }

    return nil
}