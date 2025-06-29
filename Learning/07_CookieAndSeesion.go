package Learning

//- http协议，有 3 个版本：
//  - http/1.0 版：无状态，短连接。
//  - http/1.1 版：可以记录状态。—— 默认支持。
//  - http/2.0 版：可以支持长连接。 协议头：Connection: keep-alive 。
//
//###
//### Cookie
//
//- 最早的 http/1.0 版，提供 Cookie 机制， 但是没有 Session。
//- Cookie 作用：一定时间内， 存储用户的连接信息。如：用户名、登录时间 ... 不敏感信息。
//- Cookie 出身：http自带机制。Session不是！
//- Cookie 存储：Cookie 存储在 客户端 (浏览器) 中。—— 浏览器可以存储数据。少
//  - 存储形式：key - value
//  - 可以在浏览器中查看。
//  - Cookie 不安全。直接将数据存储在浏览器上。

////### Session
////
////- ”会话“：在一次会话交流中，产生的数据。不是http、浏览器自带。
////- Session 作用：一定时间内， 存储用户的连接信息。
////- Session 存储：在服务器中。一般为 临时 Session。—— 会话结束 (浏览器关闭) ， Session被干掉！

//### 对比 Cookie 和 Session
//
//1.  Cookie 存储在 浏览器， 在哪生成呢？
//2.  Session 存储在 服务器，在哪生成呢？
//3.  什么时候生成Cookie ， 什么时候生成 Session？
