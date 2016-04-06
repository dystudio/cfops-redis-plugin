package main

import (
	"io"

	"github.com/pivotalservices/cfbackup"
	"github.com/pivotalservices/cfops-redis-plugin/generated"
	cfopsplugin "github.com/pivotalservices/cfops/plugin/cfopsplugin"
	"github.com/pivotalservices/gtils/command"
	"github.com/xchapter7x/lo"
)

var (
	//NewRemoteExecuter -
	NewRemoteExecuter = command.NewRemoteExecutor
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
	lo.G.Debug("starting backup of redis-tile")
	var sshConfigs []command.SshConfig
	if sshConfigs, err = s.getSSHConfig(sharedPlanJobName); err == nil {
		//take first node to execute restore on
		sshConfig := sshConfigs[0]

		s.GetScriptBackup(sshConfig, sharedPlanOutputFileName, "scripts/backup_shared.sh")

	}
	lo.G.Debug("done backup of redis-tile", err)
	return
}

func (s *RedisPlugin) runBackupScript(sshConfig command.SshConfig, outputFileName, scriptName string) (err error) {
	var remoteExecuter command.Executer
	var writer io.WriteCloser
	var scriptBytes []byte
	if remoteExecuter, err = NewRemoteExecuter(sshConfig); err == nil {
		if writer, err = s.PivotalCF.NewArchiveWriter(outputFileName); err == nil {
			defer writer.Close()
			if scriptBytes, err = generated.Asset(scriptName); err == nil {
				err = remoteExecuter.Execute(writer, string(scriptBytes))
			}
		}
	}
	return
}

//Restore - method to execute restore
func (s *RedisPlugin) Restore() (err error) {
	panic("Restore not implemented")
}

func (s *RedisPlugin) getSSHConfig(jobName string) (sshConfig []command.SshConfig, err error) {
	var IPs []string
	var vmCredentials cfbackup.VMCredentials

	if IPs, err = s.InstallationSettings.FindIPsByProductAndJob(productName, jobName); err == nil {
		if vmCredentials, err = s.InstallationSettings.FindVMCredentialsByProductAndJob(productName, jobName); err == nil {
			for _, ip := range IPs {
				sshConfig = append(sshConfig, command.SshConfig{
					Username: vmCredentials.UserID,
					Password: vmCredentials.Password,
					Host:     ip,
					Port:     defaultSSHPort,
					SSLKey:   vmCredentials.SSLKey,
				})
			}
		}
	}
	return
}

const (
	pluginName               = "redis-tile"
	productName              = "p-redis"
	sharedPlanJobName        = "cf-redis-broker"
	sharedPlanOutputFileName = "sharedVMPlan.tar"

	defaultSSHPort int = 22
)

//NewRedisPlugin - Contructor helper
func NewRedisPlugin() *RedisPlugin {
	redisPlugin := &RedisPlugin{
		Meta: cfopsplugin.Meta{
			Name: pluginName,
		},
	}
	redisPlugin.GetScriptBackup = redisPlugin.runBackupScript
	return redisPlugin
}

//RedisPlugin - structure
type RedisPlugin struct {
	PivotalCF            cfopsplugin.PivotalCF
	InstallationSettings cfbackup.InstallationSettings
	Meta                 cfopsplugin.Meta
	GetScriptBackup      func(command.SshConfig, string, string) error
}
