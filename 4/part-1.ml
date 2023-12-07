let fname = "4.txt";;

let readfile(filename: string)() = 
  let fcontent = open_in filename in
  let try_read() = 
    try Some (input_line fcontent)
    with End_of_file -> None 
  in
  
  let rec reading_loop(acc: string list) = 
    match try_read() with
    | Some line -> reading_loop(line :: acc)
    | None -> acc 
  in

  reading_loop []
;;

let slist = List.rev (readfile fname ());;

let parse (slist: string list) = 
  List.map (
    fun s ->
      let split = String.split_on_char ':' s in
      let rest = String.trim(List.nth split 1) in
      let pipe = String.split_on_char '|' rest in
      let winning = List.filter (fun s -> s <> "") (String.split_on_char ' ' (String.trim (List.nth pipe 0))) in
      let numbers =  List.filter (fun s -> s <> "") (String.split_on_char ' ' (String.trim (List.nth pipe 1))) in
      [winning; numbers]
  ) slist
;;

let cleaned = parse slist;;


let win clist =
  List.map (
    fun l -> (
      let winners = List.nth l 0 in
      let numbers = List.nth l 1 in

      let won = List.fold_left ( fun acc s -> if List.mem s winners then acc + 1 else acc) 0 numbers in

      2.**((float_of_int won) -. 1.)
    )
  ) clist
;;

let sum l = List.fold_left (
  fun acc num ->
    if num >= 1. then acc +. num
    else acc
) 0. l;;

let winlist = win cleaned;;

let ans = sum winlist;;