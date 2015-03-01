class Solution:
    # @return a list of lists of length 3, [[val1,val2,val3]]
    def threeSum(self, num):
        results = []
        num = sorted(num)
        for x_index in range(len(num)-2):
            x = num[x_index]
            # !!! x_index != 0
            if x_index != 0 and x == num[x_index - 1]:
                continue
            expertVal = 0 - x
            y_index = x_index + 1
            z_index = len(num) - 1
            while z_index > y_index:
                realVal = num[y_index] + num[z_index]
                if realVal == expertVal:
                    a_result = [x, num[y_index], num[z_index]]
                    results.append(a_result)
                    y_index += 1
                    z_index -= 1
                    # !!! avoid duplicate solutions
                    # version 1
                    #while z_index > y_index and num[y_index] == num[y_index-1] and num[z_index] == num[z_index+1]:
                    #    y_index += 1
                    #    z_index -= 1
                    # version 2
                    while z_index > y_index and num[y_index] == num[y_index-1]:
                        y_index += 1
                    while z_index > z_index and num[z_index] == num[z_index+1]:
                        z_index -= 1
                elif realVal > expertVal:
                    z_index -= 1
                else:
                    y_index += 1
        return results


if __name__ == "__main__":
    solu = Solution()
    num = [-1,0,1,2,-1,-4]
    results = solu.threeSum(num)
    print results
