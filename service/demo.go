package service

import (
	"fmt"

	"github.com/knight0zh/demo_config/config"
	"github.com/knight0zh/demo_server/models"
)

func DemoService() {
	var o models.TbDemo
	fmt.Println(config.OmsMysql.Model(&o).First(&o).Error)
}
