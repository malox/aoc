#include <iostream>
#include <string>
#include <cmath>
#include <vector>
#include <map>
#include <algorithm>
#include <numeric>
#include <utility>

#define NL std::cout << std::endl << " ------------------------------------------ " << std::endl

#define DD(e) " " << #e << "[" << e << "]"
#define DB(e) " " << #e << "[" << std::boolalpha << e << "]"

/*
void check( int rown, int coln, std::vector<std::string> array, int i, int j, std::map<int, std::pair<int,int>> m)
{
    int iit = i, jit = j;
    const int rowf = rown -1; const int colf = coln -1;
    

    array[i][j] = 'X';

    //for(int inc = 1; inc < rown ; ++inc)
    /*{
    // Straight rows / columns    
    {  bool found = false;
       for(int k = j+1; k < coln ; k=k+inc)
         found = check(array, iit, k, found);
    }
    {  bool found = false;
       for(int k = j-1; k >= 0 ; k = k-inc)
         found = check(array, iit, k, found);
    }  
    {  bool found = false;
       for(int k = i+1; k < rown ; k=k+inc)
         found = check(array, k, jit, found);
    }
    {  bool found = false;
       for(int k = i-1; k >= 0 ; k = k-inc)
         found = check(array, k, jit, found);
    }
    } ?/
   {  bool found = false;
       for(int k = i+1; k  < rown ; ++k)
       {
          for(int t = j+1; t < coln ; ++t)
          {
              if(found)
              {
                  int f = std::gcd(iit, jit);
                  check(array, k, t, found);
              }
              else {
                 found = check(array, k, t, found);
                 iit = k-i;
                 jit = t-j;
              }
          }
       }
    }

     
    std::cout << "#### CHECKING FOR IDX " << DD(i) << DD(j) << DD(count(array)+1) << std::endl;     
    print(array);
}*/

class Matrix
{
public:

   size_t x;
   size_t y;
   
   const size_t rown;
   const size_t coln;
   
   typedef std::vector<std::string> grid_t;
   grid_t grid;
   typedef std::map<float, std::pair<int, int>> res_t;
   res_t points;
   
   Matrix(const grid_t& v, size_t i, size_t j) 
   : grid(v), x(i), y(j), rown(v.size()), coln(v.at(0).size())
   { grid[x][y] = 'X'; }
   
   static void Print(const grid_t& grid)
   {
       NL;
       for(auto &s : grid)
          std::cout << s << std::endl;
       NL; NL;
   }

   static size_t Count(const grid_t& grid)
   {
       size_t cc = 0;
       for(auto &s : grid)
          cc += std::count(s.begin(), s.end(), '#');
       return cc;
   }
   
   void check(int i, int j)
   {
      if(grid[i][j]=='#')
      {
          const float dif = ((j-y)!=0 ? (i-x)/(j-y) : 0.);

          auto found = points.emplace(std::piecewise_construct,
                       std::forward_as_tuple(dif), std::forward_as_tuple(i, j));
          std::cout << DD(i) << DD(j) << DD(dif) << DB(found.second) <<std::endl;
          if(!found.second)
              grid[i][j]='0';
      }
   }
   
   void clearPoints()
   {
       points.clear(); std::cout << " - - - " << std::endl;
   }

   void check()
   {
       clearPoints();
       for(int k = x-1; k >= 0 ; --k)
       {
          for(int t = y-1; t >= 0 ; --t)
              check(k, t);
          clearPoints();
          for(int t = y; t < coln ; ++t)
              check(k, t);
       }

       clearPoints();
       for(int k = x; k  < rown ; ++k)
       {
          for(int t = y-1; t >= 0 ; --t)
              check(k, t);
          clearPoints();
          for(int t = y; t < coln ; ++t)
              check(k, t);
       }
     
       std::cout << "#### CHECKING FOR IDX " << DD(x) << DD(y) << DD(Count(grid)+1) << std::endl;     
       Print(grid);
   }
};

void partone(Matrix::grid_t& v)
{
    
    const int rown = v.size();
    const int coln = v[0].size();
  
    Matrix::Print(v);
   
    //std::map<int, std::pair<int,int>> m;
    for(int i = 0; i < rown; ++i)
        for (int j = 0; j < coln; ++j)
            if(v[i][j]=='#')
            {
                Matrix m(v, i, j);
                m.check();
            }           
    
}

int main()
{
  std::string s;
  Matrix::grid_t v;
  while (std::cin >> s) 
  {
    //std::cout << ch << " ";
    v.push_back(s);
  }

  partone(v);
     
  return 0;
}
