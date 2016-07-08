# Cirrius

Cirrius is a scalable, and highly-available, web and microservice platform.

[![](https://img.shields.io/circleci/token/abf9e47762afcbbd936490819683ad44594f67b5/project/abcum/cirrius/master.svg?style=flat-square)](https://circleci.com/gh/abcum/cirrius) [![](https://img.shields.io/badge/status-alpha-ff00bb.svg?style=flat-square)](https://github.com/abcum/cirrius) [![](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/abcum/cirrius) [![](https://goreportcard.com/badge/github.com/abcum/cirrius?style=flat-square)](https://goreportcard.com/report/github.com/abcum/cirrius) [![](https://img.shields.io/badge/license-Apache_License_2.0-00bfff.svg?style=flat-square)](https://github.com/abcum/cirrius) 

#### Features

- Scalable web and microservice platform written in [Go](http://golang.org)
- Administrative **project tools**
	- Easily create new websites
	- Easily create new microservices
	- Track request and response analytics
	- Accessible and intuitive web interface
- Multiple **supported languages**
	- Build services using CSS
	- Build services using LESS
	- Build services using SCSS
	- Build services using SASS
	- Build services using GCSS
	- Build services using JSON
	- Build services using HTML
	- Build services using Markdown
	- Build services using Javascript
	- Build services using Handlebars
- Automatic **optimisation** of resources
	- All assets are minified
	- All assets are optimised
	- All assets are compressed
- **Dynamic scripting** out-of-the-box
	- Write node.js event-driven scripts
	- Use the built-in versioned npm modules
	- All script logs available in the admin console
- Built-in **multi-tenancy database** for data storage
	- Enable multi-tenancy within your node.js scripts
	- Store state in a scalable, distributed, secure database
- Automatic **security** by default
	- All microservices accessible over SSL

#### Installation

```bash
go get github.com/abcum/cirrius
```

#### Running

```bash
cirrius start --port-http 80 --port-https 443 --log-level debug
```

#### Deployment

```bash
docker run --name cirrius abcum/cirrius start --port-http 80 --port-https 443 --log-level debug
```