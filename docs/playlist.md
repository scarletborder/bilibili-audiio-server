# playlist

## detail 

查询收藏夹详情

查询参数

- `mlid=123` 表示收藏夹的id（默认）

- `page=0` 从0开始页码

如果`page`过大不再有新的视频，那么返回值中`data`键为`null`

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

## list_all

展开一个用户的所有收藏夹

GET `/playlist/list_all`

`http://127.0.0.1:3000/playlist/list_all`

```json
[
   {
    "count": 7,
    "id": 34123123,
    "title": "a playlist"
  },
  {
    "count": 65,
    "id": 32123123,
    "title": "another playlist"
  },
]
```

## get

**请避免滥用该方法，请不要频繁请求该接口**

GET `/playlist/get`

请求收藏夹的全部视频，自动展开分p,跳过失效视频

`http://127.0.0.1:3000/playlist/get?mlid=34123123123`

```json
{
  "code": 0,
  "data": [
    {
      "album": "电棍不语♿只是一味欧内好汉的手",
      "artists": null,
      "cover": "http://i1.hdslb.com/bfs/archive/4ff24199fb028fd5b21eb036a3a8a9ca6c19983f.jpg",
      "id": "BV1TBf6YjECp_28091943354",
      "name": "电棍不语♿只是一味欧内好汉的手",
      "url": "很长的有效的链接"
    },
    {
      "album": "电棍不语♿只是一味欧内好汉的手",
      "artists": null,
      "cover": "http://i1.hdslb.com/bfs/archive/4ff24199fb028fd5b21eb036a3a8a9ca6c19983f.jpg",
      "id": "BV1TBf6YjECp_28102035971",
      "name": "击杀音效",
      "url": "很长的有效的链接"
    },
    ...,
    {
      "album": "今夜电棍闪闪:我才是奶龙！",
      "artists": null,
      "cover": "http://i0.hdslb.com/bfs/archive/bcb35db26fce47303571d0ece31f747a9fa98618.jpg",
      "id": "BV1FEzoYvEDh_27122141682",
      "name": "今夜电棍闪闪：我才是奶龙！",
      "url": "很长的有效的链接"
    },
    ...
  ],
  "msg": "success"
}
```

改方法会请求较长时间

请求一个数量为约90的收藏夹在无缓存的情况下大约会请求10+秒，即使做了缓存全部命中也占用1秒以上
