class Solution {
public:
    int jumpFloorII(int number) {
        //递归
        if (number <= 0) {
            return -1; //不用跳
        } else if (number == 1) {
            return 1;  //跳一次即可
        } else {
            return 2 * jumpFloorII(number-1);
        }
    }
};