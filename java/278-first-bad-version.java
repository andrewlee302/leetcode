/* The isBadVersion API is defined in the parent class VersionControl.
      boolean isBadVersion(int version); */

public class Solution extends VersionControl {
    public int firstBadVersion(int n) {
        if (!isBadVersion(n)) return -1;
        int l = 1, r = n, mid;
        while (l <= r) { 
            if (l == r) { return l; }
            mid = l + (r - l) / 2;
            if (isBadVersion(mid)) {
                r = mid ;
            } else {
                l = mid + 1;
            }
        }
        return -1;
    }
}
