<div align="center">
  <img src="../docs/images/logo.png" alt="KrillinAI" height="90">


 # أداة ترجمة ودبلجة الصوت والفيديو بالذكاء الاصطناعي

<a href="https://trendshift.io/repositories/13360" target="_blank"><img src="https://trendshift.io/api/badge/repositories/13360" alt="krillinai%2FKrillinAI | Trendshift" style="width: 250px; height: 55px;" width="250" height="55"/></a>

  **[English](../README.md)｜[简体中文](../docs/README_zh.md)｜[日本語](../docs/README_jp.md)｜[한국어](../docs/README_kr.md)｜[Français](../docs/README_fr.md)｜[Deutsch](../docs/README_de.md)｜[Español](../docs/README_es.md)｜[Português](../docs/README_pt.md)｜[Русский](../docs/README_rus.md)｜[اللغة العربية](../docs/README_ar.md)**

  [![Twitter](https://img.shields.io/badge/Twitter-KrillinAI-orange?logo=twitter)](https://x.com/KrillinAI)
[![Bilibili](https://img.shields.io/badge/dynamic/json?label=Bilibili&query=%24.data.follower&suffix=%20followers&url=https%3A%2F%2Fapi.bilibili.com%2Fx%2Frelation%2Fstat%3Fvmid%3D242124650&logo=bilibili&color=00A1D6&labelColor=FE7398&logoColor=FFFFFF)](https://space.bilibili.com/242124650)

</div>

 ### إصدار جديد لنظامي ويندوز وماك - مرحبًا باختباره وتقديم الملاحظات

## نظرة عامة

كريلين AI هو حل متكامل لتحسين وتوطين الفيديوهات بسهولة. هذه الأداة البسيطة لكن القوية تتعامل مع كل شيء من الترجمة والدبلجة إلى استنساخ الأصوات، وإعادة التنسيق - حيث تقوم بتحويل الفيديوهات بسلاسة بين الوضع الأفقي والعمودي لعرض مثالي على جميع منصات المحتوى (يوتيوب، تيك توك، بيلبلي، دويين، قناة وي تشات، ريد نوت، كوايشو). من خلال سير العمل الشامل، يحوّل كريلين AI اللقطات الخام إلى محتوى نهائي وجاهز للنشر ببضع نقرات فقط.

الميزات الرئيسية:
🎯 بدء بنقرة واحدة - ابدأ سير العمل فورًا، النسخة الجديدة لسطح المكتب متاحة الآن - أسهل في الاستخدام!

📥 تنزيل الفيديو - يدعم yt-dlp ورفع الملفات المحلية

📜 ترجمات دقيقة - تعتمد على Whisper للتعرف عالي الدقة

🧠 تقسيم ذكي - تجزئة المحاذاة التلقائية للترجمات بناءً على نماذج اللغات الكبيرة (LLM)

🌍 ترجمة احترافية - ترجمة على مستوى الفقرات للحفاظ على الاتساق

🔄 استبدال المصطلحات - تبديل المفردات المتخصصة بنقرة واحدة

🎙️ الدبلجة واستنساخ الأصوات - اختيار أصوات CosyVoice أو استنساخ الأصوات

🎬 تكوين الفيديو - إعادة التنسيق التلقائي للوضع الأفقي/العمودي

## عرض توضيحي
الصورة التالية توضح النتيجة بعد إدراج ملف الترجمة - الذي تم إنشاؤه بنقرة واحدة بعد استيراد فيديو محلي مدته 46 دقيقة - في المسار. لم يتم إجراء أي تعديل يدوي على الإطلاق. لا توجد ترجمات ناقصة أو متداخلة، وتقسيم الجمل طبيعي، وجودة الترجمة عالية جدًا.
![Alignment](../docs/images/alignment.png)

<table>
<tr>
<td width="33%">

### ترجمة الترجمة النصية
---
https://github.com/user-attachments/assets/bba1ac0a-fe6b-4947-b58d-ba99306d0339

</td>
<td width="33%">

### الدبلجة
---
https://github.com/user-attachments/assets/0b32fad3-c3ad-4b6a-abf0-0865f0dd2385

</td>

<td width="33%">

### الوضع العمودي
---
https://github.com/user-attachments/assets/c2c7b528-0ef8-4ba9-b8ac-f9f92f6d4e71

</td>

</tr>
</table>

## 🌍 اللغات المدعومة
لغات الإدخال: الصينية، الإنجليزية، اليابانية، الألمانية، التركية (مع إضافة المزيد من اللغات قريبًا)
لغات الترجمة: 56 لغة مدعومة، بما في ذلك الإنجليزية، الصينية، الروسية، الإسبانية، الفرنسية، وغيرها.

## معاينة الواجهة
![ui preview](../docs/images/ui_desktop.png)

## 🚀 بدء سريع
### الخطوات الأساسية

أولاً، قم بتنزيل ملف الإصدار التنفيذي الذي يتوافق مع نظام جهازك. اتبع التعليمات أدناه للاختيار بين نسخة سطح المكتب أو النسخة العادية، ثم ضع البرنامج في مجلد فارغ. عند تشغيل البرنامج، سيتم إنشاء بعض المجلدات تلقائياً، لذا فإن وضعه في مجلد فارغ يجعل إدارته أسهل.

[For the desktop version (release files with "desktop" in the name), refer here]  
_The desktop version is newly released to address the difficulty beginners face in editing configuration files correctly. It still has some bugs and is being continuously updated._  

انقر نقرًا مزدوجًا على الملف لبدء استخدامه.

[لنسخة غير سطح المكتب (ملفات الإصدار التي لا تحتوي على "desktop" في الاسم)، ارجع إلى هنا]  
_تعتبر نسخة غير سطح المكتب هي الإصدار الأصلي، تتميز بإعدادات أكثر تعقيدًا ولكن بوظائف مستقرة. وهي مناسبة أيضًا للنشر على الخوادم، حيث توفر واجهة مستخدم تعمل عبر الويب._

قم بإنشاء مجلد `config` في الدليل، ثم أنشئ ملف `config.toml` بداخله. انسخ محتويات ملف `config-example.toml` من مجلد `config` في الكود المصدقي إلى ملف `config.toml` الخاص بك وقم بملء تفاصيل الإعدادات. (إذا كنت ترغب في استخدام نماذج OpenAI ولكنك لا تعرف كيفية الحصول على مفتاح، يمكنك الانضمام إلى المجموعة للحصول على وصول تجريبي مجاني.)

انقر نقرًا مزدوجًا على الملف التنفيذي أو قم بتشغيله في الطرفية لبدء الخدمة.

افتح متصفحك وأدخل http://127.0.0.1:8888 لبدء استخدامه. (استبدل 8888 برقم المنفذ الذي حددته في ملف الإعدادات.)

### إلى: مستخدمي نظام macOS
[لنسخة سطح المكتب (أي ملفات الإصدار التي تحتوي على "desktop" في الاسم)، ارجع هنا]
طريقة التغليف الحالية لنسخة سطح المكتب لا تدعم التشغيل المباشر بالنقر المزدوج أو التثبيت عبر DMG بسبب مشاكل التوقيع. يتطلب ذلك إعداد الثقة يدوياً كما يلي:

1. افتح المجلد الذي يحتوي على الملف التنفيذي (لنفترض أن اسم الملف هو KrillinAI_1.0.0_desktop_macOS_arm64) في Terminal

2. نفّذ الأوامر التالية بالتسلسل:


```
sudo xattr -cr ./KrillinAI_1.0.0_desktop_macOS_arm64  
sudo chmod +x ./KrillinAI_1.0.0_desktop_macOS_arm64  
./KrillinAI_1.0.0_desktop_macOS_arm64  
```

[للنسخة العادية (ملفات الإصدار التي لا تحتوي على "desktop" في الاسم)، راجع هنا]
هذا البرنامج غير موقّع، لذا بعد إكمال إعداد الملفات وفق "الخطوات الأساسية"، ستحتاج إلى منح الثقة يدوياً للتطبيق على نظام macOS. اتبع هذه الخطوات:
1. افتح Terminal وانتقل إلى المجلد الذي يحتوي على الملف التنفيذي (لنفترض أن اسم الملف هو KrillinAI_1.0.0_macOS_arm64).

2. نفّذ الأوامر التالية بالتسلسل:

```
sudo xattr -rd com.apple.quarantine ./KrillinAI_1.0.0_macOS_arm64
sudo chmod +x ./KrillinAI_1.0.0_macOS_arm64
./KrillinAI_1.0.0_macOS_arm64
```
سيؤدي هذا إلى بدء تشغيل الخدمة.

### النشر باستخدام Docker

هذا المشروع يدعم النشر عبر يُرجى الرجوع إلى [Docker Deployment Instructions](../docs/docker.md).

### تعليمات إعداد Cookie

إذا واجهت فشلًا في تنزيل الفيديو، يُرجى الرجوع إلى تعليمات إعداد Cookie لتهيئة معلومات الـ Cookie الخاصة بك.

(ملاحظة: تم الحفاظ على نفس تنسيق الروابط والعناوين كما في النص الأصلي)


### مساعدة في الإعدادات
أسرع وأكثر طريقة ملائمة للإعداد:
* د openai لكل من transcription_provider و llm_provider. بهذه الطريقة، تحتاج فقط إلى ملء openai.apikey في الفئات الثلاث الرئيسية لبنود الإعداد التالية، وهي openai، local_model، و aliyun، ثم يمكنك إجراء ترجمة الترجمة النصية. (املأ app.proxy، model و openai.base_url حسب حالتك الخاصة.)

طريقة الإعداد لاستخدام نموذج التعرف على الكلام المحلي (غير مدعوم على macOS في الوقت الحالي) (خيار يأخذ في الاعتبار التكلفة والسرعة والجودة):

* املأ fasterwhisper لـ transcription_provider و openai لـ llm_provider. بهذه الطريقة، تحتاج فقط إلى ملء openai.apikey و local_model.faster_whisper في الفئتين الرئيسيتين لبنود الإعداد التالية، وهما openai و local_model، ثم يمكنك إجراء ترجمة الترجمة النصية. سيتم تنزيل النموذج المحلي تلقائيًا. (ينطبق نفس الأمر على app.proxy و openai.base_url كما ذكر أعلاه.)

حالات الاستخدام التي تتطلب إعدادات علي بابا السحابية
* إذا تم تعيين llm_provider إلى aliyun، فهذا يعني أنه سيتم استخدام خدمة النماذج الكبيرة من علي بابا السحابية. وبالتالي، يجب إعداد عنصر aliyun.bailian في الإعدادات.
* إذا تم تعيين transcription_provider إلى aliyun، أو إذا تم تمكين وظيفة "الدبلجة الصوتية" عند بدء المهمة، فسيتم استخدام خدمة الصوت من علي بابا السحابية. لذلك، يجب ملء عنصر aliyun.speech في الإعدادات.
* إذا تم تمكين وظيفة "الدبلجة الصوتية" وتم تحميل ملفات صوتية محلية لاستنساخ نبرة الصوت في نفس الوقت، فسيتم أيضًا استخدام خدمة التخزين السحابي OSS من علي بابا السحابية. وبالتالي، يجب ملء عنصر aliyun.oss في الإعدادات.
دليل الإعدادات: [Alibaba Cloud Configuration Instructions](../docs/aliyun.md)

## الأسئلة الشائعة
يُرجى الرجوع إلى [Frequently Asked Questions](../docs/faq.md)

## إرشادات المساهمة

- لا تقم بإرسال ملفات غير ضرورية مثل .vscode، .idea، وغيرها. يُرجى استخدام .gitignore بشكل صحيح لتصفيتها.
- لا تقم بإرسال ملف config.toml؛ بدلاً من ذلك، قم بإرسال ملف config-example.toml.
## تاريخ النجوم

[![Star History Chart](https://api.star-history.com/svg?repos=krillinai/KrillinAI&type=Date)](https://star-history.com/#krillinai/KrillinAI&Date)

