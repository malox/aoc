MY_OUTPUT="$1"

shift

echo -e "INPUT [ $@ ] \n"
echo -e "EXECUTING [ g++ -std=c++17 -o ${MY_OUTPUT} -g $@ ] \n"

g++ -std=c++17 -o ${MY_OUTPUT} -g $@

echo -e "BUILD HOPEFULLY SUCCESFUL -- check this file [ ./${MY_OUTPUT} ] \n"

