package settings

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
)

var Env = map[string]interface{}{
	"GinMode": "release",
	"DB":      map[string]string{"default": "qa:NTQ0NjU5YjU0@(rm-ks-qa.mysql.rds.aliyuncs.com:3306)/gaea?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true", "extra1": "tidb"},
	"Port":    8888,
	"Redis":   "redis://:@redis-master.redis.svc.cluster.local:6379/3",
}

func loadLocalEnvAsMap(filename string) map[string]interface{} {
	resultMap := make(map[string]interface{})
	content, _ := os.ReadFile(filename)
	if content == nil {
		return nil
	}
	err := yaml.Unmarshal(content, &resultMap)
	if err != nil {
		log.Fatal("加载env_settings.yaml失败")
	}
	return resultMap
}

func loadEnvironmentAsMap() map[string]interface{} {
	resultMap := make(map[string]interface{})
	envs := os.Environ()
	for _, e := range envs {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) != 2 {
			continue
		} else {
			resultMap[parts[0]] = parts[1]
		}
	}
	return resultMap
}

func updateEnv(m map[string]interface{}) {
	for k := range Env {
		if v, exist := m[k]; exist {
			Env[k] = v
		}
	}
}

func init() {
	updateEnv(loadLocalEnvAsMap("env_settings.yaml"))
	updateEnv(loadEnvironmentAsMap())
}
