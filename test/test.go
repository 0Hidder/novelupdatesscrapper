package main

import (
	"fmt"

	"github.com/0Hidder/novelupdatesscrapperv1/connection"
)

func main() {
	fmt.Println(connection.TestPrint())
	fmt.Println(connection.GetTitleFromString("https://www.novelupdates.com/series/skill-takers-world-domination-building-a-slave-harem-from-scratch/"))
}
