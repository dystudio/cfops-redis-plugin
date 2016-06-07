package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pivotalservices/cfbackup"
	"github.com/pivotalservices/cfops-redis-plugin/generated"
	cfopsplugin "github.com/pivotalservices/cfops/plugin/cfopsplugin"
	"github.com/pivotalservices/gtils/command"
	"github.com/pivotalservices/gtils/osutils"
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
		lo.G.Debug("starting backup of shared plan")
		if err = s.GetRunScript(sshConfig, "scripts/backupShared.sh"); err == nil {
			s.GetTarFile(sshConfig, sharedPlanOutputFileName, "cd /var/vcap/store/cf-redis-broker/ && tar cz redis-data")
		}
		lo.G.Debug("done backup of shared plan")
		s.GetStateFileJSON(sshConfig)
	}

	if sshConfigs, err = s.getSSHConfig(dedicatedPlanJobName); err == nil {
		for _, sshConfig := range sshConfigs {
			ip := sshConfig.Host
			outputFileName := fmt.Sprintf(dedicatedPlanOutputFileName, ip)
			lo.G.Debug(fmt.Sprintf("starting backup of dedicated plan on %s", ip))
			if err = s.GetRunScript(sshConfig, "scripts/backupDedicated.sh"); err == nil {
				s.GetTarFile(sshConfig, outputFileName, "cd /var/vcap/store/ && tar cz redis")
			}
			lo.G.Debug(fmt.Sprintf("done backup of dedicated plan on %s", ip))
		}
	}
	lo.G.Debug("done backup of redis-tile", err)
	return
}

func (s *RedisPlugin) getStateFileJSON(sshConfig command.SshConfig) (err error) {
	var writer io.WriteCloser
	var remoteExecuter command.Executer
	if remoteExecuter, err = NewRemoteExecuter(sshConfig); err == nil {
		if writer, err = s.PivotalCF.NewArchiveWriter(statefileOutputFileName); err == nil {
			defer writer.Close()
			lo.G.Info("getting statefile.json file from ip ->", sshConfig.Host)
			err = remoteExecuter.Execute(writer, "cat /var/vcap/store/cf-redis-broker/statefile.json")
			lo.G.Info("done getting statefile.json file from ip ->", sshConfig.Host, err)
		}
	}
	return
}

func (s *RedisPlugin) uploadStateFileJSON(sshConfig command.SshConfig) (err error) {
	var reader io.ReadCloser
	if reader, err = s.PivotalCF.NewArchiveReader(statefileOutputFileName); err == nil {
		defer reader.Close()
		err = s.GetUploadFile(sshConfig, reader, "/var/vcap/store/cf-redis-broker/statefile.json")
	}
	return
}

func (s *RedisPlugin) getTarFile(sshConfig command.SshConfig, outputFileName string, cmd string) (err error) {
	var writer io.WriteCloser
	var remoteExecuter command.Executer
	if remoteExecuter, err = NewRemoteExecuter(sshConfig); err == nil {
		if writer, err = s.PivotalCF.NewArchiveWriter(outputFileName); err == nil {
			defer writer.Close()
			lo.G.Info("creating tar file from backup on ip ->", sshConfig.Host)
			err = remoteExecuter.Execute(writer, cmd)
			lo.G.Info("done creating tar file from backup on ip ->", sshConfig.Host, err)
		}
	}
	return
}

func (s *RedisPlugin) uploadFile(sshConfig command.SshConfig, lfile io.Reader, path string) (err error) {
	remoteOps := osutils.NewRemoteOperationsWithPath(sshConfig, path)
	err = remoteOps.UploadFile(lfile)
	return
}

func (s *RedisPlugin) runScript(sshConfig command.SshConfig, scriptName string) (err error) {
	var remoteExecuter command.Executer
	var writer io.WriteCloser
	var scriptBytes []byte
	if remoteExecuter, err = NewRemoteExecuter(sshConfig); err == nil {
		writer = os.Stdout
		if scriptBytes, err = generated.Asset(scriptName); err == nil {
			reader := strings.NewReader(string(scriptBytes))
			lo.G.Info("uploading script on ip ->", sshConfig.Host)
			if err = s.GetUploadFile(sshConfig, reader, remoteScriptPath); err == nil {
				lo.G.Info("Running script on ip ->", sshConfig.Host)
				var commandToRun = fmt.Sprintf("chmod +x %s && echo %s | sudo -S %s", remoteScriptPath, sshConfig.Password, remoteScriptPath)
				err = remoteExecuter.Execute(writer, commandToRun)
				lo.G.Info("Done running script on ip ->", sshConfig.Host, err)
			}
		}
	}
	return
}

