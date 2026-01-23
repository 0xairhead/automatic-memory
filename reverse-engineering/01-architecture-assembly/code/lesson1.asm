
lesson1:     file format elf64-x86-64


Disassembly of section .init:

000000000040033c <_init>:
  40033c:	f3 0f 1e fa          	endbr64
  400340:	48 83 ec 08          	sub    rsp,0x8
  400344:	48 8b 05 95 2c 00 00 	mov    rax,QWORD PTR [rip+0x2c95]        # 402fe0 <__gmon_start__@Base>
  40034b:	48 85 c0             	test   rax,rax
  40034e:	74 02                	je     400352 <_init+0x16>
  400350:	ff d0                	call   rax
  400352:	48 83 c4 08          	add    rsp,0x8
  400356:	c3                   	ret

Disassembly of section .plt:

0000000000400360 <puts@plt-0x10>:
  400360:	ff 35 8a 2c 00 00    	push   QWORD PTR [rip+0x2c8a]        # 402ff0 <_GLOBAL_OFFSET_TABLE_+0x8>
  400366:	ff 25 8c 2c 00 00    	jmp    QWORD PTR [rip+0x2c8c]        # 402ff8 <_GLOBAL_OFFSET_TABLE_+0x10>
  40036c:	0f 1f 40 00          	nop    DWORD PTR [rax+0x0]

0000000000400370 <puts@plt>:
  400370:	ff 25 8a 2c 00 00    	jmp    QWORD PTR [rip+0x2c8a]        # 403000 <puts@GLIBC_2.2.5>
  400376:	68 00 00 00 00       	push   0x0
  40037b:	e9 e0 ff ff ff       	jmp    400360 <_init+0x24>

0000000000400380 <printf@plt>:
  400380:	ff 25 82 2c 00 00    	jmp    QWORD PTR [rip+0x2c82]        # 403008 <printf@GLIBC_2.2.5>
  400386:	68 01 00 00 00       	push   0x1
  40038b:	e9 d0 ff ff ff       	jmp    400360 <_init+0x24>

Disassembly of section .text:

0000000000400390 <_start>:
  400390:	f3 0f 1e fa          	endbr64
  400394:	31 ed                	xor    ebp,ebp
  400396:	49 89 d1             	mov    r9,rdx
  400399:	5e                   	pop    rsi
  40039a:	48 89 e2             	mov    rdx,rsp
  40039d:	48 83 e4 f0          	and    rsp,0xfffffffffffffff0
  4003a1:	50                   	push   rax
  4003a2:	54                   	push   rsp
  4003a3:	45 31 c0             	xor    r8d,r8d
  4003a6:	31 c9                	xor    ecx,ecx
  4003a8:	48 c7 c7 90 04 40 00 	mov    rdi,0x400490
  4003af:	ff 15 23 2c 00 00    	call   QWORD PTR [rip+0x2c23]        # 402fd8 <__libc_start_main@GLIBC_2.34>
  4003b5:	f4                   	hlt
  4003b6:	66 2e 0f 1f 84 00 00 	cs nop WORD PTR [rax+rax*1+0x0]
  4003bd:	00 00 00 

00000000004003c0 <_dl_relocate_static_pie>:
  4003c0:	f3 0f 1e fa          	endbr64
  4003c4:	c3                   	ret
  4003c5:	66 2e 0f 1f 84 00 00 	cs nop WORD PTR [rax+rax*1+0x0]
  4003cc:	00 00 00 
  4003cf:	90                   	nop

00000000004003d0 <deregister_tm_clones>:
  4003d0:	b8 18 30 40 00       	mov    eax,0x403018
  4003d5:	48 3d 18 30 40 00    	cmp    rax,0x403018
  4003db:	74 13                	je     4003f0 <deregister_tm_clones+0x20>
  4003dd:	b8 00 00 00 00       	mov    eax,0x0
  4003e2:	48 85 c0             	test   rax,rax
  4003e5:	74 09                	je     4003f0 <deregister_tm_clones+0x20>
  4003e7:	bf 18 30 40 00       	mov    edi,0x403018
  4003ec:	ff e0                	jmp    rax
  4003ee:	66 90                	xchg   ax,ax
  4003f0:	c3                   	ret
  4003f1:	0f 1f 40 00          	nop    DWORD PTR [rax+0x0]
  4003f5:	66 66 2e 0f 1f 84 00 	data16 cs nop WORD PTR [rax+rax*1+0x0]
  4003fc:	00 00 00 00 

