package deps

import (
	"fmt"
	"krillin-ai/config"
	"krillin-ai/internal/storage"
	"krillin-ai/log"
	"krillin-ai/pkg/util"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"go.uber.org/zap"
)

func CheckDependency() error {
	err := checkAndDownloadFfmpeg()
	if err != nil {
		log.GetLogger().Error("ffmpeg环境准备失败", zap.Error(err))
		return err
	}
	err = checkAndDownloadFfprobe()
	if err != nil {
		log.GetLogger().Error("ffprobe环境准备失败", zap.Error(err))
		return err
	}
	err = checkAndDownloadYtDlp()
	if err != nil {
		log.GetLogger().Error("yt-dlp环境准备失败", zap.Error(err))
		return err
	}
	if config.Conf.Transcribe.Provider == "fasterwhisper" {
		err = checkFasterWhisper()
		if err != nil {
			log.GetLogger().Error("fasterwhisper环境准备失败", zap.Error(err))
			return err
		}
		err = checkModel("fasterwhisper")
		if err != nil {
			log.GetLogger().Error("本地模型环境准备失败", zap.Error(err))
			return err
		}
	}
	if config.Conf.Transcribe.Provider == "whisperkit" {
		if err = checkWhisperKit(); err != nil {
			log.GetLogger().Error("whisperkit环境准备失败", zap.Error(err))
			return err
		}
		err = checkModel("whisperkit")
		if err != nil {
			log.GetLogger().Error("本地模型环境准备失败", zap.Error(err))
			return err
		}
	}
	if config.Conf.Transcribe.Provider == "whisperx" {
		err = checkWhisperX()
		if err != nil {
			log.GetLogger().Error("whisperx环境准备失败", zap.Error(err))
			return err
		}
		err = checkModel("whisperx")
		if err != nil {
			log.GetLogger().Error("本地模型环境准备失败", zap.Error(err))
			return err
		}
	}
	if config.Conf.Transcribe.Provider == "whispercpp" {
		if err = checkWhispercpp(); err != nil {
			log.GetLogger().Error("whispercpp环境准备失败", zap.Error(err))
			return err
		}
		err = checkModel("whispercpp")
		if err != nil {
			log.GetLogger().Error("whispercpp本地模型环境准备失败", zap.Error(err))
			return err
		}
	}

	return nil
}

