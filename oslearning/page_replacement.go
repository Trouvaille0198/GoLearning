package oslearning

import (
	"fmt"
	"goLearning/utils"
	"math/rand"
	"time"
)

//  getWeightedChooser
//  @Description: 获取指令地址分布权重表
//  @return *utils.Chooser
func getWeightedChooser() *utils.Chooser {
	chooser, _ := utils.NewChooser(
		utils.NewChoice("front", 25),
		utils.NewChoice("end", 25),
		utils.NewChoice("prev", 50),
	)
	return chooser
}

// AddressStream
// @description 地址流
type AddressStream struct {
	length    int
	addresses []int
}

// NewAddressStream
// @param addressNum 指令个数
// @param virtualMemorySize 虚存大小 以k为单位
// @isRandom 地址流随机生成或使用真实进程的地址流
func NewAddressStream(addressNum int, virtualMemorySize int, isRandom bool) *AddressStream {
	addresses := make([]int, addressNum)
	virtualMemorySize *= 1024
	if isRandom {
		// 随机生成地址流
		rs := rand.New(rand.NewSource(time.Now().UTC().UnixNano())) // 随机数种子
		chooser := getWeightedChooser()                             // 权重提取器

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
	} else {
		// TODO 采用真实地址流
	}
	return &AddressStream{length: addressNum, addresses: addresses}
}

// GetAddressPages
// @description 获取指令地址对应的页号
// @receiver a
// @param pageSize 页大小 以k为单位
// @return []int 指令的页号列表
func (a *AddressStream) GetAddressPages(pageSize int) []int {
	addressPages := make([]int, len(a.addresses))
	for i, address := range a.addresses {
		addressPages[i] = address/(pageSize*1024) + 1
	}
	return addressPages
}

// PageReplacer
// @description 请求页式管理 其中实现了多种页面置换算法
type PageReplacer struct {
	pageSize     int
	internalSize int
	addressPages []int
}

// NewPageReplacer
// @param addressNum 指令个数
// @param virtualMemorySize 虚存大小 以k为单位
// @param pageSize 页大小 以k为单位
// @param internalSize 内存能容纳页的个数
func NewPageReplacer(addressNum int, virtualMemorySize int, pageSize int, internalSize int, isRandom bool) *PageReplacer {
	addressStream := NewAddressStream(addressNum, virtualMemorySize, isRandom)
	addressPage := addressStream.GetAddressPages(pageSize)
	return &PageReplacer{pageSize: pageSize, addressPages: addressPage, internalSize: internalSize}
}

func (r PageReplacer) OPT() float32 {
	internalPages := make([]int, 0) // 记录内存中的页面
	failCount := 0                  // 未命中次数
	for i, pageNo := range r.addressPages {
		if !utils.IsContainInt(internalPages, pageNo) {
			// 未命中
			failCount++
			if len(internalPages) == r.internalSize {
				// 内存已满 欲淘汰一页
				latestPageIndex := -1 // 记录内存页号列表中在未来最迟访问的页号索引

				internalPagesCopy := make([]int, r.internalSize)
				copy(internalPagesCopy, internalPages)
				for _, afterPageNo := range r.addressPages[i+1:] {
					if utils.IsContainInt(internalPagesCopy, afterPageNo) {
						latestPageIndex, _ = utils.GetIndexInt(internalPages, afterPageNo) // 记录当前最迟访问到的页号索引
						internalPagesCopy, _ = utils.DeleteByElemInt(internalPagesCopy, afterPageNo)
					}
				}
				if len(internalPagesCopy) != 0 {
					// 内存中存在一页 其之后都不会被访问到
					internalPages, _ = utils.DeleteByElemInt(internalPages, internalPagesCopy[0])
				} else {
					internalPages, _ = utils.DeleteByIndexInt(internalPages, latestPageIndex)
				}
			}
			internalPages = append(internalPages, pageNo)
		}
	}
	return 1 - float32(failCount)/float32(len(r.addressPages))
}

func (r PageReplacer) LRU() float32 {
	internalPages := make([]int, 0) // 记录内存中的页面
	failCount := 0                  // 未命中次数
	for _, pageNo := range r.addressPages {
		if !utils.IsContainInt(internalPages, pageNo) {
			// 未命中
			failCount++
			if len(internalPages) == r.internalSize {
				// 内存已满 将栈底页面淘汰
				internalPages, _ = utils.DeleteByIndexInt(internalPages, 0)
			}
			internalPages = append(internalPages, pageNo)
		} else {
			// 命中 将命中页面提升至栈顶
			internalPages, _ = utils.DeleteByElemInt(internalPages, pageNo)
			internalPages = append(internalPages, pageNo)
		}
	}
	return 1 - float32(failCount)/float32(len(r.addressPages))
}

func (r PageReplacer) FIFO() float32 {
	internalPages := make([]int, 0) // 记录内存中的页面
	failCount := 0                  // 未命中次数
	for _, pageNo := range r.addressPages {
		if !utils.IsContainInt(internalPages, pageNo) {
			// 未命中
			failCount++
			if len(internalPages) == r.internalSize {
				// 内存已满 将队头页面淘汰
				internalPages, _ = utils.DeleteByIndexInt(internalPages, 0)
			}
			internalPages = append(internalPages, pageNo)
		}
	}
	return 1 - float32(failCount)/float32(len(r.addressPages))
}

// DisplayResult
// @description 展示几种页面置换算法的命中率
func (r PageReplacer) DisplayResult() {
	optHitPercent := r.OPT()
	lruHitPercent := r.LRU()
	fifoHitPercent := r.FIFO()

	fmt.Printf("%-10d %-10d %-10.2f %-10.2f %-10.2f \n",
		r.pageSize, r.internalSize, optHitPercent*100, lruHitPercent*100, fifoHitPercent*100)
}
