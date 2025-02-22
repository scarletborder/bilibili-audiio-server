# bilibili audio server

像其他音乐api那样提供bilibili视频的音频流url

## 声明

**请勿滥用，本项目仅用于学习和测试！利用本项目提供的接口、文档等造成不良影响及后果与本人无关。**


## 快速开始
### 安装

`make`

`cd build/`

`audio_server`


### [可选]登录

为了读取私密收藏夹你需要进行登录，由于项目定位，登录状态会对整个服务生效，请别在公开部署的服务中进行登陆。

目前支持二维码登录.

## 使用

[目录](./docs/content.md)


## 引用

[CuteReimu/bilibili](https://github.com/CuteReimu/bilibili) 提供了bilibili sdk

[scarletborder/private_blueprint](https://github.com/scarletborder/private_blueprint) 提供了获取bilibili视频音频url的最佳实践