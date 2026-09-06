import "slices"
func carFleet(target int, position []int, speed []int) int {
	// 0,1,7,4
	// 1,2,1,2
	// 10, 5, 3, 3

	// 7,4,1,0
	// 1,2,2,1
	// 3,3, 5, 10


	// <= top, then keep increasing
	cars := make([]Car, len(position))
	for i:=0;i<len(cars);i+=1{
		cars[i] = Car{position[i], speed[i]}
	}

	slices.SortFunc(cars, func(a, b Car) int {
		if a.position > b.position {
			return -1
		} else {
			return 1
		}
	})
	fmt.Println(cars)
	ans := 1
	currentTime := getTime(target, cars[0])
	for i:=1;i<len(cars);i+=1{
		time := getTime(target, cars[i])
		if time > currentTime {
			ans += 1
			currentTime = time
		}
	}
	return ans

}

func getTime(target int, c Car) float32 {
	diff := float32(target - c.position)
	return diff/float32(c.speed)
}

type Car struct {
	position int
	speed int
}
