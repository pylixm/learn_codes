#
# m = 1
# if m:
#     n = 2
#     print(n)  # 2
#     print(m)  # 1
#     m = 3
#
#
# print(n)  # 2
# print(m)  # 3


l = [1, 2, 3, 4, 5]
sum = 0
# 遍历列表
for i in l:
    sum += i
print(sum)

# 使用内建函数range 遍历从1到5的整数
for i in range(1, 6):
    sum += i
print(sum)

# 条件语句主要用while
while i <= 5:
    sum += i
    i += 1
print(sum)

# 无限循环
while True:
    pass

# break 和 continue的使用
while True:
    if sum > 10:
        print('sum已大于10')
        break
        # continue
    sum += 1
    print(sum)