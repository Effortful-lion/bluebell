package snowflake

import (
	"time"
	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

//初始化一个全局的node节点
func Init(startTime string, machineID int64)(err error){
	var st time.Time
	// 时间因子：从startTime开始
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

func GenID() int64{
	return node.Generate().Int64()
}