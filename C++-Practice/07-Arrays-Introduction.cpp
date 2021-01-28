#include <iostream>

using namespace std;

int main() {
  int len = 0;
  cin >> len;

  int *A = new int[len];

  for (int i = 0; i < len; i++) {
    cin >> A[i];
  }
  for (int i = len - 1; i >= 0; i--) {
    cout << A[i] << ' ';
  }

  delete[] A;
  return 0;
}