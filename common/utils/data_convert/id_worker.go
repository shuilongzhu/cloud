package data_convert

import (
	"crypto/md5"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	GenerateRandMap     map[string]string
	GenerateRandRWMutex *sync.RWMutex
	AutoId              *AutoInc
	longLetters         = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type AutoInc struct {
	start, step int
	queue       chan int
	running     bool
}

func NewAutoIncInit(start, step int) (ai *AutoInc) {
	ai = &AutoInc{
		start:   start,
		step:    step,
		running: true,
		queue:   make(chan int, 4),
	}
	go ai.process()
	return
}

func (ai *AutoInc) process() {
	defer func() { recover() }()
	for i := ai.start; ai.running; i = i + ai.step {
		ai.queue <- i
	}
}

func (ai *AutoInc) Id() int {
	return <-ai.queue
}

func (ai *AutoInc) Close() {
	ai.running = false
	close(ai.queue)
}

// NewAutoIdInc 20211102195043100000
func NewAutoIdInc(aid int, intprefix string) string {
	b := strconv.Itoa(aid)
	return intprefix + DateNowFormatC() + b
}

func init() {
	rand.Seed(time.Now().Unix())
}

// RandUp 随机字符串，包含 英文字母和数字附加=_两个符号
func RandUp(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}
	}
	for i, x := range b {
		arc = x % 61
		b[i] = longLetters[arc]
	}
	return b
}

//GetRandomStr 获取随机字符串
func GetRandomStr(length int) string {
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63())) //增大随机性
	for i := 0; i < length; i++ {
		result[i] = longLetters[r.Intn(len(longLetters))]
	}
	return string(result)
}

// GetMD5Encode 返回一个32位md5加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// GenerateRand 数据同步的请求ID
func GenerateRand(key string) string {
	GenerateRandRWMutex.Lock()
	defer func() {
		GenerateRandRWMutex.Unlock()
	}()
GenerateRand:
	Rand := string(RandUp(10))
	str := Rand + key
	value, ok := GenerateRandMap[str]
	if ok {
		log.Errorln("重复ID value:", value)
		goto GenerateRand
	}
	GenerateRandMap[str] = str
	log.Println("len(GenerateRandMap):", len(GenerateRandMap), "Rand+key:", str)
	return str
}
