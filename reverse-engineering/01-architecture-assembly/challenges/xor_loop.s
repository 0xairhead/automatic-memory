	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 26, 0	sdk_version 26, 2
	.globl	_main                           ; -- Begin function main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:
	sub	sp, sp, #80
	stp	x29, x30, [sp, #64]             ; 16-byte Folded Spill
	add	x29, sp, #64
	.cfi_def_cfa w29, 16
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	adrp	x8, ___stack_chk_guard@GOTPAGE
	ldr	x8, [x8, ___stack_chk_guard@GOTPAGEOFF]
	ldr	x8, [x8]
	stur	x8, [x29, #-8]
	str	wzr, [sp, #28]
	adrp	x8, l___const.main.data@PAGE
	add	x8, x8, l___const.main.data@PAGEOFF
	ldr	q0, [x8]
	str	q0, [sp, #32]
	ldr	w8, [x8, #16]
	str	w8, [sp, #48]
	mov	w8, #85                         ; =0x55
	str	w8, [sp, #24]
	mov	w8, #5                          ; =0x5
	str	w8, [sp, #20]
	adrp	x0, l_.str@PAGE
	add	x0, x0, l_.str@PAGEOFF
	bl	_printf
	str	wzr, [sp, #16]
	b	LBB0_1
LBB0_1:                                 ; =>This Inner Loop Header: Depth=1
	ldr	w8, [sp, #16]
	ldr	w9, [sp, #20]
	subs	w8, w8, w9
	b.ge	LBB0_4
	b	LBB0_2
LBB0_2:                                 ;   in Loop: Header=BB0_1 Depth=1
	ldrsw	x9, [sp, #16]
	add	x8, sp, #32
	ldr	w8, [x8, x9, lsl #2]
                                        ; kill: def $x8 killed $w8
	mov	x9, sp
	str	x8, [x9]
	adrp	x0, l_.str.1@PAGE
	add	x0, x0, l_.str.1@PAGEOFF
	bl	_printf
	b	LBB0_3
LBB0_3:                                 ;   in Loop: Header=BB0_1 Depth=1
	ldr	w8, [sp, #16]
	add	w8, w8, #1
	str	w8, [sp, #16]
	b	LBB0_1
LBB0_4:
	adrp	x0, l_.str.2@PAGE
	add	x0, x0, l_.str.2@PAGEOFF
	bl	_printf
	str	wzr, [sp, #12]
	b	LBB0_5
LBB0_5:                                 ; =>This Inner Loop Header: Depth=1
	ldr	w8, [sp, #12]
	ldr	w9, [sp, #20]
	subs	w8, w8, w9
	b.ge	LBB0_8
	b	LBB0_6
LBB0_6:                                 ;   in Loop: Header=BB0_5 Depth=1
	ldrsw	x8, [sp, #12]
	add	x9, sp, #32
	ldr	w8, [x9, x8, lsl #2]
	ldr	w10, [sp, #24]
	eor	w8, w8, w10
	ldrsw	x10, [sp, #12]
	str	w8, [x9, x10, lsl #2]
	b	LBB0_7
LBB0_7:                                 ;   in Loop: Header=BB0_5 Depth=1
	ldr	w8, [sp, #12]
	add	w8, w8, #1
	str	w8, [sp, #12]
	b	LBB0_5
LBB0_8:
	adrp	x0, l_.str.3@PAGE
	add	x0, x0, l_.str.3@PAGEOFF
	bl	_printf
	str	wzr, [sp, #8]
	b	LBB0_9
LBB0_9:                                 ; =>This Inner Loop Header: Depth=1
	ldr	w8, [sp, #8]
	ldr	w9, [sp, #20]
	subs	w8, w8, w9
	b.ge	LBB0_12
	b	LBB0_10
LBB0_10:                                ;   in Loop: Header=BB0_9 Depth=1
	ldrsw	x9, [sp, #8]
	add	x8, sp, #32
	ldr	w8, [x8, x9, lsl #2]
                                        ; kill: def $x8 killed $w8
	mov	x9, sp
	str	x8, [x9]
	adrp	x0, l_.str.1@PAGE
	add	x0, x0, l_.str.1@PAGEOFF
	bl	_printf
	b	LBB0_11
LBB0_11:                                ;   in Loop: Header=BB0_9 Depth=1
	ldr	w8, [sp, #8]
	add	w8, w8, #1
	str	w8, [sp, #8]
	b	LBB0_9
LBB0_12:
	adrp	x0, l_.str.2@PAGE
	add	x0, x0, l_.str.2@PAGEOFF
	bl	_printf
	ldur	x9, [x29, #-8]
	adrp	x8, ___stack_chk_guard@GOTPAGE
	ldr	x8, [x8, ___stack_chk_guard@GOTPAGEOFF]
	ldr	x8, [x8]
	subs	x8, x8, x9
	b.eq	LBB0_14
	b	LBB0_13
LBB0_13:
	bl	___stack_chk_fail
LBB0_14:
	mov	w0, #0                          ; =0x0
	ldp	x29, x30, [sp, #64]             ; 16-byte Folded Reload
	add	sp, sp, #80
	ret
	.cfi_endproc
                                        ; -- End function
	.section	__TEXT,__const
	.p2align	2, 0x0                          ; @__const.main.data
l___const.main.data:
	.long	10                              ; 0xa
	.long	20                              ; 0x14
	.long	30                              ; 0x1e
	.long	40                              ; 0x28
	.long	50                              ; 0x32

	.section	__TEXT,__cstring,cstring_literals
l_.str:                                 ; @.str
	.asciz	"Original data: "

l_.str.1:                               ; @.str.1
	.asciz	"%d "

l_.str.2:                               ; @.str.2
	.asciz	"\n"

l_.str.3:                               ; @.str.3
	.asciz	"XORed data: "

.subsections_via_symbols
