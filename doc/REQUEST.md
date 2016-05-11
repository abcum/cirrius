```yaml
context:
  - request:
  	- url:
  	  - user
  	  - pass
  	  - host
  	  - path
  	  - query
  	- head:
  	  - content-type
  	  - key
    - body
  - success
  - failure
  - display:
  	- css
  	- xml
  	- text
  	- html
  	- json
  	- pack

if content-type == 'application-xml'
  context.request.body = `marshaled xml`

if content-type == 'application-json'
  context.request.body = `marshaled json`

if content-type == 'application-pack'
  context.request.body = `marshaled msgpack`

if content-type == 'multipart/form-data'
  context.request.body = `marshaled multipart`

if content-type == 'application/x-www-form-urlencoded'
  content.request.body = `marshaled form POST`
```