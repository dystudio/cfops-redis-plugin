package main

import (
	"github.com/pivotalservices/cfbackup"
	cfopsplugin "github.com/pivotalservices/cfops/plugin/cfopsplugin"
)

func main() {
	cfopsplugin.Start(NewRedisPlugin())
}

//GetMeta - method to provide metadata
func (s *RedisPlugin) GetMeta() (meta cfopsplugin.Meta) {
	meta = s.Meta
	return
}

//Setup - on setup method
func (s *RedisPlugin) Setup(pcf cfopsplugin.PivotalCF) (err error) {
	s.PivotalCF = pcf
	s.InstallationSettings = pcf.GetInstallationSettings()
	return
}

//Backup - method to execute backup
func (s *RedisPlugin) Backup() (err error) {
	panic("Backup not implemented")
}

//Restore - method to execute restore
func (s *RedisPlugin) Restore() (err error) {
	panic("Restore not implemented")
}

const (
	pluginName  = "redis-tile"
	productName = "p-redis"
)

//NewRedisPlugin - Contructor helper
func NewRedisPlugin() *RedisPlugin {
	return &RedisPlugin{
		Meta: cfopsplugin.Meta{
			Name: pluginName,
		},
	}
}

//RedisPlugin - structure
type RedisPlugin struct {
	PivotalCF            cfopsplugin.PivotalCF
	InstallationSettings cfbackup.InstallationSettings
	Meta                 cfopsplugin.Meta
}
