section .text
global fib
fib:
	; deal with the base case
	mov eax, edi
	cmp edi, 1 
	jle return

	; move callee saved registers to the stack
	push rbx
	push rbp

	; use ebx to store the parameter n
	mov ebx, eax

	; call fib(n-1)
	lea edi, [eax - 1]
	call fib ; now eax contains fib(n-1)
	mov ebp, eax ; store fib(n-1) in ebp

	; call fib(n-2)
	lea edi, [ebx - 2]
	call fib ; now eax contains fib(n-2)

	; store return value in eax
	add eax, ebp ; add fib(n-1) to fib(n-2) 

	; restore callee saved registers from stack to register
	pop rbp
	pop rbx
return:
	ret
