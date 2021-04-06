#include "vec.h"
#include "stdio.h"
#include <time.h>

data_t dotproduct(vec_ptr u, vec_ptr v) {
    data_t acc1 = 0, acc2 = 0, acc3 = 0, acc4 = 0, acc5 = 0; // initialize accumulators
    long vector_length = vec_length(u); // take the computation out of the array
    long length_limit = vector_length - 4; // deal with array out of bound errors

    // initialize data positions
    data_t *u_pos = get_vec_start(u);
    data_t *v_pos = get_vec_start(v);
    
    long i;
    for (i = 0; i < length_limit; i+=5) { // we can assume both vectors are same length
    	// array indexing is better since it can take advantage of parallel use of FUs
      	acc1 += u_pos[i] * v_pos[i];
        acc2 += u_pos[i+1] * v_pos[i+1];
      	acc3 += u_pos[i+2] * v_pos[i+2];
      	acc4 += u_pos[i+3] * v_pos[i+3];
      	acc5 += u_pos[i+4] * v_pos[i+4];
    }   

    // do final elements in the array
    for (; i < vector_length; i++) {
        acc1 += u_pos[i] * v_pos[i];
    }
    return acc1 + acc2 + acc3 + acc4 + acc5;
}
