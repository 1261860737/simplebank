package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init(){
	rand.Seed(time.Now().UnixNano())
}

// 随机产生一个在min -max区间范围的数
func RandomInt(min, max int64) int64{
	return min+rand.Int63n(max - min + 1)
}

func RandomString(n int) string{
	var sb strings.Builder
	k := len(alphabet)

	for i:=0; i<n; i++{
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//随机生成拥有者的名字
func RandomOwner() string{
	return RandomString(6)
}

// 随机生成金额
func RandomMoney() int64{
	return RandomInt(0, 1000)
}

// 随机选择一个货币
func RandomCurrency() string{
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}