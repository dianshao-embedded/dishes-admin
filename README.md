# 上菜

[![License][license-image]][license-url]

[使用文档](https://dianshao-embedded.github.io/dishes-docs/)

## 项目特征

- 一个物联网终端 OTA 管理平台
- 基于 Wireguard VPN 可实时保持与终端设备连接
- 基于 Docker 部署，安装简单
- 很容易被集成进 CI/CD 工具以及终端开发工具中

## 依赖项目
上菜使用 Golang + Vue 前后端分离方案，其中后端基于 [gin-admin](https://github.com/LyricTian/gin-admin) 脚手架开发，前端基于 [vue-antd-admin](https://github.com/iczer/vue-antd-admin) 脚手架开发。

为了快速可靠的安装部署，上菜和相关依赖均运行于 Docker 容器之中

## 系统架构

整个系统主要由三部分组成: **dishes-admin, dishes-client, dishes-agent**. 其中 dishes-admin 为管理平台，运行于服务器中。dishes-client 运行于目标终端设备中，dishes-agent 为独立程序，方便 CI/CD 工具和开发工具集成

## MIT License

    Copyright (c) 2022 Croak

[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/npm/l/express.svg
