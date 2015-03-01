# 16 3Sum Closest (Medium)
class Solution:
    # @return an integer
    def threeSumClosest(self, num, target):
        num = sorted(num)
        diff = num[0] + num[1] + num[len(num)-1] - target
        for x_index in range(len(num)-2):
            x = num[x_index]
            if x_index != 0 and x == num[x_index - 1]:
                continue
            expertVal = target - x
            y_index = x_index + 1
            z_index = len(num) - 1
            while z_index > y_index:
                realVal = num[y_index] + num[z_index]
                tmp = realVal - expertVal
                if tmp == 0:
                    return target
                elif tmp > 0:
                    z_index -= 1
                else:
                    y_index += 1
                if abs(tmp) < abs(diff):
                    diff = tmp
        return target + diff


if __name__ == "__main__":
    solu = Solution()
    num = [-1,2,1,-4]
    closestVal = solu.threeSumClosest(num, 1)
    print closestVal