// 检测并安装ffmpeg
func checkAndDownloadFfmpeg() error {
	// 检查ffmpeg是否已经安装
	_, err := exec.LookPath("ffmpeg")
	if err == nil {
		log.GetLogger().Info("已找到ffmpeg")
		storage.FfmpegPath = "ffmpeg"
		return nil
	}

	ffmpegBinFilePath := "./bin/ffmpeg"
	if runtime.GOOS == "windows" {
		ffmpegBinFilePath += ".exe"
	}
	// 先前下载过的
	if _, err = os.Stat(ffmpegBinFilePath); err == nil {
		log.GetLogger().Info("已找到ffmpeg")
		storage.FfmpegPath = ffmpegBinFilePath
		return nil
	}

	log.GetLogger().Info("没有找到ffmpeg，即将开始自动安装")
	// 确保./bin目录存在
	err = os.MkdirAll("./bin", 0755)
	if err != nil {
		log.GetLogger().Error("创建./bin目录失败", zap.Error(err))
		return err
	}

	var ffmpegURL string
	if runtime.GOOS == "linux" {
		ffmpegURL = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/ffmpeg-6.1-linux-64.zip"
	} else if runtime.GOOS == "darwin" {
		ffmpegURL = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/ffmpeg-6.1-macos-64.zip"
	} else if runtime.GOOS == "windows" {
		ffmpegURL = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/ffmpeg-6.1-win-64.zip"
	} else {
		log.GetLogger().Error("不支持你当前的操作系统", zap.String("当前系统", runtime.GOOS))
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	// 下载
	ffmpegDownloadPath := "./bin/ffmpeg.zip"
	err = util.DownloadFile(ffmpegURL, ffmpegDownloadPath, config.Conf.App.Proxy)
	if err != nil {
		log.GetLogger().Error("下载ffmpeg失败", zap.Error(err))
		return err
	}
	err = util.Unzip(ffmpegDownloadPath, "./bin")
	if err != nil {
		log.GetLogger().Error("解压ffmpeg失败", zap.Error(err))
		return err
	}
	log.GetLogger().Info("ffmpeg解压成功")

	if runtime.GOOS != "windows" {
		err = os.Chmod(ffmpegBinFilePath, 0755)
		if err != nil {
			log.GetLogger().Error("设置文件权限失败", zap.Error(err))
			return err
		}
	}

	storage.FfmpegPath = ffmpegBinFilePath
	log.GetLogger().Info("ffmpeg安装完成", zap.String("路径", ffmpegBinFilePath))

	return nil
}

// 检测并安装ffprobe
func checkAndDownloadFfprobe() error {
	// 检查检测并安装ffprobe是否已经安装
	_, err := exec.LookPath("ffprobe")
	if err == nil {
		log.GetLogger().Info("已找到ffprobe")
		storage.FfprobePath = "ffprobe"
		return nil
	}

	ffprobeBinFilePath := "./bin/ffprobe"
	if runtime.GOOS == "windows" {
		ffprobeBinFilePath += ".exe"
	}
	// 先前下载过的
	if _, err = os.Stat(ffprobeBinFilePath); err == nil {
		log.GetLogger().Info("已找到ffprobe")
		storage.FfprobePath = ffprobeBinFilePath
		return nil
	}

	log.GetLogger().Info("没有找到ffprobe，即将开始自动安装")
	// 确保./bin目录存在
	err = os.MkdirAll("./bin", 0755)
	if err != nil {
		log.GetLogger().Error("创建./bin目录失败", zap.Error(err))
		return err
	}

	var ffprobeURL string
	if runtime.GOOS == "linux" {
		ffprobeURL = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/ffprobe-6.1-linux-64.zip"
	} else if runtime.GOOS == "darwin" {
		ffprobeURL = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/ffprobe-6.1-macos-64.zip"
	} else if runtime.GOOS == "windows" {
		ffprobeURL = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/ffprobe-6.1-win-64.zip"
	} else {
		log.GetLogger().Error("不支持你当前的操作系统", zap.String("当前系统", runtime.GOOS))
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	// 下载
	ffprobeDownloadPath := "./bin/ffprobe.zip"
	err = util.DownloadFile(ffprobeURL, ffprobeDownloadPath, config.Conf.App.Proxy)
	if err != nil {
		log.GetLogger().Error("下载ffprobe失败", zap.Error(err))
		return err
	}
	err = util.Unzip(ffprobeDownloadPath, "./bin")
	if err != nil {
		log.GetLogger().Error("解压ffprobe失败", zap.Error(err))
		return err
	}
	log.GetLogger().Info("ffprobe解压成功")

	if runtime.GOOS != "windows" {
		err = os.Chmod(ffprobeBinFilePath, 0755)
		if err != nil {
			log.GetLogger().Error("设置文件权限失败", zap.Error(err))
			return err
		}
	}

	storage.FfprobePath = ffprobeBinFilePath
	log.GetLogger().Info("ffprobe安装完成", zap.String("路径", ffprobeBinFilePath))

	return nil
}

// 检测并安装yt-dlp
func checkAndDownloadYtDlp() error {
	_, err := exec.LookPath("yt-dlp")
	if err == nil {
		log.GetLogger().Info("已找到yt-dlp")
		storage.YtdlpPath = "yt-dlp"
		return nil
	}

	ytdlpBinFilePath := "./bin/yt-dlp"
	if runtime.GOOS == "windows" {
		ytdlpBinFilePath += ".exe"
	}
	// 先前下载过的
	if _, err = os.Stat(ytdlpBinFilePath); err == nil {
		log.GetLogger().Info("已找到ytdlp")
		storage.YtdlpPath = ytdlpBinFilePath
		return nil
	}

	log.GetLogger().Info("没有找到yt-dlp，即将开始自动安装")
	err = os.MkdirAll("./bin", 0755)
	if err != nil {
		log.GetLogger().Error("创建./bin目录失败", zap.Error(err))
		return err
	}

	var ytDlpURL string
	if runtime.GOOS == "linux" {
		ytDlpURL = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/yt-dlp_linux"
	} else if runtime.GOOS == "darwin" {
		ytDlpURL = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/yt-dlp_macos"
	} else if runtime.GOOS == "windows" {
		ytDlpURL = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/yt-dlp.exe"
	} else {
		log.GetLogger().Error("不支持你当前的操作系统", zap.String("当前系统", runtime.GOOS))
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	err = util.DownloadFile(ytDlpURL, ytdlpBinFilePath, config.Conf.App.Proxy)
	if err != nil {
		log.GetLogger().Error("下载yt-dlp失败", zap.Error(err))
		return err
	}

	if runtime.GOOS != "windows" {
		err = os.Chmod(ytdlpBinFilePath, 0755)
		if err != nil {
			log.GetLogger().Error("设置文件权限失败", zap.Error(err))
			return err
		}
	}

	storage.YtdlpPath = ytdlpBinFilePath
	log.GetLogger().Info("yt-dlp安装完成", zap.String("路径", ytdlpBinFilePath))

	return nil
}

// 检测faster whisper
func checkFasterWhisper() error {
	var (
		filePath string
		err      error
	)
	if runtime.GOOS == "windows" {
		filePath = "./bin/faster-whisper/Faster-Whisper-XXL/faster-whisper-xxl.exe"
	} else if runtime.GOOS == "linux" {
		filePath = "./bin/faster-whisper/Whisper-Faster-XXL/whisper-faster-xxl"
	} else {
		return fmt.Errorf("fasterwhisper不支持你当前的操作系统: %s，请选择其它transcription provider", runtime.GOOS)
	}
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		log.GetLogger().Info("没有找到faster-whisper，即将开始自动下载，文件较大请耐心等待")
		err = os.MkdirAll("./bin", 0755)
		if err != nil {
			log.GetLogger().Error("创建./bin目录失败", zap.Error(err))
			return err
		}
		var downloadUrl string
		if runtime.GOOS == "windows" {
			downloadUrl = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/Faster-Whisper-XXL_r194.5_windows.zip"
		} else {
			downloadUrl = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/Faster-Whisper-XXL_r192.3.1_linux.zip"
		}
		err = util.DownloadFile(downloadUrl, "./bin/faster-whisper.zip", config.Conf.App.Proxy)
		if err != nil {
			log.GetLogger().Error("下载faster-whisper失败", zap.Error(err))
			return err
		}
		log.GetLogger().Info("开始解压faster-whisper")
		err = util.Unzip("./bin/faster-whisper.zip", "./bin/faster-whisper/")
		if err != nil {
			log.GetLogger().Error("解压faster-whisper失败", zap.Error(err))
			return err
		}
	}
	if runtime.GOOS != "windows" {
		err = os.Chmod(filePath, 0755)
		if err != nil {
			log.GetLogger().Error("设置文件权限失败", zap.Error(err))
			return err
		}
	}
	storage.FasterwhisperPath = filePath
	log.GetLogger().Info("faster-whisper检查完成", zap.String("路径", filePath))
	return nil
}

// 检测whisperkit
func checkWhisperKit() error {
	cmd := exec.Command("which", "whisperkit-cli")
	err := cmd.Run()
	if err != nil {
		log.GetLogger().Info("没有找到whisperkit-cli，即将开始自动安装")
		cmd = exec.Command("brew", "install", "whisperkit-cli")
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.GetLogger().Error("whisperkit-cli 安装失败", zap.String("info", string(output)), zap.Error(err))
			return err
		}
		log.GetLogger().Info("whisperkit-cli 安装成功")
	}
	storage.WhisperKitPath = "whisperkit-cli"
	log.GetLogger().Info("检测到whisperkit-cli已安装")
	return nil
}

// 检测whisperx
func checkWhisperX() error {
	var (
		filePath  string
		_filePath string
		err       error
	)
	if runtime.GOOS == "windows" {
		filePath = "whisperx"
		_filePath = ".\\bin\\whisperx\\.venv\\Scripts\\whisperx.exe"
	} else if runtime.GOOS == "linux" {
		filePath = "./bin/whisperx/.venv/bin/whisperx"
		_filePath = "./bin/whisperx/.venv/bin/whisperx"
	} else {
		return fmt.Errorf("WhisperX不支持你当前的操作系统: %s，请选择WhisperKit", runtime.GOOS)
	}

	if _, err = os.Stat(_filePath); os.IsNotExist(err) {
		var cmd *exec.Cmd
		// TODO: 下载压缩包
		// log.GetLogger().Info("没有找到WhisperX，即将开始自动下载，文件较大请耐心等待")
		// err = os.MkdirAll("./bin", 0755)
		// if err != nil {
		// 	log.GetLogger().Error("创建./bin目录失败", zap.Error(err))
		// 	return err
		// }
		// var downloadUrl string
		// if runtime.GOOS == "windows" {
		// 	downloadUrl = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/WhisperX_win.zip"
		// } else if runtime.GOOS == "darwin" {
		// 	downloadUrl = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/WhisperX_linux.zip"
		// } else {
		// 	downloadUrl = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/WhisperX_mac.zip"
		// }
		// err = util.DownloadFile(downloadUrl, "./bin/WhisperX.zip", config.Conf.App.Proxy)
		// if err != nil {
		// 	log.GetLogger().Error("下载WhisperX失败", zap.Error(err))
		// 	return err
		// }
		log.GetLogger().Info("开始解压WhisperX")
		err = util.Unzip("./bin/WhisperX.zip", "./bin/whisperx/")
		if err != nil {
			log.GetLogger().Error("解压WhisperX失败", zap.Error(err))
			return err
		}
		if runtime.GOOS == "windows" {
			cmd = exec.Command(".\\bin\\whisperx\\python\\python.exe", "-m", "venv", ".\\bin\\whisperx\\.venv")
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.GetLogger().Error("创建python虚拟环境失败", zap.String("info", string(output)), zap.Error(err))
				return err
			}
			cmd = exec.Command(".\\bin\\whisperx\\.venv\\Scripts\\activate", "&&", "pip", "install", "-r", ".\\bin\\whisperx\\requirements_win.txt")
			cmd.CombinedOutput()
		} else {
			os.Chmod("./bin/whisperx/python/bin/python3.12", 0755)
			os.Chmod("./bin/whisperx/install.sh", 0755)
			log.GetLogger().Info("开始安装WhisperX")
			cmd = exec.Command("bash", "./bin/whisperx/install.sh")
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.GetLogger().Error("WhisperX 安装失败", zap.String("info", string(output)), zap.Error(err))
				return err
			}
		}
		log.GetLogger().Info("WhisperX 安装成功")
	}

	storage.WhisperXPath = filePath
	log.GetLogger().Info("WhisperX检查完成", zap.String("路径", _filePath))
	return nil
}

