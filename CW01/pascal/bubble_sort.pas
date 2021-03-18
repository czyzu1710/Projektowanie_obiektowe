PROGRAM bubble_sort;

VAR
	list : array[0..99] of Integer;
	i, j, n, tmp: Integer;

PROCEDURE bubble;
BEGIN
	For i:= 0 To n -1  Do
	Begin
		For j:=0 To n - i -2  Do
		Begin
			if (list[j] > list[j+1]) then
				Begin
				tmp:=list[j+1];
				list[j+1] := list[j];
				list[j] := tmp;
				End;
			End;
	End;
END;


BEGIN
	Randomize();
	n := 100;	
	For i:=0 To n - 1 Do
	Begin
		list[i] := Random(n) + 1;
	End;
	writeln('Lista na początku: ');
	For i:=0 To n - 1 Do
	Begin
		write(list[i]);
		write(', ');
		if (i mod 10 = 9) then
			writeln();
	End;
	writeln();
	bubble();
	writeln('Lista po posortowaniu: ');
	For i:=0 To n - 1 Do
	Begin
		write(list[i]);
		write(', ');
		if (i mod 10 = 9) then
			writeln();
	End;
END.
