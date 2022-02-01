Interfaces:

split large csv file into smaller files:
dt split file # splits a file in 2 and appends numbers to new files, file1 and file2, placing in current location
dt split -n x file # splits a file into x files

concatenate csv files into a single file:
dt concat -name x file1 file2 ... concatenate csv files into a single file and name it x, placing in current location