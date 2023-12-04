let file = "input.txt";;

let first = ref false;;
let first_val =  ref (-1);;
let second_val = ref (-1);;

let firstWordVal = ref "";;
let firstWordIndex = ref (-1);;
let line = ref "";;

let secondWordVal = ref "";;
let secondWordIndex = ref (-1);;

let allDigitsList = ["one"; "two"; "three"; "four"; "five"; "six"; "eight"; "nine"];;
let allDigits = "one two three four five six eight nine";;

let get_actual_num some =
  match some with 
  |"one" -> 1
  |"two" -> 2
  |"three" -> 3
  |"four" -> 4
  |"five" -> 5
  |"six" -> 6
  |"seven" -> 7
  |"eight" -> 8
  |"nine" -> 9
  | _ -> 0
;;

let substring_exists haystack needle =
  let len_h = String.length haystack in
  let len_n = String.length needle in
  let rec check_at_index i =
    if i + len_n > len_h then false
    else if String.sub haystack i len_n = needle then true
    else check_at_index (i + 1)
  in
  check_at_index 0
;;



let read_file filename = 
let lines = ref [] in
let chan = open_in filename in
try
  while true; do
    lines := input_line chan :: !lines
  done; !lines
with End_of_file ->
  close_in chan;
  List.rev !lines ;;



let read_file2 file_name =
  let lines = ref [] in
  let ic = open_in file_name in
  try
    while true; do
      lines := input_line ic :: !lines
    done; !lines
  with End_of_file ->
    close_in ic;
    List.rev !lines
;;

let lines = read_file2 file ;;

let print_all line = 
  print_endline (string_of_int line)
;;


let char_is_int c =
  if c >= '0' && c <= '9' then
    int_of_string (String.make 1 c)
  else
    -1
;;






let do_something line =
  let word = ref "" in
  for i = 0 to (String.length line) - 1 do
    word := !word ^ (String.make 1 (String.get line i));

    print_endline !word;
    if char_is_int(String.get line i) <> (-1) then(
      if !first = false then(
        first := true;
        first_val := char_is_int (String.get line i);
      );
      if !first = true then(
        second_val := char_is_int (String.get line i);
      );
    if substring_exists line !word then(
      if List.mem !word allDigitsList then
        if !first = false then(
          first := true;
          first_val := get_actual_num !word;
          word := "";
        );
        if !first = true then(
          second_val := get_actual_num !word;
          word := "";
        );
    );
    if (substring_exists line !word) = false then(
      word := ""
    );
  )
  done;

  word := "";
  first := false;
  (!first_val * 10 ) + !second_val

;;
   

List.map print_all (List.map do_something lines)
