package leetcode;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

public class Pairs {
    public static int[] solve(int[] arr, int sum) {
        for (int i = 0; i < arr.length; i++) {
            for (int j = i + 1; j < arr.length; j++) {
                if (arr[i] + arr[j] == sum)
                    return new int[] { arr[i], arr[j] };
            }
        }
        return new int[0];
    }

    // { 10, 5, 2, 3, -6, 9, 11 };

    // Map : -6 : 10, -1 : 5,

    public static int[] improvedSolve(int[] arr, int sum) {
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < arr.length; i++) {
            if (map.containsKey(Integer.valueOf(arr[i])))
                return new int[] { arr[i], map.get(arr[i]) };
            map.put(Integer.valueOf(sum - arr[i]), Integer.valueOf(arr[i]));
        }

        return new int[0];
    }

    /**
     * 
     * @param arr - should be sorted
     * @param sum
     * @return
     */
    public static List<List<Integer>> sortedSolve(int[] arr, int sum, int currentElement){

        if (arr.length < 3) return new ArrayList<>();
        int left = 0;
        int right = arr.length - 1;
        List<List<Integer>> returnArray = new ArrayList<>();
        
        while(left < right) {
            if ((arr[left] + arr[right]) > sum) right--;
            else if ((arr[left] + arr[right]) < sum) left++;
            else {
                returnArray.add(Arrays.asList(currentElement, arr[left], arr[right]));
                right--;
            }
        }

        return returnArray;
    }
}
