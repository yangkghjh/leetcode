#include <cstdio>
#include <iostream>
using namespace std;

int main() {
  int start;
  int end;

  scanf("%d", &start);
  scanf("%d", &end);

  for (int i = start; i <= end; i++) {
    if (i < 1) {
      continue;
    }

    if (i > 9) {
      if (i % 2 == 0) {
        printf("even\n");
      } else {
        printf("odd\n");
      }
      continue;
    }

    if (i == 1) {
      printf("one\n");
    } else if (i == 2) {
      printf("two\n");
    } else if (i == 3) {
      printf("three\n");
    } else if (i == 4) {
      printf("four\n");
    } else if (i == 5) {
      printf("five\n");
    } else if (i == 6) {
      printf("six\n");
    } else if (i == 7) {
      printf("seven\n");
    } else if (i == 8) {
      printf("eight\n");
    } else if (i == 9) {
      printf("nine\n");
    }
  }

  return 0;
}