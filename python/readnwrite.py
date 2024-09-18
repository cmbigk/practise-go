# Open the input file for reading
with open("haha.txt", "r") as input_file:
    # Create a new output file for writing
    with open("output.txt", "w") as output_file:
        # Read each line from the input file and write it to the output file
        for line in input_file:
            output_file.write(line)
