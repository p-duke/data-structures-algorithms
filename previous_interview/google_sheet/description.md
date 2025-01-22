# Google Sheet

Today we’re going to build a basic spreadsheet application like Google sheets or Excel but
much simpler. Our spreadsheet, let’s call it OpenSheet, will only support cells which hold either
integers or formulas that sum two cells.

Example of how your setter could look
set_cell("C1", "45")
set_cell("B1", "10")
set_cell("A1", "=C1+B1")

Example of how your getter could look
get_cell("A1") # should return 55 in this case

Extra
- Figure out how to process any operator (e.g. +, -, /, x)
- Figure out how to handle more complicated expressions
