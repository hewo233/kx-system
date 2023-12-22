#include "state.h"

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "snake_utils.h"

/* Helper function definitions */
static void set_board_at(game_state_t* state, unsigned int row, unsigned int col, char ch);
static bool is_tail(char c);
static bool is_head(char c);
static bool is_snake(char c);
static char body_to_tail(char c);
static char head_to_body(char c);
static unsigned int get_next_row(unsigned int cur_row, char c);
static unsigned int get_next_col(unsigned int cur_col, char c);
static void find_head(game_state_t* state, unsigned int snum);
static char next_square(game_state_t* state, unsigned int snum);
static void update_tail(game_state_t* state, unsigned int snum);
static void update_head(game_state_t* state, unsigned int snum);

/* Task 1 */
game_state_t* create_default_state() {
  // TODO: Implement this function.
  game_state_t* default_state = malloc(sizeof(game_state_t));
  if(default_state == NULL)
  {
    printf("malloc failed\n");
    return NULL;
  }
  default_state->num_rows = 18;
  default_state->num_snakes = 1;
  default_state->board = malloc(sizeof(char*) * default_state->num_rows);
  if(default_state->board == NULL)
  {
    printf("malloc failed\n");
    return NULL;
  }
  for(int i = 0; i < default_state->num_rows; i++)
  {
    default_state->board[i] = malloc(sizeof(char) * 21);
    if(default_state->board[i] == NULL)
    {
      printf("malloc failed\n");
      return NULL;
    }
  }

  default_state->snakes = malloc(sizeof(snake_t) * default_state->num_snakes);
  if(default_state->snakes == NULL)
  {
    printf("malloc failed\n");
    return NULL;
  }
  default_state->snakes[0].tail_row = 2;
  default_state->snakes[0].tail_col = 2;
  default_state->snakes[0].head_row = 2;
  default_state->snakes[0].head_col = 4;
  default_state->snakes[0].live = true;

  for(int i = 0; i < default_state->num_rows; i++)
  {
    for(int j = 0; j < 20; j++)
    {
      default_state->board[i][j] = ' ';
    }
  }
  for(int i=0;i<20;i++)
  {
    default_state->board[0][i] = '#';
    default_state->board[17][i] = '#';
  }
  for(int i=1;i<17;i++)
  {
    default_state->board[i][0] = '#';
    default_state->board[i][19] = '#';
  }
  default_state->board[2][2] = 'd',default_state->board[2][3]='>',default_state->board[2][4]='D',default_state->board[2][9]='*';
  for (int i=0; i< default_state->num_rows; i++){
    default_state->board[i][20] = '\0';
  }
  return default_state;
}

/* Task 2 */
void free_state(game_state_t* state) {
  // TODO: Implement this function.
    for(int i = 0; i < state->num_rows; i++)
    {
        free(state->board[i]);
    }
    free(state->board);
    free(state->snakes);
    free(state);
  return;
}

/* Task 3 */
void print_board(game_state_t* state, FILE* fp) {
  // TODO: Implement this function.
  for (size_t i = 0; i < state->num_rows; i++) {
    // fputs(state->board[i], fp);
    fprintf(fp, "%s\n", state->board[i]);
  }
  return;
}

/*
  Saves the current state into filename. Does not modify the state object.
  (already implemented for you).
*/
void save_board(game_state_t* state, char* filename) {
  FILE* f = fopen(filename, "w");
  print_board(state, f);
  fclose(f);
}

/* Task 4.1 */

/*
  Helper function to get a character from the board
  (already implemented for you).
*/
char get_board_at(game_state_t* state, unsigned int row, unsigned int col) {
  return state->board[row][col];
}

/*
  Helper function to set a character on the board
  (already implemented for you).
*/
static void set_board_at(game_state_t* state, unsigned int row, unsigned int col, char ch) {
  state->board[row][col] = ch;
}

/*
  Returns true if c is part of the snake's tail.
  The snake consists of these characters: "wasd"
  Returns false otherwise.
*/
static bool is_tail(char c) {
  // TODO: Implement this function.
  if(c == 'w' || c == 'a' || c == 's' || c == 'd')
  {
    return true;
  }
  return false;
}

