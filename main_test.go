package main_test

import (
	"io"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/cfops-redis-plugin"
	"github.com/pivotalservices/cfops/plugin/cfopsplugin"
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
})

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
