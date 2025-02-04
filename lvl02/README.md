# 2nd Edition Anxiel Games

## Overview

- Welcome to the 2nd level of Anxiel's Games. In this one, a logic-based game implemented in Golang, developed by [**Raymond Martin**](https://martinraymondogwel.netlify.app). The game revolves around a sequence of four values that can be manipulated through a cycling mechanism. The objective is to reach a target state of `0000`, starting from `0001` by strategically retaining values across moves while others follow a predefined cycle.

## Rules and Mechanics

* The game consists of a **4-slot sequence**, where each slot can hold one of the values `[0, 1, 2]`.
* In each move, the player **must retain one value of choice** from the four slots.
* The retained value remains unchanged in the next move.
* 
* The remaining three values **cycle** according to the pattern:
* sh
* `0 → 1 → 2 → 0 → ...`
* ```

  ```
* The game starts with the initial state `<strong>`0001`</strong>`.
* The goal is to manipulate the sequence until it reaches the target state `<strong>0000</strong>`.

## Example Gameplay

| Move | State    | Retained Value |
| ---- | -------- | -------------- |
| 1    | `0001` | Retain index 3 |
| 2    | `1111` | Retain index 3 |
| 3    | `2221` | Retain index 3 |
| 4    | `0001` | Retain index 3 |
| ...  | ...      | ...            |
| N    | `0000` | -              |

## Implementation

The game is implemented in **Golang**, utilizing arrays to represent the sequence. It provides functionalities to:

* Manage the sequence transformation per move.
* Allow the player to choose which index to retain.
* Check if the target state is achieved.

### Core Algorithm

1. Initialize the sequence as `0001`.
2. Loop until `0000` is reached:
   * Accept user input for the index to retain.
   * Update the sequence by cycling non-retained values.
   * Print the new sequence state.
3. End the game once `0000` is reached.
4. However, a player can fail on looping back to a value that they had previously run on through in the loop.

## How to Run

### Prerequisites

* Golang installed (Go 1.x+)

### Steps

1. Clone the repository:

   ```
   git clone https://github.com/anxielray/Impossible-game-series.git Game
   cd Game
   ```
2. Run the game:

   ```
   ./run.sh
   ```

## Future Improvements

* Add a graphical interface.
* Ideation for 3rd Edition.

## License

This project is open-source and available under the [**MIT License**](LICENSE).
