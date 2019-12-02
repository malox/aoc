#include <iostream>
#include <string>
#include <cmath>

int main()
{
   int aFuel, aTotal=0;
   while (std::cin >> aFuel) {
      //std::cout << aFuel << std::endl;
      aTotal += (std::floor((double)aFuel/3) - 2);
  }

  std::cout << "Total " << aTotal << std::endl;
  return 0;
}
