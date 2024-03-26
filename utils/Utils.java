package utils;

public class Utils {
    public static void PrintArray(int[] arr){
        String comma = ", ";
        System.out.print("[");
        for (int i = 0; i < arr.length; i++){
            if (i == arr.length - 1) comma = "";
            System.out.print(arr[i] + comma);
        }
        System.out.println("]");
    }

    public static void swap(int[] arr, int from, int to){
        if (from == to || arr[from] == arr[to]) return;
        int temp = arr[from];
        arr[from] = arr[to];
        arr[to] = temp;
    }

    public static int[] generateRandomArray(int size){
        int[] arr = new int[size];
        for (int i = 0; i < size; i++){
            arr[i] = (int) (Math.random() * 100);
        }
        return arr;
    }
}
