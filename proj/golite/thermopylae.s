		.arch armv8-a
		.text

		.type fact,%function
		.global fact
		.p2align 2
fact:
		sub sp,sp,#16
		stp x29,x30, [sp]
		mov x29,sp
		sub sp,sp, #80
		mov x1,#1
		str x1,[x29,#-8]
		mov x1,#-1
		str x1,[x29,#-16]
		ldr x2,[x29,#-8]
		cmp x0,x2
		b.le skipMov_L7
		mov x1,#1
		str x1, [x29,#-16]
skipMov_L7:
		mov x1,#1
		str x1,[x29,#-24]
		ldr x1,[x29,#-16]
		ldr x2,[x29,#-24]
		cmp x1,x2
		b.eq elseLabel_L1
ifLabel_L0:
		mov x1,#1
		str x1,[x29,#-32]
		ldr x1,[x29,#-32]
		mov x0,x1
		add sp,sp,#80
		ldp x29,x30,[sp]
		add sp,sp, 16
		ret
		b done_L2
elseLabel_L1:
		mov x1,#1
		str x1,[x29,#-40]
		ldr x2,[x29,#-40]
		sub x3, x0, x2
		str x3,[x29,#-48]
		str x0, [x29, #24]
		sub sp, sp, #32
		ldr x1,[x29,#-48]
		mov x0, x1
		bl fact
		mov x1,x0
		str x1,[x29,#-56]
		ldr x0, [x29, #24]
		add sp, sp, #32
		ldr x2,[x29,#-56]
		mov x1,x2
		str x1,[x29,#-64]
		ldr x2,[x29,#-64]
		mul x3, x0, x2
		str x3,[x29,#-72]
		ldr x1,[x29,#-72]
		mov x0,x1
		add sp,sp,#80
		ldp x29,x30,[sp]
		add sp,sp, 16
		ret
done_L2:
		.size fact, (. - fact)

		.type main,%function
		.global main
		.p2align 2
main:
		sub sp,sp,#16
		stp x29,x30, [sp]
		mov x29,sp
		sub sp,sp, #112
		mov x0,#-1
		str x0,[x29,#-40]
		ldr x1,[x29,#-40]
		mov x0,x1
		str x0,[x29,#-8]
		mov x0,#0
		str x0,[x29,#-48]
		ldr x1,[x29,#-48]
		mov x0,x1
		str x0,[x29,#-16]
		b testCond_L4
loopBody_L3:
		ldr x1,[x29,#-16]
		adrp x2, .READ
		add x2, x2, :lo12:.READ
		add x1, x29,#-16
		mov x0, x2
		bl scanf
		sub sp, sp, #32
		ldr x1,[x29,#-16]
		mov x0, x1
		bl fact
		mov x1,x0
		str x1,[x29,#-56]
		add sp, sp, #32
		ldr x1,[x29,#-56]
		mov x0,x1
		str x0,[x29,#-64]
		ldr x1,[x29,#-64]
		mov x0,x1
		str x0,[x29,#-32]
		ldr x1,[x29,#-32]
		adrp x2, .PRINT_LN
		add x2, x2, :lo12:.PRINT_LN
		mov x0, x2
		mov x1, x1
		bl printf
		ldr x1,[x29,#-24]
		adrp x2, .READ
		add x2, x2, :lo12:.READ
		add x1, x29,#-24
		mov x0, x2
		bl scanf
		mov x0,#0
		str x0,[x29,#-72]
		mov x0,#-1
		str x0,[x29,#-80]
		ldr x0,[x29,#-24]
		ldr x1,[x29,#-72]
		cmp x0,x1
		b.eq skipMov_L8
		mov x0,#1
		str x0,[x29,#-80]
skipMov_L8:
		mov x0,#1
		str x0,[x29,#-88]
		ldr x0,[x29,#-80]
		ldr x1,[x29,#-88]
		cmp x0,x1
		b.eq done_L6
ifLabel_L5:
		mov x0,#1
		str x0,[x29,#-96]
		ldr x1,[x29,#-96]
		mov x0,x1
		str x0,[x29,#-8]
done_L6:
testCond_L4:
		ldr x1,[x29,#-8]
		mov x0,x1
		str x0,[x29,#-104]
		ldr x0,[x29,#-104]
		neg x0,x0
		str x0,[x29,#-104]
		mov x0,#1
		str x0,[x29,#-112]
		ldr x0,[x29,#-104]
		ldr x1,[x29,#-112]
		cmp x0,x1
		b.eq loopBody_L3
		add sp,sp,#112
		ldp x29,x30, [sp]
		add sp,sp, 16
		ret
		.size main, (. - main)

.PRINT_LN:
		.asciz "%ld\n"
		.size .PRINT_LN,5

.PRINT:
		.asciz "%ld"
		.size .PRINT,4

.READ:
		.asciz "%ld"
		.size .READ,4
