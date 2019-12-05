// Example program

#include <iostream>
#include <sstream>
#include <fstream>
#include <string>

#include <map>
#include <vector>
#include <cstdlib>
#include <limits>
#include <algorithm>

#define NL std::cout << "\n---------------------------------------------------\n"
#define DD(e) " " << #e << "[" << e << "]"
#define DS(e) " " << #e << "[" << static_cast<int>(e) << "]"
#define DB(e) " " << #e << "[" << std::boolalpha << e << "]"
#define PP(e) std::cout << e << std::endl

const size_t ksize = 
  256u;
//  5u; // { 0, 1, 2, 3, 4}
static const std::vector<int> inp =
    {34,88,2,222,254,93,150,0,199,255,39,32,137,136,1,167};
//    {3, 4, 1, 5};

size_t getPos(size_t pos, size_t offset, size_t size)
{
    size_t res = pos + offset;
    while(res >= size) res -= size;
    return res;
}

void dump(const std::vector<int>& v)
{
    for(auto i : v)
        std::cout << " " << i;
    std::cout << std::endl;
    
}

void t_main()
{
    std::vector<int> v(ksize);
    for(size_t i =0; i<ksize; ++i)
        v[i] = i;
    
    size_t skip = 0;
    size_t pos = 0;
    
    for(size_t len : inp)
    {
        std::vector<int> tmp(len);
        for(size_t j=0; j<len; ++j)
            tmp.push_back(v[getPos(pos, j, ksize)]);

        std::reverse(tmp.begin(), tmp.end());
        for(size_t j=0; j<len; ++j)
            v[getPos(pos, j, ksize)] = tmp[j];
        
        dump(v);
        pos = getPos(pos, len + skip++, ksize);
        PP(DD(pos) << DD(skip));
    }

}

int main(int argc, char *argv[])
{
  NL;
  t_main();
  NL;
  return 0;
}



