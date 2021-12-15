               .arch armv8-a
               .comm p1,8,8
               .text

               .type main,%function
               .global main
               .p2align 2
main:
               sub sp,sp, #16
               stp x29,x30, [sp]
               mov x29,sp
               sub sp,sp, #96
               mov x0, #16
               bl malloc
               str x0,[x29,#-16]
               ldr x2,[x29,#-16]
               adrp x1, p1
               add x1, x1, :lo12:p1
               str x2, [x1]
               mov x1,#3
               str x1,[x29,#-24]
               adrp x1, p1
               add x1, x1, :lo12:p1
               ldr x1, [x1]
               str x1,[x29,#-32]
               ldr x1,[x29,#-32]
               ldr x2,[x29,#-24]
               str x2, [x1,#0]
               mov x1,#4
               str x1,[x29,#-40]
               adrp x1, p1
               add x1, x1, :lo12:p1
               ldr x1, [x1]
               str x1,[x29,#-48]
               ldr x1,[x29,#-48]
               ldr x2,[x29,#-40]
               str x2, [x1,#8]
               adrp x1, p1
               add x1, x1, :lo12:p1
               ldr x1, [x1]
               str x1,[x29,#-56]
               ldr x2,[x29,#-56]
               ldr x1, [x2,#0]
               str x1,[x29,#-64]
               ldr x2,[x29,#-64]
               mov x1,x2
               str x1,[x29,#-8]
               ldr x1,[x29,#-8]
               adrp x2, .PRINT_LN
               add x2, x2, :lo12:.PRINT_LN
               mov x1, x1
               mov x0, x2
               bl printf
               adrp x1, p1
               add x1, x1, :lo12:p1
               ldr x1, [x1]
               str x1,[x29,#-72]
               ldr x2,[x29,#-72]
               ldr x1, [x2,#8]
               str x1,[x29,#-80]
               ldr x2,[x29,#-80]
               mov x1,x2
               str x1,[x29,#-8]
               ldr x1,[x29,#-8]
               adrp x2, .PRINT
               add x2, x2, :lo12:.PRINT
               mov x1, x1
               mov x0, x2
               bl printf
               adrp x1, p1
               add x1, x1, :lo12:p1
               ldr x1, [x1]
               str x1,[x29,#-88]
               ldr x1,[x29,#-88]
               mov x0, x1
               bl free
               add sp,sp,#96
               ldp x29,x30,[sp]
               add sp,sp, 16
               ret
               .size main, (. - main)

.PRINT:
               .asciz "%ld"
               .size .PRINT, 4

.PRINT_LN:
               .asciz "%ld\n"
               .size .PRINT_LN, 5
