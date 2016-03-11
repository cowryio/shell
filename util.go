// Contains convenience functions
package stone

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"os"
	"encoding/json"
	"strconv"
	"errors"
	"time"
	"bytes"
	"bufio"

	"github.com/nu7hatch/gouuid"
	"github.com/jackpal/bencode-go"
)

func Println(any ...interface{}) {
	fmt.Println(any...)
}

// convert a byte array to string
func ByteArrToString(byteArr []byte) string {
	return fmt.Sprintf("%s", byteArr)
}

// Generate sha1 hash
func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

// generate random numbers between a range
func RandNum(min, max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(max-min) + min
}

// Generate an id to be used as a stone id.
func NewID() string {
	curTime := int(time.Now().Unix())
	u4, err := uuid.NewV4()
	id := fmt.Sprintf("%s:%d", u4.String(), curTime)
	if err != nil {
		return ""
	}
	return Sha1(id)
}

// Get the keys of a map
func GetMapKeys(m map[string]interface{}) []string {
	mk := make([]string, len(m))
	i := 0
	for key, _ := range m {
		mk[i] = key
		i++
	}
	return mk
}

// Generate a canonical string representation of a map using 
// bencode.
func CanonicalMap(m map[string]interface{}) (string, error) {
	var b bytes.Buffer
    w := bufio.NewWriter(&b)
    err := bencode.Marshal(w, m)
    if err != nil {
    	return "", errors.New("Unable to bencode map: " + err.Error())
    }
    w.Flush()
    return b.String(), nil
}

// checks if a key exists in a map
func HasKey(m map[string]interface{}, key string) bool {
	for k, _ := range m {
		if k == key {
			return true
		}
	}
	return false
}

// checks that a variable value type is string
func IsStringValue(any interface{}) bool {
	switch any.(type) {
	case string:
		return true
	default:
		return false
	}
}

// checks that a variable value type is a map of any value
func IsMapOfAny(any interface{}) bool {
	switch any.(type) {
	case map[string]interface{}:
		return true
		break
	default:
		return false
		break
	}
	return false
}

// checks that a variable value type is a slice
func IsSlice(any interface{}) bool {
	switch any.(type) {
	case []interface{}:
		return true
		break
	default:
		return false
		break
	}
	return false
}

// checks that a slice contains map[string]interface{} type
func ContainsOnlyMapType(s []interface{}) bool {
	for _, v := range s {
		switch v.(type) {
		case map[string]interface{}:
			continue
			break
		default:
			return false
		}
	}
	return true
}

// checks that a string slice contains a string value
func InStringSlice(ss []string, val string) bool {
	for _, v := range ss {
		if v == val {
			return true
		}
	}
	return false
}

// convert a unix time to time object
func UnixToTime(i int64) time.Time {
	return time.Unix(i, 0)
}

// copy the contents in a map of interface{} to another similar map
func CloneMapInterface(m map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

// clone a slice of map with key type of string and value of interface{}
func CloneSliceMapInterface(sm []map[string]interface{}) []map[string]interface{} {
	newSliceMap := make([]map[string]interface{}, len(sm))
	for i, m := range sm {
		newSliceMap[i] = CloneMapInterface(m)
	}
	return newSliceMap
}

// clone slice of interface{}
func CloneSliceOfInterface(s []interface{}) []interface{} {
	newSlice := make([]interface{}, len(s))
	for i, v := range s {
		newSlice[i] = v
	} 
	return newSlice
}

// check whether the value passed is int, float64, float32 or int64
func IsNumberValue(val interface{}) bool {
	switch val.(type) {
	case int, int64, float32, float64:
		return true
	default:
		return false
	}
}

// checks whether the value passed is an integer
func IsInt(val interface{}) bool {
	switch val.(type) {
	case int, int64:
		return true
	default:
		return false
	}
}

// checks whether the value passed is a json.Number type
func IsJSONNumber(val interface{}) bool {
	switch val.(type) {
	case json.Number:
		return true
	default:
		return false
	}
}

// cast int value to float64
func IntToFloat64(num interface{}) float64 {
	switch v := num.(type) {
	case int:
		return float64(v)
	case int64:
		return float64(v)
	default:
		panic("failed to cast unsupported type to float64")
	}
}

// converts int, float32 and float64 to int64
func ToInt64(num interface{}) int64 {
	switch v := num.(type) {
	case int:
		return int64(v)
		break
	case int64:
		return v
		break
	case float64:
		return int64(v)
		break
	case string:
		val, _ := strconv.ParseInt(v, 10, 64)
		return val
		break
	default:
		panic("type is unsupported")
	}
	return 0
}

// get environment variable or return a default value when no set
func Env(key, defStr string) string {
	val := os.Getenv(key)
	if val == "" && defStr != "" {
		return defStr
	}
	return val
}

// check if a map is empty
func IsMapEmpty(m map[string]interface{}) bool {
	return len(GetMapKeys(m)) == 0
}

// converts int to string
func IntToString(v int64) string {
	return fmt.Sprintf("%d", v)
}