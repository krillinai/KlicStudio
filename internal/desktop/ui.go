package desktop

import (
	"fmt"
	"image/color"
	"krillin-ai/config"
	"krillin-ai/internal/deps"
	"krillin-ai/log"
	"path/filepath"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"go.uber.org/zap"
)

// 创建配置界面
func CreateConfigTab(window fyne.Window) fyne.CanvasObject {
	// 创建页面标题
	pageTitle := TitleText("应用配置")

	// app 配置
	appGroup := createAppConfigGroup()
	localModelGroup := createLocalModelGroup()
	openaiGroup := createOpenAIConfigGroup()
	whisperGroup := createWhisperConfigGroup()
	aliyunOssGroup := createAliyunOSSConfigGroup()
	aliyunSpeechGroup := createAliyunSpeechConfigGroup()
	aliyunBailianGroup := createAliyunBailianConfigGroup()

	// 保存按钮
	saveButton := createSaveButton(window)

	// 创建一个背景效果
	background := canvas.NewRectangle(color.NRGBA{R: 248, G: 250, B: 253, A: 255})

	// 添加一些视觉分隔和间距
	spacer1 := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 0})
	spacer1.SetMinSize(fyne.NewSize(0, 10))
	spacer2 := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 0})
	spacer2.SetMinSize(fyne.NewSize(0, 10))

	// 创建滚动容器
	configContainer := container.NewVBox(
		container.NewPadded(pageTitle),
		spacer1,
		container.NewPadded(appGroup),
		container.NewPadded(localModelGroup),
		container.NewPadded(openaiGroup),
		container.NewPadded(whisperGroup),
		container.NewPadded(aliyunOssGroup),
		container.NewPadded(aliyunSpeechGroup),
		container.NewPadded(aliyunBailianGroup),
		spacer2,
		container.NewPadded(saveButton),
	)

	scroll := container.NewScroll(configContainer)

	// 使用一个Stack将背景和滚动内容组合
	configStack := container.NewStack(background, scroll)

	return container.NewPadded(configStack)
}

// 创建字幕任务界面
func CreateSubtitleTab(window fyne.Window) fyne.CanvasObject {
	sm := NewSubtitleManager(window)

	// 创建标题
	title := TitleText("视频翻译配音")

	// 创建视频输入区域
	videoInputContainer := createVideoInputContainer(sm)

	// 创建字幕设置区域
	subtitleSettingsCard := createSubtitleSettingsCard(sm)

	// 创建配音设置区域
	voiceSettingsCard := createVoiceSettingsCard(sm)

	// 创建字幕嵌入设置区域
	embedSettingsCard := createEmbedSettingsCard(sm)

	// 创建进度和下载区域
	progress, downloadContainer, tipsLabel := createProgressAndDownloadArea(sm)

	// 创建开始按钮
	startButton := createStartButton(window, sm, videoInputContainer, embedSettingsCard, progress, downloadContainer)
	startButtonContainer := container.NewHBox(layout.NewSpacer(), startButton, layout.NewSpacer())

	// 创建一个背景效果
	background := canvas.NewRectangle(color.NRGBA{R: 248, G: 250, B: 253, A: 255})

	// 添加一些视觉分隔和间距
	spacer1 := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 0})
	spacer1.SetMinSize(fyne.NewSize(0, 10))
	spacer2 := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 0})
	spacer2.SetMinSize(fyne.NewSize(0, 10))
	spacer3 := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 0})
	spacer3.SetMinSize(fyne.NewSize(0, 10))

	// 创建进度区域容器
	progressArea := container.NewVBox(progress)

	// 创建主布局
	mainContent := container.NewVBox(
		container.NewPadded(title),
		spacer1,
		container.NewPadded(videoInputContainer),
		container.NewPadded(subtitleSettingsCard),
		container.NewPadded(voiceSettingsCard),
		container.NewPadded(embedSettingsCard),
		spacer2,
		container.NewPadded(startButtonContainer),
		spacer3,
		progressArea,
		downloadContainer,
		tipsLabel,
	)

	scroll := container.NewScroll(mainContent)

	// 使用一个Stack将背景和滚动内容组合
	contentStack := container.NewStack(background, scroll)

	return container.NewPadded(contentStack)
}

