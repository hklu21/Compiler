                .arch armv8-a
                .comm p2,8,8
                .comm p1,8,8
                .text

                .type AddPoint,%function
                .global AddPoint
                .p2align 2
AddPoint:
                sub sp,sp, #16
                stp x29,x30, [sp]
                mov x29,sp
                sub sp,sp, #96
                str x0,[x29,#24]
                str x1,[x29,#24]
                mov x0, #16
                bl malloc
                str x0,[x29,#-16]
                ldr x0,[x29,#24]
                ldr x1,[x29,#24]
                ldr x3,[x29,#-16]
                mov x2,x3
                str x2,[x29,#-8]
                adrp x2, p1
                add x2, x2, :lo12:p1
                ldr x2, [x2]
                str x2,[x29,#-24]
                ldr x3,[x29,#-24]
                ldr x2, [x3,#0]
                str x2,[x29,#-32]
                adrp x2, p2
                add x2, x2, :lo12:p2
                ldr x2, [x2]
                str x2,[x29,#-40]
                ldr x3,[x29,#-40]
                ldr x2, [x3,#0]
                str x2,[x29,#-48]
                ldr x2,[x29,#-32]
                ldr x3,[x29,#-48]
                add x4,x2,x3
                str x4,[x29,#-56]
                ldr x2,[x29,#-8]
                ldr x3,[x29,#-56]
                str x3, [x2,#0]
                adrp x2, p1
                add x2, x2, :lo12:p1
                ldr x2, [x2]
                str x2,[x29,#-64]
                ldr x3,[x29,#-64]
                ldr x2, [x3,#8]
                str x2,[x29,#-72]
                adrp x2, p2
                add x2, x2, :lo12:p2
                ldr x2, [x2]
                str x2,[x29,#-80]
                ldr x3,[x29,#-80]
                ldr x2, [x3,#8]
                str x2,[x29,#-88]
                ldr x2,[x29,#-72]
                ldr x3,[x29,#-88]
                add x4,x2,x3
                str x4,[x29,#-96]
                ldr x2,[x29,#-8]
                ldr x3,[x29,#-96]
                str x3, [x2,#8]
                ldr x2,[x29,#-8]
                mov x0,x2
                add sp,sp,#96
                ldp x29,x30,[sp]
                add sp,sp, 16
                ret
                .size AddPoint, (. - AddPoint)

                .type MakePoint,%function
                .global MakePoint
                .p2align 2
MakePoint:
                sub sp,sp, #16
                stp x29,x30, [sp]
                mov x29,sp
                sub sp,sp, #32
                str x0,[x29,#16]
                str x1,[x29,#24]
                mov x0, #16
                bl malloc
                str x0,[x29,#-24]
                ldr x0,[x29,#16]
                ldr x1,[x29,#24]
                ldr x3,[x29,#-24]
                mov x2,x3
                str x2,[x29,#-8]
                ldr x2,[x29,#-8]
                str x0, [x2,#0]
                ldr x2,[x29,#-8]
                str x1, [x2,#8]
                ldr x2,[x29,#-8]
                mov x0,x2
                add sp,sp,#32
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
                sub sp,sp, #128
                mov x1,#3
                str x1,[x29,#-32]
                mov x1,#4
                str x1,[x29,#-40]
                sub  sp, sp, #16
                ldr x1,[x29,#-32]
                mov x0, x1
                ldr x2,[x29,#-40]
                mov x1, x2
                bl MakePoint
                mov x1,x0
                str x1,[x29,#-48]
                add  sp, sp, #16
                ldr x2,[x29,#-48]
                adrp x1, p1
                add x1, x1, :lo12:p1
                str x2, [x1]
                mov x1,#5
                str x1,[x29,#-56]
                mov x1,#6
                str x1,[x29,#-64]
                sub  sp, sp, #16
                ldr x1,[x29,#-56]
                mov x0, x1
                ldr x2,[x29,#-64]
                mov x1, x2
                bl MakePoint
                mov x1,x0
                str x1,[x29,#-72]
                add  sp, sp, #16
                ldr x2,[x29,#-72]
                adrp x1, p2
                add x1, x1, :lo12:p2
                str x2, [x1]
                adrp x1, p1
                add x1, x1, :lo12:p1
                ldr x1, [x1]
                str x1,[x29,#-80]
                adrp x1, p2
                add x1, x1, :lo12:p2
                ldr x1, [x1]
                str x1,[x29,#-88]
                sub  sp, sp, #16
                ldr x1,[x29,#-80]
                mov x0, x1
                ldr x2,[x29,#-88]
                mov x1, x2
                bl AddPoint
                mov x1,x0
                str x1,[x29,#-96]
                add  sp, sp, #16
                ldr x2,[x29,#-96]
                mov x1,x2
                str x1,[x29,#-16]
                ldr x2,[x29,#-16]
                ldr x1, [x2,#0]
                str x1,[x29,#-104]
                ldr x2,[x29,#-104]
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
                str x1,[x29,#-112]
                ldr x2,[x29,#-112]
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
                str x1,[x29,#-120]
                ldr x1,[x29,#-120]
                mov x0, x1
                bl free
                adrp x1, p2
                add x1, x1, :lo12:p2
                ldr x1, [x1]
                str x1,[x29,#-128]
                ldr x1,[x29,#-128]
                mov x0, x1
                bl free
                ldr x1,[x29,#-16]
                mov x0, x1
                bl free
                add sp,sp,#128
                ldp x29,x30,[sp]
                add sp,sp, 16
                ret
                .size main, (. - main)

.PRINT_LN:
                .asciz "%ld\n"
                .size .PRINT_LN, 5
