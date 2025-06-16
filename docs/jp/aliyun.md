## 前提条件
[阿里云](https://www.aliyun.com)のアカウントを作成し、本人確認を行う必要があります。ほとんどのサービスには無料枠があります。

## 阿里云`access_key_id`と`access_key_secret`の取得
1. [阿里云AccessKey管理ページ](https://ram.console.aliyun.com/profile/access-keys)にアクセスします。
2. AccessKeyを作成するためにクリックし、必要に応じて使用方法を選択し、「ローカル開発環境で使用」を選択します。
![阿里云access key](/docs/images/aliyun_accesskey_1.png)
3. 大切に保管し、できればローカルファイルにコピーして保存します。

## 阿里云音声サービスの開通
1. [阿里云音声サービス管理ページ](https://nls-portal.console.aliyun.com/applist)にアクセスし、初めての場合はサービスを開通させる必要があります。
2. プロジェクトを作成するためにクリックします。
![阿里云speech](/docs/images/aliyun_speech_1.png)
3. 機能を選択して開通させます。
![阿里云speech](/docs/images/aliyun_speech_2.png)
4. 「ストリーミングテキスト音声合成（CosyVoice大モデル）」は商用版にアップグレードする必要があります。他のサービスは無料体験版を使用できます。
![阿里云speech](/docs/images/aliyun_speech_3.png)
5. app keyをコピーします。
![阿里云speech](/docs/images/aliyun_speech_4.png)

## 阿里云OSSサービスの開通
1. [阿里云オブジェクトストレージサービスコンソール](https://oss.console.aliyun.com/overview)にアクセスし、初めての場合はサービスを開通させる必要があります。
2. 左側のBucketリストを選択し、次に作成をクリックします。
![阿里云OSS](/docs/images/aliyun_oss_1.png)
3. クイック作成を選択し、要件に合ったBucket名を入力し、**上海**地域を選択して作成を完了します（ここに入力した名前が設定項目`aliyun.oss.bucket`の値になります）。
![阿里云OSS](/docs/images/aliyun_oss_2.png)
4. 作成が完了したらBucketに入ります。
![阿里云OSS](/docs/images/aliyun_oss_3.png)
5. 「公共アクセスをブロック」スイッチをオフにし、読み書き権限を「公共読み取り」に設定します。
![阿里云OSS](/docs/images/aliyun_oss_4.png)
![阿里云OSS](/docs/images/aliyun_oss_5.png)