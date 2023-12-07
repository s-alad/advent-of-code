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
