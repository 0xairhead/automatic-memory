# Walkthrough: Analyzing Lesson 1

This walkthrough guides you through the `lesson1` binary we compiled. Open `code/lesson1.asm` (or use `objdump -d -M intel code/lesson1`) and `code/lesson1.c` side-by-side.

## 1. The `main` Function

Find `<main>:` in the disassembly.

```assembly
0000000000400490 <main>:
  ...
  400490:	55                   	push   rbp
  400491:	48 89 e5             	mov    rbp,rsp
  400494:	48 83 ec 10          	sub    rsp,0x10
```
**Analysis:**
*   **Prologue:** `push rbp` and `mov rbp, rsp` set up the stack frame.
*   **Stack Allocation:** `sub rsp, 0x10` reserves 16 bytes on the stack for local variables (`int x`, `int y`, `int result`).

### Variable Initialization
```assembly
  400498:	c7 45 fc 0a 00 00 00 	mov    DWORD PTR [rbp-0x4],0xa   ; x = 10
  40049f:	c7 45 f8 14 00 00 00 	mov    DWORD PTR [rbp-0x8],0x14  ; y = 20
  4004a6:	c7 45 f4 00 00 00 00 	mov    DWORD PTR [rbp-0xc],0x0   ; result = 0
```
*   `[rbp-0x4]` corresponds to `x`. It's assigned `0xa` (decimal 10).
*   `[rbp-0x8]` corresponds to `y`. It's assigned `0x14` (decimal 20).

## 2. Calling `add_numbers`

Look at how the arguments are passed before the call. This follows the **System V AMD64 ABI**.

```assembly
  4004b7:	8b 55 f8             	mov    edx,DWORD PTR [rbp-0x8]   ; Load y (20) into edx
  4004ba:	8b 45 fc             	mov    eax,DWORD PTR [rbp-0x4]   ; Load x (10) into eax
  4004bd:	89 d6                	mov    esi,edx                   ; Move y into esi (2nd arg)
  4004bf:	89 c7                	mov    edi,eax                   ; Move x into edi (1st arg)
  4004c1:	e8 b0 ff ff ff       	call   400476 <add_numbers>
```
*   **1st Argument (`a`)**: Goes into **EDI**.
*   **2nd Argument (`b`)**: Goes into **ESI**.
*   **Call**: The CPU jumps to `add_numbers`.

### The Return Value
```assembly
  4004c6:	89 45 f4             	mov    DWORD PTR [rbp-0xc],eax
```
*   The result of `add_numbers` is stored in **EAX**.
*   The instruction moves that value from `eax` into `[rbp-0xc]` (our `result` variable).

## 3. The `add_numbers` Function

Now look at `<add_numbers>`:

```assembly
0000000000400476 <add_numbers>:
  ...
  40047a:	89 7d ec             	mov    DWORD PTR [rbp-0x14],edi  ; Store 1st arg (a) on stack
  40047d:	89 75 e8             	mov    DWORD PTR [rbp-0x18],esi  ; Store 2nd arg (b) on stack
```
*   Wait, why move registers back to memory? This is because `gcc` without optimization (`-O0`) tends to be very verbose and puts everything on the stack for debugging.

```assembly
  400480:	8b 55 ec             	mov    edx,DWORD PTR [rbp-0x14]
  400483:	8b 45 e8             	mov    eax,DWORD PTR [rbp-0x18]
  400486:	01 d0                	add    eax,edx                   ; eax = eax + edx
```
*   `add eax, edx` performs the actual math. The result stays in **EAX**, which is the return register.

## Conclusion

You have successfully traced:
1.  Stack setup.
2.  Variable assignment.
3.  Function arguments (EDI/ESI).
4.  Function return (EAX).
