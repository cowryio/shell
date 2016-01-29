// Contains validation method for shell
package shell

import "errors"
import "encoding/json"
import "fmt"
import "time"

// Validate `meta` block
// * A valid meta block must contain shell_id, shell_type and created_at properties
// * shell_id must be string and 40 characters in length
// * shell_type must be string
// * created_at must be a valid unix date in the past but not beyond a start/launch time
func ValidateMetaBlock(meta map[string]interface{}) error {

	// must have expected properties
	props := []string{"shell_id", "shell_type", "created_at"}
	for _, prop := range props {
		if !HasKey(meta, prop) {
			return errors.New(fmt.Sprintf("`meta` block is missing `%s` property", prop))
		} 
	}

	// shell id must be a string and should be 40 characters in length
	if !IsStringValue(meta["shell_id"]) {
		return errors.New("`meta.shell_id` value type is invalid. Expects string value")
	} else {
		if len(meta["shell_id"].(string)) != 40 {
			return errors.New("`meta.shell_id` must have 40 characters. Preferrable a UUIDv4 SHA1 hashed string")
		}
	}

	// shell_type must be string
	if !IsStringValue(meta["shell_type"]) {
		return errors.New("`meta.shell_type` value type is invalid. Expects string value")
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
// TODO: `attributes` property must be string type if set
// TODO: `ownership` property must be string type if set
func ValidateSignaturesBlock(signatures map[string]interface{}) error {

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
	
	return nil
}

// Validate a shell. This function ensures 
// the existence of mandatory shell properties and attributes.
// TODO: if ownership block exists, signatures block must have ownership property
func Validate(shellData interface{}) error {

	// parse shell data to map[string]interface{} shellData is string
	var data map[string]interface{}
	switch d := shellData.(type) {
	case string:
		if err := json.Unmarshal([]byte(d), &data); err != nil {
	        return errors.New("unable to parse json string");
	    }
	    break;
	case map[string]interface{}:
		data = d
		break
	default:
		errors.New("unsupported shell data type. Requires shell data in JSON string or golang map[string]interface{}");
	}

    // must have `meta` block
    if data["meta"] == nil {
    	return errors.New("missing `meta` block")
    } else {
    	switch meta := data["meta"].(type) {
    	case map[string]interface{}:
    		if err := ValidateMetaBlock(meta); err != nil {
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
    		// `signatures` block must have `ownership` property
    		var signatures = data["signatures"].(map[string]interface{})
    		if signatures["ownership"] == nil {
    			return errors.New("missing `ownership` property in `signatures` block")
    		}
    	}
    }

    // if `attributes` block exists, it must be a map
    if data["attributes"] != nil {
    	if !IsMapOfAny(data["attributes"]) {
    		return errors.New("`attributes` block value type is invalid. Expects a JSON object")
    	} else {
    		// `attributes` block must have `ownership` property
    		var signatures = data["signatures"].(map[string]interface{})
    		if signatures["attributes"] == nil {
    			return errors.New("missing `attributes` property in `signatures` block")
    		}
    	}
    }

    return nil
}