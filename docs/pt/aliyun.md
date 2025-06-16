## Pré-requisitos
É necessário ter uma conta do [Alibaba Cloud](https://www.aliyun.com) e passar pela verificação de identidade, a maioria dos serviços tem uma cota gratuita.

## Obtenção do `access_key_id` e `access_key_secret` do Alibaba Cloud
1. Acesse a [página de gerenciamento de AccessKey do Alibaba Cloud](https://ram.console.aliyun.com/profile/access-keys).
2. Clique em criar AccessKey, se necessário, escolha o modo de uso e selecione "Usar no ambiente de desenvolvimento local".
![Alibaba Cloud access key](/docs/images/aliyun_accesskey_1.png)
3. Guarde com segurança, é melhor copiar para um arquivo local.

## Ativação do serviço de voz do Alibaba Cloud
1. Acesse a [página de gerenciamento do serviço de voz do Alibaba Cloud](https://nls-portal.console.aliyun.com/applist), na primeira vez que entrar, será necessário ativar o serviço.
2. Clique em criar projeto.
![Alibaba Cloud speech](/docs/images/aliyun_speech_1.png)
3. Selecione as funcionalidades e ative.
![Alibaba Cloud speech](/docs/images/aliyun_speech_2.png)
4. A "síntese de voz de texto em fluxo (modelo grande CosyVoice)" precisa ser atualizada para a versão comercial, outros serviços podem ser usados na versão de experiência gratuita.
![Alibaba Cloud speech](/docs/images/aliyun_speech_3.png)
5. Copie a chave do aplicativo.
![Alibaba Cloud speech](/docs/images/aliyun_speech_4.png)

## Ativação do serviço OSS do Alibaba Cloud
1. Acesse o [console de serviços de armazenamento de objetos do Alibaba Cloud](https://oss.console.aliyun.com/overview), na primeira vez que entrar, será necessário ativar o serviço.
2. Selecione a lista de Buckets à esquerda e clique em criar.
![Alibaba Cloud OSS](/docs/images/aliyun_oss_1.png)
3. Selecione criação rápida, preencha um nome de Bucket que atenda aos requisitos e escolha a região **Xangai**, complete a criação (o nome preenchido aqui será o valor da configuração `aliyun.oss.bucket`).
![Alibaba Cloud OSS](/docs/images/aliyun_oss_2.png)
4. Após a criação, acesse o Bucket.
![Alibaba Cloud OSS](/docs/images/aliyun_oss_3.png)
5. Desative o interruptor "Impedir acesso público" e defina as permissões de leitura e escrita como "Leitura pública".
![Alibaba Cloud OSS](/docs/images/aliyun_oss_4.png)
![Alibaba Cloud OSS](/docs/images/aliyun_oss_5.png)