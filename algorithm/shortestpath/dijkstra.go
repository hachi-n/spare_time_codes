package shortestpath

import (
	"fmt"
	"math"
)

type PathInfo struct {
	NodesMap map[string]map[string]int
}

type ShortestPath struct {
	Sum   int
	Nodes []string
}

func NewPathInfo(nodesMap map[string]map[string]int) *PathInfo {
	return &PathInfo{
		NodesMap: nodesMap,
	}

}

func (sp *ShortestPath) String() string {
	return fmt.Sprintf(
		"ShortestPath sum is %d \n Start to Goal Nodes %v",
		sp.Sum, sp.Nodes,
	)

}

func (p *PathInfo) FindShortestPath(startNode, endNode string) *ShortestPath {
	shortestPath := new(ShortestPath)
	statusMap := p.createStatusMap()

	statusMap[startNode] = 0
	currentNode := startNode
	shortestNode := startNode

	for currentNode != endNode {
		nextNodes := p.NodesMap[shortestNode]
		for k, v := range nextNodes {
			if statusMap[k] > v {
				statusMap[k] = v
			}
		}

		shortestPath.Sum += statusMap[currentNode]
		shortestPath.Nodes = append(shortestPath.Nodes, currentNode)
		delete(statusMap, currentNode)

		shortestNode, _ = findShortestPathKeyValue(statusMap)
	}

	shortestPath.Sum += statusMap[endNode]
	shortestPath.Nodes = append(shortestPath.Nodes, endNode)

	return shortestPath
}

func (p *PathInfo) createStatusMap() map[string]int {
	statusMap := make(map[string]int)
	for k := range p.NodesMap {
		statusMap[k] = math.MaxInt64
	}
	return statusMap
}

func findShortestPathKeyValue(m map[string]int) (string, int) {
	minimumKey := ""
	minimumValue := math.MaxInt64
	fmt.Println(m)

	for k, v := range m {
		if minimumValue >= v {
			minimumKey = k
			minimumValue = v
		}
	}

	if minimumKey == "" {
		// log.Fatal("system error: maximum weight is %d", math.MaxInt64)
	}

	return minimumKey, minimumValue
}
