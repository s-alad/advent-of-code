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

let find_part_numbers(i: int)(j: int)(pslist: string list)(base: char) =

  let rec try_find_left j s = 
    if j == 0 then s
    else
      let left = String.get (List.nth pslist i) (j - 1) in
      match left with
      | '0' .. '9' -> 
        try_find_left (j - 1) (String.make 1 left ^ s)
      | _ -> s 
  in 

  let rec try_find_right j s = 
    if j == (String.length (List.nth pslist i) - 1) then s
    else
      let right = String.get (List.nth pslist i) (j + 1) in
      match right with
      | '0' .. '9' -> 
        try_find_right (j + 1) (s ^ String.make 1 right)
      | _ -> s
  in

  let left_side = try_find_left j "" in
  let right_side = try_find_right j "" in

  left_side ^ String.make 1 base ^ right_side
;;



let engine_schematic(pslist: string list) = 
  let rec engine_loop(s: string)(i: int)(j: int)(current_gears: 'a list) =

      (*print i, j*)
      print_int i;
      print_string ",";
      print_int j;
      print_string "\n";


      if j == (String.length s - 1) then 
        current_gears
      else 
        let topleft_s = ((i-1), (j-1), String.get (List.nth pslist (i - 1)) (j - 1)) in
        let top_s = ((i-1), (j), String.get (List.nth pslist (i - 1)) j) in
        let topright_s = ((i-1), (j+1), String.get (List.nth pslist (i - 1)) (j + 1)) in
        let left_s = ((i), (j-1), String.get s (j - 1)) in
        let right_s = ((i), (j+1), String.get s (j + 1)) in
        let bottomleft_s = ((i+1), (j-1), String.get (List.nth pslist (i + 1)) (j - 1)) in
        let bottom_s = ((i+1), (j), String.get (List.nth pslist (i + 1)) j) in
        let bottomright_s = ((i+1), (j+1), String.get (List.nth pslist (i + 1)) (j + 1)) in

        let surrounding = [
          topleft_s; top_s; topright_s;
          left_s; right_s;
          bottomleft_s; bottom_s; bottomright_s
        ] in

        let current_s = (i, j, String.get s j) in

        (* print the surrounding list *)
        List.iter (fun (i,j,s) -> print_char s; print_char ' ') [topleft_s; top_s; topright_s;];
        print_string "\n";
        List.iter (fun (i,j,s) -> print_char s; print_char ' ') [left_s; current_s; right_s;];
        print_string "\n";
        List.iter (fun (i,j,s) -> print_char s; print_char ' ') [bottomleft_s; bottom_s; bottomright_s;];
        print_string "\n";
        print_string "-------\n";

        match current_s with
        | (i, j, '*') ->
          
          let surrounding_valid_numbers = List.filter (
            fun (i, j, s) -> is_digit s
          ) surrounding in

          (*print surrounding valid numbers*)
          print_string "surrounding valid numbers: ";
          List.iter (fun (i,j,s) -> print_char s; print_char ' ') surrounding_valid_numbers;
          print_string "\n";

          (* if (i=-1, j=-1) is a number or (i=-1, j=+1) is a number, omit (i=-1, j=0). If (i=-1, j=0) is a number, omit the others *)
          
          let rec cleaner (surr_nums: (int * int * int * int * char) list) (keep: (int * int * char) list)  (middleflag: bool) (rightflag: bool) (i:int) (j:int) = 
            match surr_nums with
            | [] -> keep  
            | (-1, -1, x, y, s) :: tl ->      cleaner tl ((x, y, s) :: keep) true rightflag i j
            | (-1, 0, x, y, s) :: tl ->
              if middleflag then              cleaner tl keep middleflag true i j
              else                            cleaner tl ((x, y, s) :: keep) middleflag true i j
            | (-1, 1, x, y, s) :: tl ->
              if rightflag then               cleaner tl keep false false i j
              else                            cleaner tl ((x, y, s) :: keep) false false i j

            | (1, -1, x, y, s) :: tl ->       cleaner tl ((x, y, s) :: keep) true rightflag i j
            | (1, 0, x, y, s) :: tl -> 
              if middleflag then              cleaner tl keep middleflag true i j
              else                            cleaner tl ((x, y, s) :: keep) middleflag true i j
            | (1, 1, x, y, s) :: tl -> 
              if rightflag then               cleaner tl keep false false i j
              else                            cleaner tl ((x, y, s) :: keep) false false i j

            | (_, _, x, y, s) :: tl ->        cleaner tl ((x, y, s) :: keep) middleflag rightflag i j
          in
          
          let adjusted_surr_nums = 
            List.map(
              fun (x,y,s) -> (x-i, y-j, x, y, s) 
            ) surrounding_valid_numbers 
          in

          (*print adjusted surrounding numbers*)
          print_string "adjusted surrounding numbers: ";
          List.iter (fun (i,j,x,y,s) -> print_int i; print_char ','; print_int j; print_char '|'; print_int x; print_char ','; print_int y; print_char '|'; print_char s; print_char ' ') adjusted_surr_nums;
          print_string "\n";

          let cleaned = cleaner adjusted_surr_nums [] false false i j in

          (*print cleaned*)
          print_string "cleaned: ";
          List.iter (fun (i,j,s) -> print_int i; print_char ','; print_int j; print_char '|'; print_char s; print_char ' ') cleaned;
          print_string "\n";

          let surrounding_found_gears = List.map (
            fun (i, j, s) -> find_part_numbers i j pslist s
          ) cleaned in

          (*print surrounding found gears*)
          print_string "surrounding found gears: ";
          List.iter (fun s -> print_string s; print_char ' ') surrounding_found_gears;
          print_string "\n";
          
          if List.length surrounding_found_gears == 2 then 
            engine_loop s i (j + 1) ([surrounding_found_gears] @ current_gears)
          else
            engine_loop s i (j + 1) current_gears

        | (i, j, _) -> engine_loop s i (j + 1) current_gears
  in 

  List.mapi (
    fun i s ->
      print_string s;
      print_string "\n";
      if i == 0 || i == ((List.length pslist) - 1) then []
      else
        engine_loop s i 1 []
  ) pslist
;;

let totals = engine_schematic pslist;;

let muls = List.map (
  fun l -> 
    if List.length l == 0 then 0
    else
      List.fold_left (
        fun acc inner_list ->
          let muls = int_of_string (List.nth inner_list 0) * int_of_string (List.nth inner_list 1) in
          acc + muls
      ) 0 l
) totals;;

let added = List.fold_left (
  fun acc i -> acc + i
) 0 muls;;