# go_string

1. 字符串转换
   - ToInt64(s string) int64 - 字符串转int64
   - ToInt(s string) int - 字符串转int
   - ToInt32(s string) int32 - 字符串转int32
   - ToFloat64(s string) float64 - 字符串转float64
2. 字符串拼接
   - Join(sep string, strs ...string) string - 用分隔符拼接多个字符串
3. 字符串截取
   - Left(s string, n int) string - 从左边取n个字符
   - Right(s string, n int) string - 从右边取n个字符
   - TrimLeft(s string, n int) string - 从左边删除n个字符
   - TrimRight(s string, n int) string - 从右边删除n个字符
4. 去除空格
   - Trim(s string) string - 去除首尾空格
5. 前缀/后缀/包含判断
   - HasPrefix(s, prefix string) bool - 判断前缀
   - HasSuffix(s, suffix string) bool - 判断后缀
   - Contains(s, substr string) bool - 判断是否包含

6. 其他实用函数
   - IsEmpty(s string) bool - 判断是否为空或只有空格
   - Reverse(s string) string - 反转字符串
   - ToUpper(s string) string - 转大写
   - ToLower(s string) string - 转小写
   - Replace(s, old, new string) string - 替换所有匹配
   - (s, sep string) []string - 按分隔符分割
