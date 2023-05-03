package snowflake

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1) // TODO
	if err != nil {
		log.Panicln(err)
	}
}

func Generate() uint {
	return uint(node.Generate().Int64())
}
