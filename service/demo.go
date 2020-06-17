package service

import (
	"fmt"

	"github.com/knight0zh/demo_server/base"
	"github.com/knight0zh/demo_server/models"
)

func DemoService() {
	var o models.TbDemo
	fmt.Println(base.OmsMysql.Model(&o).First(&o).Error)
}
