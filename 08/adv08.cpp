// Example program

#include <iostream>
#include <sstream>
#include <fstream>
#include <string>

#include <map>
#include <vector>
#include <cstdlib>
#include <limits>

#define NL std::cout << "\n---------------------------------------------------\n"
#define DD(e) " " << #e << "[" << e << "]"
#define DS(e) " " << #e << "[" << static_cast<int>(e) << "]"
#define DB(e) " " << #e << "[" << std::boolalpha << e << "]"
#define PP(e) std::cout << e << std::endl


typedef std::map<std::string, int> map_t;

bool Check(int Int, const std::string& Op, int Ref)
{
    if ( Op == "!=" ) {  return Int != Ref; } else 
    if ( Op == "<"  ) {  return Int <  Ref; } else 
    if ( Op == "<=" ) {  return Int <= Ref; } else 
    if ( Op == ">"  ) {  return Int >  Ref; } else 
    if ( Op == ">=" ) {  return Int >= Ref; } else 
    if ( Op == "==" ) {  return Int == Ref; } else 
    {  return false; }
}

int& getRegValue(map_t& imap, const std::string& reg)
{
    const auto it = imap.find(reg);
    if(it==imap.end())
    {
        int& val = imap[reg];
        val = 0;
        return val;
    }
    return it->second;
}

void t_main(std::string filename)
{
  NL;
  std::ifstream infile(filename.c_str());

  // b inc 5 if a > 1
  std::string var;
  std::string op;
  int inc = 0;
  std::string iff;
  std::string var2;
  std::string check;
  int ref = 0;
  
  map_t m;

  int gmax = std::numeric_limits<int>::min();

  while (infile >> var >> op >> inc >> iff >> var2 >> check >> ref)
  {
      int& regVal = getRegValue(m, var);
      int oldRegVal = regVal;
      int& checkVar = getRegValue(m, var2);
      
      const bool aCheck = Check(checkVar, check, ref);
      if(aCheck)
      {
          if(op=="inc") {
              regVal += inc;
          } else if(op=="dec") {
              regVal -= inc;
          }
      }
      
      gmax = std::max(gmax, regVal);
      
      PP(DD(var) << DD(op) << DD(inc) << DD(var2) << DD(check) << DD(ref) << " --" << DD(oldRegVal) << DD(regVal) << DD(ref) << DD(checkVar) << DB(aCheck));
  }
  
  int max = std::numeric_limits<int>::min();
  for(map_t::iterator it = m.begin(); it != m.end(); ++it)
      max = std::max(max, it->second);
  
  PP(DD(max) << DD(gmax));
      
  NL;
}

int main(int argc, char *argv[])
{
  if( argc > 1)
      t_main(argv[1]); 
  return 0;
}



