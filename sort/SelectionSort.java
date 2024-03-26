package sort;

import utils.Utils;

public class SelectionSort {

    public SelectionSort() {
        System.out.println("Initiated Selection sort");
    }

    // [4, 2, 7, 1, 3]

    public void sort(int[] arr) {
        for (int lastunsortedIndex = arr.length - 1; lastunsortedIndex >= 0; lastunsortedIndex--) {
            int localMax = Integer.MIN_VALUE;
            int localMaxIndex = 0;
            for (int i = 0; i <= lastunsortedIndex; i++){
                if (arr[i] > localMax){
                    localMax = arr[i];
                    localMaxIndex = i;
                }
            }
            Utils.swap(arr, lastunsortedIndex, localMaxIndex);
        }

        Utils.PrintArray(arr);
    }
}
