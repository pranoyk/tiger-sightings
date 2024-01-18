package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

func DecodeCursor(cursor string) (res time.Time, uuid string, err error) {
	if cursor == "" {
		currTime := time.Now().Format("2006-01-02T15:04:05Z")
		res, _ = time.Parse("2006-01-02T15:04:05Z", currTime)
		uuid = "00000000-0000-0000-0000-000000000000"
		return
	}
	byt, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return
	}

	arrStr := strings.Split(string(byt), ",")
	if len(arrStr) != 2 {
		err = errors.New("cursor is invalid")
		return
	}
	fmt.Printf("arrStr: %v\n", arrStr)
	res, err = time.Parse("2006-01-02T15:04:05Z", arrStr[0])
	if err != nil {
		return
	}
	uuid = arrStr[1]
	return
}

func Encode(lastSeen time.Time, uuid string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(lastSeen.Format("2006-01-02T15:04:05Z") + "," + uuid))
	return encoded
}
