# Brainfuck

[Brainfuck][BF] is an [esoteric programming language][EPL] created in [1993 by Urban Müller][UM].

Notable for its extreme minimalism, the language consists of only eight simple commands and an
[instruction pointer][PC]. While it is fully [Turing complete][TC], it is not intended for practical
use, but to challenge and amuse [programmers][PR]. Brainfuck simply requires one to break commands
into microscopic steps.

The language's name is a reference to the slang term [brainfuck][bf], which refers to things so
complicated or unusual that they exceed the limits of one's understanding.

## The Instructions

Brainfuck is tiny. It consists of eight different instructions. These instructions can be used to
manipulate the state of the Brainfuck machine:

 * `>` - Increment the data pointer by 1.
 * `<` - Decrement the data pointer by 1.
 * `+` - Increment the value in the current cell (the cell the data pointer is pointing to).
 * `-` - Decrement the value in the current cell.
 * `.` - Take the integer in the current cell, treat it as an ASCII char and print it on the output
stream.
 * `,` - Read a character from the input stream, convert it to an integer and save it to the current
cell.
 * `[` - This always needs to come with a matching `]`. If the current cell contains a zero, set
the instruction pointer to the index of the instruction after the matching `]`.
 * `]` - If the current cell does not contain a zero, set the instruction pointer to the index of
the instruction after the matching `[`.

That’s all of it, the complete Brainfuck language.

[BF]: https://en.wikipedia.org/wiki/Brainfuck
[bf]: https://en.wiktionary.org/wiki/brainfuck
[UM]: https://dx.doi.org/10.1080/07350198.2020.1727096
[EPL]: https://en.wikipedia.org/wiki/Esoteric_programming_language
[PC]: https://en.wikipedia.org/wiki/Program_counter
[TC]: https://en.wikipedia.org/wiki/Turing_completeness
[PR]: https://en.wikipedia.org/wiki/Programmer
