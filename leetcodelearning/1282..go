package leetcodelearning

type room []int
type rooms []room

func groupThePeople(groupSizes []int) (res [][]int) {
	hash := map[int]rooms{} // 房间大小:房间数组
	for i, size := range groupSizes {
		if _, ok := hash[size]; !ok {
			hash[size] = rooms{room{i}}
		} else {
			rooms := hash[size]
			roomsAllFull := true
			for j := 0; j < len(rooms); j++ {
				if len(rooms[j]) < size {
					// 没满 把i放进去
					rooms[j] = append(rooms[j], i)
					roomsAllFull = false
					break
				}
			}
			if roomsAllFull {
				// 房间都满了 再加一个
				hash[size] = append(hash[size], room{i})
			}
		}
	}
	for _, rooms := range hash {
		for _, room := range rooms {
			res = append(res, room)
		}
	}
	return res
}
