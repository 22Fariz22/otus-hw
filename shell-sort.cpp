#include <iostream>
using namespace std;

void shellSort(int arr[],int n){

  for(int gap=n/2;gap>0;gap/=2){
    for(int i =gap;i<n;i+=1){
        int tmp = arr[i];
        int j;
        for(j=i;j>=gap && arr[j-gap]>tmp;j-=gap){
              arr[j] = arr[j-gap];
        }
        arr[j]=tmp;
    }
  }

}


int main(){
  int arr[] = {99,8,7,6,5,4,3,2,1};
  int n = sizeof(arr)/sizeof(arr[0]);

  shellSort(arr,n);
}