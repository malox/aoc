// Example program
#include <iostream>
#include <string>
#include <vector>
#include <cstdlib>
#include <fstream>
#include <vector>
#include <algorithm>

#define NL std::cout << "\n---------------------------------------------------\n"
#define DD(e) " " << #e << "[" << e << "]"
#define DS(e) " " << #e << "[" << static_cast<int>(e) << "]"
#define DB(e) " " << #e << "[" << std::boolalpha << e << "]"
#define PP(e) std::cout << e << std::endl

typedef std::vector<int> VI;
typedef std::vector<VI> VV;

void print(VI& v)
{
   for(size_t i=0; i<v.size();++i)
      std::cout << v[i] << " ";
   std::cout << std::endl;
}

void print(VV& vv)
{
   for(size_t i=0; i<vv.size();++i) {
      PP("vv" <<DD(i) << DD(vv[i].size()));
      print(vv[i]);
   }
}

void fill(VI& v)
{
    //int a[] = {0,2,7,0};
    int a[] = {4,10,4,1,8,4,9,14,5,1,14,15,0,15,3,5};
    size_t s = sizeof(a)/sizeof(int);
    for(size_t i=0; i<s;++i)
        v.push_back(a[i]);
}
size_t max(VI& v)
{
    size_t m = 0;
    for (size_t i=1; i<v.size();++i)
        if(v[i]>v[m])
            m=i;
    return m;
}

bool check(VV& vv, VI& vi, size_t aIter)
{
    bool aFound = false;
    size_t i=0;
    for(; i<vv.size() and not aFound;++i)
        aFound = std::equal(vi.begin(), vi.end(), vv[i].begin());
    if (aFound)
        PP("check" << DD(vv.size()) << DD(i) << DD(aIter-i+1) << DB(aFound));
    return aFound;
}

void t_main()
{
  NL;
  VI vi;
  VV vv;
  fill(vi);
  print(vi);
  //PP(DD(max(vi)) << DB(check(vv,vi)));
  //vv.push_back(vi);
  //PP(DB(check(vv,vi)));

  
  size_t aIter = 0;
  while(not check(vv,vi,aIter))
  {
      vv.push_back(vi);
      size_t m = max(vi);
      size_t blocks = vi[m];
     // PP(DD(m)<<DD(blocks));
      vi[m]=0;
      while(blocks>0)
      {
          if(m==(vi.size()-1))
              m=-1; // controlled overflow
          vi[++m]++;
          --blocks;
      }
      ++aIter;
  }
  NL;
  //vv.push_back(vi);
  //print(vv);
  //NL;
  PP(DD(aIter));
  NL;
}

int main()
{
  t_main();
  return 0;
}


