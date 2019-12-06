#include <iostream>
#include <string>
#include <cmath>
#include <vector>

#define DD(e) " " << #e << "[" << e << "]"
#define DB(e) " " << #e << "[" << std::boolalpha << e << "]"

int& getParam(int idx, bool param, std::vector<int>& v)
{
    return (param) ? v[v[idx]] : v[idx];
}

void printval(std::vector<int> v, int input, int& output)
{
  int vsize = v.size();
  int it = 0;

  while(it < vsize)
  {
    int curr = v[it];
    if(curr == 99) { std::cout << "found 99 at idx " << it << std::endl; break; }

    std::string s = std::to_string(curr);
    
    int op=0;
    bool p1 = true, p2 = true, p3 = true;
    switch(s.size())
    {
        case 1:
        case 2: op = curr; break;
        case 3: p1 = false; op = s[2]-'0'; break;
        case 4: p2 = false; p1 = (s[1]=='0') ; op = s[3]-'0'; break;
        case 5: p3 = false; p2 = (s[1]=='0') ; p1 = (s[2]=='0') ; op = s[4]-'0'; break;
    }

    switch(op)
    {
        case 1: v[v[it+3]] = getParam(it+1, p1, v) + getParam(it+2, p2, v); it+=4; break;
        case 2: v[v[it+3]] = getParam(it+1, p1, v) * getParam(it+2, p2, v); it+=4; break;
        case 3: v[v[it+1]] = input; it += 2; break;
        case 4: output = getParam(it+1, p1, v) ; it += 2 ; break;
        case 5: (getParam(it+1, p1, v)) ? it = getParam(it+2, p2, v) : it+=3; break;
        case 6: (!getParam(it+1, p1, v)) ? it = getParam(it+2, p2, v) : it+=3; break;
        case 7: (getParam(it+1, p1, v) < getParam(it+2, p2, v)) ? v[v[it+3]]=1 : v[v[it+3]] = 0; it+=4; break;
        case 8: (getParam(it+1, p1, v) == getParam(it+2, p2, v)) ? v[v[it+3]]=1 : v[v[it+3]] = 0; it+=4; break;
        default:
           std::cout << "invalid" << DD(op) << " at idx " << it << std::endl;
           ++it;
    }
    
    //std::cout << DD(s) << DD(s.size()) << DD(op) << DB(p1) << DB(p2) << DB(p3) << DD(input) << DD(output) << DD(it) << std::endl; 
  }

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

  printval(v, 1, aInt);
  std::cout << "first : " << aInt << std::endl;
  printval(v, 5, aInt);
  std::cout << "second : " << aInt << std::endl;
     
  return 0;
}
