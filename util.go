// Contains convenience functions
package shell

import (
	"crypto/sha1"
	"fmt"
	"github.com/cowryio/shell/Godeps/_workspace/src/github.com/nu7hatch/gouuid"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
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

// Generate an id to be used as a shell id.
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

// Generate a canonical string representation of a map.
// ValueType that is not int, string or map[string]interface{}
// will be ignored
func GetCanonicalMapString(m map[string]interface{}) string {
	var cannonicalStr = []string{}
	var keys = GetMapKeys(m)
	sort.Strings(keys)
	for _, key := range keys {
		val := m[key]
		switch d := val.(type) {
		case int:
			cannonicalStr = append(cannonicalStr, key+":"+strconv.Itoa(d))
			break
		case int32:
			cannonicalStr = append(cannonicalStr, key+":"+strconv.Itoa(int(d)))
			break
		case int64:
			cannonicalStr = append(cannonicalStr, key+":"+strconv.Itoa(int(d)))
			break
		case float32:
			cannonicalStr = append(cannonicalStr, key+":" + fmt.Sprintf("%.3f", d))
			break
		case float64:
			cannonicalStr = append(cannonicalStr, key+":" + fmt.Sprintf("%.3f", d))
			break
		case string:
			cannonicalStr = append(cannonicalStr, key+":"+d)
			break
		case map[string]interface{}:
			cannonicalStr = append(cannonicalStr, key+":"+GetCanonicalMapString(d))
			break
		}
	}
	return strings.Join(cannonicalStr, ":")
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
	}
	return false
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

// check if a value supplied is int, float64, float32 or int64
func IsNumberValue(val interface{}) bool {
	switch val.(type) {
	case int, int64, float32, float64:
		return true
	default:
		return false
	}
}

// converts int, float32 and float64 to int64
func ToInt64(num interface{}) int64 {
	switch v := num.(type) {
	case int:
	case int64:
		return int64(v)
		break
	case float32:
	case float64:
		return int64(v)
	case string:
		val, _ := strconv.ParseInt(v, 10, 64)
		return val
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
