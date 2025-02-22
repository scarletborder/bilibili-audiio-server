# 登录

`/login/*`

## 开始二维码登录事务

Get `/login/qrcode`

Example : `http://127.0.0.1:3000/login/qrcode`

```json
{
  "code": 0,
  "msg": "获取二维码成功",
  "transaction_id": "3cb1371b-3046-4748-bdbd-97c7669bbe53"
}
```

事务id在服务端软件的缓存中保存，过期时间为3分钟，超时需要重新开始登录流程

## 获取事务的二维码

Get `/login/qrcode_img`

Example : `http://127.0.0.1:3000/login/qrcode_img?tid=3cb1371b-3046-4748-bdbd-97c7669bbe53`

返回图片(`image/png`)

获取的二维码可以使用bilibili移动客户端扫码登录

## 请求登录信息

Get `/login/qrcode_status`

当扫码登录后可以请求该url查询之前的登录状态，该方法会阻塞直到扫码成功或者已经无法扫码

Example : `http://127.0.0.1:3000/login/qrcode_status?tid=3cb1371b-3046-4748-bdbd-97c7669bbe5`

```json
{
  "url": "https://passport.biligame.com/x/passport-login/web/crossDomain?DedeUserID=你的b站UID&DedeUserID__ckMd5=一串散列值&Expires=一串秒单位的时间戳&SESSDATA=很长的一串SESSION&bili_jct=一段令牌&gourl=https%3A%2F%2Fwww.bilibili.com&first_domain=.bilibili.com",
  "refresh_token": "一个refresh token",
  "timestamp": number毫秒时间戳,
  "code": 0,
  "message": ""
}
```

如果在之后的业务中要重新获取token，请见 [user/cookies](./user.md)