// 创建应用配置组
func createAppConfigGroup() *fyne.Container {
	appSegmentDurationEntry := StyledEntry("字幕分段处理时长(分钟)")
	appSegmentDurationEntry.Bind(binding.IntToString(binding.BindInt(&config.Conf.App.SegmentDuration)))
	appSegmentDurationEntry.Validator = func(s string) error {
		val, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("请输入数字")
		}
		if val < 1 || val > 30 {
			return fmt.Errorf("请输入1-30之间的数字")
		}
		return nil
	}

	appTranslateParallelNumEntry := StyledEntry("翻译并行数量")
	appTranslateParallelNumEntry.Bind(binding.IntToString(binding.BindInt(&config.Conf.App.TranslateParallelNum)))
	appTranslateParallelNumEntry.Validator = func(s string) error {
		val, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("请输入数字")
		}
		if val < 1 || val > 10 {
			return fmt.Errorf("请输入1-10之间的数字")
		}
		return nil
	}

	appProxyEntry := StyledEntry("网络代理地址")
	appProxyEntry.Bind(binding.BindString(&config.Conf.App.Proxy))

	appTranscribeProviderEntry := StyledSelect([]string{"openai", "fasterwhisper", "whisperkit", "aliyun"}, func(s string) {
		config.Conf.App.TranscribeProvider = s
	})
	appTranscribeProviderEntry.SetSelected(config.Conf.App.TranscribeProvider)

	appLlmProviderEntry := StyledSelect([]string{"openai", "aliyun"}, func(s string) {
		config.Conf.App.LlmProvider = s
	})
	appLlmProviderEntry.SetSelected(config.Conf.App.LlmProvider)

	// 格式化表单项以使其更美观
	form := widget.NewForm(
		widget.NewFormItem("字幕分段处理时长(分钟)", appSegmentDurationEntry),
		widget.NewFormItem("翻译并行数量", appTranslateParallelNumEntry),
		widget.NewFormItem("网络代理地址", appProxyEntry),
		widget.NewFormItem("语音识别服务源", appTranscribeProviderEntry),
		widget.NewFormItem("LLM服务源", appLlmProviderEntry),
	)

	return GlassCard("应用配置", "基本设置参数", form)
}

// 创建本地模型配置组
func createLocalModelGroup() *fyne.Container {
	localModelFasterwhisperEntry := StyledSelect([]string{"tiny", "medium", "large-v2"}, func(s string) {
		config.Conf.LocalModel.Fasterwhisper = s
	})
	localModelFasterwhisperEntry.SetSelected(config.Conf.LocalModel.Fasterwhisper)

	localModelWhisperkitEntry := StyledSelect([]string{"large-v2"}, func(s string) {
		config.Conf.LocalModel.Whisperkit = s
	})
	localModelWhisperkitEntry.SetSelected(config.Conf.LocalModel.Whisperkit)

	form := widget.NewForm(
		widget.NewFormItem("Fasterwhisper模型", localModelFasterwhisperEntry),
		widget.NewFormItem("Whisperkit模型", localModelWhisperkitEntry),
	)

	return StyledCard("本地模型配置", form)
}

