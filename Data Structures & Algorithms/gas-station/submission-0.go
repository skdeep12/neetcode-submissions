func canCompleteCircuit(gas []int, cost []int) int {
    start := 0
	fuel := 0

	totalFuel := 0
	totalCost := 0
	for i:=0;i<len(cost);i+=1{
		totalCost += cost[i]
		totalFuel += gas[i]
		fuel += gas[i] - cost[i]
		if fuel < 0 {
			start = i+1
			fuel = 0
		}
	}
	if totalCost > totalFuel{
		return -1
	}
	return start
}
