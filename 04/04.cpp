#include <iostream>
#include <string>
#include <map>

 
bool t_check(int i)
{
    bool adouble = false;
    const std::string s = std::to_string(i);
    for ( size_t j = 0; j < s.size()-1; ++j )
    {
        if( s[j+1] <  s[j]) { return false; }
        if( s[j+1] == s[j]) { adouble = true; }  
    }
    return adouble;
}

void u_check(int i)
{
    std::cout << " checking " << i << " : " << std::boolalpha << t_check(i) << std::endl;
}


bool z_check(int i)
{
    std::map<int,int> m;
    const std::string s = std::to_string(i);
    for ( size_t j = 0; j < s.size()-1; ++j )
    {
        if( s[j+1] <  s[j]) { return false; }
        if( s[j+1] == s[j]) { m[(int)s[j]]++; }  
    }
    
    for(auto j : m)
        if(j.second==1)
            return true;
    return false;
}


void w_check(int i)
{
    std::cout << " z checking " << i << " : " << std::boolalpha << z_check(i) << std::endl;
}

int main()
{
    u_check(111111);
    u_check(223450);
    u_check(123789);

    w_check(112233);
    w_check(123444);
    w_check(111122);

    size_t count = 0, count_bis=0;
    for (size_t j=353096; j<=843212; ++j)
    {
        if(t_check(j))  
            count++;
        if(z_check(j))
            count_bis++;
    }

    std::cout << "count " << count << " - bis " << count_bis << std::endl;
    return 0;
}