// 创建视频输入容器
func createVideoInputContainer(sm *SubtitleManager) *fyne.Container {
	// 视频输入方式选择
	inputTypeRadio := widget.NewRadioGroup([]string{"本地视频", "视频链接"}, nil)
	inputTypeRadio.Horizontal = true
	inputTypeContainer := container.NewHBox(
		widget.NewLabel("输入方式:"),
		inputTypeRadio,
	)

	// 视频链接输入框
	urlEntry := StyledEntry("请输入视频链接")
	urlEntry.Hide()
	urlEntry.OnChanged = func(text string) {
		sm.SetVideoUrl(text)
	}

	// 视频选择按钮
	selectButton := PrimaryButton("选择视频文件", theme.FolderOpenIcon(), sm.ShowFileDialog)
	selectedVideoLabel := widget.NewLabel("")
	selectedVideoLabel.Hide()

	// 设置视频选择回调
	sm.SetVideoSelectedCallback(func(path string) { // 设置视频地址+控制信息展示
		if path != "" {
			sm.SetVideoUrl(path)
			selectedVideoLabel.SetText("已选择: " + filepath.Base(path))
			selectedVideoLabel.Show()
		} else {
			selectedVideoLabel.Hide()
		}
	})

	// 视频输入容器
	videoInputContainer := container.NewVBox()
	videoInputContainer.Objects = []fyne.CanvasObject{selectButton, selectedVideoLabel}

	// 切换输入方式
	inputTypeRadio.SetSelected("本地视频")
	inputTypeRadio.OnChanged = func(value string) {
		if value == "本地视频" {
			urlEntry.Hide()
			selectButton.Show()
			selectedVideoLabel.Show()
			videoInputContainer.Objects = []fyne.CanvasObject{selectButton, selectedVideoLabel}
			sm.SetVideoUrl("")
		} else {
			selectButton.Hide()
			selectedVideoLabel.Hide()
			urlEntry.Show()
			videoInputContainer.Objects = []fyne.CanvasObject{urlEntry}
		}
		videoInputContainer.Refresh()
	}

	// 创建语言选择容器
	langContainer := container.NewGridWithColumns(2,
		container.NewHBox(
			widget.NewLabel("源语言:"),
			StyledSelect([]string{
				"简体中文", "英文", "日文", "土耳其语", "德语", "韩语", "俄语",
			}, func(value string) {
				langMap := map[string]string{
					"简体中文": "zh_cn", "英文": "en", "日文": "ja",
					"土耳其语": "tr", "德语": "de", "韩语": "ko", "俄语": "ru",
				}
				sm.SetSourceLang(langMap[value])
			}),
		),
		container.NewHBox(
			widget.NewLabel("目标语言:"),
			StyledSelect([]string{
				"简体中文", "繁体中文", "英语", "日语", "韩语", "法语", "德语", "俄语",
				"西班牙语", "葡萄牙语", "意大利语", "阿拉伯语", "土耳其语",
			}, func(value string) {
				langMap := map[string]string{
					"简体中文": "zh_cn", "繁体中文": "zh_tw", "英语": "en",
					"日语": "ja", "韩语": "ko", "法语": "fr", "德语": "de",
					"俄语": "ru", "西班牙语": "es", "葡萄牙语": "pt",
					"意大利语": "it", "阿拉伯语": "ar", "土耳其语": "tr",
				}
				sm.SetTargetLang(langMap[value])
			}),
		),
	)

	// 设置默认语言
	langContainer.Objects[0].(*fyne.Container).Objects[1].(*widget.Select).SetSelected("简体中文")
	langContainer.Objects[1].(*fyne.Container).Objects[1].(*widget.Select).SetSelected("简体中文")

	// 创建容器
	content := container.NewVBox(
		container.NewPadded(inputTypeContainer),
		container.NewPadded(videoInputContainer),
		container.NewPadded(langContainer),
	)

	return GlassCard("1. 视频源设置", "选择视频和语言", content)
}

