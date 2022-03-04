#include <iostream>

enum Operator { OP_MIN, OP_PLUS, OP_PUSH, OP_DUMP };

Operator min() { return (OP_MIN); }
Operator plus() { return (OP_PLUS); }
Operator push(int value) { return (OP_PUSH); }
Operator dump() { return (OP_DUMP); }

void compile() {}
void interpret() {
  std::cout << "Test function called!" << std::endl;
  while (true) {
    std::string a;
    std::cout << "$uni-> ";
    std::cin >> a;
    std::cout << a << std::endl;
  }
}

int main(int argc, char const *argv[]) {
  if (argc >= 2) {
    if (argv[1] == "compile") {
      std::cout << "Test function called!" << std::endl;
      compile();
    } else if (argv[1] == "test") {
      interpret();
    }
  }
  return 0;
}
