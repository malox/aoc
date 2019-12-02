#include <iostream>
#include <string>
#include <cmath>
#include <vector>


void printval(std::vector<int> v, int one, int two)
{
  size_t vsize = v.size();
  size_t it = 0;
  v[1]=one; v[2]=two;

  while(it < vsize)
  {
    int curr = v[it];
    if(curr == 99) { break; }

    int first = v[v[it+1]];
    int second = v[v[it+2]];
    
    v[v[it+3]] = (curr == 1) ? first + second : first * second;
    
    it+=4;  
  }

  std::cout << "v[0] for one=" << one << " two= " << two << " : " << v[0] << std::endl;
}


int main()
{
  int aInt;
  std::vector<int> v;
  while (std::cin >> aInt) 
  {
    if (std::cin.peek() == ',')
      std::cin.ignore();

    //std::cout << aInt << " ";
    v.push_back(aInt);
  }

  size_t vsize= v.size();
  size_t it = 0;

  for(int i=0; i < 100; ++i)
    for(int j=0; j <100; ++j)
      printval(v, i, j);
  
  return 0;
}
