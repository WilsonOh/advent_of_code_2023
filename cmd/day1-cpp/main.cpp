#include <fmt/core.h>
#include <fstream>
#include <string>
#include <string_view>
#include <vector>

int main() {
  auto input_file = std::ifstream("./input.txt");
  std::vector<std::string> letterDigits = { "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" };
  int ans = 0;
  for (std::string line; std::getline(input_file, line);) {
    std::vector<std::string> digits;
    for (std::size_t i = 0; i < line.length(); ++i) {
      if (std::isdigit(line[i])) { digits.push_back(std::string(1, line[i])); }
      // part 2
      auto ss = std::string_view{ line }.substr(i);
      for (std::size_t j = 0; j < letterDigits.size(); ++j) {
        if (ss.starts_with(letterDigits[j])) { digits.push_back(std::to_string(j + 1)); }
      }
    }
    ans += std::stoi(digits.front() + digits.back());
    digits.clear();
  }
  fmt::print("Part 1: {}\n", ans);
}
