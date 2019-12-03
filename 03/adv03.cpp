#include <iostream>
#include <sstream>
#include <fstream>
#include <string>
#include <map>
#include <vector>
#include <algorithm> 
#include <cmath>
#include <tuple>

typedef std::tuple<int, int, int> point_t;
class Grid
{
public:
   int _steps = 0;
   std::vector<point_t> _points;
   
   void addPoint(int x, int y, const std::string& p) 
   {
        //std::cout << "Processing [" << p <<"] new point [" << x << " : " << y << "]" << std::endl;
       _points.emplace_back(x, y, _steps++);
   }

   void addPoint(const std::string& p) 
   {

       const auto cur = _points.crbegin();
       const int row = std::get<0>(*cur), col = std::get<1>(*cur);
       const char& c = p[0];
       int j = std::stoi(p.substr(1));
       for(int i = 1; i <= j ; ++i)
       switch(c) 
       {
           case 'R' : addPoint(row, col+i, p); break;
           case 'L' : addPoint(row, col-i, p); break;
           case 'U' : addPoint(row-i, col, p); break;
           case 'D' : addPoint(row+i, col, p); break;
       }
   }
   
   Grid() { addPoint(0,0, "origin");}

   void dump() {
     std::cout << "--" << std::endl;
     for(const point_t& p : _points) 
        std::cout << "Pos " << std::get<0>(p) << " " << std::get<1>(p) << " - " << std::get<2>(p) << std::endl;
     std::cout << "--" << std::endl;
   }
};


void t_main(const std::string& fname)
{
  std::ifstream infile(fname.c_str());

  std::string instr;
  std::getline(infile, instr);
  std::istringstream one(instr);
  std::getline(infile, instr);
  std::istringstream two(instr);

  Grid ga, gb;
  while (std::getline(one, instr, ','))
      ga.addPoint(instr);

  std::cout << "--" << std::endl;

  while (std::getline(two, instr, ','))
      gb.addPoint(instr);

  //ga.dump();
  //gb.dump();

  std::cout << "--" << std::endl;

  std::vector<point_t> result;
  for(const point_t& p : ga._points)
      for(const point_t& q : gb._points)
          if(std::get<0>(p)==std::get<0>(q) &&std::get<1>(p)==std::get<1>(q)
      && std::get<0>(p)!=0 &&std::get<1>(p)!=0)
              result.emplace_back(std::get<0>(p), std::get<1>(p), std::get<2>(p) + std::get<2>(q));
  
  int best=1234567890;
  int steps = best;
  for(const point_t& p : result) {
     int locbest = std::abs(std::get<0>(p)) + std::abs(std::get<1>(p));
     int locstep = std::get<2>(p);
     std::cout << "intersect at " << std::get<0>(p) << "  " << std::get<1>(p)  << " : sum " << locbest << " steps " << locstep << std::endl;
     best = std::min(locbest, best);
     steps = std::min(locstep, steps);
  }
 
   std::cout << "--" << std::endl;
   std::cout << "best : " << best << " - steps " << steps << std::endl; 
}

int main(int argc, char *argv[])
{
  if( argc > 1)
      t_main(argv[1]); 
  return 0;
}