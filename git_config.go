package git

import (
	"net/http"

	"github.com/abdfnx/gh/api"
	"github.com/abdfnx/gh/core/ghrepo"
	"github.com/abdfnx/gh/pkg/cmdutil"
	"github.com/abdfnx/gh/pkg/iostreams"
	"github.com/abdfnx/gh/core/ghinstance"
	"github.com/abdfnx/gh/pkg/cmd/factory"
)

type ConfStruct struct {
	HttpClient func() (*http.Client, error)
	IO         *iostreams.IOStreams
	BaseRepo   func() (ghrepo.Interface, error)
}

func f(f *cmdutil.Factory) ConfStruct {
	var opts = ConfStruct{
		HttpClient: f.HttpClient,
	}

	return opts
}

var httpClient, _ = f(factory.New()).HttpClient()

var apiClient = api.NewClientFromHTTP(httpClient)
var currentUser, _ = api.CurrentLoginName(apiClient, ghinstance.Default())

var cmdFactory = factory.New()
var configRootCmd = NewCmdConfigRoot(cmdFactory)
var cfg, _ = cmdFactory.Config()

var expandedArgs = []string{}

var cmd, _, _ = configRootCmd.Traverse(expandedArgs)

func GitConfig() string {
	if cmdutil.IsAuthCheckEnabled(cmd) && !cmdutil.CheckAuth(cfg) {
		return ":username"
	} else {
		return currentUser
	}
}

func GitConfigWithMsg(msg1 string, msg2 string) string {
	if cmd != nil && cmdutil.IsAuthCheckEnabled(cmd) && !cmdutil.CheckAuth(cfg) {
		return msg1 + ":username" + msg2
	} else {
		return msg1 + currentUser + msg2
	}
}
