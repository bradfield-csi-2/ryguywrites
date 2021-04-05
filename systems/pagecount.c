#include <stdint.h>
#include <stdio.h>
#include <time.h>

#define TEST_LOOPS 3333333 // divided by 3 to get rid of mod operation

uint64_t pagecount(uint64_t memory_size, uint64_t page_shift) {
  memory_size = memory_size >> page_shift;
  return memory_size;
}

int main (int argc, char** argv) {
  clock_t baseline_start, baseline_end, test_start, test_end;
  uint64_t memory_size, page_size;
  double clocks_elapsed, time_elapsed;
  int i, ignore = 0;

  uint64_t msizes[] = {1L << 32, 1L << 40, 1L << 52};
  uint64_t pshifts[] = {12, 16, 32}; // edited to allow bit shifting instead of division

  baseline_start = clock();
  for (i = 0; i < TEST_LOOPS; i++) {
    memory_size = msizes[i % 3];
    page_size = pshifts[i % 3];
    ignore += 1 + memory_size +
              page_size; // so that this loop isn't just optimized away
  }
  baseline_end = clock();

  test_start = clock();
  for (i = 0; i < TEST_LOOPS; i++) {
    // give up modularity to remove mod operations
    memory_size = msizes[0];
    page_size = pshifts[0];
    ignore += pagecount(memory_size, page_size) + memory_size + page_size;

    memory_size = msizes[1];
    page_size = pshifts[1];
    ignore += pagecount(memory_size, page_size) + memory_size + page_size;

    memory_size = msizes[2];
    page_size = pshifts[2];
    ignore += pagecount(memory_size, page_size) + memory_size + page_size;
  }
  // deal with the remainder (10000000 % 3 == 1)
  memory_size = msizes[0];
  page_size = pshifts[0];
  ignore += pagecount(memory_size, page_size) + memory_size + page_size;
  test_end = clock();

  clocks_elapsed = test_end - test_start - (baseline_end - baseline_start);
  time_elapsed = clocks_elapsed / CLOCKS_PER_SEC;

  printf("%.2fs to run %d tests (%.2fns per test)\n", time_elapsed, TEST_LOOPS,
         time_elapsed * 1e9 / TEST_LOOPS);
  return ignore;
}
