                .arch armv8-a
                .text

                .type fib1,%function
                .global fib1
                .p2align 2
fib1:
                sub sp,sp, #16
                stp x29,x30, [sp]
                mov x29,sp
                sub sp,sp, #80
                mov x1,#2
                str x1,[x29,#-8]
                mov x1,#0
                str x1,[x29,#-16]
                ldr x1,[x29,#-8]
                cmp x0,x1
                b.ge skipMov_L2
                mov x1,#1
                str x1,[x29,#-16]
skipMov_L2:
                ldr x1,[x29,#-16]
                cmp x1,#0
                b.eq L3
                mov x0,x0
                add sp,sp,#80
                ldp x29,x30,[sp]
                add sp,sp, 16
                ret
                b L4
L3:
                mov x1,#1
                str x1,[x29,#-24]
                ldr x1,[x29,#-24]
                subs x2,x0,x1
                str x2,[x29,#-32]
                str x0,[x29,#24]
                sub  sp, sp, #16
                ldr x1,[x29,#-32]
                mov x0, x1
                bl fib1
                mov x1,x0
                str x1,[x29,#-40]
                ldr x0,[x29,#24]
                add  sp, sp, #16
                mov x1,#2
                str x1,[x29,#-48]
                ldr x1,[x29,#-48]
                subs x2,x0,x1
                str x2,[x29,#-56]
                str x0,[x29,#24]
                sub  sp, sp, #16
                ldr x1,[x29,#-56]
                mov x0, x1
                bl fib1
                mov x1,x0
                str x1,[x29,#-64]
                ldr x0,[x29,#24]
                add  sp, sp, #16
                ldr x1,[x29,#-40]
                ldr x2,[x29,#-64]
                add x3,x1,x2
                str x3,[x29,#-72]
                ldr x1,[x29,#-72]
                mov x0,x1
                add sp,sp,#80
                ldp x29,x30,[sp]
                add sp,sp, 16
                ret
L4:
                add sp,sp,#80
                ldp x29,x30,[sp]
                add sp,sp, 16
                ret
                .size fib1, (. - fib1)

                .type fib2,%function
                .global fib2
                .p2align 2
fib2:
                sub sp,sp, #16
                stp x29,x30, [sp]
                mov x29,sp
                sub sp,sp, #96
                mov x1,#0
                str x1,[x29,#-40]
                ldr x2,[x29,#-40]
                mov x1,x2
                str x1,[x29,#-8]
                mov x1,#1
                str x1,[x29,#-48]
                ldr x2,[x29,#-48]
                mov x1,x2
                str x1,[x29,#-16]
                b L14
L13:
                mov x1,#1
                str x1,[x29,#-56]
                ldr x1,[x29,#-56]
                subs x2,x0,x1
                str x2,[x29,#-64]
                ldr x2,[x29,#-64]
                mov x1,x2
                mov x0,x1
                ldr x1,[x29,#-8]
                ldr x2,[x29,#-16]
                add x3,x1,x2
                str x3,[x29,#-72]
                ldr x2,[x29,#-72]
                mov x1,x2
                str x1,[x29,#-24]
                ldr x2,[x29,#-16]
                mov x1,x2
                str x1,[x29,#-8]
                ldr x2,[x29,#-24]
                mov x1,x2
                str x1,[x29,#-16]
L14:
                mov x1,#0
                str x1,[x29,#-80]
                mov x1,#0
                str x1,[x29,#-88]
                ldr x1,[x29,#-80]
                cmp x0,x1
                b.eq skipMov_L22
                mov x1,#1
                str x1,[x29,#-88]
skipMov_L22:
                ldr x1,[x29,#-88]
                cmp x1,#1
                b.eq L13
                ldr x1,[x29,#-8]
                mov x0,x1
                add sp,sp,#96
                ldp x29,x30,[sp]
                add sp,sp, 16
                ret
                .size fib2, (. - fib2)

                .type main,%function
                .global main
                .p2align 2
main:
                sub sp,sp, #16
                stp x29,x30, [sp]
                mov x29,sp
                sub sp,sp, #80
                mov x0, #16
                bl malloc
                str x0,[x29,#-48]
                ldr x2,[x29,#-48]
                mov x1,x2
                str x1,[x29,#-8]
                ldr x1,[x29,#-32]
                adrp x2, .READ
                add x2, x2, :lo12:.READ
                add x1, x29,#-32
                mov x0, x2
                bl scanf
                ldr x1,[x29,#-8]
                ldr x2,[x29,#-32]
                str x2, [x1,#0]
                ldr x1,[x29,#-32]
                adrp x2, .READ
                add x2, x2, :lo12:.READ
                add x1, x29,#-32
                mov x0, x2
                bl scanf
                ldr x1,[x29,#-8]
                ldr x2,[x29,#-32]
                str x2, [x1,#8]
                ldr x2,[x29,#-8]
                ldr x1, [x2,#0]
                str x1,[x29,#-56]
                sub  sp, sp, #16
                ldr x1,[x29,#-56]
                mov x0, x1
                bl fib1
                mov x1,x0
                str x1,[x29,#-64]
                add  sp, sp, #16
                ldr x2,[x29,#-64]
                mov x1,x2
                str x1,[x29,#-16]
                ldr x2,[x29,#-8]
                ldr x1, [x2,#8]
                str x1,[x29,#-72]
                sub  sp, sp, #16
                ldr x1,[x29,#-72]
                mov x0, x1
                bl fib2
                mov x1,x0
                str x1,[x29,#-80]
                add  sp, sp, #16
                ldr x2,[x29,#-80]
                mov x1,x2
                str x1,[x29,#-24]
                ldr x1,[x29,#-8]
                mov x0, x1
                bl free
                ldr x1,[x29,#-16]
                adrp x2, .PRINT_LN
                add x2, x2, :lo12:.PRINT_LN
                mov x1, x1
                mov x0, x2
                bl printf
                ldr x1,[x29,#-24]
                adrp x2, .PRINT_LN
                add x2, x2, :lo12:.PRINT_LN
                mov x1, x1
                mov x0, x2
                bl printf
                add sp,sp,#80
                ldp x29,x30,[sp]
                add sp,sp, 16
                ret
                .size main, (. - main)

.PRINT_LN:
                .asciz "%ld\n"
                .size .PRINT_LN, 5

.READ:
                .asciz "%ld"
                .size .READ,4
