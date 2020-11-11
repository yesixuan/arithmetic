/***
 * 在有序数组 arr 0 - n 中 寻找 target
 */
function binarySearch(arr, r = arr.length - 1, target) {
  let l = 0
  while (l <= r) {
    // const mid = l + Math.floor((r - l) / 2)
    const mid = Math.floor((l + r) / 2)
    if (arr[mid] === target) return mid
    if (arr[mid] < target) l = mid + 1
    if (arr[mid] > target) r = mid - 1
  }
  return -1
}

/**
 * 递归解法
 */
function recursionBinarySearch(arr, l = 0, r = arr.length - 1, target) {
  if (l > r) return -1
  const mid = l + Math.floor((r - l) / 2)
  if (arr[mid] === target) return mid
  if (arr[mid] < target) return recursionBinarySearch(arr, mid + 1, r, target)
  if (arr[mid] > target) return recursionBinarySearch(arr, l, mid - 1, target)
}

function swap(i, j, arr) {
  if (i === j) return
  var tmp = arr[i]
  arr[i] = arr[j]
  arr[j] = tmp
}

// 215
var findKthLargest = function(nums, k) {
  var L = 0, R = nums.length - 1
  while(true) {
    var l = L, r = R, i = l + 1
    // 一次左右分边
    while(i <= r) {
      if (nums[l] > nums[i]) {
        swap(i, r, nums)
        r--
      } else {
        swap(l, i, nums)
        l++
        i++
      }
    }
    if (l + 1 === k) break
    // while(nums[l] === nums[l - 1]) {
    //   l--
    //   if (l + 1 === k) break
    // }
    // if (l + 1 === k) break
    if (l + 1 < k) { // 往后找
      L = l + 1
      R = nums.length - 1
    } else { // 往前找
      R = r - 1
      L = 0
    }
  }
  return nums[l]
};

function judge (word1, word2) {
  let unMatch = 0
  if (Math.abs(word1.length - word2.length) > 1) return false
  for (let i = 0; i < word1.length; i++) {
    if (word1[i] !== word2[i]) {
      unMatch++
      if (unMatch > 1) {
        return false
      }
    }
  }
  if (unMatch === 0) return true
  return true
}

// 127
var ladderLength = function(beginWord, endWord, wordList) {
  const queue = [
    [beginWord, 0]
  ]
  const visited = {
    [beginWord]: true
  }
  while (queue.length) {
    const [word, time] = queue.shift()
    if (word === endWord) return time + 1
    for (let i = 0; i < wordList.length; i++) {
      if (judge(word, wordList[i])) {
        if (!visited[wordList[i]]) {
          queue.push([wordList[i], time + 1])
          visited[wordList[i]] = true
        }
      }
    }
  }
  return 0
};

console.log(
  ladderLength(
    "hit",
    "cog",
    ["hot","dot","dog","lot","log"]
  )
  // isPalindrome("AmannamA")
  // binarySearch([1, 2, 3, 4, 5], 4, 1),
  // recursionBinarySearch([1, 2, 3, 4, 5], 0, 4, 5),
  // findKthLargest(
  //   [3,2,3,1,2,4,5,5,6],
  //   4
  // )
)