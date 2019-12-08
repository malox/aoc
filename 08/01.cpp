#include <iostream>
#include <string>
#include <cmath>
#include <vector>
#include <deque>
#include <map>
#include <algorithm>

#define DD(e) " " << #e << "[" << e << "]"
#define DB(e) " " << #e << "[" << std::boolalpha << e << "]"

const int kLimit = 25*6;

struct Layer
{
    void add(int code)
    {
        m[code]++;
        v.push_back(code);
    }

    std::vector<int> v;
    std::map<int, int> m;
};

void partone(std::vector<Layer>& layers)
{
  std::sort(layers.begin(), layers.end(), [](Layer& one, Layer& two) {return one.m[0]<two.m[0];});

  for(int it = 0 ; it < layers.size() ; ++it)
  {
      std::cout << "Part one : Layer[" << it << "] val[" << layers[it].m[1]*layers[it].m[2] << "]";
      for(const auto& pix : layers[it].m)
          std::cout << " -" << DD(pix.first) << DD(pix.second);
      std::cout << std::endl;
      break;
  }
}

void partwo(std::vector<Layer>& layers)
{
    std::vector<int> img(kLimit);
    for(int it = 0 ; it < kLimit ; ++it)
        for(int curr = 0; curr < layers.size(); ++curr)
            if(layers[curr].v[it]!=2) 
            {
                //std::cout << "found at" << DD(it) << DD(curr) << std::endl;
                img[it] = layers[curr].v[it]; 
                break;
            }

    std::cout << std::endl;
    for(int it = 0 ; it < kLimit ; ++it)
        std::cout << ((img[it]==1)? "0" : " ") << (((it+1)%25==0)? "\n": "");
    std::cout << std::endl;
}

int main()
{
  char ch;
  std::vector<int> v;
  while (std::cin >> ch) 
  {
    if (std::cin.peek() == ',')
      std::cin.ignore();

    //std::cout << aInt << " ";
    v.push_back(ch-'0');
  }

  const int vsize = v.size();
//  std::cout << DD(v[0]) << DD(vsize) << DD(v[vsize-1]) << std::endl;


  std::vector<Layer> layers(vsize/kLimit);
  for(int it = 0 ; it < vsize ; ++it)
  {
      Layer& lay = layers[it/kLimit];
      lay.add(v[it]);
  }
  

  partwo(layers);
  partone(layers);

  return 0;
}
