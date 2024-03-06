1. 禁用F12+右键 => 可以在浏览器的更多工具选项卡打开开发人员工具
2. 反调试 => 根据debugger触发的堆栈找到代码所在位置，在本地替换版本中删除这部分反调代码
3. js混淆 => 寻找并发现字符串映射规律，判断出字符串是统一存在一个数组里，取字符串时先算一个偏移量后从数组取出
4. 从Game Over！字符串定位到只有游戏通过才会执行的的可疑代码
``` JAVASCRIPT
g[h(432)][h(469)] = function(x) {
    var n = h
      , e = x ? "game-won" : n(443)
      , t = x ? s0(n(439), "V+g5LpoEej/fy0nPNivz9SswHIhGaDOmU8CuXb72dB1xYMrZFRAl=QcTq6JkWK4t3") : n(453);
    this[n(438)][n(437)].add(e),
    this[n(438)][n(435)]("p")[-1257 * -5 + 9 * 1094 + -5377 * 3].textContent = t
}
```
最不用动脑的办法是在可疑函数里打断点，猜测形参x是胜利与否，故意输掉并修改函数传入的x为true，假装获胜触发游戏通关的逻辑