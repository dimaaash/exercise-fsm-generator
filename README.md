# exercise-fsm-generator

This project provides a generic Finite State Machine (FSM) implementation in Go, along with a sample client application for computing the remainder of a binary number modulo 3 using an FSM.

## Structure

- `fsm/` — Generic FSM library package
- `cmd/mod3/` — Client application demonstrating a mod-3 FSM
- `Makefile` — Build, run, test, and clean targets


## FSM Library

The FSM library is generic and reusable. It allows you to define:
- States
- Input alphabet
- Initial state
- Final states
- Transition function

It includes robust error handling:
- Panics on invalid input symbols or states
- Panics if the transition function returns an invalid state
- Detects and panics on possible infinite loops (configurable max steps)

You can use this library to model any finite automaton by providing the appropriate configuration.


## Mod-3 Example

The mod-3 client application (`cmd/mod3/main.go`) demonstrates how to use the FSM library to compute the remainder of a binary number (given as a string) modulo 3. It outputs both the final state (S0, S1, S2) and the actual remainder value (0, 1, 2).

Example output:

```
Input: 1011
Final State: S1
Remainder (mod 3): 1
```

## Usage

Build and run the mod-3 application:

```sh
make build
make run
```


Run unit tests (including error handling and edge cases):

```sh
make test
```


## Customization & Testing

You can create new FSMs by implementing your own transition functions and state sets using the `fsm` package.

Unit tests cover:
- Correct transitions and outputs for the mod-3 FSM
- Panics on invalid input symbols
- Panics on transition to invalid states
- Infinite loop detection

