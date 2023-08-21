# Compiling Go - A Brief Explanation

## What is a compiler, what is compiling?

- Compiling is transforming source code into a binary executable program;
- The Go compiler compiles .go source files into machine code;
- This machine code can then be executed on any platform;

## Go Compilation Key Points

- Go uses a custom compiler called `gc`;
- It compiled extremely fast compared to other languages;
- Go build times are measured in milliseconds compared to seconds in other compiled languages;
- The compiler does optimizations but focuses on fast compilation;
- Go object files use a custom `.o` format;
- Go compiles to a single static binary executable;
- Dynamic linking is optional and can be disabled;
- Go binaries are statically linked by default;
- Cross-compilation is build-in and easy - compile on one platform, run on another;

## Advantages of Compiled Languages

- **Performance** : runs closer to native machine code speed;
  - Strong typing and safety checking at compile time.
- **Dependency Management** : executable contains all dependencies;
- **Distribution** : can distribute executables without language runtime;
- **Security** : harder to inject malicious code at runtime;