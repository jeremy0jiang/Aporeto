/**
 *@author Shenghao Jiang
 *@since 04-17-2016
 *@version 1.0
 */

/**
 *The script/program should output a file with name provided in 
 *the args with all duplicate lines removed.
 */

package main

import (
	"flag"
	"bufio"
	"fmt"
	"log"
	"os"
)

/**
 *method:reomve
 *This memthod will remove dupulicate lines in the input string
 *and output strings that duplicates have been removed
 *@param stirng
 *@return string 
 */
func remove(elements []string) []string {
    encountered := map[string]bool{}

    // Create a map of all unique elements.
    for v:= range elements {
	encountered[elements[v]] = true
    }

    // Place all keys from the map into a slice.
    result := []string{}
    for key, _ := range encountered {
	result = append(result, key)
    }
    return result
}

/**
 *method: readLines
 *This method reads a file line by line and return stored lines
 *@param path
 *@return lines
 */
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

/**
 *method writeLines
 *This method writes the lines to the given file.
 *@param lines
 */
func writeLines(lines []string, path string) error {
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }
  return w.Flush()
}

/**
 *method main
 *This method call functions and perfrom the task
 */
func main(){
	
	var infile string              // input file name
  flag.StringVar(&infile,"file","input_file","a file name")

	var outfile string             // ouput file name
	
	// read flags 
	flag.StringVar(&outfile,"output","output_file","a file name")
	
	verbosePtr := flag.Bool("verbose",false, "a bool")

	flag.Parse()
	
	//read inputfile and report error if any
		lines, err := readLines(infile)
    if err != nil {
    log.Fatalf("readLines: %s", err)
  }
	
	//check verbose option
	if *verbosePtr == true{
		fmt.Println("File Successfully Open")
  }
	
	//check verbose option
	result :=remove(lines)
	if *verbosePtr == true{
		fmt.Println("Duplicate Successfully Removed")
  }
	
	//write file and report err if any
  if err := writeLines(result, outfile); err != nil {
    log.Fatalf("writeLines: %s", err)
  }
  if *verbosePtr == true{
		fmt.Println("File Successfully Written")
  }
}
