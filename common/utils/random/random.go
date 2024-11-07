package random

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"time"
)

var node *snowflake.Node

const startTime = "2024-04-16"
const machineID = 1

func init() {

	var st time.Time
	st, err := time.Parse("2006-01-02", startTime)

	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		log.Panicln("Snowflake Init failed", zap.Error(err))
		return
	}
}

// 生成 64 位的 雪花 ID
func GetSnowIDbyInt64() int64 {
	return node.Generate().Int64()
}
func GetSnowIDbyStr() string {
	return node.Generate().String()
}

// 生成36 位的字符串的uuid
func GetUUid() string {
	return uuid.NewString()
}

// GetRandomNum 生成指定长度的随机数
func GetRandomNum(length int) (num string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		num += fmt.Sprintf("%d", r.Intn(9)+1)
	}
	return num
}
