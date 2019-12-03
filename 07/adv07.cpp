// Example program

#include <iostream>
#include <sstream>
#include <fstream>
#include <string>

#include <map>
#include <vector>
#include <cstdlib>


#define NL std::cout << "\n---------------------------------------------------\n"
#define DD(e) " " << #e << "[" << e << "]"
#define DS(e) " " << #e << "[" << static_cast<int>(e) << "]"
#define DB(e) " " << #e << "[" << std::boolalpha << e << "]"
#define PP(e) std::cout << e << std::endl

class Node
{
public:
    std::string _name;
    int _weight;
    Node* _parent;
    std::vector<Node*> _childs;
    Node() : _weight(0), _parent(0) {}
    void dump();
    void calc();
    int sum();
};

typedef std::map<std::string, Node> mnodes_t;

void Node::dump()
{
    std::string childs;
    for(int i =0; i < _childs.size() ; ++i)
        childs+= _childs[i]->_name + " ";
    std::string parent = _parent ? _parent->_name : std::string("");
    PP(" Node -" << DD(_name) << DD(_weight) << DD(parent) << DD(_childs.size()) << DD(childs));
}

void Node::calc()
{
    for(int i =0; i < _childs.size() ; ++i)
        PP(" Child -" << DD(_childs[i]->_name) << DD(_childs[i]->_weight) << DD(_childs[i]->sum()));
}

int Node::sum()
{
    int aSum = _weight;
    for(int i =0; i < _childs.size() ; ++i)
        aSum += _childs[i]->sum();
    return aSum;
}


void print(mnodes_t& mnodes, Node*& head)
{
    mnodes_t::iterator it = mnodes.begin();
    mnodes_t::iterator end = mnodes.end();
    for( ; it != end ; ++it)
    {
        if(it->second._parent==0)
            head = &it->second;
        it->second.dump();
    }
}

Node& getNode(mnodes_t& inodes, const std::string& iname)
{
    mnodes_t::iterator it = inodes.find(iname);
    if(it != inodes.end())
    {
        return inodes[iname];
    }
    else
    {
        Node& anode = inodes[iname];
        anode._name = iname;
        return anode;
    }
}

void t_main(std::string filename)
{
  NL;

  Node* head = 0;
  mnodes_t mnodes;

  std::ifstream infile(filename.c_str());

  std::string line;
  while (std::getline(infile, line))
  {
//      NL; PP(DD(line));
      std::istringstream ssline(line);

      std::string node;
      int weight=0;
      std::string arrow;
      std::string childs;
  
      if (ssline >> node >> weight >> arrow >> childs) 
      {
//           PP(DD(node) << DD(weight) << DD(arrow) << DD(childs));
           Node& n = getNode(mnodes, node);
           n._weight = weight;

           std::istringstream sschilds(childs);
           std::string schild;
           while (std::getline(sschilds, schild, ','))
           {
//               PP(DD(schild));
               Node& cn = getNode(mnodes, schild);
               cn._parent = &n;
               n._childs.push_back(&cn);
           }
      }
      else 
      {
          std::istringstream ssline(line);
          if(ssline >> node >> weight)
          {
//              PP(DD(node) << DD(weight) );
              Node& n = getNode(mnodes, node);
              n._weight = weight;
              if(weight<=0)
                  PP("GREP " <<DD(node));
          }
      }
//      PP(" Final -" << DD(node) << DD(weight) << DD(arrow) << DD(childs));
  }

  NL;
  print(mnodes, head);
  
  NL;
  head->dump();
  head->calc();
  NL;
  PP("it should be hmvwl");
}

int main(int argc, char *argv[])
{
  if( argc > 1)
      t_main(argv[1]); // hmvwl 
  return 0;
}


