package data;

import utils.Utils;

public class Data {
    private static final int[] sortArr = {20, 35, -15, 7, 55, 1, 22};

    public static int[] getSortArr(){
        return sortArr;
    }

    public static int[] generateRandomArray(int size){
        return Utils.generateRandomArray(size);
    }

}
