section .text
global index
index:
	; rdi: matrix
	; rsi: rows
	; rdx: cols
	; rcx: rindex
	; r8: cindex

	mov rax, rdx ; move cols to rax (to calculate row offset)
	mul rcx ; get row offset in rax
	add rax, r8 ; add col offset to rax
	imul rax, 4 ; convert offset units from index to bytes
	mov eax, dword [rdi + rax] ; write the value of the int into our return register
	ret
