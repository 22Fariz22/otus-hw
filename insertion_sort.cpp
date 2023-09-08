#include <iostream>
using namespace std;

void insertionSort(int arr[], int n)
{
  int i,key,j;
  for (i = 1;i<n;i++){
    key = arr[i];
    j = i-1;

    while(j>=0 && arr[j]>key){
      arr[j+1]= arr[j];
      j = j-1;
    }
    arr[j+1]=key;

  }
}

int main(){
  int arr[] = {99,8,7,6,5,4,3,2,1};
  int n = sizeof(arr) / sizeof(arr[0]);
  insertionSort(arr,n);
}