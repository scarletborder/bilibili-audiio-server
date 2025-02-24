# media

## bvid+cid获取

bvid和cid需要用`_`连接，作为一个路径参数

GET `/media/:bvidcid`

`http://127.0.0.1:3000/media/BV1TBf6YjECp_28091943354`

下载文件，文件名为`bvid_cid.m4s`

例如 `BV1TBf6YjECp_28091943354.m4s`

## by proxy

代理一个url，返回m4s音频流文件

这里的url是通过`/song/url`请求到的MPEG-DASH 流媒体的音频流文件url，即域名为`xxx.bilivideo.com`

GET `/media/proxy`

查询参数（url的base64）或payload(url)两种传入参数方式

**payload**

```json
{
    "url" : "https://xxxx.bilivideo.com/upgcxcode/54xxxxx"
}
```

文件名为一个随机的`uuid.m4s`