// 创建字幕设置卡片
func createSubtitleSettingsCard(sm *SubtitleManager) *fyne.Container {
	// 创建更美观的双语位置选择器
	bilingualCheck := widget.NewCheck("启用双语字幕", func(checked bool) {
		sm.SetBilingualEnabled(checked)
	})
	bilingualCheck.SetChecked(true)

	// 使用更长的选项文本，强制下拉框显示更宽
	positionSelect := widget.NewSelect([]string{
		"翻译后字幕在上方    ", // 添加额外空格增加宽度
		"翻译后字幕在下方    ", // 添加额外空格增加宽度
	}, func(value string) {
		if value == "翻译后字幕在上方    " {
			sm.SetBilingualPosition(1)
		} else {
			sm.SetBilingualPosition(2)
		}
	})
	positionSelect.SetSelected("翻译后字幕在上方    ")

	fillerCheck := widget.NewCheck("启用语气词过滤", func(checked bool) {
		sm.SetFillerFilter(checked)
	})
	fillerCheck.SetChecked(true)

	// 使用更合理的布局
	content := container.NewVBox(
		container.NewHBox(bilingualCheck, fillerCheck),
		positionSelect, // 直接让它占据整行以获得足够空间
	)

	return StyledCard("2. 字幕设置", content)
}

// 创建配音设置卡片
func createVoiceSettingsCard(sm *SubtitleManager) *fyne.Container {
	// 创建配音启用复选框和性别选择
	voiceoverCheck := widget.NewCheck("启用配音", func(checked bool) {
		sm.SetVoiceoverEnabled(checked)
	})

	genderSelect := StyledSelect([]string{"男声", "女声"}, func(value string) {
		if value == "男声" {
			sm.SetVoiceoverGender(2)
		} else {
			sm.SetVoiceoverGender(1)
		}
	})
	genderSelect.SetSelected("男声")

	// 创建音频选择按钮 - 使用普通按钮，蓝色文字
	audioSampleButton := SecondaryButton("选择音色克隆样本", theme.MediaMusicIcon(), sm.ShowAudioFileDialog)

	// 使用漂亮的网格布局
	grid := container.NewGridWithColumns(2,
		container.NewHBox(voiceoverCheck, genderSelect),
		audioSampleButton,
	)

	return StyledCard("3. 配音设置", grid)
}

// 创建字幕嵌入设置卡片
func createEmbedSettingsCard(sm *SubtitleManager) *fyne.Container {
	// 创建字幕嵌入复选框
	embedCheck := widget.NewCheck("合成字幕嵌入视频", nil)

	// 创建视频类型选择
	embedTypeSelect := StyledSelect([]string{
		"横屏视频", "竖屏视频", "横屏+竖屏视频",
	}, func(value string) {
		switch value {
		case "横屏视频":
			sm.SetEmbedSubtitle("horizontal")
		case "竖屏视频":
			sm.SetEmbedSubtitle("vertical")
		case "横屏+竖屏视频":
			sm.SetEmbedSubtitle("all")
		}
	})
	embedTypeSelect.Disable()

	// 创建标题输入区域
	mainTitleEntry := StyledEntry("请输入主标题")
	subTitleEntry := StyledEntry("请输入副标题")

	titleInputContainer := container.NewVBox(
		container.NewGridWithColumns(2,
			widget.NewLabel("主标题:"),
			mainTitleEntry,
		),
		container.NewGridWithColumns(2,
			widget.NewLabel("副标题:"),
			subTitleEntry,
		),
	)
	titleInputContainer.Hide()

	// 设置复选框行为
	embedCheck.OnChanged = func(checked bool) {
		if checked {
			embedTypeSelect.Enable()
			embedTypeSelect.SetSelected("横屏视频")
			sm.SetEmbedSubtitle("horizontal")
			if embedTypeSelect.Selected == "竖屏视频" || embedTypeSelect.Selected == "横屏+竖屏视频" {
				titleInputContainer.Show()
			}
		} else {
			embedTypeSelect.Disable()
			sm.SetEmbedSubtitle("none")
			titleInputContainer.Hide()
		}
	}

	// 更新标题输入区域的显示状态
	embedTypeSelect.OnChanged = func(value string) {
		if value == "竖屏视频" || value == "横屏+竖屏视频" {
			titleInputContainer.Show()
		} else {
			titleInputContainer.Hide()
		}
	}

	// 创建顶部控制区域
	topContainer := container.NewHBox(embedCheck, embedTypeSelect)

	// 创建主容器
	mainContainer := container.NewVBox(
		topContainer,
		container.NewPadded(titleInputContainer),
	)

	return StyledCard("字幕嵌入设置", mainContainer)
}

