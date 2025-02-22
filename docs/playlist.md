# playlist

GET `/playlist/detail`

`http://127.0.0.1:3000/playlist/detail?mlid=??????`

```json
{
  "code": 0,
  "data": [
    {
      "aid": 113894160929245,
      "bvid": "BV1TBf6YjECp",
      "title": "电棍不语♿只是一味欧内好汉的手",
      "cover": "http://i1.hdslb.com/bfs/archive/4ff24199fb028fd5b21eb036a3a8a9ca6c19983f.jpg",
      "artist": {
        "id": 266476883,
        "name": "胃痛狂喝双氧水"
      },
      "has_part": 3
    },
    {
      "aid": 113576115177555,
      "bvid": "BV1FEzoYvEDh",
      "title": "今夜电棍闪闪:我才是奶龙！",
      "cover": "http://i0.hdslb.com/bfs/archive/bcb35db26fce47303571d0ece31f747a9fa98618.jpg",
      "artist": {
        "id": 1434794716,
        "name": "获乱时期的爱情"
      },
      "has_part": 1
    },
    ......
  ],
  "msg": "success"
}
```

如果是私密收藏夹，当用有权限账号（比如自己已经扫码登录过），返回和上相同，否则api的状态码为500

错误为

```
错误码: -403, 错误信息: 访问权限不足
```