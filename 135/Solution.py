class Solution:
    # @param ratings, a list of integer
    # @return an integer
    def candy(self, ratings):
        # peak_index represents the index of the last rate in some peak
        if ratings:
            former = ratings[0]
        else:
            return 0
        down = True
        isPeak = True
        peak_index = 0
        candys = [ 0 for i in range(len(ratings))]
        candys[0] = 1
        for index in range(1, len(ratings) + 1):
            if index == len(ratings):
                rate = None
            else:
                rate = ratings[index]
            if rate != None and rate < former:
                if not down:
                    peak_index = index - 1
                    isPeak = True
                else:
                    isPeak = False
                candys[index] = candys[index-1] - 1
                down = True
            elif rate != None and rate > former or not rate:
                if down:
                    candys[index-1] = 1
                    print "index=%d" %index
                    for i in range(1, index-peak_index):
                        if ratings[index-i] < ratings[index-1-i]:
                            tmp = candys[index-i] + 1
                            if index-1-i == peak_index:
                                if tmp > candys[peak_index]:
                                    candys[peak_index] = tmp
                            else:
                                candys[index-1-i] = tmp
                        else:
                            candys[index-1-i] = 1
                if rate != None and rate > former:
                    candys[index] = candys[index-1] + 1
                    down = False
            else:
                candys[index] = 1
            former = rate
        total = 0
        for addend in candys:
            total += addend
        return total

if __name__ == "__main__":
    #ratings = [100, 90, 80, 80, 80, 110, 110, 120, 130, 130, 10, 9, 8, 7, 6, 5, 5]
    #ratings = [3, 3, 4, 5, 7, 7, 6]
    #ratings = [1, 2]
    #ratings = [1, 0]
    ratings = [2, 3, 0, 2]
    solu = Solution()
    minimumSum= solu.candy(ratings)
    print minimumSum

