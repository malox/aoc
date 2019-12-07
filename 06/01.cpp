#include <iostream>
#include <string>
#include <cmath>
#include <vector>
#include <map>
#include <set>

#define DD(e) " " << #e << "[" << e << "]"
#define DB(e) " " << #e << "[" << std::boolalpha << e << "]"


struct obj
{
    std::string n;
    obj* p;
    std::vector<obj*> c;

};

std::map<std::string, obj>& get()
{
    static std::map<std::string, obj> m;
    return m;
}

obj* get(std::string name)
{
    std::map<std::string, obj>& m = get();
    std::map<std::string, obj>::iterator it = m.find(name);
    if(it == m.end())
    {
        obj* o = &m[name];
        o->n = name;
        return o;
    }
    else
        return &m[name];
}

int find(obj* from, std::string to, std::set<std::string>& s)
{
    //std::cout << "I am " << from->n << " looking for " << to << DD(s.size()) << std::endl;
    std::set<int> ret;
    if(from->n==to) { /*std::cout << " ... and I just found myself!" << std::endl;*/ return 0;}

    s.insert(from->n);
    if(from->p && s.find(from->p->n)==s.end())
    {
       int pf = find(from->p, to, s);
       if(pf>=0) ret.insert(pf);
    }
    for(obj* ch : from->c)
    {
       if(s.find(ch->n)!=s.end()) {continue;}
       int cf = find(ch, to, s);
       if(cf>=0) ret.insert(cf);
    }
    s.erase(from->n);
    if(!ret.empty()) return *ret.begin()+1;
    return -1;
}



size_t tot()
{
    size_t count = 0;
    for(auto it : get())
    {
        obj* h = &it.second;
        while(h->p!=0) {
          h = h->p;
          count++;
       }
    }
    return count;
}

int main()
{
  std::string s;
  obj* h = 0;
  while (std::cin >> s) 
  {
    std::string::size_type pos = s.find(')', 0);

    obj* one = get(s.substr(0,pos));
    obj* two = get(s.substr(pos+1));
    one->c.push_back(two);
    two->p=one;
  }

  std::set<std::string> ss;
  std::cout << DD(tot()) << DD((find(get("YOU"),"SAN", ss )-2)) << std::endl;
     
  return 0;
}
