let fname = "2.txt";;

let readfile (filename: string)() = 
  let fcontent = open_in filename in
  
  let try_read () = 
    try Some (input_line fcontent) 
    with End_of_file -> None 
  in
  
  let rec reading_loop(acc: string list) = 
    match try_read () with
    | Some line -> reading_loop (line :: acc)
    | None -> acc 
  in

  reading_loop []
;;

let slist = List.rev (readfile fname ());;

let clean_slist (slist: string list): (string * string list) list =
  List.map (
    fun s ->
      let split = String.split_on_char ':' s in 
      let head = List.hd split in
      let rest = List.nth split 1 in

      let id = List.nth (String.split_on_char ' ' head) 1 in
      let leftover = String.split_on_char ';' rest in

      let cleaned_leftovers = List.map (fun s -> String.trim s) leftover in

      (id, cleaned_leftovers)
  ) slist
;;

let parse(l: (string * string list) list) =
  List.map (
    fun (id, slist) -> (
      id, 
      List.map (
        fun s -> 
          let split = String.split_on_char ',' s in
          let cleaned = List.map (fun s -> String.trim s) split in
          let parsed = List.map (
            fun s -> 
              let split_cube = String.split_on_char ' ' s in
              let number = List.nth split_cube 0 in
              let color = List.nth split_cube 1 in
              (number, color)
          ) cleaned in
          parsed
      ) slist
    ) 
  ) l
;;

let cube_games = parse (clean_slist slist);;

let process_game (cubes: ((string * string) list) list): bool =
  List.for_all (
    fun cube_game ->
      let rec satisfy (cube_game: (string * string) list) (reds: int) (blues: int) (greens: int) =
        match cube_game with
        | [] -> reds <= 12 && blues <= 14 && greens <= 13
        | (number, color) :: rest -> 
          match color with
          | "red" -> satisfy rest (reds + int_of_string number) blues greens
          | "blue" -> satisfy rest reds (blues + int_of_string number) greens
          | "green" -> satisfy rest reds blues (greens + int_of_string number)
          | _ -> false
      in satisfy cube_game 0 0 0
  ) cubes
;;

let play_game (cubes: (string * (string * string) list list) list) =
  let rec loop (cubes: (string * (string * string) list list) list) (ans: (string * bool) list) = 
    match cubes with
    | [] -> ans
    | (id, cube_list_of_games) :: rest -> 
      let process = process_game cube_list_of_games in
      if process then
        loop rest ((id, true) :: ans)
      else
        loop rest ((id, false) :: ans)
  in 
  loop cubes [] 
;;

let played = play_game cube_games;;

let get_true_total (l: (string * bool) list) =
  List.fold_left (
    fun acc (id, bool) -> 
      if bool then acc + (int_of_string id) else acc
  ) 0 l
;;

let ans = get_true_total played;;