// 创建进度和下载区域
func createProgressAndDownloadArea(sm *SubtitleManager) (*widget.ProgressBar, *fyne.Container, *fyne.Container) {
	// 创建进度条
	progress := widget.NewProgressBar()
	progress.Hide()

	//进度百分比标签
	percentLabel := widget.NewLabel("0%")
	percentLabel.Hide()
	percentLabel.Alignment = fyne.TextAlignTrailing

	// 进度条容器
	progressContainer := container.NewBorder(nil, nil, nil, percentLabel, progress)
	progressContainer.Hide()

	// 添加半透明背景和阴影
	progressBg := canvas.NewRectangle(color.NRGBA{R: 240, G: 245, B: 250, A: 230})
	progressBg.SetMinSize(fyne.NewSize(0, 40))
	progressBg.CornerRadius = 8

	progressShadow := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 20})
	progressShadow.Move(fyne.NewPos(2, 2))
	progressShadow.SetMinSize(fyne.NewSize(0, 40))
	progressShadow.CornerRadius = 8

	progressWithBg := container.NewStack(
		progressShadow,
		progressBg,
		container.NewPadded(progressContainer),
	)
	progressWithBg.Hide()

	// 设置进度条和标签
	sm.SetProgressBar(progress)
	sm.SetProgressLabel(percentLabel)

	// 创建下载容器背景
	downloadBg := canvas.NewRectangle(color.NRGBA{R: 240, G: 250, B: 255, A: 230})
	downloadBg.CornerRadius = 10

	// 创建下载容器
	downloadContainer := container.NewVBox()
	downloadContainer.Hide()
	sm.SetDownloadContainer(downloadContainer)

	// 包装下载容器和背景
	downloadWithBg := container.NewStack(
		downloadBg,
		container.NewPadded(downloadContainer),
	)
	downloadWithBg.Hide()

	// 创建提示标签
	tipsLabel := widget.NewLabel("")
	tipsLabel.Hide()
	tipsLabel.Alignment = fyne.TextAlignCenter
	tipsLabel.Wrapping = fyne.TextWrapWord
	sm.SetTipsLabel(tipsLabel)

	tipsBg := canvas.NewRectangle(color.NRGBA{R: 255, G: 250, B: 230, A: 200})
	tipsBg.CornerRadius = 6

	tipsWithBg := container.NewStack(
		tipsBg,
		container.NewPadded(tipsLabel),
	)
	tipsWithBg.Hide()

	return progress, downloadWithBg, tipsWithBg
}