0000000000400400 <register_tm_clones>:
  400400:	be 18 30 40 00       	mov    esi,0x403018
  400405:	48 81 ee 18 30 40 00 	sub    rsi,0x403018
  40040c:	48 89 f0             	mov    rax,rsi
  40040f:	48 c1 ee 3f          	shr    rsi,0x3f
  400413:	48 c1 f8 03          	sar    rax,0x3
  400417:	48 01 c6             	add    rsi,rax
  40041a:	48 d1 fe             	sar    rsi,1
  40041d:	74 11                	je     400430 <register_tm_clones+0x30>
  40041f:	b8 00 00 00 00       	mov    eax,0x0
  400424:	48 85 c0             	test   rax,rax
  400427:	74 07                	je     400430 <register_tm_clones+0x30>
  400429:	bf 18 30 40 00       	mov    edi,0x403018
  40042e:	ff e0                	jmp    rax
  400430:	c3                   	ret
  400431:	0f 1f 40 00          	nop    DWORD PTR [rax+0x0]
  400435:	66 66 2e 0f 1f 84 00 	data16 cs nop WORD PTR [rax+rax*1+0x0]
  40043c:	00 00 00 00 

0000000000400440 <__do_global_dtors_aux>:
  400440:	f3 0f 1e fa          	endbr64
  400444:	80 3d c9 2b 00 00 00 	cmp    BYTE PTR [rip+0x2bc9],0x0        # 403014 <completed.0>
  40044b:	75 13                	jne    400460 <__do_global_dtors_aux+0x20>
  40044d:	55                   	push   rbp
  40044e:	48 89 e5             	mov    rbp,rsp
  400451:	e8 7a ff ff ff       	call   4003d0 <deregister_tm_clones>
  400456:	c6 05 b7 2b 00 00 01 	mov    BYTE PTR [rip+0x2bb7],0x1        # 403014 <completed.0>
  40045d:	5d                   	pop    rbp
  40045e:	c3                   	ret
  40045f:	90                   	nop
  400460:	c3                   	ret
  400461:	0f 1f 40 00          	nop    DWORD PTR [rax+0x0]
  400465:	66 66 2e 0f 1f 84 00 	data16 cs nop WORD PTR [rax+rax*1+0x0]
  40046c:	00 00 00 00 

0000000000400470 <frame_dummy>:
  400470:	f3 0f 1e fa          	endbr64
  400474:	eb 8a                	jmp    400400 <register_tm_clones>

0000000000400476 <add_numbers>:
  400476:	55                   	push   rbp
  400477:	48 89 e5             	mov    rbp,rsp
  40047a:	89 7d ec             	mov    DWORD PTR [rbp-0x14],edi
  40047d:	89 75 e8             	mov    DWORD PTR [rbp-0x18],esi
  400480:	8b 55 ec             	mov    edx,DWORD PTR [rbp-0x14]
  400483:	8b 45 e8             	mov    eax,DWORD PTR [rbp-0x18]
  400486:	01 d0                	add    eax,edx
  400488:	89 45 fc             	mov    DWORD PTR [rbp-0x4],eax
  40048b:	8b 45 fc             	mov    eax,DWORD PTR [rbp-0x4]
  40048e:	5d                   	pop    rbp
  40048f:	c3                   	ret

0000000000400490 <main>:
  400490:	55                   	push   rbp
  400491:	48 89 e5             	mov    rbp,rsp
  400494:	48 83 ec 10          	sub    rsp,0x10
  400498:	c7 45 fc 0a 00 00 00 	mov    DWORD PTR [rbp-0x4],0xa
  40049f:	c7 45 f8 14 00 00 00 	mov    DWORD PTR [rbp-0x8],0x14
  4004a6:	c7 45 f4 00 00 00 00 	mov    DWORD PTR [rbp-0xc],0x0
  4004ad:	bf b8 11 40 00       	mov    edi,0x4011b8
  4004b2:	e8 b9 fe ff ff       	call   400370 <puts@plt>
  4004b7:	8b 55 f8             	mov    edx,DWORD PTR [rbp-0x8]
  4004ba:	8b 45 fc             	mov    eax,DWORD PTR [rbp-0x4]
  4004bd:	89 d6                	mov    esi,edx
  4004bf:	89 c7                	mov    edi,eax
  4004c1:	e8 b0 ff ff ff       	call   400476 <add_numbers>
  4004c6:	89 45 f4             	mov    DWORD PTR [rbp-0xc],eax
  4004c9:	8b 45 f4             	mov    eax,DWORD PTR [rbp-0xc]
  4004cc:	89 c6                	mov    esi,eax
  4004ce:	bf d1 11 40 00       	mov    edi,0x4011d1
  4004d3:	b8 00 00 00 00       	mov    eax,0x0
  4004d8:	e8 a3 fe ff ff       	call   400380 <printf@plt>
  4004dd:	b8 00 00 00 00       	mov    eax,0x0
  4004e2:	c9                   	leave
  4004e3:	c3                   	ret

Disassembly of section .fini:

00000000004004e4 <_fini>:
  4004e4:	f3 0f 1e fa          	endbr64
  4004e8:	48 83 ec 08          	sub    rsp,0x8
  4004ec:	48 83 c4 08          	add    rsp,0x8
  4004f0:	c3                   	ret
