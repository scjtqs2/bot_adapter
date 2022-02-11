package config

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Path 配置文件路径
var Path = "config.yaml"

// defaultConfig 默认配置文件
//go:embed default_config.yml
var defaultConfig string

// Config 配置类
type Config struct {
	ServerType  ServerType          `json:"server_type" yaml:"server_type"`
	HTTPConfig  HTTPConfig          `json:"http_config" yaml:"http_config"`
	Plugins     []*PluginConfig     `json:"plugins" yaml:"plugins"`
	GrpcConfig  GrpcConfig          `json:"grpc_config" yaml:"grpc_config"`
	Permissions []*PermissionConfig `json:"permissions" yaml:"permissions"`
}

// GrpcConfig grpc的配置
type GrpcConfig struct {
	Addr                  string `json:"addr" yaml:"addr"`
	MaxConnectionAgeGrace int    `json:"max_connection_age_grace" yaml:"max_connection_age_grace"`
	MaxConnectionIdle     int    `json:"max_connection_idle" yaml:"max_connection_idle"`
	Timeout               int    `json:"timeout" yaml:"timeout"`
}

// HTTPConfig http的配置信息
type HTTPConfig struct {
	ServerAddr string `json:"server_addr" yaml:"server_addr"`
	Token      string `json:"token" yaml:"token"`
	LocalHost  string `json:"local_host" yaml:"local_host"` // 本地监听 addr
	Secret     string `json:"secret" yaml:"secret"`
	MaxTries   int    `json:"max_tries" yaml:"max_tries"` // 重试次数
	Timeout int `json:"timeout" yaml:"timeout"` // http请求cqhttp的超时时间 秒
}

// PluginConfig 插件配置列表
type PluginConfig struct {
	AppID            string              `json:"app_id" yaml:"app_id"`                         // 随便一个字符串就行，用来生成token的
	AppSecret        string              `json:"app_secret" yaml:"app_secret"`                 // 随便一个字符串就行，用来生成token的
	EncryptKey       string              `json:"encrypt_key" yaml:"encrypt_key"`               // 推送数据的加密key
	PostAddr         string              `json:"post_addr" yaml:"post_addr"`                   // 接收推送的地址
	PluginName       string              `json:"plugin_name" yaml:"plugin_name"`               // app名称
	PluginDesc       string              `json:"plugin_desc" yaml:"plugin_desc"`               // app描述
	RegisterKeyWords []*RegisterKeyWords `json:"register_key_words" yaml:"register_key_words"` // 注册用的拦截开头字符。注册了后，排在它后面的插件将无法收到对应消息的推送
}

// RegisterKeyWords 拦截触发器
type RegisterKeyWords struct {
	MsgType        string   `json:"msg_type" yaml:"msg_type"`                 // private,group
	PrefixKeyWords []string `json:"prefix_key_words" yaml:"prefix_key_words"` // 消息字符串的前缀注册器
}

// PermissionConfig 权限配置
type PermissionConfig struct {
	AppID           string   `json:"app_id" yaml:"app_id"`                     // 插件app的appid
	IsAdmin         bool     `json:"is_admin" yaml:"is_admin"`                 // 是否管理员，管理员具备所有权限。配置此项将忽略 api_permissions 和 push_permissions
	IsOnlyCqhttp    bool     `json:"is_only_cqhttp" yaml:"is_only_cqhttp"`     // 快读配置所有的cqhttp的api的权限。无 bot_adapter自己的接口权限。具备全部推送权限。配置此项忽略 api_permissions 和 push_permissions
	ApiPermissions  []string `json:"api_permissions" yaml:"api_permissions"`   // 将所有的权限字符串以数组的方式填入。前两个填false，在这里自定义定义权限
	PushPermissions []string `json:"push_permissions" yaml:"push_permissions"` // 推送权限配置
}

// generateConfig 生成默认配置文件
func generateConfig() {
	fmt.Println("未找到配置文件，正在为您生成配置文件中！")
	sb := strings.Builder{}
	sb.WriteString(defaultConfig)
	_ = os.WriteFile(Path, []byte(sb.String()), 0o644)
	fmt.Println("默认配置文件已生成，请修改 config.yml 后重新启动!")
	os.Exit(1)
}

// Save 将当前配置文件写入文件
func (c *Config) Save() error {
	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	_ = os.WriteFile(Path, b, 0o644)
	return nil
}

// Parse 从文件中读取配置信息
func Parse() *Config {
	file, err := os.ReadFile(Path)
	config := &Config{}
	if err == nil {
		err = yaml.Unmarshal(file, &config)
		if err != nil {
			log.Fatalf("配置文件不合法！err:%v", err)
		}
	} else {
		generateConfig()
	}
	/**
	todo os.GetEnv 替换配置文件
	*/
	if os.Getenv("HTTP_LISTEN") != "" {
		config.HTTPConfig.LocalHost = os.Getenv("HTTP_LISTEN")
	}
	if os.Getenv("HTTP_TOKEN") != "" {
		config.HTTPConfig.Token = os.Getenv("HTTP_TOKEN")
	}
	if os.Getenv("HTTP_SECRET") != "" {
		config.HTTPConfig.Secret = os.Getenv("HTTP_SECRET")
	}
	if os.Getenv("HTTP_CQHTTP_ADDR") != "" {
		config.HTTPConfig.ServerAddr = os.Getenv("HTTP_CQHTTP_ADDR")
	}

	return config
}
