package leetcode;

import java.util.HashSet;
import java.util.Set;

public class LongestConsecutive {
    public LongestConsecutive(){
        System.out.println("Initialized LongestConsecutive");
    }

    public static int solve(int[] nums) {
        Set<Integer> set = new HashSet<>();

        // add all the elements in set
        for (int i = 0; i < nums.length; i++)
            set.add(nums[i]);

        int maxBand = 0;
        // iterate over elements
        for (int i = 0; i< nums.length; i++){
            int temp = nums[i];
            int currBand = 0;

            if (set.contains(temp - 1))
                continue;
            while(set.contains(temp++))
                currBand++;
            maxBand = Math.max(maxBand, currBand);
        }

        return maxBand;
    }
}
