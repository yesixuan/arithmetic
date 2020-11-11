## 匹配括号是否配对（20）

利用栈的特性

## 利用队列进行层序遍历计算最短路径

```js
var numSquares = function(n) {
  // 利用队列的性质
  var queue = [
    [n, 0]
  ]
  var cache = {
    [n]: true
  }
  while (queue.length) {
    let [ num, time ] = queue.shift()

    if (num === 0) return time

    for (let i = 1; num - i * i >= 0; i++) {
      if (!cache[num - i * i]) {
        queue.push([num - i * i, time + 1])
        cache[num - i * i] = true
      }
    }
  }
  return '无解'
};
```