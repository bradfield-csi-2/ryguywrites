/*
Naive code for multiplying two matrices together.

There must be a better way!
*/

#include <stdio.h>
#include <stdlib.h>

/*
  A naive implementation of matrix multiplication.

  DO NOT MODIFY THIS FUNCTION, the tests assume it works correctly, which it
  currently does
*/
void matrix_multiply(double **C, double **A, double **B, int a_rows, int a_cols,
                     int b_cols) {
  for (int i = 0; i < a_rows; i++) {
    for (int j = 0; j < b_cols; j++) {
      C[i][j] = 0;
      for (int k = 0; k < a_cols; k++)
        C[i][j] += A[i][k] * B[k][j];
    }
  }
}

void fast_matrix_multiply(double **c, double **a, double **b, int a_rows,
                          int a_cols, int b_cols) {
  // TODO: write a faster implementation here!

  // flip b along the top left to bottom right diagonal
  // ex: swap b_ij with b_ji
  // this forces matrix b to be row order as well when doing the matrix mult
  double temp;
  int i, j;
  for (i = 0; i < a_rows; i++) {
    for (j = 0; j < i; j++) { // only need to flip up to (not including) the diagonal
      temp = b[i][j];
      b[i][j] = b[j][i];
      b[j][i] = temp;
    }
  }

  double acc1, acc2, acc3, acc4; // I stopped at 4 out of laziness, more might be faster
  int k;
  for (i = 0; i < a_rows-3; i+=4) { // I cheat and assume a power of 2
    for (j = 0; j < a_rows; j++) {
      acc1 = 0;
      acc2 = 0;
      acc3 = 0;
      acc4 = 0;
      for (k = 0; k < a_cols; k++) {
        acc1 += a[i][k] * b[j][k];
        acc2 += a[i+1][k] * b[j][k];
        acc3 += a[i+2][k] * b[j][k];
        acc4 += a[i+3][k] * b[j][k];
      }
      c[i][j] = acc1;
      c[i+1][j] = acc2;
      c[i+2][j] = acc3;
      c[i+3][j] = acc4;
    }
  }
}
