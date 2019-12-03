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

class Parser
{
public:

   void parse(char c)
   {
       if(!_skip)
       {
           switch(c)
           {
               case '!' : _skip = true; break;
               case '{' : if(!_ignore) ++_groups ; else _garbage++; break;
               case '}' : if(!_ignore) _score += _groups-- ; else _garbage++; break;
               case '<' : if(!_ignore) _ignore = true ; else _garbage++; break;
               case '>' : _ignore = false; break;
               default : if(_ignore) _garbage++;
           }
           
       }
       else
           _skip = false;
   }
   
   bool _skip = false;
   bool _ignore = false;
   size_t _groups = 0;
   size_t _score = 0;
   size_t _garbage = 0;
   
};

int main(int argc, char *argv[])
{
  NL;
  char c;
  Parser p;
  while (std::cin >> c)
  {
    p.parse(c) ;
  }
  PP(DD(p._score) << DD(p._garbage));
  NL;
  return 0;
}



