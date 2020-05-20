# TwitterAnalysis

This is a sample template for TwitterAnalysis - Below is a brief explanation of what we have generated for you:

```bash
.
├── Makefile                    <-- Make to automate build
├── README.md                   <-- This instructions file
├── hello-world                 <-- Source code for a lambda function
│   ├── main.go                 <-- Lambda function code
│   └── main_test.go            <-- Unit tests
└── template.yaml
```

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)

## Setup process

### Installing dependencies

In this example we use the built-in `go get` and the only dependency we need is AWS Lambda Go SDK:

```shell
go get -u github.com/aws/aws-lambda-go/...
```

**NOTE:** As you change your application code as well as dependencies during development, you might want to research how to handle dependencies in Golang at scale.

### Building

Golang is a statically compiled language, meaning that in order to run it you have to build the executable target.

You can issue the following command in a shell to build it:

```shell
GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world
```

**NOTE**: If you're not building the function on a Linux machine, you will need to specify the `GOOS` and `GOARCH` environment variables, this allows Golang to build your function for another system architecture and ensure compatibility.

### Local development

**Invoking function locally through local API Gateway**

```bash
sam local start-api
```

If the previous command ran successfully you should now be able to hit the following local endpoint to invoke your function `http://localhost:3000/hello`

**SAM CLI** is used to emulate both Lambda and API Gateway locally and uses our `template.yaml` to understand how to bootstrap this environment (runtime, where the source code is, etc.) - The following excerpt is what the CLI will read in order to initialize an API and its routes:

```yaml
...
Events:
    HelloWorld:
        Type: Api # More info about API Event Source: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
        Properties:
            Path: /hello
            Method: get
```

## Packaging and deployment

AWS Lambda Golang runtime requires a flat folder with the executable generated on build step. SAM will use `CodeUri` property to know where to look up for the application:

```yaml
...
    FirstFunction:
        Type: AWS::Serverless::Function
        Properties:
            CodeUri: hello_world/
            ...
```

To deploy your application for the first time, run the following in your shell:

```bash
sam deploy --guided
```

The command will package and deploy your application to AWS, with a series of prompts:

* **Stack Name**: The name of the stack to deploy to CloudFormation. This should be unique to your account and region, and a good starting point would be something matching your project name.
* **AWS Region**: The AWS region you want to deploy your app to.
* **Confirm changes before deploy**: If set to yes, any change sets will be shown to you before execution for manual review. If set to no, the AWS SAM CLI will automatically deploy application changes.
* **Allow SAM CLI IAM role creation**: Many AWS SAM templates, including this example, create AWS IAM roles required for the AWS Lambda function(s) included to access AWS services. By default, these are scoped down to minimum required permissions. To deploy an AWS CloudFormation stack which creates or modified IAM roles, the `CAPABILITY_IAM` value for `capabilities` must be provided. If permission isn't provided through this prompt, to deploy this example you must explicitly pass `--capabilities CAPABILITY_IAM` to the `sam deploy` command.
* **Save arguments to samconfig.toml**: If set to yes, your choices will be saved to a configuration file inside the project, so that in the future you can just re-run `sam deploy` without parameters to deploy changes to your application.

You can find your API Gateway Endpoint URL in the output values displayed after deployment.

### Testing

We use `testing` package that is built-in in Golang and you can simply run the following command to run our tests:

```shell
go test -v ./hello-world/
```
# Appendix

### Golang installation

Please ensure Go 1.x (where 'x' is the latest version) is installed as per the instructions on the official golang website: https://golang.org/doc/install

A quickstart way would be to use Homebrew, chocolatey or your linux package manager.

#### Homebrew (Mac)

Issue the following command from the terminal:

```shell
brew install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
brew update
brew upgrade golang
```

#### Chocolatey (Windows)

Issue the following command from the powershell:

```shell
choco install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
choco upgrade golang
```

## Bringing to the next level

Here are a few ideas that you can use to get more acquainted as to how this overall process works:

* Create an additional API resource (e.g. /hello/{proxy+}) and return the name requested through this new path
* Update unit test to capture that
* Package & Deploy

Next, you can use the following resources to know more about beyond hello world samples and how others structure their Serverless applications:

