const express = require("express");
const axios = require("axios");
const bodyParser = require("body-parser");
const path = require("path");
const fs = require("fs");
const { v4: uuidv4 } = require("uuid");
const session = require("express-session");

const app = express();
const port = 3000;
const session_name = "my-webvpn-session-id-" + uuidv4().toString();

app.set("view engine", "pug");
app.set("trust proxy", false);
app.use(express.static(path.join(__dirname, "public")));
app.use(
  session({
    name: session_name,
    secret: uuidv4().toString(),
    secure: false,
    resave: false,
    saveUninitialized: true,
  })
);
app.use(bodyParser.json());
var userStorage = {
  username: {
    password: "password",
    info: {
      age: 18,
    },
    strategy: {
      "baidu.com": true,
      "google.com": false,
    },
  },
};

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

app.use("/proxy", async (req, res) => {
  const { username } = req.session;
  if (!username) {
    res.sendStatus(403);
  }

  let url = (() => {
    try {
      return new URL(req.query.url);
    } catch {
      res.status(400);
      res.end("invalid url.");
      return undefined;
    }
  })();

  if (!url) return;

  if (!userStorage[username].strategy[url.hostname]) {
    res.status(400);
    res.end("your url is not allowed.");
  }

  try {
    const headers = req.headers;
    headers.host = url.host;
    headers.cookie = headers.cookie.split(";").forEach((cookie) => {
      var filtered_cookie = "";
      const [key, value] = cookie.split("=", 1);
      if (key.trim() !== session_name) {
        filtered_cookie += `${key}=${value};`;
      }
      return filtered_cookie;
    });
    const remote_res = await (() => {
      if (req.method == "POST") {
        return axios.post(url, req.body, {
          headers: headers,
        });
      } else if (req.method == "GET") {
        return axios.get(url, {
          headers: headers,
        });
      } else {
        res.status(405);
        res.end("method not allowed.");
        return;
      }
    })();
    res.status(remote_res.status);
    res.header(remote_res.headers);
    res.write(remote_res.data);
  } catch (e) {
    res.status(500);
    res.end("unreachable url.");
  }
});

app.post("/user/login", (req, res) => {
  const { username, password } = req.body;
  if (
    typeof username != "string" ||
    typeof password != "string" ||
    !username ||
    !password
  ) {
    res.status(400);
    res.end("invalid username or password");
    return;
  }
  if (!userStorage[username]) {
    res.status(403);
    res.end("invalid username or password");
    return;
  }
  if (userStorage[username].password !== password) {
    res.status(403);
    res.end("invalid username or password");
    return;
  }
  req.session.username = username;
  res.send("login success");
});

// under development
app.post("/user/info", (req, res) => {
  if (!req.session.username) {
    res.sendStatus(403);
  }
  update(userStorage[req.session.username].info, req.body);
  res.sendStatus(200);
});

app.get("/home", (req, res) => {
  if (!req.session.username) {
    res.sendStatus(403);
    return;
  }
  res.render("home", {
    username: req.session.username,
    strategy: ((list)=>{
      var result = [];
      for (var key in list) {
        result.push({host: key, allow: list[key]});
      }
      return result;
    })(userStorage[req.session.username].strategy),
  });
});

// demo service behind webvpn
app.get("/flag", (req, res) => {
  if (
    req.headers.host != "127.0.0.1:3000" ||
    req.hostname != "127.0.0.1" ||
    req.ip != "127.0.0.1" 
  ) {
    res.sendStatus(400);
    return;
  }
  const data = fs.readFileSync("/flag");
  res.send(data);
});

app.listen(port, '0.0.0.0', () => {
  console.log(`app listen on ${port}`);
});
