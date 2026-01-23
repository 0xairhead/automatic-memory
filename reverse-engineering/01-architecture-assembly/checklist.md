# Checklist: Architecture & Assembly

## Basics
- [ ] **Registers:** Understand EAX, EBX, ECX, EDX, ESI, EDI, ESP, EBP vs RAX, RBX...
- [ ] **EIP/RIP:** Understand how the instruction pointer works.
- [ ] **Flags:** Understand Zero Flag (ZF), Carry Flag (CF), Sign Flag (SF).

## Memory
- [ ] **Stack:** operations (push/pop), stack frames (prologue/epilogue).
- [ ] **Heap:** dynamic allocation concepts.
- [ ] **Memory Segments:** .text, .data, .bss.

## Function Calls
- [ ] **cdecl:** arguments on stack, caller cleanup.
- [ ] **stdcall:** arguments on stack, callee cleanup.
- [ ] **System V AMD64 (Linux):** arguments in RDI, RSI, RDX, RCX, R8, R9.
