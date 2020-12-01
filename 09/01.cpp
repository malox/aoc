#include <iostream>
#include <string>
#include <cmath>
#include <vector>
#include <deque>
#include <algorithm>

#define DD(e) " " << #e << "[" << e << "]"
#define DB(e) " " << #e << "[" << std::boolalpha << e << "]"

typedef enum parammode
{
    kPos = 0,
    kImm,
    kRel
} pm_t;

typedef enum opmode
{
    kRead = 0,
    kWrite
} opm_t;

struct Ampli
{
  std::vector<long> instr;
  std::deque<long> in;
  std::deque<long> out;
  long vsize ;
  long it = 0;
  long offset = 0;
  bool stop = false;

  Ampli(std::vector<long> v, std::deque<long> dq) : instr(v), in(dq), vsize(v.size()) {}
  
  long& getParam(long idx, pm_t param, opm_t opm) 
  {
      switch(param)
      {
          case kPos: return instr[instr[idx]];
          case kImm: return kWrite == opm ? instr[instr[idx]] : instr[idx];
          case kRel: return instr[instr[idx+offset]];
      }
  }

  long pop(std::deque<long>& dq) const
  {
    long a = std::move(dq.front());
    dq.pop_front();
    return a;
  }

  void halt()
  {
     stop = true;
  }

   pm_t extrMode(const char ch) const
   {
       //std::cout << "extrMode " << ch << std::endl;
       switch(ch)
       {
           case '0': return kPos;
           case '1': return kImm;
           case '2': return kRel;
       }
   }
   
   int extrOpCode(const char ch) const
   {
       return ch-'0';
   }
   

void run()
{
  while(it < vsize)
  {
    long curr = instr[it];
    if(curr == 99) { /* std::cout << "found 99 at " << DD(it) << std::endl; */ halt(); return ; }
    std::string s = std::to_string(curr);
    
    long op=0;
    pm_t p1 = kPos, p2 = kPos, p3 = kPos;
      
    switch(s.size())
    {
        case 1:
        case 2: op = curr; break;
        case 3: p1 = extrMode(s[0]); op = extrOpCode(s[2]); break;
        case 4: p2 = extrMode(s[0]); p1 = extrMode(s[1]); op = extrOpCode(s[3]); break;
        case 5: p3 = extrMode(s[0]); p2 = extrMode(s[1]) ; p1 = extrMode(s[2]) ; op = extrOpCode(s[4]); break;
    }

    //std::cout << DD(s) << DD(s.size()) << DD(op) << DD(p1) << DD(p2) << DD(p3) << DD(in.size()) << DD(out.size()) << DD(it) << std::endl; 
    switch(op)
    {
        case 1: getParam(it+3, kPos, kWrite) = getParam(it+1, p1, kRead) + getParam(it+2, p2, kRead); it+=4; break;
        case 2: getParam(it+3, kPos, kWrite) = getParam(it+1, p1, kRead) * getParam(it+2, p2, kRead); it+=4; break;
        case 3: if(in.size()) { getParam(it+1, p1, kWrite) = pop(in); it += 2;} else {return;}; break;
        case 4: /*std::cout << getParam(it+1, p1, kRead) << std::endl ;*/ out.push_back(getParam(it+1, p1, kRead)) ; it += 2 ; break;
        case 5: (getParam(it+1, p1, kRead)) ? it = getParam(it+2, p2, kRead) : it+=3; break;
        case 6: (!getParam(it+1, p1, kRead)) ? it = getParam(it+2, p2, kRead) : it+=3; break;
        case 7: getParam(it+3, p3, kWrite) = (getParam(it+1, p1, kRead) <  getParam(it+2, p2, kRead)) ? 1 : 0; it+=4; break;
        case 8: getParam(it+3, p3, kWrite) = (getParam(it+1, p1, kRead) == getParam(it+2, p2, kRead)) ? 1 : 0; it+=4; break;
        case 9: offset = getParam(it+1, p1, kRead) ; it += 2 ; break;
        default:
           std::cout << "invalid" << DD(op) << " at idx " << it << std::endl;
           ++it;
    }
    //std::cout << DD(s) << DD(s.size()) << DD(op) << DD(p1) << DD(p2) << DD(p3) << DD(in.size()) << DD(out.size()) << DD(it) << std::endl; 
  }
  std::cout << "didn't find the halt instr" << std::endl;
  halt();
}

};

void oldpartone(std::vector<long> v)
{
  std::cout << "oldpartone" << std::endl;
  long amax = 0;
  std::vector<long> in = {0,1,2,3,4};
  do {
      long out = 0;
      for(long ph : in)
      {
          std::deque<long> dq = {ph,out};
          Ampli amp(v, dq);
          amp.run();
          if(amp.out.size()!=1) std::cout << "warning1 " << DD(amp.out.size()) << std::endl;
          out = amp.pop(amp.out);
      }
      amax = std::max(amax, out);
  } while (std::next_permutation(in.begin(),in.end()));
  
  std::cout << " max " << amax << std::endl;
}

void oldparttwo(std::vector<long> v)
{
  std::cout << "oldparttwo" << std::endl;
  long amax = 0;
  std::vector<long> in = {5,6,7,8,9};

  do 
  {
      std::vector<Ampli> amps;
      for(size_t i = 0; i < 5 ; ++ i) { amps.push_back(Ampli(v, {in.at(i)})); }
      amps[0].in.push_back(0);
      
      std::deque<long> out;
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
              
              amp.run();

              if(amp.out.size()!=1) std::cout << "warning2 " << DD(amp.out.size()) << std::endl;
              out.push_back( amp.pop(amp.out) );
          }
      }
      amax = std::max(amax, out.front());
  } while (std::next_permutation(in.begin(),in.end()));
  
  std::cout << " max2 " << amax << std::endl; // 84088865
}


void partone(std::vector<long> v)
{
  std::cout << "partone" << std::endl;

  std::deque<long> dq = {};
  Ampli amp(v, dq);
  //while(!amp.stop)
  {
  //   amp.in.insert(amp.in.end(), amp.out.begin(), amp.out.end());

     amp.out.clear();
     amp.run();
  }
  



  for( const long& it : amp.out)
    std::cout << it ;
  std::cout << std::endl;
  
  std::cout << DD(amp.out.size()) << std::endl;
}


int main()
{
  long along;
  std::vector<long> v;
  while (std::cin >> along) 
  {
    if (std::cin.peek() == ',')
      std::cin.ignore();

    //std::cout << along << " ";
    v.push_back(along);
  }

  oldparttwo(v);
  oldpartone(v);

  //partone(v);
     
  return 0;
}
