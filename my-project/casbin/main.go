package casbin

import (
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
	// 模型和策略文件路径
	modelPath := "./model.conf"
	policyPath := "./policy.csv"

	// 加载模型和策略
	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		log.Fatalf("Failed to load Casbin enforcer: %v", err)
	}

	// 添加角色
	_, err = enforcer.AddGroupingPolicy("user1", "admin", "read")
	if err != nil {
		log.Fatalf("Failed to add group: %v", err)
	}
	// 保存模型和策略
	err = enforcer.SavePolicy()
	if err != nil {
		log.Fatalf("Failed to save Casbin enforcer: %v", err)
	}
	// 执行权限检查
	result, err := enforcer.Enforce("user1", "admin", "read")
	if err != nil {
		log.Fatalf("Failed to enforce permission: %v", err)
	}
	if result {
		log.Println("Permission granted")
	} else {
		log.Println("Permission denied")
	}

}
