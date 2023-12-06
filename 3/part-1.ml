let fname = "3.txt";;

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

let is_digit digit = match digit with '0' .. '9' -> true | _ -> false;; 

let pad_schematic_with_dots(slist: string list) =
  let left_right = List.map (fun s -> "." ^ s ^ ".") slist in
  let top_bottom = String.make (String.length (List.hd slist) + 2) '.' in
  [top_bottom] @ left_right @ [top_bottom]
;;

let pslist = pad_schematic_with_dots slist;;

let engine_schematic(pslist: string list) = 
  let rec engine_loop(s: string)(i: int)(j: int)(current_number: string)(toadd: bool)(total: int) =

      (*print i, j*)
      print_int i;
      print_string ",";
      print_int j;
      print_string "\n";


      if j == (String.length s - 1) then 
        if toadd then
          let new_total = total + (int_of_string current_number) in
          new_total
        else
          total
      else 
        let topleft_s = String.get (List.nth pslist (i - 1)) (j - 1) in
        let top_s = String.get (List.nth pslist (i - 1)) j in
        let topright_s = String.get (List.nth pslist (i - 1)) (j + 1) in
        let left_s = String.get (s) (j - 1) in
        let right_s = String.get (s) (j + 1) in
        let bottomleft_s = String.get (List.nth pslist (i + 1)) (j - 1) in
        let bottom_s = String.get (List.nth pslist (i + 1)) j in
        let bottomright_s = String.get (List.nth pslist (i + 1)) (j + 1) in

        let surrounding = [
          topleft_s; top_s; topright_s;
          left_s; right_s;
          bottomleft_s; bottom_s; bottomright_s
        ] in

        let current_s = String.get s j in

        (* print the surrounding list *)
        List.iter (fun s -> print_char s; print_char ' ') [topleft_s; top_s; topright_s;];
        print_string "\n";
        List.iter (fun s -> print_char s; print_char ' ') [left_s; current_s; right_s;];
        print_string "\n";
        List.iter (fun s -> print_char s; print_char ' ') [bottomleft_s; bottom_s; bottomright_s;];
        print_string "\n";
        print_string "-------\n";

        match current_s with
        | '0' .. '9' ->

          let checktoadd = List.exists (fun s -> s != '.' && not (is_digit s)) surrounding in

          print_string "checktoadd: ";
          print_string (string_of_bool checktoadd);
          print_string "\n";
          print_string "toadd: ";
          print_string (string_of_bool toadd);
          print_string "-------\n";

          engine_loop s i (j + 1) (current_number ^ (String.make 1 current_s)) (checktoadd || toadd) total
        | _ ->
          if toadd then
            let new_total = total + (int_of_string current_number) in
            engine_loop s i (j + 1) "" false new_total
          else
            engine_loop s i (j + 1) "" false total
  in 

  List.mapi (
    fun i s ->
      print_string s;
      print_string "\n";
      if i == 0 || i == ((List.length pslist) - 1) then 0
      else
        engine_loop s i 1 "" false 0
  ) pslist
;;

let totals = engine_schematic pslist;;

let sums = List.fold_left (fun acc x -> acc + x) 0 totals;;