/*
  Returns true if c is part of the snake's head.
  The snake consists of these characters: "WASDx"
  Returns false otherwise.
*/
static bool is_head(char c) {
  // TODO: Implement this function.
  if(c == 'W' || c == 'A' || c == 'S' || c == 'D' || c == 'x')
  {
    return true;
  }
  return false;
}

/*
  Returns true if c is part of the snake.
  The snake consists of these characters: "wasd^<v>WASDx"
*/
static bool is_snake(char c) {
  // TODO: Implement this function.
  if(is_head(c) || is_tail(c) ) return true;
  if(c == '^' || c == '<' || c == 'v' || c == '>') return true;
  return false;
}

/*
  Converts a character in the snake's body ("^<v>")
  to the matching character representing the snake's
  tail ("wasd").
*/
static char body_to_tail(char c) {
  // TODO: Implement this function.
   if(c == '^') return 'w';
   if(c == '<') return 'a';
   if(c == 'v') return 's';
   if(c == '>') return 'd';
   return '?';
}

/*
  Converts a character in the snake's head ("WASD")
  to the matching character representing the snake's
  body ("^<v>").
*/
static char head_to_body(char c) {
  // TODO: Implement this function.
    if(c == 'W') return '^';
    if(c == 'A') return '<';
    if(c == 'S') return 'v';
    if(c == 'D') return '>';
    return '?';
}

/*
  Returns cur_row + 1 if c is 'v' or 's' or 'S'.
  Returns cur_row - 1 if c is '^' or 'w' or 'W'.
  Returns cur_row otherwise.
*/
static unsigned int get_next_row(unsigned int cur_row, char c) {
  // TODO: Implement this function.
  if(c == 'v' || c == 's' || c == 'S') return cur_row + 1;
  if(c == '^' || c == 'w' || c == 'W') return cur_row - 1;
  return cur_row;
}

/*
  Returns cur_col + 1 if c is '>' or 'd' or 'D'.
  Returns cur_col - 1 if c is '<' or 'a' or 'A'.
  Returns cur_col otherwise.
*/
static unsigned int get_next_col(unsigned int cur_col, char c) {
  // TODO: Implement this function.
  if(c == '>' || c == 'd' || c == 'D') return cur_col + 1;
  if(c == '<' || c == 'a' || c == 'A') return cur_col - 1;
  return cur_col;
}

/*
  Task 4.2

  Helper function for update_state. Return the character in the cell the snake is moving into.

  This function should not modify anything.
*/
static char next_square(game_state_t* state, unsigned int snum) {
  // TODO: Implement this function.
  // save_board(state, "test.txt");
  // free_state(state);
  // state = create_default_state();
  // save_board(state, "test2.txt");
  unsigned int cur_row = state->snakes[snum].head_row;
  unsigned int cur_col = state->snakes[snum].head_col;
  unsigned int next_row = get_next_row(cur_row, state->board[cur_row][cur_col]);
  unsigned int next_col = get_next_col(cur_col, state->board[cur_row][cur_col]);
  return state->board[next_row][next_col];
}

/*
  Task 4.3

  Helper function for update_state. Update the head...

  ...on the board: add a character where the snake is moving

  ...in the snake struct: update the row and col of the head

  Note that this function ignores food, walls, and snake bodies when moving the head.
*/
static void update_head(game_state_t* state, unsigned int snum) {
  // TODO: Implement this function.
  unsigned int cur_row = state->snakes[snum].head_row;
  unsigned int cur_col = state->snakes[snum].head_col;
  unsigned int next_row = get_next_row(cur_row, state->board[cur_row][cur_col]);
  unsigned int next_col = get_next_col(cur_col, state->board[cur_row][cur_col]);
  state->board[next_row][next_col] = state->board[cur_row][cur_col];
  state->board[cur_row][cur_col] = head_to_body(state->board[cur_row][cur_col]);
  state->snakes[snum].head_row = next_row;
  state->snakes[snum].head_col = next_col;
  return;
}

