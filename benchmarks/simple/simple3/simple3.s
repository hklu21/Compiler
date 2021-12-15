               .arch armv8-a
               .text

               .type MakePoint,%function
               .global MakePoint
               .p2align 2
MakePoint:
               sub sp,sp, #16
               stp x29,x30, [sp]
               mov x29,sp
               sub sp,sp, #16
               str x0,[x29,#16]
               str x1,[x29,#24]
               mov x0, #16
               bl malloc
               str x0,[x29,#-16]
               ldr x0,[x29,#16]
               ldr x1,[x29,#24]
               ldr x3,[x29,#-16]
               mov x2,x3
               str x2,[x29,#-8]
               ldr x2,[x29,#-8]
               str x0, [x2,#0]
               ldr x2,[x29,#-8]
               str x1, [x2,#8]
               ldr x2,[x29,#-8]
               mov x0,x2
               add sp,sp,#16
               ldp x29,x30,[sp]
               add sp,sp, 16
               ret
               .size MakePoint, (. - MakePoint)

               .type main,%function
               .global main
               .p2align 2
main:
               sub sp,sp, #16
               stp x29,x30, [sp]
               mov x29,sp
               sub sp,sp, #80
               mov x1,#128
               str x1,[x29,#-32]
               mov x1,#-1
               str x1,[x29,#-40]
               ldr x1,[x29,#-32]
               ldr x2,[x29,#-40]
               mul x3,x1,x2
               str x3,[x29,#-48]
               mov x1,#64
               str x1,[x29,#-56]
               sub  sp, sp, #16
               ldr x1,[x29,#-48]
               mov x0, x1
               ldr x2,[x29,#-56]
               mov x1, x2
               bl MakePoint
               mov x1,x0
               str x1,[x29,#-64]
               add  sp, sp, #16
               ldr x2,[x29,#-64]
               mov x1,x2
               str x1,[x29,#-16]
               ldr x2,[x29,#-16]
               ldr x1, [x2,#0]
               str x1,[x29,#-72]
               ldr x2,[x29,#-72]
               mov x1,x2
               str x1,[x29,#-8]
               ldr x1,[x29,#-8]
               adrp x2, .PRINT_LN
               add x2, x2, :lo12:.PRINT_LN
               mov x1, x1
               mov x0, x2
               bl printf
               ldr x2,[x29,#-16]
               ldr x1, [x2,#8]
               str x1,[x29,#-80]
               ldr x2,[x29,#-80]
               mov x1,x2
               str x1,[x29,#-8]
               ldr x1,[x29,#-8]
               adrp x2, .PRINT_LN
               add x2, x2, :lo12:.PRINT_LN
               mov x1, x1
               mov x0, x2
               bl printf
               ldr x1,[x29,#-16]
               mov x0, x1
               bl free
               add sp,sp,#80
               ldp x29,x30,[sp]
               add sp,sp, 16
               ret
               .size main, (. - main)

.PRINT_LN:
               .asciz "%ld\n"
               .size .PRINT_LN, 5