// 检测whispercpp
func checkWhispercpp() error {
	var (
		filePath string
		err      error
	)
	if runtime.GOOS == "windows" {
		filePath = filepath.Join("bin", "whispercpp", "whisper-cli.exe")
	} else {
		return fmt.Errorf("whisper.cpp不支持你当前的操作系统: %s，请选择其它transcription provider", runtime.GOOS)
	}
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		log.GetLogger().Info("没有找到whispercpp，即将开始自动下载，文件较大请耐心等待")
		err = os.MkdirAll("bin", 0755)
		if err != nil {
			log.GetLogger().Error("创建./bin目录失败", zap.Error(err))
			return err
		}
		var downloadUrl string
		if runtime.GOOS == "windows" {
			downloadUrl = "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/whispercpp-windows-cuda.zip"
		}
		zipFilePath := filepath.Join("bin", "whispercpp-windows-cuda.zip")
		err = util.DownloadFile(downloadUrl, zipFilePath, config.Conf.App.Proxy)
		if err != nil {
			log.GetLogger().Error("下载whispercpp失败", zap.Error(err))
			return err
		}
		log.GetLogger().Info("开始解压whispercpp")
		err = util.Unzip(zipFilePath, filepath.Join("bin", "whispercpp")+string(filepath.Separator))
		if err != nil {
			log.GetLogger().Error("解压whispercpp失败", zap.Error(err))
			return err
		}
	}
	if runtime.GOOS != "windows" {
		err = os.Chmod(filePath, 0755)
		if err != nil {
			log.GetLogger().Error("设置文件权限失败", zap.Error(err))
			return err
		}
	}
	storage.WhispercppPath = filePath
	log.GetLogger().Info("whispercpp检查完成", zap.String("路径", filePath))
	return nil
}

