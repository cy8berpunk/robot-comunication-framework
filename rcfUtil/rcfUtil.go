/*
Package rcfutil implements basic parsing and de-encoding for rcf_node & rcf_node_client
*/
package rcfUtil

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

// serializable protocol msg
type Smsg struct {
	Type string
	Name string
	Id int
	Operation string
	Payload []byte
}

// naming convention whitelist
// every topic, action, service name is compared to that list. Characters which conflict with the protocl are removed
var namingSchemeWhitelist string = "abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789"

// basic logger declarations
// loggers are initiated by node or client
var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func EncodeMsg(msg *Smsg) ([]byte, error) {
	serializedMsg, err := json.Marshal(&msg)
	if err != nil {
		return []byte{}, err
	}
	return serializedMsg, nil
}


func DecodeMsg(msg *Smsg, data []byte) error {
	err := json.Unmarshal(data, msg)
	if err != nil {
		return err
	}
	return nil
}


// CompareSlice compares two slices for equality
// slices must be of same length
// returns false if slices are not equal
func CompareSlice(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	InfoLogger.Println("CompareSlice called")
	return true
}

// TopicsContainTopic checks if the topics map contains a certain topic(name)
// returns false if topic(name) is not included in the list
func TopicsContainTopic(imap map[string][][]byte, key string) bool {
	if _, ok := imap[key]; ok {
		return true
	}
	InfoLogger.Println("TopicsContainTopic called")
	return false
}

// GenRandomIntID generates random id
// returns generated random id
func GenRandomIntID() int {
	InfoLogger.Println("GenRandomIntID called")
	pullReqID := rand.Intn(1000000000)
	if pullReqID == 0 || pullReqID == 2 {
		pullReqID = rand.Intn(100000000)
	}
	return pullReqID
}
