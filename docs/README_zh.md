<div align="center">
  <img src="./images/logo.png" alt="KrillinAI" height="90">

  # 极简部署AI视频翻译配音工具

  **[English](../README.md)｜[简体中文](../docs/README_zh.md)｜[日本語](../docs/README_jp.md)｜[Русский](./docs/README_rus.md)**

 [![QQ 群](https://img.shields.io/badge/QQ%20群-754069680-green?logo=tencent-qq)](https://jq.qq.com/?_wv=1027&k=754069680)
[![哔哩哔哩](https://img.shields.io/badge/哔哩哔哩-KrillinAI-red?logo=bilibili)](https://space.bilibili.com/242124650)

</div>

### 📢win&mac桌面端新发布 欢迎测试反馈

 ## 项目简介  

Krillin AI 是一款全能型音视频本地化与增强解决方案。这款简约而强大的工具，集音视频翻译、配音、语音克隆于一身，支持横竖屏格式输出，确保在所有主流平台（哔哩哔哩，小红书，抖音，视频号，快手，YouTube，TikTok等）都能完美呈现。通过端到端的工作流程，Krillin AI 仅需点击几次，就能将原始素材转化为精美即用的跨平台内容。

## 主要特点与功能：
🎯 **一键启动**：无需复杂的环境配置，自动安装依赖，立即投入使用 

📥 **视频获取**：支持yt-dlp下载或本地文件上传

📜 **精准识别**：基于Whisper的高准确度语音识别

🧠 **智能分段**：使用LLM进行字幕分段和对齐

🔄 **术语替换**：一键替换专业领域词汇 

🌍 **专业翻译**：基于LLM，段落级翻译保持语义连贯性

🎙️ **配音克隆**：提供CosyVoice精选音色或自定义音色克隆

🎬 **视频合成**：自动处理横竖版视频和字幕排版


## 效果展示
下图为46分钟的本地视频导入，一键执行后生成的字幕文件入轨后的效果，无任何手动调整。无缺失、重叠，断句自然，翻译质量也非常高。
![对齐效果](./images/alignment.png)

<table>
<tr>
<td width="50%">

### 字幕翻译
---
https://github.com/user-attachments/assets/bba1ac0a-fe6b-4947-b58d-ba99306d0339

</td>
<td width="50%">



### 配音
---
https://github.com/user-attachments/assets/0b32fad3-c3ad-4b6a-abf0-0865f0dd2385

</td>
</tr>
</table>


## 语言支持
输入语言支持：中文，英文，日语，德语，土耳其语（持续增加中）

翻译语言支持：英文，中文，俄语，西班牙语，法语等56种语言

## 界面预览
![界面预览](./images/ui_desktop.png)


## 🚀 快速开始
### 基本步骤
1. 下载[Release](https://github.com/krillinai/KrillinAI/releases)中与你设备系统匹配的可执行文件，放入空文件夹
2. 在文件夹内创建`config`文件夹，然后在`config`文件夹创建`config.toml`文件，复制源代码`config`目录下的`config-example.toml`文件的内容填入`config.toml`，并对照填写你的配置信息。（想用openai模型但不会获取key的可以加群免费试用）
3. 双击，或在终端执行可执行文件，启动服务
4. 打开浏览器，输入`http://127.0.0.1:8888`，开始使用 (8888替换成你在配置文件中填写的端口)

### To: macOS用户
本软件没有做签名，因此在macOS上运行时，在完成“基本步骤”中的文件配置后，还需要手动信任应用，方法如下：
1. 在终端打开可执行文件（假设文件名是KrillinAI_1.0.0_macOS_arm64）所在目录
2. 依次执行以下命令：
   ```
    sudo xattr -rd com.apple.quarantine ./KrillinAI_1.0.0_macOS_arm64
    sudo chmod +x ./KrillinAI_1.0.0_macOS_arm64
    ./KrillinAI_1.0.0_macOS_arm64
    ```
    即可启动服务

### Docker部署
本项目支持Docker部署，请参考[Docker部署说明](./docker.md)

### Cookie配置说明(非必选)

如果你遇到视频下载失败的情况

请参考 [Cookie 配置说明](./get_cookies.md) 配置你的Cookie信息。

### 配置帮助（必看）
最快速便捷的配置方式：
* `transcription_provider`和`llm_provider`都选择`openai`，这样在下方`openai`、`local_model`、`aliyun`三个配置项大类里只需要填写`openai.apikey`就可以进行字幕翻译。(`app.proxy`、`model`和`openai.base_url`按自己情况选填)

使用本地语言识别模型（暂不支持macOS）的配置方式（兼顾成本、速度与质量的选择）
* `transcription_provider`填写`fasterwhisper`，`llm_provider`填写`openai`，这样在下方`openai`、`local_model`三个配置项大类里只需要填写`openai.apikey`和`local_model.faster_whisper`就可以进行字幕翻译，本地模型会自动下载。(`app.proxy`和`openai.base_url`同上)

以下几种使用情况，需要进行阿里云的配置：
* 如果`llm_provider`填写了`aliyun`，需要使用阿里云的大模型服务，因此需要配置`aliyun.bailian`项的配置
* 如果`transcription_provider`填写了`aliyun`，或者在启动任务时开启了“配音”功能，都需要使用阿里云的语音服务，因此需要填写`aliyun.speech`项的配置
* 如果开启了“配音”功能，同时上传了本地的音频做音色克隆，则还需要使用阿里云的OSS云存储服务，因此需要填写`aliyun.oss`项的配置  
阿里云配置帮助：[阿里云配置说明](./aliyun.md)

## 常见问题

请移步[常见问题](./faq.md)

## 贡献规范
1. 不要提交无用文件，如.vscode、.idea等，请善于使用.gitignore过滤
2. 不要提交config.toml，而是使用config-example.toml提交

## 联系我们
1. 加入我们的QQ群，解答问题：754069680
2. 关注我们的社交媒体账号，[哔哩哔哩](https://space.bilibili.com/242124650)，每天分享AI科技领域优质内容

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=krillinai/KrillinAI&type=Date)](https://star-history.com/#krillinai/KrillinAI&Date)
