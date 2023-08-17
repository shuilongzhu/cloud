package data_convert

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var s = new(SnowFlake)

/*
   雪花算法(snowFlake)的具体实现方案:
*/

type SnowFlake struct {
	mu sync.Mutex
	//雪花算法开启时的起始时间戳
	twepoch int64

	//每一部分占用的位数
	workerIdBits     int64 //每个数据中心的工作机器的编号位数
	datacenterIdBits int64 //数据中心的编号位数
	sequenceBits     int64 //每个工作机器每毫秒递增的位数

	//每一部分最大的数值
	maxWorkerId     int64
	maxDatacenterId int64
	maxSequence     int64

	//每一部分向左移动的位数
	workerIdShift     int64
	datacenterIdShift int64
	timestampShift    int64

	//当前数据中心ID号
	datacenterId int64
	//当前机器的ID号
	workerId int64
	//序列号
	sequence int64
	//上一次生成ID号前41位的毫秒时间戳
	lastTimestamp int64
}

/*
   获取毫秒的时间戳
*/
func timeGen() int64 {
	return time.Now().UnixMilli()
}

/*
   获取比lastTimestamp大的当前毫秒时间戳
*/
func tilNextMills() int64 {
	timeStampMill := timeGen()
	for timeStampMill <= s.lastTimestamp {
		timeStampMill = timeGen()
	}
	return timeStampMill
}
func NextId() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	nowTimestamp := timeGen() //获取当前的毫秒级别的时间戳
	if nowTimestamp < s.lastTimestamp {
		//系统时钟倒退,倒退了s.lastTimestamp-nowTimestamp
		log.Errorln(fmt.Sprintf("clock moved backwards, Refusing to generate id for %d milliseconds", s.lastTimestamp-nowTimestamp))
		return -1
	}
	if nowTimestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & s.maxSequence
		if s.sequence == 0 {
			//tilNextMills中有一个循环等候当前毫秒时间戳到达lastTimestamp的下一个毫秒时间戳
			nowTimestamp = tilNextMills()
		}
	} else {
		s.sequence = 0
	}
	s.lastTimestamp = nowTimestamp
	return (nowTimestamp-s.twepoch)<<s.timestampShift | //时间戳差值部分
		s.datacenterId<<s.datacenterIdShift | //数据中心部分
		s.workerId<<s.workerIdShift | //工作机器编号部分
		s.sequence //序列号部分

}

func NewSnowFlake(workerId int64, datacenterId int64) (*SnowFlake, error) {
	mySnow := new(SnowFlake)
	mySnow.twepoch = time.Now().Unix() //返回当前时间的时间戳(时间戳是指北京时间1970年01月01日8时0分0秒到此时时刻的总秒数)
	if workerId < 0 || datacenterId < 0 {
		return nil, errors.New("workerId or datacenterId must not lower than 0 ")
	}
	/*
	   标准的雪花算法
	*/
	mySnow.workerIdBits = 5
	mySnow.datacenterIdBits = 5
	mySnow.sequenceBits = 12

	mySnow.maxWorkerId = -1 ^ (-1 << mySnow.workerIdBits)         //64位末尾workerIdBits位均设为1,其余设为0
	mySnow.maxDatacenterId = -1 ^ (-1 << mySnow.datacenterIdBits) //64位末尾datacenterIdBits位均设为1,其余设为0
	mySnow.maxSequence = -1 ^ (-1 << mySnow.sequenceBits)         //64位末尾sequenceBits位均设为1,其余设为0

	if workerId >= mySnow.maxWorkerId || datacenterId >= mySnow.maxDatacenterId {
		return nil, errors.New("workerId or datacenterId must not higher than max value ")
	}
	mySnow.workerIdShift = mySnow.sequenceBits
	mySnow.datacenterIdShift = mySnow.sequenceBits + mySnow.workerIdBits
	mySnow.timestampShift = mySnow.sequenceBits + mySnow.workerIdBits + mySnow.datacenterIdBits

	mySnow.lastTimestamp = -1
	mySnow.workerId = workerId
	mySnow.datacenterId = datacenterId

	return mySnow, nil
}

func init() {
	tempS, err := NewSnowFlake(0, 0)
	if nil != err {
		log.Errorln("生成生成雪花算法失败")
		return
	}
	s = tempS
}
