package sort;

import utils.Utils;

public class InsertionSort {

    public InsertionSort(){
        System.out.println("Initiated Insertion sort");
    }

    public static void sort(int[] arr){
        for (int sortedIndex = 1; sortedIndex < arr.length ; sortedIndex++){
            int newElement = arr[sortedIndex];
            int i = sortedIndex;
            while(i > 0 && arr[i-1] > newElement){
                arr[i] = arr[i-1];
                i--;
            }
            arr[i] = newElement;
        }
        Utils.PrintArray(arr);
    }
}
