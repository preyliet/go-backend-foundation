# Go Backend Foundation 🚀
A structured journey through the fundamentals of Go, following "The Go Programming Language" by Donovan & Kernighan.

## 📂 Chapter 1: Tutorial Recap
This chapter covered the essential syntax and standard libraries required to build functional tools.

- **1.1 & 1.2**: Basic syntax and Command-Line Arguments using `os.Args`.
- **1.3 & 1.4**: Data structures and generating animated GIFs.
- **1.5 Fetching URLs**: Created a mini-CURL tool using `net/http`.
- **1.6 Concurrency**: Learned to fetch multiple URLs simultaneously using **Goroutines** and **Channels**.
- **1.7 Web Server**: Built a local echo server to handle HTTP requests.
- **1.8 Loose Ends**: Mastered **Pointers**, **Structs**, and **Switch** statements for cleaner logic.

## Chapter 2: Program Structure
This chapter dives into the mechanics of how Go programs are organized, covering naming conventions, variable declarations, and memory management.

### 2.1 Names
* Learned Go's naming rules (Unicode letters, underscores, and digits).
* **Visibility:** Explored how capitalized names are exported (public) while lowercase names are unexported (private).

### 2.2 Declarations
* Practiced the four main declaration types: `var`, `const`, `type`, and `func`.

### 2.3 Variables & Pointers
* **Zero Values:** Understanding how Go initializes variables to a default state (e.g., `0`, `""`, `nil`).
* **Pointers:** Using `&` to get addresses and `*` to dereference, allowing direct memory access.
* **The `flag` Package:** Built `echo4.go` to handle command-line flags via pointers.
* **Memory Lifecycle:** Learned about the `new` function and how variables "escape" to the heap.
