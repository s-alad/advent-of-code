let fname = "1.txt";;

let readfile filename () = 
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

let convert_to_char_list (slist: string list) = 
  List.map ( fun s -> List.init (String.length s) (String.get s) ) slist
;;

let rec find_valid_digits (clist: char list) =
  List.filter (
    fun c -> 
      match c with
      | '0' .. '9' -> true
      | _ -> false
  ) clist
;;

let rec find_valid_numbers (clist: char list) (l: char list) =
  match clist with
  | [] -> l
  | 'o' :: 'n' :: 'e' :: _ -> 
    find_valid_numbers (List.tl clist) ('1' :: l)
  | 't' :: 'w' :: 'o' :: _ -> 
    find_valid_numbers (List.tl clist) ('2' :: l)
  | 't' :: 'h' :: 'r' :: 'e' :: 'e' :: _ -> 
    find_valid_numbers (List.tl clist) ('3' :: l)
  | 'f' :: 'o' :: 'u' :: 'r' :: _ -> 
    find_valid_numbers (List.tl clist) ('4' :: l)
  | 'f' :: 'i' :: 'v' :: 'e' :: _ -> 
    find_valid_numbers (List.tl clist) ('5' :: l)
  | 's' :: 'i' :: 'x' :: _ -> 
    find_valid_numbers (List.tl clist) ('6' :: l)
  | 's' :: 'e' :: 'v' :: 'e' :: 'n' :: _ -> 
    find_valid_numbers (List.tl clist) ('7' :: l)
  | 'e' :: 'i' :: 'g' :: 'h' :: 't' :: _ -> 
    find_valid_numbers (List.tl clist) ('8' :: l)
  | 'n' :: 'i' :: 'n' :: 'e' :: _ -> 
    find_valid_numbers (List.tl clist) ('9' :: l)
  | 'z' :: 'e' :: 'r' :: 'o' :: _ -> 
    find_valid_numbers (List.tl clist) ('0' :: l)
  | '0' .. '9' :: rest -> 
    find_valid_numbers rest (List.hd clist :: l)
  | _ :: rest -> 
    find_valid_numbers rest l
;;

let compute (clist: char list list) = 
  List.fold_left (fun (acc: int) (current_char_list: char list) ->
    let found = List.rev (find_valid_numbers current_char_list []) in
    let first = List.hd found in
    let last = if List.length found > 1 then List.hd (List.rev found) else first in
    let total = ((int_of_char first - int_of_char '0') * 10) 
                + (int_of_char last - int_of_char '0') in
    total + acc
  ) 0 clist
;;

let ans = compute (convert_to_char_list slist);;