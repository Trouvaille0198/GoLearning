package osLearning

import (
	"goLearning/utils"
	"math/rand"
	"time"
)

func getWeightedChooser() *utils.Chooser {
	chooser, _ := utils.NewChooser(
		utils.NewChoice("front", 25),
		utils.NewChoice("end", 25),
		utils.NewChoice("prev", 50),
	)
	return chooser
}

func GetRandomAddresses(addressNum int, virtualMemorySize int) []int {
	// 获取随机指令地址
	virtualMemorySize *= 1024
	rs := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	addresses := make([]int, addressNum)
	chooser := getWeightedChooser()
	addresses[0] = rs.Intn(virtualMemorySize)
	for i := 1; i < addressNum; i++ {
		initForm := chooser.PickSource(rs)
		switch initForm {
		case "front":
			randAddress := rs.Intn(virtualMemorySize / 2)
			addresses[i] = randAddress
		case "end":
			randAddress := rs.Intn(virtualMemorySize/2) + virtualMemorySize/2
			addresses[i] = randAddress
		case "prev":
			addresses[i] = addresses[i-1] + 1
		}
	}
	return addresses
}

func GetAddressPages(addresses []int, pageSize int) []int {
	addressPages := make([]int, len(addresses))
	for i, address := range addresses {
		addressPages[i] = address/(pageSize*1024) + 1
	}
	return addressPages
}

type RequestPageManager struct {
	_pageSize     int
	_internalSize int
}

func NewRequestPageManager(pageSize int, internalSize int) *RequestPageManager {
	r := RequestPageManager{_pageSize: pageSize, _internalSize: internalSize}
	return &r
}

func main() {

}
