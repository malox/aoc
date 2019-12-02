#include <iostream>
#include <string>
#include <cmath>

void t_main(int aMass, int & aTotal)
{
  int aFuel = std::floor((double)aMass/3) - 2;
  if(aFuel > 0) {
      t_main(aFuel, aTotal);
      aTotal += aFuel;
  }
}

int main()
{
   int aMass, aTotal=0;
   while (std::cin >> aMass) {
      t_main(aMass, aTotal);
  }

  std::cout << "Total " << aTotal << std::endl;
  return 0;
}