// 检测本地模型
func checkModel(whisperType string) error {
	var err error
	if _, err = os.Stat("./models/whisperkit"); os.IsNotExist(err) {
		err = os.MkdirAll("./models/whisperkit", 0755)
		if err != nil {
			log.GetLogger().Error("创建./models目录失败", zap.Error(err))
			return err
		}
	}
	// 模型文件
	var model string
	var modelPath string // cli中使用的model path
	switch whisperType {
	case "fasterwhisper":
		model = config.Conf.Transcribe.Fasterwhisper.Model
		modelPath = fmt.Sprintf("./models/faster-whisper-%s/model.bin", model)
		if _, err = os.Stat(modelPath); os.IsNotExist(err) {
			// 下载
			log.GetLogger().Info(fmt.Sprintf("没有找到模型文件%s,即将开始自动下载", modelPath))
			downloadUrl := fmt.Sprintf("https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/faster-whisper-%s.zip", model)
			err = util.DownloadFile(downloadUrl, fmt.Sprintf("./models/faster-whisper-%s.zip", model), config.Conf.App.Proxy)
			if err != nil {
				log.GetLogger().Error("下载fasterwhisper模型失败", zap.Error(err))
				return err
			}
			err = util.Unzip(fmt.Sprintf("./models/faster-whisper-%s.zip", model), fmt.Sprintf("./models/faster-whisper-%s/", model))
			if err != nil {
				log.GetLogger().Error("解压模型失败", zap.Error(err))
				return err
			}
			log.GetLogger().Info("模型下载完成", zap.String("路径", modelPath))
		}
	//case "whisperx":
	//	// TODO: upload models
	//	model = config.Conf.Transcribe.Whisperx.Model
	//	modelDir := fmt.Sprintf("./models/whisperx/models--Systran--faster-whisper-%s", model)
	//	if _, err = os.Stat(modelDir); os.IsNotExist(err) {
	//		log.GetLogger().Info(fmt.Sprintf("没有找到WhisperX模型%s,即将开始自动下载", modelDir))
	//		// downloadUrl := fmt.Sprintf("https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/WhisperX_models_%s.zip", model)
	//		// err = util.DownloadFile(downloadUrl, fmt.Sprintf("./models/WhisperX_models_%s.zip", model), config.Conf.App.Proxy)
	//		// if err != nil {
	//		// 	log.GetLogger().Info("下载WhisperX模型失败", zap.Error(err))
	//		// 	return err
	//		// }
	//		err = util.Unzip(fmt.Sprintf("./models/WhisperX_models_%s.zip", model), "./models/whisperx/")
	//		if err != nil {
	//			log.GetLogger().Error("解压模型失败", zap.Error(err))
	//			return err
	//		}
	//		log.GetLogger().Info("WhisperX模型下载完成", zap.String("路径", modelDir))
	//	}
	case "whispercpp":
		model = config.Conf.Transcribe.Whispercpp.Model
		modelPath = fmt.Sprintf("./models/whispercpp/ggml-%s.bin", model)
		if _, err = os.Stat(modelPath); os.IsNotExist(err) {
			log.GetLogger().Info(fmt.Sprintf("没有找到whisper.cpp模型%s,即将开始自动下载", modelPath))
			downloadUrl := fmt.Sprintf("https://gitcode.com/hf_mirrors/ai-gitcode/whisper.cpp/blob/main/ggml-%s.bin", model)
			err = util.DownloadFile(downloadUrl, fmt.Sprintf("./models/whispercpp/ggml-%s.bin", model), config.Conf.App.Proxy)
			if err != nil {
				log.GetLogger().Info("下载whisper.cpp模型失败", zap.Error(err))
				return err
			}
			log.GetLogger().Info("whisper.cpp模型下载完成", zap.String("路径", modelPath))
		}
	case "whisperkit":
		model = config.Conf.Transcribe.Whisperkit.Model
		modelPath = fmt.Sprintf("./models/whisperkit/openai_whisper-%s", model)
		files, _ := os.ReadDir(modelPath)
		if len(files) == 0 {
			log.GetLogger().Info("没有找到whisperkit模型，即将开始自动下载")
			downloadUrl := "https://modelscope.cn/models/Maranello/KrillinAI_dependency_cn/resolve/master/whisperkit-large-v2.zip"
			err = util.DownloadFile(downloadUrl, "./models/whisperkit/openai_whisper-large-v2.zip", config.Conf.App.Proxy)
			if err != nil {
				log.GetLogger().Info("下载whisperkit模型失败", zap.Error(err))
				return err
			}
			err = util.Unzip("./models/whisperkit/openai_whisper-large-v2.zip", "./models/whisperkit/")
			if err != nil {
				log.GetLogger().Error("解压whisperkit模型失败", zap.Error(err))
				return err
			}
			log.GetLogger().Info("whisperkit模型下载完成", zap.String("路径", modelPath))
		}
	}

	log.GetLogger().Info("模型检查完成", zap.String("路径", modelPath))
	return nil
}
