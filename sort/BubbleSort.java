package sort;

import utils.Utils;

public class BubbleSort {

    public BubbleSort() {
        System.out.println("Initiated Bubble sort");
    }

    public void sort(int[] arr) {

        for (int lastunsortedIndex = arr.length - 1; lastunsortedIndex >= 0; lastunsortedIndex--) {
            for (int i = 0; i < lastunsortedIndex - 1; i++) {
                if (arr[i] > arr[i + 1]) {
                    Utils.swap(arr, i, i + 1);
                }
            }
        }
        Utils.PrintArray(arr);
    }

}