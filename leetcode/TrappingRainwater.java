package leetcode;

public class TrappingRainwater {
    public static int solve(int[] height){
        int processElements = 0;
        int area = 0;
        int h = 0;
        for (int i = 0; i < height.length; i++){
            for (int j = i; j < height.length; j++){
                if (height[j] >= height[i]){
                    processElements = j;
                    h = Math.min(height[j], height[i]);
                }
            }
            while((i + 1) < processElements){
                i++;
                area += h - height[i];
            }
        }
        return area;
    }
}