// 创建开始按钮
func createStartButton(window fyne.Window, sm *SubtitleManager, videoInputContainer *fyne.Container, embedSettingsCard *fyne.Container, progress *widget.ProgressBar, downloadContainer *fyne.Container) *widget.Button {
	btn := widget.NewButtonWithIcon("开始任务", theme.MediaPlayIcon(), nil)
	btn.Importance = widget.HighImportance

	btn.OnTapped = func() {
		// 按钮动画效果替换为简单的刷新
		originalImportance := btn.Importance
		btn.Importance = widget.DangerImportance
		btn.Refresh()

		go func() {
			time.Sleep(300 * time.Millisecond)
			btn.Importance = originalImportance
			btn.Refresh()
		}()

		var mainTitle, subTitle string

		if embedSettingsCard != nil && len(embedSettingsCard.Objects) > 1 {
			if titleContainer, ok := embedSettingsCard.Objects[1].(*fyne.Container); ok && titleContainer != nil && len(titleContainer.Objects) >= 2 {
				// 获取主标题
				if mainTitleRow, ok := titleContainer.Objects[0].(*fyne.Container); ok && mainTitleRow != nil && len(mainTitleRow.Objects) >= 2 {
					if mainTitleEntry, ok := mainTitleRow.Objects[1].(*widget.Entry); ok {
						mainTitle = mainTitleEntry.Text
					}
				}

				// 获取副标题
				if subTitleRow, ok := titleContainer.Objects[1].(*fyne.Container); ok && subTitleRow != nil && len(subTitleRow.Objects) >= 2 {
					if subTitleEntry, ok := subTitleRow.Objects[1].(*widget.Entry); ok {
						subTitle = subTitleEntry.Text
					}
				}
			}
		}

		sm.SetVerticalTitles(mainTitle, subTitle)

		// 显示进度条并隐藏下载容器
		progress.Show()
		sm.progressBar.SetValue(0) // 直接访问进度条
		downloadContainer.Hide()

		// 检查是否有视频URL
		if sm.GetVideoUrl() == "" {
			inputType := "本地视频" // 默认值

			if videoInputContainer != nil && len(videoInputContainer.Objects) > 0 {
				for i := 0; i < len(videoInputContainer.Objects); i++ {
					// 如果对象是Container，查找其中的RadioGroup
					if container, ok := videoInputContainer.Objects[i].(*fyne.Container); ok {
						for j := 0; j < len(container.Objects); j++ {
							if radio, ok := container.Objects[j].(*widget.RadioGroup); ok {
								inputType = radio.Selected
								break
							}
						}
					}
				}
			}

			if inputType == "本地视频" {
				dialog.ShowError(fmt.Errorf("请先选择视频文件"), window)
			} else {
				dialog.ShowError(fmt.Errorf("请输入视频链接"), window)
			}
			progress.Hide()
			return
		}

		// 桌面端的启动之前要check config
		err := config.CheckConfig()
		if err != nil {
			dialog.ShowError(fmt.Errorf("配置不正确: %v", err), window)
			log.GetLogger().Error("配置不正确", zap.Error(err))
			progress.Hide()
			return
		}

		err = deps.CheckDependency()
		if err != nil {
			dialog.ShowError(fmt.Errorf("依赖环境准备失败: %v", err), window)
			log.GetLogger().Error("依赖环境准备失败", zap.Error(err))
			progress.Hide()
			return
		}

		log.GetLogger().Info("配置内容", zap.Any("config", config.Conf))

		if err = sm.StartTask(); err != nil {
			dialog.ShowError(err, window)
			progress.Hide()
			return
		}

		downloadContainer.Show()
		sm.progressBar.Refresh()
	}

	return btn
}

// 创建OpenAI配置组
func createOpenAIConfigGroup() *fyne.Container {
	openaiBaseUrlEntry := StyledEntry("OpenAI API base url")
	openaiBaseUrlEntry.Bind(binding.BindString(&config.Conf.Openai.BaseUrl))

	openaiModelEntry := StyledEntry("OpenAI模型名称")
	openaiModelEntry.Bind(binding.BindString(&config.Conf.Openai.Model))

	openaiApiKeyEntry := StyledPasswordEntry("OpenAI API密钥")
	openaiApiKeyEntry.Bind(binding.BindString(&config.Conf.Openai.ApiKey))

	form := widget.NewForm(
		widget.NewFormItem("API base url", openaiBaseUrlEntry),
		widget.NewFormItem("模型名称", openaiModelEntry),
		widget.NewFormItem("API密钥", openaiApiKeyEntry),
	)

	return StyledCard("OpenAI配置", form)
}

