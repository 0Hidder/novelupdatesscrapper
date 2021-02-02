package main

import (
	"fmt"

	"github.com/0Hidder/novelupdatesscrapperv1/connection"
	"github.com/0Hidder/novelupdatesscrapperv1/mongodb"
)

func main() {
	fmt.Println(connection.TestPrint())
	fmt.Println(connection.GetChapterFromString("https://www.novelupdates.com/series/skill-takers-world-domination-building-a-slave-harem-from-scratch/"))
	mongodb.Test()
	mongodb.Test1()
}
