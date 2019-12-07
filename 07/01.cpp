#include <iostream>
#include <string>
#include <cmath>
#include <vector>
#include <deque>
#include <algorithm>

#define DD(e) " " << #e << "[" << e << "]"
#define DB(e) " " << #e << "[" << std::boolalpha << e << "]"

struct Ampli
{
  std::vector<int> instr;
  std::deque<int> in;
  std::deque<int> out;
  int vsize ;
  int it = 0;
  bool stop = false;

  Ampli(std::vector<int> v, std::deque<int> dq) : instr(v), in(dq), vsize(v.size()) {}
  
int& getParam(int idx, bool param, std::vector<int>& v)
{
    return (param) ? v[v[idx]] : v[idx];
}

int pop(std::deque<int>& dq) const
{
    int a = std::move(dq.front());
    dq.pop_front();
    return a;
}

void halt()
{
   stop = true;
}

void printval()
{
  while(it < vsize)
  {
    int curr = instr[it];
    if(curr == 99) { /* std::cout << "found 99 at " << DD(it) << std::endl;*/ halt(); return ; }
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
        case 1: instr[instr[it+3]] = getParam(it+1, p1, instr) + getParam(it+2, p2, instr); it+=4; break;
        case 2: instr[instr[it+3]] = getParam(it+1, p1, instr) * getParam(it+2, p2, instr); it+=4; break;
        case 3: if(in.size()) { instr[instr[it+1]] = pop(in); it += 2;} else {return;}; break;
        case 4: out.push_back(getParam(it+1, p1, instr)) ; it += 2 ; break;
        case 5: (getParam(it+1, p1, instr)) ? it = getParam(it+2, p2, instr) : it+=3; break;
        case 6: (!getParam(it+1, p1, instr)) ? it = getParam(it+2, p2, instr) : it+=3; break;
        case 7: (getParam(it+1, p1, instr) < getParam(it+2, p2, instr)) ? instr[instr[it+3]]=1 : instr[instr[it+3]] = 0; it+=4; break;
        case 8: (getParam(it+1, p1, instr) == getParam(it+2, p2, instr)) ? instr[instr[it+3]]=1 : instr[instr[it+3]] = 0; it+=4; break;
        default:
           std::cout << "invalid" << DD(op) << " at idx " << it << std::endl;
           ++it;
    }
    //std::cout << DD(s) << DD(s.size()) << DD(op) << DB(p1) << DB(p2) << DB(p3) << DD(input) << DD(output) << DD(it) << std::endl; 
  }
  std::cout << "didn't find the halt instr" << std::endl;
  halt();
}

};

void partone(std::vector<int> v)
{
  int amax = 0;
  std::vector<int> in = {0,1,2,3,4};
  do {
      int out = 0;
      for(int ph : in)
      {
          std::deque<int> dq = {ph,out};
          Ampli amp(v, dq);
          amp.printval();
          if(amp.out.size()!=1) std::cout << "warning1 " << DD(amp.out.size()) << std::endl;
          out = amp.pop(amp.out);
      }
      amax = std::max(amax, out);
  } while (std::next_permutation(in.begin(),in.end()));
  
  std::cout << " max " << amax << std::endl;
}

void parttwo(std::vector<int> v)
{
  int amax = 0;
  std::vector<int> in = {5,6,7,8,9};

  do 
  {
      std::vector<Ampli> amps;
      for(size_t i = 0; i < 5 ; ++ i) { amps.push_back(Ampli(v, {in.at(i)})); }
      amps[0].in.push_back(0);
      
      std::deque<int> out;
      bool aRunning = true;
      while(aRunning)
      {
          aRunning = false;
          for(size_t i = 0; i < 5 ; ++ i)
          {
              Ampli & amp = amps[i];
              if(amp.stop) continue;
              aRunning = true;

              while(out.size()) { amp.in.push_back(amp.pop(out)); }
              
              amp.printval();
              if(amp.out.size()!=1) std::cout << "warning2 " << DD(amp.out.size()) << std::endl;
              out.push_back( amp.pop(amp.out) );
          }
      }
      amax = std::max(amax, out.front());
  } while (std::next_permutation(in.begin(),in.end()));
  
  std::cout << " max2 " << amax << std::endl; // 84088865
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

  partone(v);
  parttwo(v);

     
  return 0;
}
