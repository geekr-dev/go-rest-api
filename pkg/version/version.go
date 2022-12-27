package version

import (
	"fmt"
	"runtime"
)

// 版本信息包含的内容
type Info struct {
	GitTag       string `json:"gitTag"`       // Git 标签
	GitCommit    string `json:"gitCommit"`    // Git 提交
	GitTreeState string `json:"gitTreeState"` // 版本树状态
	BuildDate    string `json:"buildDate"`    // 构建日期
	GoVersion    string `json:"goVersion"`    // Go 版本
	Compiler     string `json:"compiler"`     // 编译器
	Platform     string `json:"platform"`     // 平台架构
}

func (i Info) String() string {
	return i.GitTag
}

func Get() Info {
	return Info{
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
