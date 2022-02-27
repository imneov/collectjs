# Thing of JSON


## Background
当我们用Go处理JSON时，经常需要不断的maral、unmarl，这里带来了巨大的损耗。因此，我们构建了一个直接处理json对象的库。

### 为了方便



## JSON

**JSON**(JavaScript Object Notation) 是一种轻量级的数据交换格式。 易于人阅读和编写。同时也易于机器解析和生成。 它基于[JavaScript Programming Language](http://www.crockford.com/javascript), Standard ECMA-262 3rd Edition - December 1999的一个子集。 JSON采用完全独立于语言的文本格式，但是也使用了类似于C语言家族的习惯（包括C, C++, C#, Java, JavaScript, Perl, Python等）。 这些特性使JSON成为理想的数据交换语言。

JSON建构于两种结构：

- “名称/值”对的集合（A collection of name/value pairs）。不同的语言中，它被理解为*对象（object）*，纪录（record），结构（struct），字典（dictionary），哈希表（hash table），有键列表（keyed list），或者关联数组 （associative array）。
- 值的有序列表（An ordered list of values）。在大部分语言中，它被理解为数组（array）。

这些都是常见的数据结构。事实上大部分现代计算机语言都以某种形式支持它们。这使得一种数据格式在同样基于这些结构的编程语言之间交换成为可能。

JSON具有以下这些形式：

对象是一个无序的“‘名称/值’对”集合。一个对象以 `{`左括号 开始， `}`右括号 结束。每个“名称”后跟一个 `:`冒号 ；“‘名称/值’ 对”之间使用 `,`逗号 分隔。

![img](ThingOfJSON.assets/object.png)

数组是值（value）的有序集合。一个数组以 `[`左中括号 开始， `]`右中括号 结束。值之间使用 `,`逗号 分隔。

![img](ThingOfJSON.assets/array.png)

值（*value*）可以是双引号括起来的字符串（*string*）、数值(number)、`true`、`false`、 `null`、对象（object）或者数组（array）。这些结构可以嵌套。

![img](ThingOfJSON.assets/value.png)

字符串（*string*）是由双引号包围的任意数量Unicode字符的集合，使用反斜线转义。一个字符（character）即一个单独的字符串（character string）。

字符串（*string*）与C或者Java的字符串非常相似。

![img](ThingOfJSON.assets/string.png)

数值（*number*）也与C或者Java的数值非常相似。除去未曾使用的八进制与十六进制格式。除去一些编码细节。

![img](ThingOfJSON.assets/number.png)

空白可以加入到任何符号之间。 以下描述了完整的语言。

![img](ThingOfJSON.assets/whitespace.png)
