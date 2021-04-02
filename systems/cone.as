default rel

section .text
global volume
volume:
	; v = 1/3*pi*r^2*h
	vmulss xmm0, xmm0, xmm0 ; get r^2 in return register
	vmulss xmm0, xmm0, xmm1 ; multiply h to r^2
	vmulss xmm0, xmm0, [pi_over_3] ; multiply pi/3 to r^2*h
 	ret

section .rodata
pi_over_3: dd 1.04719666667	
