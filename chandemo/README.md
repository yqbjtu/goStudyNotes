channel默认阻塞的，当发送数据到channel时，会发生阻塞，直到其他goroutine从该channel读取数据

发送方如果数据写入完毕，需要close channel， 用于通知接收方数据传递完毕。
往关闭的channel中写入数据，会报panic send on closed channel错误。
但是可以从关闭的channel中读取数据， 返回数据的默认值和false

通常用法是，发送方主动关闭channel，接收方通过多重返回值判断channel是否关闭，如果返回值是false，表示channel关闭了
