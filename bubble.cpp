#include <iostream>
#include <vector>

using namespace std;

void swap(vector<int>& s,int i){
  int x = s[i];
  s[i]=s[i+1];
  s[i+1]=x;
}

int main(){
  vector<int> s {99,8,7,6,5,4,3,2,1};
  int n = s.size();

  for (int j=1;j<n;++j){
    for(int i=j-1;i>=0 && s[i]>s[i+1];i--){
      swap(s,i);
    }
  }
}

