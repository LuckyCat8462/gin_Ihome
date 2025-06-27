package Learning

//go语言中有两个特殊函数：  —— 首字母小写，包外可见 。
//
//1. main()   —— 项目的入口函数
//2. init() —— 当导包，但没有在程序中使用。 在main() 调用之前，自动被调用。
//   - 查看：光标置于 MySQL包的 “mysql” 上。 使用 Ctrl-鼠标左键。 看到源码。 在 driver.go 底部包含 init() 函数的 定义。
//   - init() 作用：实现注册 MySQL 驱动。

//init()函数会先于main（）函数被调用