/*
  Task 4.4

  Helper function for update_state. Update the tail...

  ...on the board: blank out the current tail, and change the new
  tail from a body character (^<v>) into a tail character (wasd)

  ...in the snake struct: update the row and col of the tail
*/
static void update_tail(game_state_t* state, unsigned int snum) {
  // TODO: Implement this function.
  unsigned int cur_row = state->snakes[snum].tail_row;
  unsigned int cur_col = state->snakes[snum].tail_col;
  unsigned int next_row = get_next_row(cur_row, state->board[cur_row][cur_col]);
  unsigned int next_col = get_next_col(cur_col, state->board[cur_row][cur_col]);
  state->board[next_row][next_col] = body_to_tail(state->board[next_row][next_col]);
  state->board[cur_row][cur_col] = ' ';
  state->snakes[snum].tail_row = next_row;
  state->snakes[snum].tail_col = next_col;
  return;
}

/* Task 4.5 */
void update_state(game_state_t* state, int (*add_food)(game_state_t* state)) {
  // TODO: Implement this function.
  for(unsigned int i = 0; i < state->num_snakes; i++)
  {
    char next_head_square = next_square(state, i);
    if(next_head_square == ' ')
    {
      update_head(state, i);
      update_tail(state, i);   
    }
    else if(next_head_square == '*')
    {
      update_head(state, i);
      add_food(state);
    }
    else
    {
      state->snakes[i].live = false;
      state->board[state->snakes[i].head_row][state->snakes[i].head_col] = 'x';
    }
  }
  return;
}

/* Task 5 */
game_state_t* load_board(FILE* fp) {
  // TODO: Implement this function.
  game_state_t* load_state = malloc(sizeof(game_state_t));
  load_state->num_rows = 0;
  load_state->num_snakes = 0;
  load_state->snakes = NULL;
  load_state->board = NULL;
  
  char *line = malloc(sizeof(char) * 1000000);
  while(fgets(line, 1000000, fp) != NULL)
  {
    line[strlen(line) - 1] = '\0';
    load_state->num_rows++;
    load_state->board= realloc(load_state->board, (sizeof *load_state->board) * load_state->num_rows);
    load_state->board[load_state->num_rows - 1] = malloc(sizeof(char) * ( strlen(line) + 1));
    strcpy(load_state->board[load_state->num_rows - 1], line);
  }
  free(line);
  return load_state;
}

/*
  Task 6.1

  Helper function for initialize_snakes.
  Given a snake struct with the tail row and col filled in,
  trace through the board to find the head row and col, and
  fill in the head row and col in the struct.
*/
static void find_head(game_state_t* state, unsigned int snum) {
  // TODO: Implement this function.
  unsigned int cur_row = state->snakes[snum].tail_row;
  unsigned int cur_col = state->snakes[snum].tail_col;
  while(!is_head(state->board[cur_row][cur_col]))
  {
    if(state->board[cur_row][cur_col] == ' ')
    {
      printf("Snake %d is missing a head!\n", snum);
      exit(1);
    }
    unsigned int next_row = get_next_row(cur_row, state->board[cur_row][cur_col]);
    unsigned int next_col = get_next_col(cur_col, state->board[cur_row][cur_col]);
    cur_row = next_row;
    cur_col = next_col;
  }
  state->snakes[snum].head_row = cur_row;
  state->snakes[snum].head_col = cur_col;
  return;
}

/* Task 6.2 */
game_state_t* initialize_snakes(game_state_t* state) {
  // TODO: Implement this function.
  for(unsigned int i=0;i<state->num_rows;i++)
  {
    for(unsigned int j=0; j < strlen(state->board[i]); j++)
    {
      if(is_tail(state->board[i][j]))
      {
        state->num_snakes++;
        state->snakes = realloc(state->snakes, sizeof(snake_t) * state->num_snakes);
        state->snakes[state->num_snakes - 1].tail_row = i;
        state->snakes[state->num_snakes - 1].tail_col = j;
        find_head(state, state->num_snakes - 1);
        state->snakes[state->num_snakes - 1].live = true;
      }
    }
  }
  return state;
}