//Restore - method to execute restore
func (s *RedisPlugin) Restore() (err error) {
	lo.G.Debug("starting restore of redis-tile")
	var sshConfigs []command.SshConfig
	var reader io.ReadCloser
	if sshConfigs, err = s.getSSHConfig(sharedPlanJobName); err == nil {
		//take first node to execute restore on
		sshConfig := sshConfigs[0]
		lo.G.Debug("starting restore of shared plan")
		if reader, err = s.PivotalCF.NewArchiveReader(sharedPlanOutputFileName); err == nil {
			defer reader.Close()
			if err = s.GetUploadFile(sshConfig, reader, remoteArchivePath); err == nil {
				err = s.GetRunScript(sshConfig, "scripts/restoreShared.sh")
			}
			s.UploadStateFileJSON(sshConfig)
		} else {
			lo.G.Info("Skipping restore of shared VM plan as file does not exist")
		}
		lo.G.Debug("done restore of shared plan")
	}

	if sshConfigs, err = s.getSSHConfig(dedicatedPlanJobName); err == nil {
		for _, sshConfig := range sshConfigs {
			ip := sshConfig.Host
			outputFileName := fmt.Sprintf(dedicatedPlanOutputFileName, ip)
			lo.G.Debug(fmt.Sprintf("starting restore of dedicated plan on %s using file %s", ip, outputFileName))
			if reader, err = s.PivotalCF.NewArchiveReader(outputFileName); err == nil {
				lo.G.Debug(fmt.Sprintf("uploading file %s to ip %s", outputFileName, ip))
				defer reader.Close()
				if err = s.GetUploadFile(sshConfig, reader, remoteArchivePath); err == nil {
					lo.G.Debug(fmt.Sprintf("running script on ip %s", ip))
					err = s.GetRunScript(sshConfig, "scripts/restoreDedicated.sh")
					lo.G.Debug(fmt.Sprintf("finished script on ip %s", ip), err)
				}
			} else {
				lo.G.Info(fmt.Sprintf("Skipping restore of dedicated VM plan on %s as file does not exist", ip))
			}
			lo.G.Debug(fmt.Sprintf("done restore of dedicated plan on %s", ip))
		}
	}
	lo.G.Debug("done restore of redis-tile", err)
	return
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
	pluginName                      = "redis-tile"
	productName                     = "p-redis"
	sharedPlanJobName               = "cf-redis-broker"
	sharedPlanOutputFileName        = pluginName + "-sharedVMPlan.tar"
	statefileOutputFileName         = pluginName + "-statefile.json"
	dedicatedPlanJobName            = "dedicated-node"
	dedicatedPlanOutputFileName     = pluginName + "-%s-dedicatedVMPlan.tar"
	remoteArchivePath               = "/var/vcap/store/tmp_backup/redis-tile.tar"
	remoteScriptPath                = "/var/vcap/store/tmp_backup/execute.sh"
	defaultSSHPort              int = 22
)

//NewRedisPlugin - Contructor helper
func NewRedisPlugin() *RedisPlugin {
	redisPlugin := &RedisPlugin{
		Meta: cfopsplugin.Meta{
			Name: pluginName,
		},
	}
	redisPlugin.GetRunScript = redisPlugin.runScript
	redisPlugin.GetUploadFile = redisPlugin.uploadFile
	redisPlugin.GetTarFile = redisPlugin.getTarFile
	redisPlugin.GetStateFileJSON = redisPlugin.getStateFileJSON
	redisPlugin.UploadStateFileJSON = redisPlugin.uploadStateFileJSON
	return redisPlugin
}

//RedisPlugin - structure
type RedisPlugin struct {
	PivotalCF            cfopsplugin.PivotalCF
	InstallationSettings cfbackup.InstallationSettings
	Meta                 cfopsplugin.Meta
	GetRunScript         func(command.SshConfig, string) error
	GetUploadFile        func(sshConfig command.SshConfig, lfile io.Reader, path string) error
	GetTarFile           func(sshConfig command.SshConfig, outputFileName string, cmd string) (err error)
	GetStateFileJSON     func(sshConfig command.SshConfig) (err error)
	UploadStateFileJSON  func(sshConfig command.SshConfig) (err error)
}
