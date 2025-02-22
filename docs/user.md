# user

## cookies

Get `/user/cookies`

Example : `http://127.0.0.1:3000/user/cookies`

```json
{
  "code": 0,
  "cookies": [
    {
      "Name": "buvid3",
      "Value": "20A?????????oc",
      "Quoted": false,
      "Path": "/",
      "Domain": ".bilibili.com",
      "Expires": "2026-?????Z",
      "RawExpires": "?????2026 ??? GMT",
      "MaxAge": 0,
      "Secure": false,
      "HttpOnly": false,
      "SameSite": 0,
      "Partitioned": false,
      "Raw": "buvid3=20A4????????oc; path=/; expires=??????2026??????? GMT; domain=.bilibili.com",
      "Unparsed": null
    },
    {
      "Name": "b_nut",
      "Value": "????????",
      "Quoted": false,
      "Path": "/",
      "Domain": ".bilibili.com",
      "Expires": "2026??????Z",
      "RawExpires": "??????2026 ????????? GMT",
      "MaxAge": 0,
      "Secure": false,
      "HttpOnly": false,
      "SameSite": 0,
      "Partitioned": false,
      "Raw": "b_nut=??????????; path=/; expires=???????2026??????GMT; domain=.bilibili.com",
      "Unparsed": null
    }
  ]
}
```

以上返回值示例是一个典型的游客cookies，在没有登录时会自动使用游客cookies
可以保证的是至少需要一年该cookies才会过期，所以本软件没有写自动续期的功能