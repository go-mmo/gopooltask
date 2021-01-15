# gopooltask
using go and channel to limit task running

###Why
>有时候我们想并发的去做一些事情,比如并发加载文件,但是CPU会爆掉,我们又想限制
>任务的运行个数,自己再去写一个类似的channel与go的组合,又太麻烦,所以就有了这个项目.
>
>本来也是可以当作 gopool 池来用的,但是我发现性能貌似并不是很好,所以还是当作任务池来用吧.
>




#### notice
> This project using gomod(vgO) ,so there is a go.mod file in the root directory.
> 
------

