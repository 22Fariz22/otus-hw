#include <iostream>
#include <vector>
#include <sstream>

using namespace std;

string builder(int i, int n)
{
  stringstream ss;
  string str_i = to_string(i);

  if (str_i.size() < n)
  {
    for (int j = 0; j < n - str_i.size(); ++j)
    {
      ss << "0";
    }
  }

  string s = ss.str() + str_i;

  return s;
}

long int builderNForGen(int fullLen)
{
  stringstream ss;
  ss << "1";

  for (int i = 0; i < fullLen; ++i)
  {
    ss << "0";
  }

  long intNGen = stol(ss.str());

  return intNGen;
}

bool checkIfIsLucky(int fullLen, string ticket)
{
  int leftSum = 0;
  int rightSum = 0;

  for (int i = 0; i < fullLen; ++i)
  {
    char el(ticket[i]);

    if (i < fullLen / 2)
    {
      int leftInt = (int)el;
      leftSum += leftInt;
    }
    else
    {
      int rightInt = (int)el;
      rightSum += rightInt;
    }
  }

  if (leftSum == rightSum)
  {
    return true;
  }
  else
  {
    return false;
  }
}

int main()
{
  int n;
  cout << "введите N: " << endl;
  cin >> n;
  int fullLen = n * 2;
  long int cnt = 0;

  long int intNGen = builderNForGen(fullLen);

  for (int i = 0; i < intNGen; ++i)
  {
    string s = builder(i, fullLen);

    if (checkIfIsLucky(fullLen, s) == true)
    {
      cnt += 1;
      cout << cnt << endl;
    }
  }
  cout << "cnt happy tickets: " << cnt << endl;
}
