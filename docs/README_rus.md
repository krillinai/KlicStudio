<div align="center">
  <img src="./images/logo.png" alt="KrillinAI" height="90">

  # AI инструмент для перевода и озвучки аудио и видео

  <a href="https://trendshift.io/repositories/13360" target="_blank"><img src="https://trendshift.io/api/badge/repositories/13360" alt="krillinai%2FKrillinAI | Trendshift" style="width: 250px; height: 55px;" width="250" height="55"/></a>

  **[English](../README.md)｜[简体中文](../docs/README_zh.md)｜[日本語](../docs/README_jp.md)｜[한국어](../docs/README_kr.md)｜[Français](../docs/README_fr.md)｜[Deutsch](../docs/README_de.md)｜[Español](../docs/README_es.md)｜[Português](../docs/README_pt.md)｜[Русский](../docs/README_rus.md)｜[اللغة العربية](../docs/README_ar.md)**
  
  [![Twitter](https://img.shields.io/badge/Twitter-KrillinAI-orange?logo=twitter)](https://x.com/KrillinAI)
[![Bilibili](https://img.shields.io/badge/dynamic/json?label=Bilibili&query=%24.data.follower&suffix=%20подписчиков&url=https%3A%2F%2Fapi.bilibili.com%2Fx%2Frelation%2Fstat%3Fvmid%3D242124650&logo=bilibili&color=00A1D6&labelColor=FE7398&logoColor=FFFFFF)](https://space.bilibili.com/242124650)

</div>

## Обзор

Krillin AI — это универсальное решение для простой локализации и улучшения видео. Этот минималистичный, но мощный инструмент выполняет всё: от перевода и дубляжа до клонирования голоса и адаптации формата — легко преобразует видео между горизонтальным и вертикальным режимами для идеального отображения на любых платформах (YouTube, TikTok, Bilibili, Douyin, WeChat Channel, RedNote, Kuaishou). Благодаря сквозному рабочему процессу Krillin AI превращает исходные материалы в готовый к публикации контент всего за несколько кликов.

## Ключевые возможности:
🎯 **Запуск в один клик** - мгновенное начало работы

📥 **Загрузка видео** - поддержка yt-dlp и локальных файлов

📜 **Точные субтитры** - распознавание с высокой точностью на основе Whisper

🧠 **Умное разделение** - логическая разбивка и выравнивание субтитров с помощью LLM

🌍 **Профессиональный перевод** -  согласованный перевод на уровне абзацев

🔄 **Замена терминов** - смена специализированной лексики в один клик 

🎙️ **Озвучка и клонирование голоса** - выбор голосов CosyVoice или создание копий

🎬 **Видеомонтаж** - автоматическое форматирование для горизонтальных и вертикальных форматов

## Пример работы
На изображении ниже показан результат автоматической вставки субтитров в видео после однокликового запуска обработки 46-минутного локального видео. Никаких ручных корректировок не производилось.
![Alignment](../docs/images/alignment.png)

<table>
<tr>
<td width="33%">

### Перевод субтитров
---
https://github.com/user-attachments/assets/bba1ac0a-fe6b-4947-b58d-ba99306d0339

</td>
<td width="33%">

### Озвучка
---
https://github.com/user-attachments/assets/0b32fad3-c3ad-4b6a-abf0-0865f0dd2385

</td>

<td width="33%">

### портретный режим
---
https://github.com/user-attachments/assets/c2c7b528-0ef8-4ba9-b8ac-f9f92f6d4e71

</td>
</tr>
</table>

## 🌍 Поддерживаемые языки
Входные языки: китайский, английский, японский, немецкий, турецкий (добавляются новые языки)
Языки перевода: 56 языков, включая английский, китайский, русский, испанский, французский и др.

## Предпросмотр интерфейса
![Предпросмотр интерфейса](../docs/images/ui_desktop.png)


## 🚀 Быстрый старт
### Основные шаги
1. Скачайте исполняемый файл, соответствующий вашей операционной системе, из раздела релизов и поместите его в пустую папку.
2. Создайте папку config внутри этой папки, затем создайте файл config.toml в папке config. Скопируйте содержимое файла config-example.toml из директории config исходного кода в config.toml и заполните вашу конфигурационную информацию соответствующим образом.
3. Дважды щелкните на исполняемом файле, чтобы запустить сервис.
4. Откройте браузер и введите http://127.0.0.1:8888, чтобы начать использование (замените 8888 на порт, который вы указали в файле config.toml).

### Для пользователей macOS
Это программное обеспечение не подписано, поэтому после завершения настройки файлов в "Основных шагах" вам потребуется вручную подтвердить доверие к приложению в macOS. Выполните следующие действия:
1. Откройте терминал и перейдите в директорию, где находится исполняемый файл (предположим, имя файла `KrillinAI_1.0.0_macOS_arm64`).
2. Выполните следующие команды по порядку:
```
sudo xattr -rd com.apple.quarantine ./KrillinAI_1.0.0_macOS_arm64
sudo chmod +x ./KrillinAI_1.0.0_macOS_arm64
./KrillinAI_1.0.0_macOS_arm64
```
Это запустит сервис.

### Инструкции по настройке Cookie
Этот проект поддерживает развертывание через Docker. Пожалуйста, обратитесь к [Docker Deployment Instructions](./docs/docker.md).

### Cookie Configuration Instructions

Если вы столкнулись с ошибками при загрузке видео, пожалуйста, обратитесь к [Cookie Configuration Instructions](./docs/get_cookies.md) для настройки информации о ваших cookie.

### Помощь по настройке
Самый быстрый и удобный способ настройки:
* Выберите openai для transcription_provider и llm_provider. Таким образом, вам нужно будет заполнить только openai.apikey в следующих трех основных категориях конфигурации, а именно openai, local_model и aliyun, и затем вы сможете выполнять перевод субтитров. (Заполните app.proxy, model и openai.base_url в соответствии с вашей ситуацией.)

Способ настройки для использования локальной модели распознавания речи (временно не поддерживается на macOS) (выбор, учитывающий стоимость, скорость и качество):
* Заполните fasterwhisper для transcription_provider и openai для llm_provider. Таким образом, вам нужно будет заполнить только openai.apikey и local_model.faster_whisper в следующих двух основных категориях конфигурации, а именно openai и local_model, и затем вы сможете выполнять перевод субтитров. Локальная модель будет загружена автоматически. (То же самое относится к app.proxy и openai.base_url, как упоминалось выше.)

Следующие ситуации использования требуют настройки Alibaba Cloud:
* Если llm_provider заполнен как aliyun, это означает, что будет использоваться сервис больших моделей Alibaba Cloud. Следовательно, необходимо настроить параметр aliyun.bailian.
* Если transcription_provider заполнен как aliyun, или если функция "озвучки" включена при запуске задачи, будет использоваться голосовой сервис Alibaba Cloud. Поэтому необходимо заполнить параметр aliyun.speech.
* Если функция "озвучки" включена и одновременно загружаются локальные аудиофайлы для клонирования тембра голоса, также будет использоваться сервис облачного хранилища OSS от Alibaba Cloud. Следовательно, необходимо заполнить параметр aliyun.oss.
Руководство по настройке: [Alibaba Cloud Configuration Instructions](./docs/aliyun.md)

## Часто задаваемые вопросы
Пожалуйста, обратитесь к [Frequently Asked Questions](./docs/faq.md)

## Рекомендации по внесению вклада

- Не отправляйте ненужные файлы, такие как .vscode, .idea и т.д. Пожалуйста, используйте .gitignore для их фильтрации.
- Не отправляйте config.toml; вместо этого отправляйте config-example.toml.

## История звезд

[![Star History Chart](https://api.star-history.com/svg?repos=krillinai/KrillinAI&type=Date)](https://star-history.com/#krillinai/KrillinAI&Date)
