package main_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotalservices/cfbackup"
	. "github.com/pivotalservices/cfops-redis-plugin"
	"github.com/pivotalservices/cfops/plugin/cfopsplugin"
	"github.com/pivotalservices/cfops/tileregistry"
	"github.com/pivotalservices/gtils/command"
)

var _ = Describe("Given RedisPlugin", func() {
	var redisPlugin *RedisPlugin
	Describe("given a Meta() method", func() {
		Context("called on a plugin with valid meta data", func() {
			var meta cfopsplugin.Meta
			BeforeEach(func() {
				redisPlugin = NewRedisPlugin()
				meta = redisPlugin.GetMeta()
			})

			It("then it should return a meta data object with all required fields", func() {
				Ω(meta.Name).ShouldNot(BeEmpty())
			})
		})
	})
	testInstallationSettings("./fixtures/installation-settings-redis.json")
})

func testInstallationSettings(installationSettingsPath string) {
	var redisPlugin *RedisPlugin
	Describe(fmt.Sprintf("given a installationSettingsFile %s", installationSettingsPath), func() {
		Describe("given a Backup() method", func() {
			Context("when called on a properly setup RedisPlugin object", func() {
				var err error
				var controlTmpDir string
				var counter int
				BeforeEach(func() {
					controlTmpDir, _ = ioutil.TempDir("", "unit-test")
					redisPlugin = &RedisPlugin{
						Meta: cfopsplugin.Meta{
							Name: "redis-tile",
						},
						GetScriptBackup: func(sshConfig command.SshConfig, outputFileName, scriptName string) (err error) {
							counter++
							return
						},
					}
					configParser := cfbackup.NewConfigurationParser(installationSettingsPath)
					pivotalCF := cfopsplugin.NewPivotalCF(configParser.InstallationSettings, tileregistry.TileSpec{
						ArchiveDirectory: controlTmpDir,
					})
					redisPlugin.Setup(pivotalCF)
					err = redisPlugin.Backup()
				})

				AfterEach(func() {
					os.RemoveAll(controlTmpDir)
				})

				It("then it should have created right number of archive files", func() {
					Ω(err).ShouldNot(HaveOccurred())
					Ω(counter).Should(BeEquivalentTo(6))
				})
			})
		})

		Describe("given a Setup() method", func() {
			Context("when called with a PivotalCF containing a Redis tile", func() {
				var pivotalCF cfopsplugin.PivotalCF
				BeforeEach(func() {
					configParser := cfbackup.NewConfigurationParser(installationSettingsPath)
					pivotalCF = cfopsplugin.NewPivotalCF(configParser.InstallationSettings, tileregistry.TileSpec{})
					redisPlugin.Setup(pivotalCF)
				})

				It("then it should extract Mysql username required for backup/restore", func() {
					Ω(redisPlugin.PivotalCF).ShouldNot(BeNil())
				})

			})
		})
	})
}

func testRunBackupScript(sshConfig command.SshConfig, outputFileName, scriptName string) (err error) {

	return
}
func IsEmpty(name string) bool {
	f, err := os.Open(name)
	if err != nil {
		return false
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true
	}
	return false // Either not empty or error, suits both cases
}