* [AWS Serverless Application Repository](https://aws.amazon.com/serverless/serverlessrepo/)
````
{
  "tweet": [
    {
      "id": "12685042",
      "screen_name": "mouri45",
      "name": "もーり",
      "created_at": "2020-05-20T11:42:39+09:00",
      "text": "本日13:00からプログラミングなんでも オンラインもくもく会開催！途中参加OK！まだまだ参加者募集中！https://t.co/EW7YOO20zB#java #javascript #python #ruby… https://t.co/Kx0QPZylrI"
    },
    {
      "id": "9275912",
      "screen_name": "hiroaki_ohkawa",
      "name": "大川~生きづらさの解消に挑む~",
      "created_at": "2020-05-20T09:29:03+09:00",
      "text": "#エンジニア 向け情報 #golang #rustlangここからのhttps://t.co/GzalkqlqQkここ。https://t.co/YsJ2JaJCwDhttps://t.co/rsY56FF2hT… https://t.co/00PAN444yH"
    },
    {
      "id": "873428101769736192",
      "screen_name": "ReEngines",
      "name": "RE:ENGINES",
      "created_at": "2020-05-20T02:44:29+09:00",
      "text": "【過去記事ツイート】  Go言語の基礎〜基本構文その１〜 https://t.co/DpnGvQpqha #golang"
    },
    {
      "id": "48335348",
      "screen_name": "objectxplosive",
      "name": "眼力 玉壱號",
      "created_at": "2020-05-20T02:18:54+09:00",
      "text": "組み込み(?) な slice と map だけが polymorphic に振る舞えて、user defined types には許していないのが少し辛い事はある。諦めて interface {} に潰してしまうか codege… https://t.co/lVAQnyRCMQ"
    },
    {
      "id": "48335348",
      "screen_name": "objectxplosive",
      "name": "眼力 玉壱號",
      "created_at": "2020-05-20T02:14:47+09:00",
      "text": "Generics があった方が直感的な場所はある気はする。 https://t.co/3ZA8ZHSoMP の戻り値が kitchen sink な interface{} に潰れちゃうので手動で coerce  せざるをえなくて… https://t.co/LHLFSImf7a"
    },
    {
      "id": "1166338150433579008",
      "screen_name": "xiyegen",
      "name": "Nishino Wataru",
      "created_at": "2020-05-19T21:25:24+09:00",
      "text": "インタフェースの実装パターン #golang https://t.co/69EprjwcW4 #Qiita"
    },
    {
      "id": "258348576",
      "screen_name": "grove_twtr",
      "name": "grove",
      "created_at": "2020-05-19T19:05:51+09:00",
      "text": "Oh…非 Go で if の括弧を忘れる現象に。#golang"
    },
    {
      "id": "59957316",
      "screen_name": "mkamimura",
      "name": "kamimura",
      "created_at": "2020-05-19T18:12:07+09:00",
      "text": "“Go - 並行プログラミング - チャレンジ:火星で生きるもの - 発見を報告する - 通信、チャンネル、forループ、select、case、timeパッケージ、Sleep関数、After関数、競合状態、syncパッケージ、L… https://t.co/XhVJ8IKYx6"
    },
    {
      "id": "12685042",
      "screen_name": "mouri45",
      "name": "もーり",
      "created_at": "2020-05-19T17:23:38+09:00",
      "text": "明日(5/20)もプログラミングなんでも オンラインもくもく会やります！まだまだ参加者募集中！https://t.co/EW7YOO20zB#java #javascript #python #ruby #golang… https://t.co/6JTH1YzJR8"
    },
    {
      "id": "12685042",
      "screen_name": "mouri45",
      "name": "もーり",
      "created_at": "2020-05-19T11:23:07+09:00",
      "text": "本日もプログラミングなんでも オンラインもくもく会開催！途中参加OK！まだまだ参加者募集中です！https://t.co/u6VZggJ01k#java #javascript #python #ruby #golang… https://t.co/3lvSnPqSEF"
    }
  ],
  "sentimental": [
    {
      "sentiment": "NEUTRAL",
      "positive": 0.011342134326696396,
      "negative": 0.00021466199541464448,
      "mixed": 0.0000016308273416143493,
      "neutral": 0.9884415864944458
    },
    {
      "sentiment": "NEUTRAL",
      "positive": 0.0003667290147859603,
      "negative": 0.000015905996406218037,
      "mixed": 0.0000020865957139903912,
      "neutral": 0.9996151924133301
    },
    {
      "sentiment": "NEUTRAL",
      "positive": 0.00016882021736819297,
      "negative": 0.000020609626517398283,
      "mixed": 0.0000015425471246999223,
      "neutral": 0.9998090863227844
    },
    {
      "sentiment": "NEGATIVE",
      "positive": 0.00021462734730448574,
      "negative": 0.9335510730743408,
      "mixed": 0.000010119967555510812,
      "neutral": 0.0662240982055664
    },
    {
      "sentiment": "NEGATIVE",
      "positive": 0.00370821263641119,
      "negative": 0.5150494575500488,
      "mixed": 0.0004263958253432065,
      "neutral": 0.4808160066604614
    },
    {
      "sentiment": "NEUTRAL",
      "positive": 0.00022753627854399383,
      "negative": 0.00003690304947667755,
      "mixed": 7.264406463036721e-7,
      "neutral": 0.9997348189353943
    },
    {
      "sentiment": "NEUTRAL",
      "positive": 0.00012306390271987766,
      "negative": 0.00017789100820664316,
      "mixed": 5.893031129744486e-7,
      "neutral": 0.9996985197067261
    },
    {
      "sentiment": "NEUTRAL",
      "positive": 0.0007572532631456852,
      "negative": 0.00001866973252617754,
      "mixed": 0.000001832857037697977,
      "neutral": 0.9992222785949707
    },
    {
      "sentiment": "NEUTRAL",
      "positive": 0.002888354007154703,
      "negative": 0.00020212891104165465,
      "mixed": 0.0000011219791531402734,
      "neutral": 0.9969084858894348
    },
    {
      "sentiment": "NEUTRAL",
      "positive": 0.032975222915410995,
      "negative": 0.00022723602887708694,
      "mixed": 0.0000021944608761259587,
      "neutral": 0.9667953848838806
    }
  ]
}
````