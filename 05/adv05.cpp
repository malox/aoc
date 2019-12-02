// Example program
#include <iostream>
#include <string>
#include <vector>
#include <cstdlib>
#include <fstream>
#include <vector>

#define NL std::cout << "\n---------------------------------------------------\n"
#define DD(e) " " << #e << "[" << e << "]"
#define DS(e) " " << #e << "[" << static_cast<int>(e) << "]"
#define DB(e) " " << #e << "[" << std::boolalpha << e << "]"
#define PP(e) std::cout << e << std::endl

void print(std::vector<int>& v)
{
   for(size_t i=0; i<v.size();++i)
      std::cout << v[i] << " ";
   std::cout << std::endl;
}

void t_main()
{
  NL;
  int a=0;
  std::vector<int> v;
  std::ifstream infile("adv05i.txt");
  while (infile >> a)
  {
    v.push_back(a);
  }
  //for(size_t i=0; i<v.size();++i)
      //PP(DD(i)<<DD(v[i]));
  PP(DD(v.size()));
  NL;

  int aCount=0;
  int aNext=0;
  bool aContinue=true;
  while(aContinue)
  {
      //PP("BEG - " << DD(aCount) << DD(aNext) << DD(v[aNext]));
      if(aNext >=0 and aNext < v.size())
      {
          int aNew = v[aNext];
          ++v[aNext];
          ++aCount;

          //PP(DD(aCount) << DD(aNext));
          aNext += aNew;
          //print(v);
      }
      else
      {
          aContinue=false;
      }
      //PP(" END- " << DD(aCount) << DD(aNext) << DD(v[aNext]));
  }
  PP(DD(aCount));
  NL;
}

int main()
{
  t_main();
  return 0;
}


