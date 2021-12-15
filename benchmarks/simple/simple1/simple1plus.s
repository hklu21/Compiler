               .arch armv8-a
               .text

               .type Add,%function
               .global Add
               .p2align 2
Add:
               sub sp,sp, #16
               stp x29,x30, [sp]
               mov x29,sp
               sub sp,sp, #16
               add x2,x0,x1
               str x2,[x29,#-8]
               ldr x2,[x29,#-8]
               mov x0,x2
               add sp,sp,#16
               ldp x29,x30,[sp]
               add sp,sp, 16
               ret
               .size Add, (. - Add)

               .type main,%function
               .global main
               .p2align 2
main:
               sub sp,sp, #16
               stp x29,x30, [sp]
               mov x29,sp
               sub sp,sp, #48
               mov x1,#129
               str x1,[x29,#-40]
               ldr x2,[x29,#-40]
               mov x1,x2
               str x1,[x29,#-8]
               ldr x1,[x29,#-16]
               adrp x2, .READ
               add x2, x2, :lo12:.READ
               add x1, x29,#-16
               mov x0, x2
               bl scanf
               sub  sp, sp, #16
               ldr x1,[x29,#-8]
               mov x0, x1
               ldr x2,[x29,#-16]
               mov x1, x2
               bl Add
               mov x1,x0
               str x1,[x29,#-48]
               add  sp, sp, #16
               ldr x2,[x29,#-48]
               mov x1,x2
               str x1,[x29,#-24]
               ldr x1,[x29,#-24]
               adrp x2, .PRINT_LN
               add x2, x2, :lo12:.PRINT_LN
               mov x1, x1
               mov x0, x2
               bl printf
               add sp,sp,#48
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
