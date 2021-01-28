#include <cstdio>
#include <iostream>

using namespace std;

int main() {
  int i;     // %d
  long l;    // %ld
  char c;    // %c
  float f;   // %f
  double d;  // %lf

  scanf("%d %ld %c %f %lf", &i, &l, &c, &f, &d);

  printf("%d\n%ld\n%c\n%f\n%lf\n", i, l, c, f, d);

  return 0;
}
