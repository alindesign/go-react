package react

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"sync"
)

var (
	cLock    sync.Mutex
	cacheMap = map[string]string{}
)

func setCache(key, value string) {
	cLock.Lock()
	cacheMap[key] = value
	cLock.Unlock()
}

func getCache(key string) string {
	return cacheMap[key]
}

func getOrSet(key string, resolver func() string) string {
	cLock.Lock()
	val, ok := cacheMap[key]
	cLock.Unlock()

	if val == "" || !ok {
		val = resolver()
		setCache(key, val)
		return val
	}

	return val
}

func checksum(namespace string, val interface{}) string {
	content, err := json.Marshal(val)
	if err != nil {
		return ""
	}

	h := md5.New()
	h.Write(content)

	hash := hex.EncodeToString(h.Sum(nil))

	if hash != "" {
		return namespace + ":" + hash
	}

	return ""
}
