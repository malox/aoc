// Example program
#include <iostream>
#include <string>
#include <vector>
#include <cstdlib>

#define NL std::cout << "\n---------------------------------------------------\n"
#define DD(e) " " << #e << "[" << e << "]"
#define DS(e) " " << #e << "[" << static_cast<int>(e) << "]"
#define DB(e) " " << #e << "[" << std::boolalpha << e << "]"
#define PP(e) std::cout << e << std::endl

const int msize=9;
int matrix[msize][msize];
const int moffset = msize/2;


typedef enum {
kRight = 1,
kUp,
kLeft,
kDown
} dir_t;


class Pos
{
public:
   Pos() : _row(moffset), _col(moffset), _last(kDown) {}

bool write(int iVal, int& iRow, int& iCol, dir_t& iDir);
bool set(int val);

dir_t getNextMove(dir_t iMove);

char* toString(dir_t iMove);
bool updatePos(int& iRow, int& iCol, dir_t iDir);


static void reset();
static void fill();
static void print();
static int square(int, int);

int _row;
int _col;
dir_t _last;

};

bool rangeCheck(int val)
{
    return (val >= 0) and (val < msize);
}

bool rangeCheck(Pos p)
{
    return rangeCheck(p._row) and rangeCheck(p._col);
}

bool Pos::updatePos(int& iRow, int& iCol, dir_t iDir)
{
    //PP("updatePos  IN -" << DD(iRow) << DD(iCol) << DS(iDir) << " Dir[" << toString(iDir) << "]");
    switch (iDir)
    {
        case kRight:
          ++iCol;
          break;
        case kUp:
          --iRow;
          break;
        case kLeft:
          --iCol;
          break;
        case kDown:
          ++iRow;
          break;
    }
    //PP("updatePos OUT -" << DD(iRow) << DD(iCol) << DS(iDir) << " Dir[" << toString(iDir) << "]");
}

int Pos::square(int iRow, int iCol)
{
    if(moffset==iRow and moffset==iCol)
        return 1;

    int aSquare = 0;

    if(rangeCheck(iRow-1) and rangeCheck(iCol-1))
        aSquare += matrix[iRow-1][iCol-1];

    if(rangeCheck(iRow-1) and rangeCheck(iCol))
        aSquare += matrix[iRow-1][iCol];

    if(rangeCheck(iRow-1) and rangeCheck(iCol+1))
        aSquare += matrix[iRow-1][iCol+1];
    
    if(rangeCheck(iRow) and rangeCheck(iCol-1))
        aSquare += matrix[iRow][iCol-1];
    
//    if(rangeCheck(iRow) and rangeCheck(iCol))
//        aSquare += matrix[iRow][iCol];
    
    if(rangeCheck(iRow) and rangeCheck(iCol+1))
        aSquare += matrix[iRow][iCol+1];
 
    if(rangeCheck(iRow+1) and rangeCheck(iCol-1))
        aSquare += matrix[iRow+1][iCol-1];
    
    if(rangeCheck(iRow+1) and rangeCheck(iCol))
        aSquare += matrix[iRow+1][iCol];
    
    if(rangeCheck(iRow+1) and rangeCheck(iCol+1))
        aSquare += matrix[iRow+1][iCol+1];
    
    return aSquare;
}

bool Pos::write(int iVal, int& iRow, int& iCol, dir_t& iDir)
{
    //print();
    //PP("\nWrite -" << DD(iVal) << DD(iRow) << DD(iCol) << " Dir[" << toString(iDir) << "]");
    if(rangeCheck(iRow) and rangeCheck(iCol) and matrix[iRow][iCol]==0)
    {
        //PP("Match -" << DD(iVal) << DD(iRow) << DD(iCol) << " Dir[" << toString(iDir) << "]");
        matrix[iRow][iCol] = square(iRow, iCol);
        return true;
    }
    if(rangeCheck(iRow) and rangeCheck(iCol) and matrix[iRow][iCol]!=0)
    {
        int aRow=iRow, aCol=iCol;
        dir_t aNext = getNextMove(iDir);
        updatePos(aRow,aCol,aNext);
        //PP("Trying Next Dir-" << DD(iVal) << DD(aRow) << DD(aCol) << " Dir[" << toString(aNext) << "]");
        if(rangeCheck(aRow) and rangeCheck(aCol) and matrix[aRow][aCol]==0)
        {
            //PP("Match Next Dir-" << DD(iVal) << DD(aRow) << DD(aCol) << " Dir[" << toString(aNext) << "]");
            matrix[aRow][aCol] = square(aRow, aCol);
            iRow=aRow ; iCol=aCol; iDir=aNext;
            return true;
        }
    }
    if(rangeCheck(iRow) and rangeCheck(iCol) and matrix[iRow][iCol]!=0)
    {
        int aRow=iRow, aCol=iCol;
        updatePos(aRow,aCol,iDir);
        //PP("Trying Same Dir -" << DD(iVal) << DD(aRow) << DD(aCol) << " Dir[" << toString(iDir) << "]");
        if(rangeCheck(aRow) and rangeCheck(aCol) and matrix[aRow][aCol]==0)
        {
            //PP("Match Same Dir-" << DD(iVal) << DD(aRow) << DD(aCol) << " Dir[" << toString(iDir) << "]");
            matrix[aRow][aCol] = square(aRow, aCol);
            iRow=aRow ; iCol=aCol;
            return true;
        }
    }
    return false;
}

bool Pos::set(int val)
{
    return write(val, _row, _col, _last);     
}


dir_t Pos::getNextMove(dir_t iMove)
{
    switch (iMove)
    {
        case kRight:
          return kUp;
        case kUp:
          return kLeft;
        case kLeft:
          return kDown;
        case kDown:
          return kRight;
    }
}

char* Pos::toString(dir_t iMove)
{
    switch (iMove)
    {
        case kLeft:
          return "Left";
        case kUp:
          return "Up";
        case kRight:
          return "Right";
        case kDown:
          return "Down";
    }
}


void Pos::reset()
{
    for(int i = 0; i < msize; ++i)
        for(int j = 0; j < msize; ++j)
            matrix[i][j] = 0;
}

void Pos::fill()
{
    Pos p;
    
    for(int i = 0; i < (msize)*(msize); ++i) {
        //PP( DD(i) << DD(p._row) << DD(p._col) );
        if (not p.set( i+1 ))
            break;    
    }
}

void Pos::print()
{
    for(int i = 0; i < msize; ++i)
        for(int j = 0; j < msize; ++j)
            std::cout << matrix[i][j] << ((j+1 == msize) ? "\n" : "\t");
        
}



void t_main()
{
  NL;
  Pos::reset();
  Pos::fill();
  NL;
  Pos::print();
  NL;
}

int main()
{
  t_main();
  return 0;
}


