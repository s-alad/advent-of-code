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
  let rec convert_line (s: string) = 
    match s with
    | "" -> []
    | _ -> (String.get s 0) :: (convert_line (String.sub s 1 ((String.length s) - 1)))
  in
  List.map convert_line slist
;;

let compute (clist: char list list) = 
  List.fold_left (fun (acc: int) (current_char_list: char list) ->
    let found = List.filter (
      fun c -> 
        match c with
        | '0' .. '9' -> true
        | _ -> false
    ) current_char_list in 

    let first = List.hd found in
    let last = if List.length found > 1 then List.hd (List.rev found) else first in
    let total = ((int_of_char first - int_of_char '0') * 10) + (int_of_char last - int_of_char '0') in

    print_string "first: ";
    print_char first;
    print_newline ();

    print_string "first int of char: ";
    print_int (int_of_char first - int_of_char '0');
    print_newline ();

    print_string "last: ";
    print_char last;
    print_newline ();

    print_string "last int of char: ";
    print_int (int_of_char last - int_of_char '0');
    print_newline ();

    print_string "total: ";
    print_int total;
    print_newline ();
    total + acc
  ) 0 clist
;;
    
let ans = compute (convert_to_char_list slist);;