// 创建Whisper配置组
func createWhisperConfigGroup() *fyne.Container {
	whisperBaseUrlEntry := StyledEntry("Whisper API base url")
	whisperBaseUrlEntry.Bind(binding.BindString(&config.Conf.Openai.Whisper.BaseUrl))

	whisperApiKeyEntry := StyledPasswordEntry("Whisper API密钥")
	whisperApiKeyEntry.Bind(binding.BindString(&config.Conf.Openai.Whisper.ApiKey))

	form := widget.NewForm(
		widget.NewFormItem("API base url", whisperBaseUrlEntry),
		widget.NewFormItem("API密钥", whisperApiKeyEntry),
	)

	return StyledCard("Whisper配置", form)
}

// 创建阿里云OSS配置组
func createAliyunOSSConfigGroup() *fyne.Container {
	ossAccessKeyIdEntry := StyledEntry("阿里云AccessKey ID")
	ossAccessKeyIdEntry.Bind(binding.BindString(&config.Conf.Aliyun.Oss.AccessKeyId))

	ossAccessKeySecretEntry := StyledPasswordEntry("阿里云AccessKey Secret")
	ossAccessKeySecretEntry.Bind(binding.BindString(&config.Conf.Aliyun.Oss.AccessKeySecret))

	ossBucketEntry := StyledEntry("OSS Bucket名称")
	ossBucketEntry.Bind(binding.BindString(&config.Conf.Aliyun.Oss.Bucket))

	form := widget.NewForm(
		widget.NewFormItem("AccessKey ID", ossAccessKeyIdEntry),
		widget.NewFormItem("AccessKey Secret", ossAccessKeySecretEntry),
		widget.NewFormItem("Bucket名称", ossBucketEntry),
	)

	return GlassCard("阿里云OSS配置", "对象存储服务", form)
}

// 创建阿里云语音配置组
func createAliyunSpeechConfigGroup() *fyne.Container {
	speechAppKeyEntry := StyledEntry("阿里云语音服务AppKey")
	speechAppKeyEntry.Bind(binding.BindString(&config.Conf.Aliyun.Speech.AppKey))

	form := widget.NewForm(
		widget.NewFormItem("AppKey", speechAppKeyEntry),
	)

	return StyledCard("阿里云语音配置", form)
}

// 创建阿里云百炼配置组
func createAliyunBailianConfigGroup() *fyne.Container {
	bailianApiKeyEntry := StyledPasswordEntry("阿里云百炼API密钥")
	bailianApiKeyEntry.Bind(binding.BindString(&config.Conf.Aliyun.Bailian.ApiKey))

	form := widget.NewForm(
		widget.NewFormItem("API密钥", bailianApiKeyEntry),
	)

	return StyledCard("阿里云百炼配置", form)
}

// 创建保存按钮
func createSaveButton(window fyne.Window) *widget.Button {
	// 创建保存按钮（但不设置点击事件）
	saveButton := widget.NewButtonWithIcon("保存配置", theme.DocumentSaveIcon(), nil)
	saveButton.Importance = widget.HighImportance

	// 设置点击事件
	saveButton.OnTapped = func() {
		// 创建loading对话框
		progress := dialog.NewProgress("保存中", "正在保存配置...", window)
		progress.Show()

		// 模拟保存进度
		go func() {
			for i := 0.0; i <= 1.0; i += 0.1 {
				time.Sleep(50 * time.Millisecond)
				progress.SetValue(i)
			}

			// 保存配置
			err := config.SaveConfig()
			progress.Hide()

			if err != nil {
				dialog.ShowError(fmt.Errorf("保存配置失败: %v", err), window)
				log.GetLogger().Error("保存配置失败", zap.Error(err))
				return
			}

			// 重新加载配置
			config.LoadConfig()

			successDialog := dialog.NewInformation("成功", "配置已保存", window)
			successDialog.SetDismissText("确定")
			successDialog.Show()
		}()
	}

	return saveButton
}
